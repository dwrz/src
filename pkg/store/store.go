package store

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"code.dwrz.net/src/pkg/log"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Parameters struct {
	Path string
	Log  *log.Logger
}

type Store struct {
	log  *log.Logger
	path string

	mu          sync.Mutex
	collections map[string]*Collection
}

func New(p Parameters) (*Store, error) {
	var s = &Store{
		collections: map[string]*Collection{},
		log:         p.Log,
		path:        p.Path,
	}

	// Create the directory, if it doesn't exist.
	if _, err := os.Stat(p.Path); os.IsNotExist(err) {
		if err := os.MkdirAll(p.Path, os.ModeDir|fm); err != nil {
			return nil, fmt.Errorf(
				"failed to create directory: %w", err,
			)
		}
	}

	// Seed existing collections.
	entries, err := os.ReadDir(s.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		if err := s.NewCollection(e.Name()); err != nil {
			return nil, fmt.Errorf(
				"failed to create new collection: %w", err,
			)
		}
	}

	return s, nil
}
