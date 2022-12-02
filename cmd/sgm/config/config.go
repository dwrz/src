package config

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
)

var (
	//go:embed config.json
	configuration []byte

	attachment = flag.String("a", "", "attachment path")
	from       = flag.String("f", "", "sender address")
	subject    = flag.String("s", "", "subject")
	text       = flag.String("m", "", "message text")
	to         = flag.String("t", "", "recipient address, comma delimited")
)

type Config struct {
	Attachment string
	From       string
	Pass       string
	Subject    string
	Text       string
	To         string
	User       string
}

func New() (*Config, error) {
	flag.Parse()

	var cfg = &Config{
		Attachment: *attachment,
		From:       *from,
		Subject:    *subject,
		Text:       *text,
		To:         *to,
	}
	// Load the username and password from the embedded configuration.
	if err := json.Unmarshal(configuration, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return cfg, nil
}
