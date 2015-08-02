package zapi

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	path   string
	client *Client
}

type User struct {
	Id                   int64                  `json:"id,omitempty"`            // 	yes no 	Automatically assigned when the user is created
	Url                  string                 `json:"url,omitempty"`           //	  yes no 	The user's API url
	Name                 string                 `json:"name"`                    //	  no 	yes The user's name
	ExternalId           *string                `json:"external_id"`             // 	no 	no 	A unique id you can specify for the user
	Alias                *string                `json:"alias"`                   // 	no 	no 	An alias displayed to end users
	CreatedAt            *Date                  `json:"created_at,omitempty"`    // 	yes no 	The time the user was created
	UpdatedAt            *Date                  `json:"updated_at,omitempty"`    // 	yes no 	The time the user was last updated
	Active               bool                   `json:"active"`                  // 	yes no 	false if the user has been deleted
	Verified             bool                   `json:"verfied"`                 // 	no 	no 	If the user's identity has been verified or not
	Shared               bool                   `json:"shared"`                  // 	yes no 	If the user is shared from a different Zendesk. Ticket sharing accounts only
	SharedAgent          bool                   `json:"shared_agent"`            // 	yes no 	If the user is a shared agent from a different Zendesk. Ticket sharing accounts only
	Locale               string                 `json:"locale"`                  // 	yes no 	The user's locale
	LocaleId             int64                  `json:"locale_id"`               // 	no 	no 	The user's language identifier
	TimeZone             string                 `json:"time_zone"`               // 	no 	no 	The user's time zone. See Time Zone below
	LastLoginAt          *Date                  `json:"last_login_at"`           // 	yes no 	The last time the user signed in to Zendesk
	TwoFactorAuthEnabled *bool                  `json:"two_factor_auth_enabled"` // 	yes no 	If two factor authentication is enabled.
	Email                string                 `json:"email"`                   // 	no 	no 	The user's primary email address
	Phone                *string                `json:"phone"`                   // 	no 	no 	The user's primary phone number
	Signature            *string                `json:"signature"`               // 	no 	no 	The user's signature. Only agents and admins can have signatures
	Details              string                 `json:"details"`                 // 	no 	no 	Any details you want to store about the user, such as an address
	Notes                string                 `json:"notes"`                   // 	no 	no 	Any notes you want to store about the user
	OrganizationId       int64                  `json:"organization_id"`         // 	no 	no 	The id of the organization the user is associated with
	Role                 string                 `json:"role,omitempty"`          // 	no 	no 	The user's role. Possible values are "end-user", "agent", or "admin"
	CustomRoleId         *int64                 `json:"custom_role_id"`          // 	no 	no 	A custom role if the user is an agent on the Enterprise plan
	Moderator            bool                   `json:"moderator"`               // 	no 	no 	Designates whether the user has forum moderation capabilities
	TicketRestriction    *string                `json:"ticket_restriction"`      // 	no 	no 	Specifies which tickets the user has access to. Possible values are: "organization", "groups", "assigned", "requested", null
	OnlyPrivateComments  bool                   `json:"only_private_comments"`   // 	no 	no 	true if the user can only create private comments
	Tags                 []string               `json:"tags"`                    // 	no 	no 	The user's tags. Only present if your account has user tagging enabled
	Suspended            bool                   `json:"suspended"`               // 	no 	no 	If the agent is suspended. Tickets from suspended users are also suspended, and these users cannot sign in to the end user portal
	RestrictedAgent      bool                   `json:"restricted_agent"`        // 	no 	no 	If the agent has any restrictions; false for admins and unrestricted agents, true for other agents
	Photo                *Attachment            `json:"photo"`                   // 	no 	no 	The user's profile picture represented as an Attachment object
	UserFields           map[string]interface{} `json:"user_fields"`             // 	no 	no 	Custom fields for the user
}

func (us *Users) List() (users []User, err error) {
	path := fmt.Sprintf("%s%s", us.path, "/users.json")
	responseBody, err := us.client.Get(path, nil)
	if err != nil {
		return []User{}, err
	}
	fmt.Printf("%s\n", string(responseBody))
	usersPager := struct {
		Count        int64  `json:"count"`
		NextPage     *int64 `json:"next_page"`
		PreviousPage *int64 `json:"previous_page"`
		Users        []User `json:"users"`
	}{}
	if err := json.Unmarshal(responseBody, &usersPager); err != nil {
		return users, err
	}
	return usersPager.Users, nil
}

func (us *Users) Get(id int64) (User, error) {
	return User{}, nil
}

func (us *Users) GetManyByIds(ids []int64) ([]User, error) {
	return []User{}, nil
}

func (us *Users) GetManyByExternalIds(externalIds []string) ([]User, error) {
	return []User{}, nil
}

func (us *Users) Create(user User) (User, error) {
	path := fmt.Sprintf("%s%s", us.path, "/users.json")
	type userWrap struct {
		WrappedUser User `json:"user"`
	}
	reqBody, err := json.MarshalIndent(&userWrap{user}, "", "    ")
	if err != nil {
		return User{}, err
	}
	fmt.Printf("request ->\n%s\n", string(reqBody))

	responseBody, err := us.client.Post(path, nil, reqBody)
	if err != nil {
		return User{}, err
	}

	fmt.Printf("response ->\n%s\n", string(responseBody))
	userWrapper := userWrap{user}
	if err := json.Unmarshal(responseBody, &userWrapper); err != nil {
		return userWrapper.WrappedUser, err
	}

	return user, nil
}

func (us *Users) CreateWithIdentities(user User, identities []Identity) (User, error) {
	return user, nil
}

func (us *Users) CreateMany(users []User) ([]User, error) {
	return users, nil
}

func (us *Users) UpdateMany(users []User) ([]User, error) {
	return users, nil
}

func (us *Users) UpdateManyByIds(ids int64, values map[string]interface{}) ([]User, error) {
	return []User{}, nil
}

func (us *Users) UpdateManyByExternalIds(externalIds int64, values map[string]interface{}) ([]User, error) {
	return []User{}, nil
}

func (us *Users) DeleteManyByIds(ids int64) error {
	return nil
}

func (us *Users) DeleteManyByExternalIds(externalIds int64) error {
	return nil
}

func (us *Users) Search(query string) ([]User, error) {
	return []User{}, nil
}

func (us *Users) SearchByExternalIds(externalIds []string) ([]User, error) {
	return []User{}, nil
}

func (us *Users) Autocomplete(name string) ([]User, error) {
	return []User{}, nil
}

type UserRelated struct {
	RequestedTickets int64
	CcdTickets       int64
	Topics           int64
	TopicComments    int64
	Votes            int64
	Subscriptions    int64
}

func (u *User) Related() (UserRelated, error) {
	return UserRelated{}, nil
}

func (u *User) Update() (User, error) {
	return *u, nil
}

func (u *User) Delete() (User, error) {
	return *u, nil
}
