package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"code.dwrz.net/src/pkg/dqs/entry"
)

func (s *Store) DeleteEntry(date string) error {
	name := fmt.Sprintf("%s/%s.json", s.Dir, date)

	if err := os.Remove(name); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAllEntries() ([]*entry.Entry, error) {
	files, err := os.ReadDir(s.Dir)
	if err != nil {
		return nil, err
	}

	var (
		echan       = make(chan *entry.Entry, len(files))
		errs        = make(chan error, len(files))
		concurrency = 24
		sem         = make(chan struct{}, concurrency)
		wg          sync.WaitGroup
	)

	wg.Add(len(files))

	for _, f := range files {
		go func(name string) {
			defer func() { <-sem }()
			defer wg.Done()

			sem <- struct{}{}

			if name == userFile {
				return
			}

			date := name[:len(name)-len(filepath.Ext(name))]

			entry, err := s.GetEntry(date)
			if err != nil {
				errs <- err
			}

			echan <- entry
		}(f.Name())
	}

	wg.Wait()
	close(errs)
	close(echan)

	if len(errs) > 0 {
		return nil, <-errs
	}

	var entries = []*entry.Entry{}
	for entry := range echan {
		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date.Before(entries[j].Date)
	})

	return entries, nil
}

func (s *Store) GetEntry(date string) (*entry.Entry, error) {
	name := fmt.Sprintf("%s/%s.json", s.Dir, date)

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open entry file: %w", err)
	}

	var e = &entry.Entry{}
	if err := json.Unmarshal(data, e); err != nil {
		return nil, fmt.Errorf("failed to parse entry file: %w", err)
	}

	return e, nil
}

func (s *Store) UpdateEntry(e *entry.Entry) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}

	name := fmt.Sprintf(
		"%s/%s.json", s.Dir, e.Date.Format(entry.DateFormat),
	)

	if err := os.WriteFile(name, data, permissions); err != nil {
		return err
	}

	return nil
}
