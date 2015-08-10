package zapi

import (
	"encoding/json"
	"fmt"
)

type FieldOption struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TicketFields struct {
	path   string
	client *Client
}

func NewTicketFields(client *Client) TicketFields {
	return TicketFields{
		path:   "",
		client: client,
	}
}

type TicketField struct {
	Class               *TicketFields `json:"-"`
	Id                  int64         `json:"id,omitempty"`          // yes 	no 	Automatically assigned upon creation
	Url                 string        `json:"url,omitempty"`         // yes 	no 	The URL for this resource
	Type                string        `json:"type"`                  // no 	yes 	The type of the ticket field: "checkbox", "date", "decimal", "int64", "regexp", "tagger", "text", or "textarea"
	Title               string        `json:"title"`                 // no 	yes 	The title of the ticket field
	RawTitle            string        `json:"raw_title"`             // no 	no 	The dynamic content placeholder, if present, or the "title" value, if not. See Dynamic Content
	Description         string        `json:"description"`           // no 	no 	The description of the purpose of this ticket field, shown to users
	RawDescription      string        `json:"raw_description"`       // no 	no 	The dynamic content placeholder, if present, or the "description" value, if not. See Dynamic Content
	Position            int64         `json:"position"`              // no 	no 	A relative position for the ticket fields, determines the order of ticket fields on a ticket
	Active              bool          `json:"active"`                // no 	no 	Whether this field is available
	Required            bool          `json:"requried"`              // no 	no 	If it's required for this field to have a value when updated by agents
	CollapsedForAgents  bool          `json:"collapsed_for_agents"`  // no 	no 	If this field should be shown to agents by default or be hidden alongside infrequently used fields. Classic interface only
	RegexpForValidation *string       `json:"regexp_for_validation"` // no 	no 	Regular expression field only. The validation pattern for a field value to be deemed valid.
	TitleInPortal       string        `json:"title_in_portal"`       // no 	no 	The title of the ticket field when shown to end users
	RawTitleInPortal    string        `json:"raw_title_in_portal"`   // no 	no 	The dynamic content placeholder, if present, or the "title_in_portal" value, if not. See Dynamic Content
	VisibleInPortal     bool          `json:"visible_in_portal"`     // no 	no 	Whether this field is available to end users
	EditableInPortal    bool          `json:"editable_in_portal"`    // no 	no 	Whether this field is editable by end users
	RequiredInPortal    bool          `json:"required_in_portal"`    // no 	no 	If it's required for this field to have a value when updated by end users
	Tag                 *string       `json:"tag"`                   // no 	no 	A tag value to set for checkbox fields when checked
	CreatedAt           *Date         `json:"created_at,omitempty"`  // yes 	no 	The time the ticket field was created
	UpdatedAt           *Date         `json:"updated_at,omitempty"`  // yes 	no 	The time of the last update of the ticket field
	SystemFieldOptions  []FieldOption `json:"system_field_options"`  // yes 	no 	Presented for a ticket field of type "tickettype", "priority" or "status"
	CustomFieldOptions  []FieldOption `json:"custom_field_options"`  // no 	yes 	Required and presented for a ticket field of type "tagger"
	Removable           bool          `json:"removeable"`            // yes 	no 	If this field is not a system basic field that must be present for all tickets on the account
}

func (tfs *TicketFields) List() (ticketFields []TicketField, err error) {
	path := fmt.Sprintf("%s%s", tfs.path, "/ticket_fields.json")
	responseBody, err := tfs.client.Get(path, nil)
	if err != nil {
		return ticketFields, err
	}
	// fmt.Printf("%s\n", string(responseBody))
	ticketFieldsWrapper := struct {
		TicketFields []TicketField `json:"ticket_Fields"`
	}{}
	if err := json.Unmarshal(responseBody, &ticketFieldsWrapper); err != nil {
		return ticketFields, err
	}
	for _, ticketField := range ticketFieldsWrapper.TicketFields {
		ticketField.Class = tfs
		ticketFields = append(ticketFields, ticketField)
	}
	return ticketFields, err
}
