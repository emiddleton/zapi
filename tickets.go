package zapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Tickets struct {
	path   string
	client *Client
}

func NewTickets(path string, client *Client) Tickets {
	return Tickets{
		path:   path,
		client: client,
	}
}

type Ticket struct {
	Class               *Tickets                 `json:"-"`                               // reference to parent class
	Comments            *Comments                `json:"-"`                               // comments
	Id                  int64                    `json:"id,omitempty"`                    // yes  no  Automatically assigned when creating tickets
	Url                 string                   `json:"url,omitempty"`                   // yes  no  The API url of this ticket
	ExternalId          *string                  `json:"external_id,omitempty"`           //  no  no  An id you can use to link Zendesk tickets to local records
	Type                string                   `json:"type"`                            //  no  no  The type of this ticket, i.e. "problem", "incident", "question" or "task"
	Subject             string                   `json:"subject"`                         //  no  no  The value of the subject field for this ticket
	RawSubject          string                   `json:"raw_subject"`                     //  no  no  The dynamic content placeholder, if present, or the "subject" value, if not. See Dynamic Content
	Description         string                   `json:"description"`                     // yes  no  The first comment on the ticket
	Priority            *string                  `json:"priority"`                        //  no  no  Priority, defines the urgency with which the ticket should be addressed: "urgent", "high", "normal", "low"
	Status              string                   `json:"status"`                          //  no  no  The state of the ticket, "new", "open", "pending", "hold", "solved", "closed"
	Recipient           *string                  `json:"recipients,omitempty"`            //  no  no  The original recipient e-mail address of the ticket
	RequesterId         *int64                   `json:"requester_id"`                    //  no yes  The user who requested this ticket
	Requester           *User                    `json:"requester,omitempty"`             //  no yes  The literal requesting user
	SubmitterId         int64                    `json:"submitter_id,omitempty"`          //  no  no  The user who submitted the ticket; The submitter always becomes the author of the first comment on the ticket.
	AssigneeId          int64                    `json:"assignee_id,omitempty"`           //  no  no  What agent is currently assigned to the ticket
	OrganizationId      *int64                   `json:"organization_id,omitempty"`       // yes  no  The organization of the requester
	GroupId             int64                    `json:"group_id,omitempty"`              //  no  no  The group this ticket is assigned to
	CollaboratorIds     []int64                  `json:"collaborator_ids,omitempty"`      //  no  no  Who are currently CC'ed on the ticket
	ForumTopicId        *int64                   `json:"forum_topic_id,omitempty"`        //  no  no  The topic this ticket originated from, if any
	ProblemId           *int64                   `json:"problem_id,omitempty"`            //  no  no  The problem this incident is linked to, if any
	HasIncidents        bool                     `json:"has_incidents"`                   // yes  no  Is true of this ticket has been marked as a problem, false otherwise
	DueAt               *Date                    `json:"due_at,omitempty"`                //  no  no  If this is a ticket of type "task" it has a due date. Due date format uses ISO 8601 format.
	Tags                []string                 `json:"tags,omitempty"`                  //  no  no  The array of tags applied to this ticket
	Via                 *ViaObject               `json:"via,omitempty"`                   // yes  no  This object explains how the ticket was created
	CustomFields        []map[string]interface{} `json:"custom_fields"`                   //  no  no  The custom fields of the ticket
	SatisfactionRating  map[string]interface{}   `json:"satisfaction_rating,omitempty"`   // yes  no  The satisfaction rating of the ticket, if it exists, or the state of satisfaction, 'offered' or 'unoffered'
	SharingAgreementIds []string                 `json:"sharing_agreement_ids,omitempty"` // yes  no  The ids of the sharing agreements used for this ticket
	FollowupIds         *[]int64                 `json:"followup_ids,omitempty"`          // yes  no  The ids of the followups created from this ticket - only applicable for closed tickets
	TicketFormId        *int64                   `json:"ticket_form_id,omitempty"`        //  no  no  The id of the ticket form to render for this ticket - only applicable for enterprise accounts
	BrandId             int64                    `json:"brand_id,omitempty"`              //  no  no  The id of the brand this ticket is associated with - only applicable for enterprise accounts
	CreatedAt           *Date                    `json:"created_at,omitempty"`            // yes  no  When this record was created
	UpdatedAt           *Date                    `json:"updated_at,omitempty"`            // yes  no  When this record last got updated
}

