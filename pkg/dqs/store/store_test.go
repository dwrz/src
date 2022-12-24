package store

import (
	"os"
	"testing"
	"time"

	"code.dwrz.net/src/pkg/dqs/diet"
	"code.dwrz.net/src/pkg/dqs/entry"
)

func BenchmarkGetAllEntries(b *testing.B) {
	b.StopTimer()

	temp, err := os.MkdirTemp(os.TempDir(), "dqs")
	if err != nil {
		b.Errorf("failed to setup temp dir: %v", err)
		return
	}

	s, err := Open(temp)
	if err != nil {
		b.Errorf("failed to open store: %v", err)
		return
	}

	var start = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	var end = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	for t := start; t.Before(end); t = t.AddDate(0, 0, 1) {
		if err := s.UpdateEntry(&entry.Entry{
			Categories: diet.Vegetarian.Template(),
			Date:       t,
			Diet:       diet.Vegetarian,
		}); err != nil {
			b.Errorf("failed to create entry: %v", err)
			return
		}
	}

	b.ResetTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := s.GetAllEntries(); err != nil {
			b.Errorf("failed to get all entries: %v", err)
			return
		}
	}
	b.StopTimer()

	// if err := os.RemoveAll(temp); err != nil {
	// 	b.Errorf("failed to remove %s directory: %v", temp, err)
	// 	return
	// }
}
