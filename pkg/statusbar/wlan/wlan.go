package wlan

import (
	"bufio"
	"bytes"
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
	return "wlan"
}

// TODO: signal strength icon.
func (b *Block) Render(ctx context.Context) (string, error) {
	out, err := exec.CommandContext(
		ctx, "iwctl", "station", "wlan0", "show",
	).Output()
	if err != nil {
		return "", fmt.Errorf("exec iwctl failed: %v", err)
	}

	var (
		scanner = bufio.NewScanner(bytes.NewReader(out))

		state, network, ip, rssi string
	)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch {
		case fields[0] == "State":
			state = fields[1]
		case fields[0] == "Connected":
			network = fields[2]
		case fields[0] == "IPv4":
			ip = fields[2]
		case fields[0] == "RSSI":
			rssi = fields[1] + fields[2]
		}
	}

	if state == "disconnected" {
		return " ", nil
	}
	return fmt.Sprintf(" %s %s %s", network, ip, rssi), nil
}