type Twicket struct {
}

type ProblemService struct {
}

type ChannelTwitter struct {
}

func (ts *Tickets) List(filters Filters) (tickets []Ticket, err error) {
	path := fmt.Sprintf("%s%s", ts.path, "/tickets.json")
	responseBody, err := ts.client.Get(path, filters.toParams(&url.Values{}))
	if err != nil {
		return []Ticket{}, err
	}
	fmt.Printf("%s\n", string(responseBody))
	ticketsPager := struct {
		Count        int64    `json:"count"`
		NextPage     *int64   `json:"next_page"`
		PreviousPage *int64   `json:"previous_page"`
		Tickets      []Ticket `json:"tickets"`
	}{}
	if err := json.Unmarshal(responseBody, &ticketsPager); err != nil {
		return tickets, err
	}
	return ticketsPager.Tickets, err
}

func (ts *Tickets) Get(id int64) (ticket Ticket, err error) {
	return ticket, err
}

func (ts *Tickets) ShowMany(ids []int64) (tickets []Ticket, err error) {
	return tickets, err
}

func (ts *Tickets) Create(ticket Ticket) (Ticket, error) {

	path := fmt.Sprintf("%s%s", ts.path, "/tickets.json")
	type ticketWrap struct {
		WrappedTicket Ticket `json:"ticket"`
	}
	reqBody, err := json.MarshalIndent(&ticketWrap{ticket}, "", "    ")
	if err != nil {
		return Ticket{}, err
	}
	fmt.Printf("request ->\n%s\n", string(reqBody))

	responseBody, err := ts.client.Post(path, nil, reqBody)
	if err != nil {
		return Ticket{}, err
	}

	fmt.Printf("response ->\n%s\n", string(responseBody))
	ticketWrapper := ticketWrap{ticket}
	if err := json.Unmarshal(responseBody, &ticketWrapper); err != nil {
		return ticketWrapper.WrappedTicket, err
	}

	return ticket, nil
}

func (ts *Tickets) CreateMany(tickets []Ticket) ([]Ticket, error) {
	return tickets, nil
}

func (ts *Tickets) Update(ticket Ticket) (Ticket, error) {
	return ticket, nil
}

func (ts *Tickets) UpdateMany(tickets []Ticket) ([]Ticket, error) {
	return tickets, nil
}

func (ts *Tickets) MarkAsSpam(ticketId int64) error {
	return nil
}

func (ts *Tickets) MarkManyAsSpam(ticketIds []int64) error {
	return nil
}

func (ts *Tickets) Merge(ticketId int64) (Ticket, error) {
	return Ticket{}, nil
}

func (ts *Tickets) Related() (Ticket, error) {
	return Ticket{}, nil
}

func Delete() (Ticket, error) {
	return Ticket{}, nil
}

func (ts *Tickets) DeleteMany(ids []int64) error {
	return nil
}

func (t *Ticket) Collaborators(ticket Ticket) ([]User, error) {
	return []User{{}}, nil
}

func (t *Ticket) Incidents(ticket Ticket) ([]Ticket, error) {
	return []Ticket{{}}, nil
}

func (ts *Tickets) Problems() ([]Ticket, error) {
	return []Ticket{{}}, nil
}

func (ts *ProblemService) Autocomplete(text string) ([]Ticket, error) {
	return []Ticket{{}}, nil
}

/*
  NOTE: user sibmitter is set to user making api call
*/
func CreateTicketFromTweet(twitterStatusMessageId, monitorTwitterHandleId int64) (Twicket, error) {
	return Twicket{}, nil
}

func (ct *Twicket) Statuses(commentIds []int64) (statuses []string, err error) {
	return statuses, err
}
