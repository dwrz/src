package wwan

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"path"
	"strings"
)

type Block struct {
	iface string
}

func New(iface string) *Block {
	return &Block{iface: iface}
}

func (b *Block) Name() string {
	return "wwan"
}

// TODO: signal strength icon.
// TODO: get IP address (net.Interfaces).
func (b *Block) Render() (string, error) {
	out, err := exec.Command("mmcli", "--list-modems").Output()
	if err != nil {
		return "", fmt.Errorf("exec mmcli failed: %v", err)
	}

	fields := strings.Fields(string(out))
	if len(fields) == 0 {
		return "", fmt.Errorf("unexpected output: %v", err)
	}
	modem := path.Base(fields[0])

	out, err = exec.Command(
		"mmcli", "-m", modem, "--output-keyvalue",
	).Output()
	if err != nil {
		return "", fmt.Errorf("exec mmcli failed: %v", err)
	}

	var (
		scanner = bufio.NewScanner(bytes.NewReader(out))

		state, signal string
	)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}

		switch {
		case fields[0] == "modem.generic.state":
			state = fields[2]
		case fields[0] == "modem.generic.signal-quality.value":
			signal = fields[2]
		}
	}

	if state == "disabled" {
		return " ", nil
	}

	// Get the IP address.
	iface, err := net.InterfaceByName(b.iface)
	if err != nil {
		return "", fmt.Errorf(
			"failed to get interface %s: %v", b.iface, err,
		)
	}

	var ip4, ip6 string
	addrs, err := iface.Addrs()
	if err != nil {
		return "", fmt.Errorf("failed to get addresses: %v", err)
	}

	for _, addr := range addrs {
		a := addr.String()
		if isIPv4(a) {
			ip4 = a
			continue
		} else {
			ip6 = a
		}
	}

	switch {
	case ip4 == "" && ip6 == "":
		return fmt.Sprintf(" %s%%", signal), nil
	case ip4 != "" && ip6 == "":
		return fmt.Sprintf(" %s %s%%", ip4, signal), nil
	case ip4 == "" && ip6 != "":
		return fmt.Sprintf(" %s %s%%", ip6, signal), nil
	default:
		return fmt.Sprintf(" %s %s %s%%", ip4, ip6, signal), nil
	}
}

func isIPv4(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return true
		case ':':
			return false
		}
	}

	return false
}
