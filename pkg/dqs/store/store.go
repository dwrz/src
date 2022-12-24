package store

import (
	"fmt"
	"os"
)

const permissions = 0700

type Store struct {
	Dir string
}

func Open(dir string) (*Store, error) {
	var store = &Store{Dir: dir}

	// Create the directory, if it doesn't exist.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModeDir|permissions); err != nil {
			return nil, fmt.Errorf(
				"failed to create dqs directory: %w", err,
			)
		}
	}

	return store, nil
}
