package message

import (
	"io"
)

type ContentType string

const (
	MailioV1 ContentType = "mailio/v1"
	Mime     ContentType = "mime"
)

// Basic email message structure
type Message struct {
	ID              string          `json:"id" validate:"required"`                               // message id
	SenderAddress   string          `json:"senderAddress,omitempty" validate:"len=42"`            // mailio sender address
	ReceiverAddress string          `json:"receiverAddress,omitempty" validate:"len=42"`          // mailio receiver address
	MessageID       string          `json:"messageId" validate:"required,min=3"`                  // RFC2822 message id
	ParentMessageID string          `json:"parentMessageId,omitempty"`                            // Parent RFC2822 message id
	OwnerAddress    string          `json:"ownerAddress" validate:"required,len=42"`              // the owner of this message (can be receivers or senders mailio address)
	Read            bool            `json:"read"`                                                 // if message has been read by the receiver
	Folder          string          `json:"folder" validate:"required"`                           // folder message recides in
	ContentType     ContentType     `json:"contentType" validate:"required,oneof=mailio/v1 mime"` // content type of the message
	Content         []byte          `json:"-" validate:"required,min=1"`                          // message content (ignored in potential json response)
	AttachmentRefs  []AttachmentRef `json:"attachmentRef,omitempty"`                              // message attachment references (pointers to files on a file system)
	Created         int64           `json:"created" validate:"required"`                          // message creation timestamp (unix epoch miliseconds)
	Modified        int64           `json:"modified,omitempty"`                                   // message modification timestamp (unix epoch miliseconds)
}

type MessageParser interface {
	Parse(reader io.Reader) (*Message, error)
}
