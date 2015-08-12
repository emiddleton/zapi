package zapi

import (
	"encoding/json"
	"fmt"
)

type UserFields struct {
	path   string
	client *Client
}

func NewUserFields(client *Client) UserFields {
	return UserFields{
		path:   "",
		client: client,
	}
}

type UserField struct {
	Class               *UserFields   `json:"-"`
	Id                  int64         // yes 	no 	Automatically assigned upon creation
	Url                 string        // yes 	no 	The URL for this resource
	Key                 string        // no 	on create 	A unique key that identifies this custom field. This is used for updating the field and referencing in placeholders.
	Type                string        // no 	yes 	Type of the custom field: "checkbox", "date", "decimal", "dropdown", "integer", "regexp", "text", or "textarea"
	Title               string        // no 	yes 	The title of the custom field
	RawTitle            string        // no 	no 	The dynamic content placeholder, if present, or the "title" value, if not. See Dynamic Content
	Description         string        // no 	no 	User-defined description of this field's purpose
	RawDescription      string        // no 	no 	The dynamic content placeholder, if present, or the "description" value, if not. See Dynamic Content
	Position            int64         // no 	no 	Ordering of the field relative to other fields
	Active              bool          // no 	no 	If true, this field is available for use
	System              bool          // yes 	no 	If true, only active and position values of this field can be changed
	RegexpForValidation string        // no 	no 	Regular expression field only. The validation pattern for a field value to be deemed valid.
	CreatedAt           Date          // yes 	no 	The time the ticket field was created
	UpdatedAt           Date          // yes 	no 	The time of the last update of the ticket field
	Tag                 string        // no 	no 	Optional for custom field of type "checkbox"; not presented otherwise.
	CustomFieldOptions  []FieldOption // no 	yes 	Required and presented for a custom field of type "dropdown"
}

func (ufs *UserFields) List() (userFields []UserField, err error) {
	path := fmt.Sprintf("%s%s", ufs.path, "/user_fields.json")
	responseBody, err := ufs.client.Get(path, nil)
	if err != nil {
		return userFields, err
	}
	// fmt.Printf("%s\n", string(responseBody))
	userFieldsWrapper := struct {
		UserFields []UserField `json:"user_Fields"`
	}{}
	if err := json.Unmarshal(responseBody, &userFieldsWrapper); err != nil {
		return userFields, err
	}
	for _, userField := range userFieldsWrapper.UserFields {
		userField.Class = ufs
		userFields = append(userFields, userField)
	}
	return userFields, err
}
