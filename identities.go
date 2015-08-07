package zapi

import (
	"encoding/json"
	"fmt"
)

type Identities struct {
	path   string
	client *Client
}

func NewIdentities(path string, client *Client) Identities {
	return Identities{
		path:   path,
		client: client,
	}
}

type Identity struct {
	Class              *Identities `json:"-"`
	Id                 int64       `json:"id,omitempty"`                  // yes 	no 	Automatically assigned upon creation
	Url                string      `json:"url,omitempty"`                 // yes 	no 	The API url of this identity
	UserId             int64       `json:"user_id"`                       // yes 	yes 	The id of the user
	Type               string      `json:"type"`                          // yes 	yes 	One of "email", "twitter", "facebook", "google", or "phone_number"
	Value              string      `json:"value"`                         // yes 	yes 	The identifier for this identity, e.g. an email address
	Verified           bool        `json:"verified"`                      // no 	no 	Is true if the identity has gone through verification
	Primary            bool        `json:"primary"`                       // no 	no 	Is true if the primary identity of the user
	CreatedAt          Date        `json:"created_at,omitempty"`          // yes 	no 	The time the identity got created
	UpdatedAt          Date        `json:"updated_at,omitempty"`          // yes 	no 	The time the identity got updated
	UndeliverableCount int64       `json:"undeliverable_count,omitempty"` // yes 	no 	The number of times a non-delivery response was received at that address (max. 50)
}

func (is *Identities) List() (identities []Identity, err error) {
	path := fmt.Sprintf("%s%s", is.path, "/identities.json")
	responseBody, err := is.client.Get(path, nil)
	if err != nil {
		return []Identity{}, err
	}

	// fmt.Printf("%s\n", string(responseBody))
	identitiesPager := struct {
		Identities []Identity `json:"identities"`
	}{}

	if err := json.Unmarshal(responseBody, &identitiesPager); err != nil {
		return identities, err
	}

	for _, identity := range identitiesPager.Identities {
		identity.Class = is
		identities = append(identities, identity)
	}

	return identities, nil
}

func (i *Identity) MakePrimary() error {
	path := fmt.Sprintf("%s/identities/%d/make_primary.json", i.Class.path, i.Id)
	//responseBody
	_, err := i.Class.client.Put(path, nil, []byte("{}"))
	if err != nil {
		return err
	}
	// fmt.Printf("%s\n", string(responseBody))
	return nil
}
