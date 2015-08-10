package zapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTicketFields(t *testing.T) {
	client, err := NewClientFromFile("./auth.json")
	if err != nil {
		t.Fatal(err)
	}

	tfs := NewTicketFields(&client)

	ticketFields, err := tfs.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, ticketField := range ticketFields {
		ticketFieldJson, err := json.MarshalIndent(&ticketField, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%s\n", ticketFieldJson)
	}

}
