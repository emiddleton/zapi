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

func TestUsers(t *testing.T) {

	data, err := ioutil.ReadFile("./auth.json")
	if err != nil {
		t.Fatal(err)
	}

	c := Client{httpClient: &http.Client{}}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}

	us := Users{
		client: &c,
	}

	users, err := us.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		userJson, err := json.MarshalIndent(&user, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%s\n", userJson)
	}

	guid, err := uuid.NewV4()
	guidStr := guid.String()
	user, err := us.Create(
		User{
			Name:       randomdata.FullName(randomdata.RandomGender),
			Email:      randomdata.Email(),
			ExternalId: &guidStr,
			Verified:   true,
		},
	)
	fmt.Printf("%#v\n", user)
}
