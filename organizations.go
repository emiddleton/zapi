package zapi

import ()

type Organizations struct {
	path   string
	client *Client
}

type Organization struct {
	Id                 int64                  // yes no   Automatically assigned when the organization is created
	Url                string                 // yes no   The API url of this organization
	ExternalId         string                 // no  no   A unique external id to associate organizations to an external record
	Name               string                 // no  yes  The name of the organization
	CreatedAt          Date                   // yes no   The time the organization was created
	UpdatedAt          Date                   // yes no   The time of the last update of the organization
	DomainNames        []string               // no  no   An array of domain names associated with this organization
	Details            string                 // no  no   Any details obout the organization, such as the address
	Notes              string                 // no  no   Any notes you have about the organization
	GroupId            int64                  // no  no   New tickets from users in this organization are automatically put in this group
	SharedTickets      bool                   // no  no   End users in this organization are able to see each other's tickets
	SharedComments     bool                   // no  no   End users in this organization are able to see each other's comments on tickets
	Tags               []string               // no  no   The tags of the organization
	OrganizationFields map[string]interface{} // no  no   Custom fields for this organization
}
