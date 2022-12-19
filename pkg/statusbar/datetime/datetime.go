package datetime

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Block struct {
	format   string
	label    string
	timezone string
}

type Parameters struct {
	Format   string
	Label    string
	Timezone string
}

func New(p Parameters) *Block {
	if p.Timezone == "" {
		p.Timezone = "UTC"
	}

	return &Block{
		format:   p.Format,
		label:    p.Label,
		timezone: p.Timezone,
	}
}

func (b *Block) Name() string {
	return "datetime"
}

func (b *Block) Render(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "date", b.format)
	cmd.Env = append(cmd.Env, fmt.Sprintf("TZ=%s", b.timezone))

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to exec: %v", err)
	}

	if b.label != "" {
		return fmt.Sprintf(
			"%s %s", b.label, strings.TrimSpace(string(out)),
		), nil
	} else {
		return strings.TrimSpace(string(out)), nil
	}
}
