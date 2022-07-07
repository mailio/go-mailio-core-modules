package mailbox

type Folder struct {
	ID             string `json:"id"`                       // id of the folder
	Name           string `json:"name" validate:"required"` // name is required
	TotalMessages  int64  `json:"totalMessages"`            // total number of messages in the folder
	UnreadMessages int64  `json:"unreadMessages"`           // total number of unread messages in the folder
}
