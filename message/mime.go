package message

// mime header
type MimeHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// common mime parts
type MimeObject struct {
	ContentType        string        `json:"contentType"`
	ContentDisposition string        `json:"contentDisposition,omitempty"`
	Headers            []*MimeHeader `json:"headers,omitempty"`
}

// mailio mime message (SMTP support)
type MailioMime struct {
	MimeObject
	Subject       string      `json:"subject,omitempty"`      // optional mime subject
	Sender        string      `json:"sender"`                 // sender email
	ReplyTo       string      `json:"replyTo,omitempty"`      // reply-to mime field
	MessageID     string      `json:"messageId"`              // mime message id
	DateAsString  string      `json:"dateAsString,omitempty"` // date (RFC ... )
	To            []string    `json:"to"`                     // to email addresses
	Cc            []*string   `json:"cc,omitempty"`           // optinal cc addresses
	Bcc           []*string   `json:"bcc,omitempty"`          // list of bcc email addresses
	AllRecipients []*string   `json:"allRecipients"`          // list of all recipinets on current system
	MimePart      *MimeObject `json:"mimePart,omitempty"`     // single mime part
	Body          MimeObject  `json:"body"`                   // email body
	Date          string      `json:"date,omitempty"`         // string date
}
