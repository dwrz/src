package eth

import (
	"context"
	"fmt"
	"net"
)

type Block struct {
	iface string
}

func New(iface string) *Block {
	return &Block{iface: iface}
}

func (b *Block) Name() string {
	return "eth"
}

func (b *Block) Render(ctx context.Context) (string, error) {
	iface, err := net.InterfaceByName(b.iface)
	if err != nil {
		if err.Error() == "route ip+net: no such network interface" {
			return fmt.Sprintf(" "), nil
		}
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
		return fmt.Sprintf(" "), nil
	case ip4 != "" && ip6 == "":
		return fmt.Sprintf(" %s", ip4), nil
	case ip4 == "" && ip6 != "":
		return fmt.Sprintf(" %s", ip6), nil
	default:
		return fmt.Sprintf(" %s %s", ip4, ip6), nil
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
