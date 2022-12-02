package digip

import (
	"context"
	"testing"
)

func TestGoogleV4(t *testing.T) {
	ctx := context.Background()

	ip, err := Google(ctx, V4)
	if err != nil {
		t.Errorf("google v4 failed: %v", err)
		return
	}

	t.Logf("got ip: %s", ip)
}

func TestGoogleV6(t *testing.T) {
	ctx := context.Background()

	ip, err := Google(ctx, V6)
	if err != nil {
		t.Errorf("google v6 failed: %v", err)
		return
	}

	t.Logf("got ip: %s", ip)
}

func TestOpenDNSV4(t *testing.T) {
	ctx := context.Background()

	ip, err := OpenDNS(ctx, A)
	if err != nil {
		t.Errorf("opendns A failed: %v", err)
		return
	}

	t.Logf("got ip: %s", ip)
}

func TestOpenDNSV6(t *testing.T) {
	ctx := context.Background()

	ip, err := OpenDNS(ctx, AAAA)
	if err != nil {
		t.Errorf("opendns AAAA failed: %v", err)
		return
	}

	t.Logf("got ip: %s", ip)
}
