package message

// attachment reference to the file on any file system
type AttachmentRef struct {
	Filename  string `json:"filename,omitempty"`             // file name
	Reference string `json:"reference" validate:"required"`  // required
	Size      int64  `json:"size" validate:"required,min=1"` // file size (at least 1 byte)
}
