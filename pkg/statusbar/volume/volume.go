package volume

import (
	"fmt"
	"os/exec"
	"strings"
)

type Block struct{}

func New() *Block {
	return &Block{}
}

func (b *Block) Name() string {
	return "volume"
}

func (b *Block) Render() (string, error) {
	out, err := exec.Command(
		"pactl", "get-sink-mute", "@DEFAULT_SINK@",
	).Output()
	if err != nil {
		return "", fmt.Errorf("exec pactl failed: %v", err)
	}

	if strings.Contains(string(out), "yes") {
		return fmt.Sprintf(""), nil
	}

	out, err = exec.Command(
		"pactl", "get-sink-volume", "@DEFAULT_SINK@",
	).Output()
	if err != nil {
		return "", fmt.Errorf("exec pactl failed: %v", err)
	}

	if fields := strings.Fields(string(out)); len(fields) < 5 {
		return fmt.Sprintf(" "), nil
	} else {
		return fmt.Sprintf(" %s", fields[4]), nil
	}
}
