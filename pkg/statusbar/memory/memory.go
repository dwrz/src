package memory

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Block struct{}

func New() *Block {
	return &Block{}
}

func (b *Block) Name() string {
	return "memory"
}

func (b *Block) Render() (string, error) {
	var output strings.Builder

	out, err := exec.Command("free", "-m").Output()
	if err != nil {
		return "", fmt.Errorf("failed to exec: %v", err)
	}

	for _, l := range strings.Split(string(out), "\n") {
		if len(l) == 0 {
			continue
		}

		fields := strings.Fields(l)
		if len(fields) < 3 {
			continue
		}

		switch fields[0] {
		case "Mem:":
			total, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return "", fmt.Errorf(
					"failed to parse total memory: %v", err,
				)
			}
			used, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				return "", fmt.Errorf(
					"failed to parse used memory: %v", err,
				)
			}

			fmt.Fprintf(&output, " %.0f%% ", (used/total)*100)
		case "Swap:":
			total, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return "", fmt.Errorf(
					"failed to parse total memory: %v", err,
				)
			}
			used, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				return "", fmt.Errorf(
					"failed to parse used memory: %v", err,
				)
			}

			fmt.Fprintf(&output, " %.0f%%", (used/total)*100)
		}
	}

	return output.String(), nil
}
