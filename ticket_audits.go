package zapi

import (
	"time"
)

type TicketAudit struct {
	Id        int64                    // Automatically assigned when creating audits
	TicketId  int64                    // The ID of the associated ticket
	Metadata  hash                     // Metadata for the audit, custom and system data
	Via       ViaObject                // This object explains how this audit was created
	CreatedAt time.Time                // The time the audit was created
	AuthorId  int64                    // The user who created the audit
	events    []map[string]interface{} // An array of the events that happened in this audit. See Audit Events
}

type CreateEvent struct {
	Id        int64    // Automatically assigned when the event is created
	Type      string   // Has the value Create
	FieldName string   // The name of the field that was set
	Value     []string // The value of the field that was set
}

type ChangeEvent struct {
	Id            integer  // Automatically assigned when the event is created
	Type          string   // Has the value Change
	FieldName     string   // The name of the field that was changed
	Value         []string // The value of the field that was changed
	PreviousValue []string // The previous value of the field that was changed
}

type CommentEvent struct {
	Id            int64        // Automatically assigned when the event is created
	Type          string       // Has the value VoiceComment
	Data          string       // A hash of properties about the call
	Public        bool         // If true, the ticket requester can see this comment
	FormattedFrom string       // A formatted version of the phone number which dialed the call
	FormattedTo   string       // A formatted version of the phone number which answered the call
	Body          string       // The actual comment made by the author
	HtmlBody      string       // The actual comment made by the author formatted to HTML
	Public        bool         // If this is a public comment or an internal agents only note
	Trusted       bool         // If this comment is trusted or marked as being potentially fraudulent
	AuthorId      int64        // The id of the author of this comment
	Attachments   []Attachment // The attachments on this comment as Attachment objects
}

type CommentPrivacyChangeEvent struct {
	Id        int64  // Automatically assigned when the event is created
	Type      string // Has the value CommentPrivacyChange
	CommentId int64  // The id if the comment that changed privacy
	Public    bool   // Tells if the comment was made public or private
}

type NotificationEvent struct {
	Id         int64     // Automatically assigned when the event is created
	Type       string    // Has the value Notification
	Subject    string    // The subject of the message sent to the recipients
	Body       string    // The message sent to the recipients
	Recipients []string  //	A array of simple objects with the ids and names of the recipients of this notification
	Via        ViaObject // The business rule that created the notification
}

type CCEvent struct {
	Id         int64     // Automatically assigned when the event is created
	Type       string    // Has the value Cc
	Recipients []string  // A array of simple objects with the ids and names of the recipients of this notification
	Via        ViaObject // A reference to the business rule that created this notification
}

type SatisfactionRatingEvent struct {
	Id         int64  // Automatically assigned when creating events
	Type       string // Has the value SatisfactionRating
	Score      string // The rating state "offered", "unoffered", "good", "bad"
	AssigneeId int64  // Who the ticket was assigned to upon rating time
	Body       string // The users comment posted during rating
}

type TicketSharingEvent struct {
	Id          int64  // Automatically assigned when creating events
	Type        string // Has the value TicketSharingEvent
	AgreementId int64  // ID of the sharing agreement
	Action      string // Either shared or unshared
}

type ErrorEvent struct {
	Id      integer // Automatically assigned when the event is creating
	Type    string  // Has the value Error
	Message string  // The error message
}

type TweetEvent struct {
	Id            int64    // Automatically assigned when the event is created
	Type          string   // Has the value Tweet
	DirectMessage bool     //	Whether this tweet was a direct message
	Body          string   // The body of the tweet
	Recipients    []string // The recipients of this tweet
}

type FacebookEvent struct {
	Id            int64             // Automatically assigned when the event is created
	Type          string            // Has the value FacebookEvent
	Page          map[string]string // The name and graph id of the Facebook Page associated with the event
	Communication int64             // The Zendesk id of the associated communication (wall post or message)
	TicketVia     string            // "post" or "message" depending on the association with a Wall post or a private message
	Body          string            // The value of the message posted to Facebook
}

type FacebookCommentEvent struct {
	Id            int64                  // Automatically assigned when the event is created
	Type          string                 // has the value FacebookComment
	Data          map[string]interface{} // Properties of the Facebook comment
	Body          string                 // The actual comment made by the author
	HtmlBody      string                 // The actual comment made by the author formatted as HTML
	Public        bool                   // If this is a public comment or an internal-agents-only note
	Trusted       bool                   // If this comment is trusted or marked as being potentially fraudulent
	AuthorId      int64                  // The id of the author of this comment
	GraphObjectId string                 // The graph object id of the associated Facebook Wall post or message
}

type ExternalEvent struct {
	Id       int64  // Automatically assigned when the event is created
	Type     string // Has the value External
	Resource string // External target id
	Body     string // Trigger message for this target event
}

type LogMelnTranscriptEvent struct {
	Id   int64  // Automatically assigned when creating events
	Type string // Has the value LogMeInTranscript
	Body string // An audit of the transcript
}

type PushEvent struct {
	Id             int64  // Automatically assigned when the event is created
	Type           string // Has the value Push
	Value          string // Data being pushed out of our system
	ValueReference string // A reference to the destination of the data
}

type ViaObject struct {
	Channel string                 // 	This tells you how the ticket or event was created
	Source  map[string]interface{} // 	For some channels a source object gives more information about how or why the ticket or event was created
}
