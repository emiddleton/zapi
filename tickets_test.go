package zapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTickets(t *testing.T) {

	data, err := ioutil.ReadFile("./auth.json")
	if err != nil {
		t.Fatal(err)
	}

	c := Client{httpClient: &http.Client{}}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}

	ts := Tickets{
		client: &c,
	}

	tickets, err := ts.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, ticket := range tickets {
		ticketJson, err := json.MarshalIndent(&ticket, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%s\n", ticketJson)
	}
}
