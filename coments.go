package zapi

import (
	"time"
)

type Comment struct {
	Id          int64                  // yes 	Automatically assigned when the comment is created
	Type        string                 // yes 	Has the value Comment
	Body        string                 // yes 	The comment string
	HtmlBody    string                 // yes 	The comment formatted as HTML
	Public      bool                   // no 	  true if a public comment; false if an internal note
	AuthorId    int64                  // yes 	The id of the comment author
	Attachments []Attachment           // yes 	Attachments, if any. See Attachment
	Via         map[string]interface{} // yes 	How the comment was created. See Via Object
	Metadata    map[string]interface{} // yes 	System information (web client, IP address, etc.)
	CreatedAt   time.Time              // yes 	The time the comment was created
}

func (t *Ticket) Comments() (comments []Comment, err error) {
	return comments, nil
}

func (c *Comment) Redact(text string) (Comment, error) {
	return c, nil
}

func (c *Comment) MakePrivate() (Comment, error) {
	return c, nil
}
