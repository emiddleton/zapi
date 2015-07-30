package zapi

import (
	"time"
)

type User struct {
	Id                   int64                  // 	yes no 	Automatically assigned when the user is created
	Url                  string                 //	yes no 	The user's API url
	Name                 string                 //	no 	yes The user's name
	ExternalId           string                 // 	no 	no 	A unique id you can specify for the user
	Alias                string                 // 	no 	no 	An alias displayed to end users
	CreatedAt            time.Time              // 	yes no 	The time the user was created
	UpdatedAt            time.Time              // 	yes no 	The time the user was last updated
	Active               bool                   // 	yes no 	false if the user has been deleted
	Verified             bool                   // 	no 	no 	If the user's identity has been verified or not
	Shared               bool                   // 	yes no 	If the user is shared from a different Zendesk. Ticket sharing accounts only
	SharedAgent          bool                   // 	yes no 	If the user is a shared agent from a different Zendesk. Ticket sharing accounts only
	Locale               string                 // 	yes no 	The user's locale
	LocaleId             int64                  // 	no 	no 	The user's language identifier
	TimeZone             string                 // 	no 	no 	The user's time zone. See Time Zone below
	LastLoginAt          time.Time              // 	yes no 	The last time the user signed in to Zendesk
	TwoFactorAuthEnabled bool                   // 	yes no 	If two factor authentication is enabled.
	Email                string                 // 	no 	no 	The user's primary email address
	Phone                string                 // 	no 	no 	The user's primary phone number
	Signature            string                 // 	no 	no 	The user's signature. Only agents and admins can have signatures
	Details              string                 // 	no 	no 	Any details you want to store about the user, such as an address
	Notes                string                 // 	no 	no 	Any notes you want to store about the user
	OrganizationId       int64                  // 	no 	no 	The id of the organization the user is associated with
	Role                 string                 // 	no 	no 	The user's role. Possible values are "end-user", "agent", or "admin"
	CustomRoleId         int64                  // 	no 	no 	A custom role if the user is an agent on the Enterprise plan
	Moderator            bool                   // 	no 	no 	Designates whether the user has forum moderation capabilities
	TicketRestriction    string                 // 	no 	no 	Specifies which tickets the user has access to. Possible values are: "organization", "groups", "assigned", "requested", null
	OnlyPrivateComments  bool                   // 	no 	no 	true if the user can only create private comments
	Tags                 []string               // 	no 	no 	The user's tags. Only present if your account has user tagging enabled
	Suspended            bool                   // 	no 	no 	If the agent is suspended. Tickets from suspended users are also suspended, and these users cannot sign in to the end user portal
	RestrictedAgent      bool                   // 	no 	no 	If the agent has any restrictions; false for admins and unrestricted agents, true for other agents
	Photo                Attachment             // 	no 	no 	The user's profile picture represented as an Attachment object
	UserFields           map[string]interface{} // 	no 	no 	Custom fields for the user
}

type Users interface {
}

func (us *Users) List() ([]User, error) {
	return []User{}, nil
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
	return user, nil
}

func (us *Users) CreateWithIdentities(user User, identities []Identities) (User, error) {
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
