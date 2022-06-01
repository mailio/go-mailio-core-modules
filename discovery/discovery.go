package discovery

import "context"

type DnsDiscoverer interface {
	Discover(ctx context.Context, domain string) (*Discovery, error)
}

// general interface for discovery
type Discoverer interface {
	DnsDiscoverer
}

// DnsDiscovery is a Discovery that scans for DNS records.
type Discovery struct {
	// The domain name to query.
	Domain string `json:"domain"`
	// Ip of the requested domain
	Ip string `json:"ip"`
	// Flag is server is mailio server
	Mailio bool `json:"mailio"`
	// Mailio base64 public key (if Mailio flag is true)
	MailioPublicKey string `json:"mailioPublicKey,omitempty"`
	// List of domains mailio server is authoritative for
	MailioDomainList string `json:"mailioDomainList,omitempty"`
}
