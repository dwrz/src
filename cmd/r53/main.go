package main

import (
	"context"
	"encoding/json"
	"os"

	"code.dwrz.net/src/pkg/digip"
	"code.dwrz.net/src/pkg/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"

	"code.dwrz.net/src/cmd/r53/config"
)

func main() {
	// Setup the logger.
	l := log.New(os.Stderr)

	// Setup the main context.
	ctx, cancel := context.WithCancel(context.Background())

	// Get the service configuration.
	cfg, err := config.New(ctx)
	if err != nil {
		l.Error.Fatalf("failed to get config: %v", err)
	}

	// Create the Route53 client.
	var svc = route53.NewFromConfig(*cfg.AWS)

	// Get the IP address.
	var ip4, ip6 string

	// Dig with OpenDNS.
	if ip, err := digip.OpenDNS(ctx, digip.A); err != nil {
		l.Error.Printf("failed to get opendns ipv4 ip: %v", err)
	} else {
		ip4 = ip.String()
	}
	if ip, err := digip.OpenDNS(ctx, digip.AAAA); err != nil {
		l.Error.Printf("failed to get opendns ipv6 ip: %v", err)
	} else {
		ip6 = ip.String()
	}

	// Fallback on Google.
	if ip4 == "" {
		if ip, err := digip.Google(ctx, digip.V4); err != nil {
			l.Error.Printf("failed to get google ipv4 ip: %v", err)
		} else {
			ip4 = ip.String()
		}
	}
	if ip6 == "" {
		if ip, err := digip.Google(ctx, digip.V6); err != nil {
			l.Error.Printf("failed to get google ipv6 ip: %v", err)
		} else {
			ip6 = ip.String()
		}
	}

	// Abort if we don't have any IP address.
	if ip4 == "" && ip6 == "" {
		l.Error.Fatalf("failed to retrieve public ip")
	}

	// Update the domains.
	for _, domain := range os.Args[1:] {
		// Assemble the change(s).
		var changes []types.Change
		if ip4 != "" {
			change := types.Change{
				Action: types.ChangeActionUpsert,
				ResourceRecordSet: &types.ResourceRecordSet{
					Name: &domain,
					ResourceRecords: []types.ResourceRecord{
						{Value: &ip4},
					},
					TTL:  &cfg.TTL,
					Type: types.RRTypeA,
				},
			}

			changes = append(changes, change)
		}
		if ip6 != "" {
			change := types.Change{
				Action: types.ChangeActionUpsert,
				ResourceRecordSet: &types.ResourceRecordSet{
					Name: &domain,
					ResourceRecords: []types.ResourceRecord{
						{Value: &ip6}},
					TTL:  &cfg.TTL,
					Type: types.RRTypeAaaa,
				},
			}
			changes = append(changes, change)
		}

		input := &route53.ChangeResourceRecordSetsInput{
			ChangeBatch: &types.ChangeBatch{
				Changes: changes,
				Comment: aws.String("DDNS"),
			},
			HostedZoneId: &cfg.HostedZoneId,
		}

		result, err := svc.ChangeResourceRecordSets(ctx, input)
		if err != nil {
			l.Error.Fatalf("failed to update route 53: %v", err)
		}

		outcome, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			l.Error.Printf("failed to unmarshal response: %v", err)

			outcome = []byte(result.ChangeInfo.Status)
		}

		l.Debug.Printf(
			"updated %s route 53 records:\n%s", domain, outcome,
		)
	}

	// Cancel the main context.
	cancel()
}
