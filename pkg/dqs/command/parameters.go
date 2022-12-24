package command

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/dqs/store"
)

type Parameters struct {
	Args  []string
	Date  time.Time
	Store *store.Store
}

func (p *Parameters) Validate() error {
	if p.Args == nil {
		return fmt.Errorf("missing arguments")
	}
	if p.Store == nil {
		return fmt.Errorf("missing store")
	}

	return nil
}
