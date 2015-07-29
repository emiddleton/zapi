package zapi

import (
	"time"
)

func Import(ticket Ticket) (Ticket, error) {
	return ticket, nil
}

func ImportMany(tickets []Ticket) ([]Ticket, error) {
	return tickets, nil
}
