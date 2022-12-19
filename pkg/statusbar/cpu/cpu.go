package cpu

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const path = "/proc/loadavg"

type Block struct{}

func New() *Block {
	return &Block{}
}

func (b *Block) Name() string {
	return "cpu"
}

func (b *Block) Render(ctx context.Context) (string, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %v", path, err)
	}

	if fields := strings.Fields(string(out)); len(fields) >= 1 {
		return fmt.Sprintf(" %s/%d", fields[0], runtime.NumCPU()), nil
	} else {
		return fmt.Sprintf(" /%d", runtime.NumCPU()), nil
	}
}
