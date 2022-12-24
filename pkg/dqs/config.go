package dqs

import "fmt"

type Config struct {
	Dir string
}

func (c *Config) Validate() error {
	if c.Dir == "" {
		return fmt.Errorf("missing dqs directory")
	}

	return nil
}
