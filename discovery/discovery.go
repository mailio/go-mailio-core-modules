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
	// Flag telling if server  at domain supports mailio protocol
	IsMailio bool `json:"isMailio"`
	// Mailio base64 public key (if Mailio flag is true)
	PublicKey string `json:"publicKey,omitempty"`
	// The type of the public key (ed25519 support only at the moment)
	PublicKeyType string `json:"publicKeyType,omitempty"`
}
