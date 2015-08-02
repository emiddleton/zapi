package zapi

import ()

type TicketComments struct {
}

type TicketComment struct {
	Id          int64                  // Automatically assigned when the comment is created (readonly)
	Type        string                 // Has the value Comment (readonly)
	Body        string                 // The comment string (readonly)
	HtmlBody    string                 // The comment formatted as HTML (readonly)
	Public      bool                   // true if a public comment; false if an internal note
	AuthorId    int64                  // The id of the comment author (readonly)
	Attachments []Attachment           // Attachments, if any. See Attachment (readonly)
	Via         ViaObject              // How the comment was created. See Via Object (readonly)
	Metadata    map[string]interface{} // System information (web client, IP address, etc.) (readonly)
	CreatedAt   Date                   // The time the comment was created (readonly)
}

func (t *Ticket) TicketComments() ([]TicketComment, error) {
	return []TicketComment{}, nil
}

func (tc *TicketComment) Redact(text string) (TicketComment, error) {
	return *tc, nil
}

func (tc *TicketComment) MakePrivate() (TicketComment, error) {
	return *tc, nil
}
