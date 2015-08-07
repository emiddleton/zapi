package zapi

import (
	"encoding/json"
	// "fmt"
)

type Comments struct {
	path   string
	client *Client
}

func NewComments(path string, client *Client) Comments {
	return Comments{
		path:   path,
		client: client,
	}
}

type Comment struct {
	Class       *Comments              `json:"-"`                   // reference to parent class
	Id          int64                  `json:"id,omitempty"`        // Automatically assigned when the comment is created (readonly)
	Type        string                 `json:"type,omitempty"`      // Has the value Comment (readonly)
	Body        string                 `json:"body,omitempty"`      // The comment string (readonly)
	HtmlBody    string                 `json:"html_body,omitempty"` // The comment formatted as HTML (readonly)
	Public      bool                   `json:"public"`              // true if a public comment; false if an internal note
	AuthorId    int64                  `json:"author_id"`           // The id of the comment author (readonly)
	Attachments []Attachment           `json:"attachments"`         // Attachments, if any. See Attachment (readonly)
	Via         ViaObject              `json:"via"`                 // How the comment was created. See Via Object (readonly)
	Metadata    map[string]interface{} `json:"metadata"`            // System information (web client, IP address, etc.) (readonly)
	CreatedAt   Date                   `json:"created_at"`          // The time the comment was created (readonly)
}

func (cs *Comments) Create(comment Comment) (Comment, error) {

	type commentWrap struct {
		WrappedComment Comment `json:"comment"`
	}
	reqBody, err := json.MarshalIndent(&commentWrap{comment}, "", "    ")
	if err != nil {
		return Comment{}, err
	}
	// fmt.Printf("request ->\n%s\n", string(reqBody))

	responseBody, err := cs.client.Put(cs.path, nil, reqBody)
	if err != nil {
		return Comment{}, err
	}

	// fmt.Printf("response ->\n%s\n", string(responseBody))
	commentWrapper := commentWrap{comment}
	if err := json.Unmarshal(responseBody, &commentWrapper); err != nil {
		return commentWrapper.WrappedComment, err
	}

	return commentWrapper.WrappedComment, nil

}

func (cs *Comments) List() ([]Comment, error) {
	return []Comment{}, nil
}

func (c *Comment) Redact(text string) (Comment, error) {
	return *c, nil
}

func (c *Comment) MakePrivate() (Comment, error) {
	return *c, nil
}
