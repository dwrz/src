package disk

import (
	"fmt"
	"os/exec"
	"strings"
)

type Block struct {
	mounts []string
}

func New(mounts ...string) *Block {
	return &Block{mounts: mounts}
}

func (b *Block) Name() string {
	return "disk"
}

func (b *Block) Render() (string, error) {
	out, err := exec.Command("df", "-h").Output()
	if err != nil {
		return "", fmt.Errorf("exec df failed: %v", err)
	}

	var mounts = map[string]string{}
	for i, l := range strings.Split(string(out), "\n") {
		if i == 0 || len(l) == 0 {
			continue
		}

		fields := strings.Fields(l)
		if len(fields) < 6 {
			continue
		}

		var icon rune = ''
		switch fields[5] {
		case "/":
			icon = '/'
		case "/home":
			icon = ''
		}

		mounts[fields[5]] = fmt.Sprintf("%c %s", icon, fields[4])
	}

	var output strings.Builder
	for i, m := range b.mounts {
		if s, exists := mounts[m]; exists {
			output.WriteString(s)
			if i < len(b.mounts)-1 {
				output.WriteRune(' ')
			}
		}
	}

	return output.String(), nil
}
