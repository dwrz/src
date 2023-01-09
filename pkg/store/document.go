package store

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Document struct {
	Created time.Time       `json:"created"`
	Id      string          `json:"id"`
	Updated time.Time       `json:"updated"`
	Data    json.RawMessage `json:"data"`

	path string
}

func (d *Document) Update(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to json marshal: %w", err)
	}

	d.Data = data
	d.Updated = time.Now()

	data, err = json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to json marshal: %w", err)
	}

	if err := os.WriteFile(d.path, data, fm); err != nil {
		return fmt.Errorf("failed to write file %s: %w", d.path, err)
	}

	return nil
}

func (d *Document) Unmarshal(v any) error {
	if err := json.Unmarshal(d.Data, v); err != nil {
		return err
	}

	return nil
}
