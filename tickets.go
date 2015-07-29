package zapi

import (
	"time"
)

type Ticket struct {
	Id                    int64                  // yes 	no 	Automatically assigned when creating tickets
	Url                   string                 // yes 	no 	The API url of this ticket
	ExternalId            string                 // no 	no 	An id you can use to link Zendesk tickets to local records
	Type                  string                 //	no 	no 	The type of this ticket, i.e. "problem", "incident", "question" or "task"
	Subject               string                 //	no 	no 	The value of the subject field for this ticket
	RawSubject            string                 // no 	no 	The dynamic content placeholder, if present, or the "subject" value, if not. See Dynamic Content
	Description           string                 //	yes 	no 	The first comment on the ticket
	priority              string                 //	no 	no 	Priority, defines the urgency with which the ticket should be addressed: "urgent", "high", "normal", "low"
	status                string                 //	no 	no 	The state of the ticket, "new", "open", "pending", "hold", "solved", "closed"
	recipient             string                 //	no 	no 	The original recipient e-mail address of the ticket
	requester_id          int64                  // no 	yes 	The user who requested this ticket
	submitter_id          int64                  // no 	no 	The user who submitted the ticket; The submitter always becomes the author of the first comment on the ticket.
	assignee_id           int64                  //	no 	no 	What agent is currently assigned to the ticket
	organization_id       int64                  // yes 	no 	The organization of the requester
	group_id              int64                  // no 	no 	The group this ticket is assigned to
	collaborator_ids      []int64                // no 	no 	Who are currently CC'ed on the ticket
	forum_topic_id        int64                  // no 	no 	The topic this ticket originated from, if any
	problem_id            int64                  // no 	no 	The problem this incident is linked to, if any
	has_incidents         bool                   // yes 	no 	Is true of this ticket has been marked as a problem, false otherwise
	due_at                time.Time              // no 	no 	If this is a ticket of type "task" it has a due date. Due date format uses ISO 8601 format.
	tags                  []string               // no 	no 	The array of tags applied to this ticket
	via                   Via                    // yes 	no 	This object explains how the ticket was created
	custom_fields         map[string]interface{} // no 	no 	The custom fields of the ticket
	satisfaction_rating   map[string]interface{} //	yes 	no 	The satisfaction rating of the ticket, if it exists, or the state of satisfaction, 'offered' or 'unoffered'
	sharing_agreement_ids []string               // yes 	no 	The ids of the sharing agreements used for this ticket
	followup_ids          []int64                // yes  	no 	The ids of the followups created from this ticket - only applicable for closed tickets
	ticket_form_id        int64                  // 	no 	no 	The id of the ticket form to render for this ticket - only applicable for enterprise accounts
	brand_id              int64                  // no 	no 	The id of the brand this ticket is associated with - only applicable for enterprise accounts
	created_at            time.Time              //	yes 	no 	When this record was created
	updated_at            time.Time              //	yes 	no 	When this record last got updated
}

type TicketService struct {
}

type ProblemService struct {
}

type ChannelTwitter struct {
}

func (ts *TicketService) List() (tickets []Ticket, page int64, offset int64, err error) {
	return tickets, page, offset, err
}

func (ts *TicketService) Get(id int64) (ticket Ticket, err error) {
	return ticket, err
}

func (ts *TicketService) ShowMany(ids []int64) (tickets []Ticket, err error) {
	return tickets, err
}

func (ts *TicketService) Create(ticket Ticket) (Ticket, error) {
	return ticket, nil
}

func (ts *TicketService) CreateMany(tickets []Ticket) ([]Ticket, error) {
	return tickets, nil
}

func (ts *TicketService) Update(ticket Ticket) (Ticket, error) {
	return ticket, nil
}

func (ts *TicketService) UpdateMany(tickets []Ticket) ([]Ticket, error) {
	return tickets, nil
}

func (ts *TicketService) MarkAsSpam(ticketId int64) error {
	return nil
}

func (ts *TicketService) MarkManyAsSpam(ticketIds []int64) error {
	return nil
}

func (ts *TicketService) Merge(ticketId int64) (Ticket, error) {
	return Ticket{}, nil
}

func (ts *TicketService) Related() (Ticket, error) {
	return Ticket{}, nil
}

func Delete() (Ticket, error) {
	return Ticket{}, nil
}

func (ts *TicketService) DeleteMany(ids []int64) error {
	return nil
}

func (t *Ticket) Collaborators(ticket Ticket) ([]User, error) {
	return []User{{}}, nil
}

func (t *Ticket) Incidents(ticket Ticket) ([]Ticket, error) {
	return []Ticket{{}}, nil
}

func (ts *TicketService) Problems() ([]Ticket, error) {
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
