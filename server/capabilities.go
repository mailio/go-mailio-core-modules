package server

const (
	// Possible protocols
	PROTOCOL_HTTPS = "https"
	PROTOCOL_PROTO = "proto"

	PROTOCOL_PURPOSE_REST_API = "restapi"
	PROTOCOL_PURPOSE_EXCHANGE = "exchange"
)

// ServerCapabilites
type ServerCapabilities struct {
	EmailDomains                     []string          `json:"emailDomains,omitempty"`                                     // supported email domains
	SupportedProtocols               []*ServerProtocol `json:"supportedProtocols" validate:"required,min=2"`               // list of supported protocols (at least REST and Proto protocols)
	MaximumMessageSizeBytes          int64             `json:"maximumMessageSizeBytes" validate:"required,min=1"`          // maximum message size in bytes
	MaximumSingleAttachmentSizeBytes int64             `json:"maximumSingleAttachmentSizeBytes" validate:"required,min=1"` // maximum single attachment size in bytes
}

// Protocol description, purpose and version
type ServerProtocol struct {
	Name    string `json:"name,omitempty"`                                     // name of the protocol
	Version string `json:"version" validate:"required"`                        // version of the protocol
	Type    string `json:"type" validate:"required,oneof=https proto"`         // type of the protocol
	Purpose string `json:"purpose" validate:"required,oneof=restapi exchange"` // purpose of the protocol
}
