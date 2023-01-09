package store

import "os"

const (
	fm  os.FileMode = 0700
	ext             = ".json"
)

func hasExt(s string) bool {
	if len(s) <= len(ext) {
		return false
	}
	if s[len(s)-len(ext):] != ext {
		return false
	}

	return true
}
