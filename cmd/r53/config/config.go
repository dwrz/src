package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	//go:embed config.json
	configuration []byte
	ttl           = flag.String("ttl", "300", "TTL")
)

type Config struct {
	AWS          *aws.Config `json:"-"`
	HostedZoneId string      `json:"hostedZoneId"`
	TTL          int64       `json:"-"`
}

func New(ctx context.Context) (*Config, error) {
	flag.Parse()

	var cfg = &Config{}
	if err := json.Unmarshal(configuration, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	awscfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve AWS config: %v", err)
	}
	cfg.AWS = &awscfg

	if n, err := strconv.ParseInt(*ttl, 10, 64); err != nil {
		return nil, fmt.Errorf("failed to parse ttl: %v", err)
	} else {
		cfg.TTL = n
	}

	return cfg, nil
}
