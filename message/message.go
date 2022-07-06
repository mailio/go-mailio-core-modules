package message

import (
	"io"
)

type ContentType string
type ContentSource string

const (
	ENCRYPTED_HTML ContentType = "enc/html/text" // encrypted HTML content
	MIME           ContentType = "mime"          // MailioMime (normally encrypted and signed on receiving by the mailio server)

	SOURCE_IPFS ContentSource = "ipfs" // content is stored in IPFS
	SOURCE_S3   ContentSource = "s3"   // content is stored in S3
)

// Basic message envelope internal structure (compatible with golang validator)
// In case content type HTML or PLAIN, Both //ContentRef and //AttachmentRefs are in the shape of ContentRef struct
// In case content type MIME, //ContentRef is in the shape of MailioMime struct
type Envelope struct {
	ID                string        `json:"id" validate:"required"`                                   // message id
	SenderPublicKey   string        `json:"senderPublicKey,omitempty"`                                // base64 senders public key
	ReceiverPublicKey string        `json:"receiverPublicKey,omitempty"`                              // base64 receivers public key
	MessageID         string        `json:"messageId" validate:"required,min=3"`                      // RFC2822 message id
	ParentMessageID   string        `json:"parentMessageId,omitempty"`                                // Parent RFC2822 message id
	OwnerAddressHex   string        `json:"ownerAddressHex" validate:"required,len=42"`               // the owner of this message on the current system (usually receivers mailio address in hex format)
	Read              bool          `json:"read"`                                                     // if message has been read by the receiver
	Folder            string        `json:"folder" validate:"required"`                               // folder message resides in
	ContentType       ContentType   `json:"contentType" validate:"required,oneof=enc/html/text mime"` // content type of the message
	ContentRef        *string       `json:"contentRef,omitempty"`                                     // base64 encrypted reference to a content object. Required BoxSharedKey to decrypt
	ContentSource     ContentSource `json:"contentSource,omitempty"`                                  // source of the content (IPFS, AMAZON S3, ...)
	BoxSharedKey      string        `json:"boxSharedKey,omitempty"`                                   // base64 (sealed curve25519+xsalsa20+poly1305 encryption key to unbox content and attachment AES CGM keys. First 24 bytes is nonce)
	AttachmentRefs    []*string     `json:"attachmentRefs,omitempty"`                                 // base64 message attachment references (encrypted pointers to files on a file system. Rquired BoxSharedKey to decrypt)
	Created           int64         `json:"created" validate:"required"`                              // message creation timestamp (unix epoch miliseconds)
	Modified          int64         `json:"modified,omitempty"`                                       // message modification timestamp (unix epoch miliseconds)
}

type ContentRef struct {
	Filename  string `json:"filename,omitempty"`             // file name (optional)
	Hash      string `json:"hash,omitempty"`                 // base64 md5 hash of the file (quick check that file hasn't changed)
	Reference string `json:"reference" validate:"required"`  // required
	Size      int64  `json:"size" validate:"required,min=1"` // file size (at least 1 byte)
	AesKey    string `json:"aesKey,omitempty"`               // base64 (AES GCM key to unlock content)
}

// mailio message parser
type MessageParser interface {
	Parse(reader io.Reader) (*Envelope, error)
}

// mailio message composer where Set functions are required, while Add functions operate on optional fields
type MessageComposer interface {
	// optional fields
	AddID(string)
	AddSenderPublicKey(string)
	AddReceiverPublicKey(string)
	AddParentMessageID(string)
	AddRead(bool)
	AddFolder(string)
	AddOwnersHexAddress(string)
	AddSignature(string)
	AddAttachments([]ContentRef)
	AddModified(int64)
	AddSealedAesKey(string)

	// required fields
	SetContentType(ContentType)
	SetBody(content []byte, contentType string)
	SetMessageID(string)
	SetCreated(int64)

	// pretty print of an object
	ToString()
}
