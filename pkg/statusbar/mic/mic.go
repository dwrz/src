package mic

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
	return "mic"
}

func (b *Block) Render() (string, error) {
	out, err := exec.Command(
		"pactl", "get-source-mute", "@DEFAULT_SOURCE@",
	).Output()
	if err != nil {
		return "", fmt.Errorf("exec pactl failed: %v", err)
	}

	if strings.Contains(string(out), "yes") {
		return fmt.Sprintf(""), nil
	}

	return fmt.Sprintf(" OPEN"), nil
}
