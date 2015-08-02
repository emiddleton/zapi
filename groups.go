package zapi

import ()

type Groups struct {
	path   string
	client *Client
}

type Group struct {
	Id        int64  // yes  no 	Automatically assigned when creating groups
	Url       string // yes  no 	The API url of this group
	Name      string // no 	yes 	The name of the group
	Deleted   bool   // yes  no 	Deleted groups get marked as such
	CreatedAt Date   // yes  no 	The time the group was created
	UpdatedAt Date   // yes  no 	The time of the last update of the group
}
