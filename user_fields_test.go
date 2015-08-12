package zapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserFields(t *testing.T) {
	client, err := NewClientFromFile("./auth.json")
	if err != nil {
		t.Fatal(err)
	}

	ufs := NewUserFields(&client)

	userFields, err := ufs.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, userField := range userFields {
		userFieldJson, err := json.MarshalIndent(&userField, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%s\n", userFieldJson)
	}

}
