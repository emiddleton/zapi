package zapi

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/nu7hatch/gouuid"
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

	tickets, err := ts.List(nil)
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

	guid, err := uuid.NewV4()
	guidStr := guid.String()
	ticket, err := ts.Create(
		Ticket{
			Subject: "Hell is freezing over I need a lighter",
			Comment: &Comment{
				Body: "Where can I find a lighter on ClassDo",
			},
			Requester: &User{
				Name:  randomdata.FullName(randomdata.RandomGender),
				Email: randomdata.Email(),
			},
			ExternalId: &guidStr,
		},
	)
	fmt.Printf("%#v\n", ticket)

}
