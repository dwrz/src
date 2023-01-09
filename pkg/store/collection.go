package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"code.dwrz.net/src/pkg/log"
)

type Collection struct {
	log  *log.Logger
	path string
	name string
}

func (s *Store) NewCollection(name string) error {
	var c = &Collection{
		log:  s.log,
		name: name,
		path: filepath.Join(s.path, name),
	}

	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		if err := os.MkdirAll(c.path, os.ModeDir|fm); err != nil {
			return fmt.Errorf(
				"failed to create directory: %w", err,
			)
		}
	}

	s.mu.Lock()
	s.collections[c.name] = c
	s.mu.Unlock()

	return nil
}

func (s *Store) Collection(name string) *Collection {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.collections[name]
}

func (c *Collection) All() ([]*Document, error) {
	names, err := c.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	var (
		wg sync.WaitGroup

		mu   sync.Mutex
		errs []error
		docs []*Document
	)
	for _, name := range names {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()

			d, err := c.FindId(name)
			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf(
					"failed to load document: %v", err,
				))
				mu.Unlock()

				return
			}

			mu.Lock()
			docs = append(docs, d)
			mu.Unlock()
		}(name)
	}

	wg.Wait()

	if len(errs) > 0 {
		for _, err := range errs {
			c.log.Error.Print(err)
		}

		return nil, fmt.Errorf("failed to load all quotes")
	}

	return docs, nil
}

func (c *Collection) Count() (int, error) {
	f, err := os.Open(c.path)
	if err != nil {
		return 0, fmt.Errorf("failed to open: %w", err)
	}
	defer f.Close()

	entries, err := f.ReadDir(0)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory: %w", err)
	}

	var count int
	for _, e := range entries {
		// Nested collections are not allowed.
		if e.IsDir() {
			continue
		}

		// Ignore any files without our file extension.
		if !hasExt(e.Name()) {
			continue
		}

		count++
	}

	return count, nil
}

func (c *Collection) FindId(id string) (*Document, error) {
	if !hasExt(id) {
		id += ext
	}

	path := filepath.Join(c.path, id)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var d = &Document{path: path}
	if err := json.Unmarshal(data, d); err != nil {
		return nil, fmt.Errorf(
			"failed to json unmarshal document %s: %w", path, err,
		)
	}

	return d, nil
}

func (c *Collection) List() ([]string, error) {
	entries, err := os.ReadDir(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var names []string
	for _, e := range entries {
		// Nested collections are not allowed.
		if e.IsDir() {
			continue
		}

		// Ignore any files without a JSON file extension.
		var name = e.Name()
		if !hasExt(name) {
			continue
		}

		names = append(names, name)
	}

	return names, nil
}

func (c *Collection) Create(id string, v any) (*Document, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to json marshal: %w", err)
	}

	now := time.Now()
	var d = &Document{
		Created: now,
		Data:    data,
		Id:      id,
		Updated: now,
	}

	data, err = json.MarshalIndent(d, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to json marshal: %w", err)
	}

	path := filepath.Join(c.path, d.Id+ext)
	if err := os.WriteFile(path, data, fm); err != nil {
		return nil, fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return d, nil
}

func (c *Collection) Delete(id string) error {
	var path = filepath.Join(c.path, id)
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to remove file: %v", err)
	}

	return nil
}

func (c *Collection) Random() (*Document, error) {
	entries, err := os.ReadDir(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("empty collection")
	}

	name := entries[random.Intn(len(entries))].Name()
	path := filepath.Join(c.path, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var d = &Document{
		path: path,
	}
	if err := json.Unmarshal(data, d); err != nil {
		return nil, fmt.Errorf(
			"failed to json unmarshal document %s: %w", path, err,
		)
	}

	return d, nil
}
