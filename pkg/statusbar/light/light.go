package light

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Block struct{}

func New() *Block {
	return &Block{}
}

func (b *Block) Name() string {
	return "light"
}

func (b *Block) Render(ctx context.Context) (string, error) {
	out, err := exec.CommandContext(ctx, "brightnessctl", "-m").Output()
	if err != nil {
		return "", fmt.Errorf("failed to exec: %v", err)
	}

	if fields := strings.Split(string(out), ","); len(fields) >= 4 {
		return fmt.Sprintf(" %s", fields[3]), nil
	} else {
		return " ", nil
	}
}
