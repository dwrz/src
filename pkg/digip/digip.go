package digip

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

type Version string

const (
	V4 = "-4"
	V6 = "-6"
)

type Record string

const (
	A    = "A"
	AAAA = "AAAA"
)

// dig -4 TXT +short o-o.myaddr.l.google.com @ns1.google.com
// dig -6 TXT +short o-o.myaddr.l.google.com @ns1.google.com
func Google(ctx context.Context, v Version) (net.IP, error) {
	cmd := exec.CommandContext(
		ctx,
		"dig",
		"TXT",
		"+short",
		string(v),
		"o-o.myaddr.l.google.com",
		"@ns1.google.com",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("dig failed: %v", err)
	}
	if string(output) == "" {
		return nil, fmt.Errorf("dig failed: no output")
	}
	if len(output) < 2 {
		return nil, fmt.Errorf(
			"dig failed: unexpected output: %s", string(output),
		)
	}

	// Google returns the IP with quotes.
	// We need to trim the quotes.
	trimmed := strings.TrimSpace(string(output))
	ip := net.ParseIP(trimmed[1 : len(trimmed)-1])

	return ip, nil
}

// dig A +short myip.opendns.com @resolver1.opendns.com
// dig AAAA +short myip.opendns.com @resolver1.opendns.com
func OpenDNS(ctx context.Context, r Record) (net.IP, error) {
	cmd := exec.CommandContext(
		ctx,
		"dig",
		"+short",
		string(r),
		"myip.opendns.com",
		"@resolver1.opendns.com",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("dig failed: %v", err)
	}
	if string(output) == "" {
		return nil, fmt.Errorf("dig failed: no output")
	}
	if len(output) < 2 {
		return nil, fmt.Errorf(
			"dig failed: unexpected output: %s", string(output),
		)
	}

	ip := net.ParseIP(strings.TrimSpace(string(output)))

	return ip, nil
}
