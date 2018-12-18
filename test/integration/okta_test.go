package integration

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/chanzuckerberg/go-okta/okta"
)

var (
	client *okta.Client
	err    error
)

func init() {
	apiKey := os.Getenv("OKTA_API_KEY")
	baseURL := "https://czi.okta.com/api/v1/"
	client, err = okta.NewClient(apiKey, baseURL, nil)
	if err != nil {
		log.Fatalf("Couldn't create an Okta Client: %v", err)
	}

}

func TestGroups(t *testing.T) {

	groupList, _, err := client.Groups.List(context.Background())
	if err != nil {
		t.Fatalf("Groups.List returned error: %v", err)
	}
	if len(groupList) == 0 {
		t.Fatalf("Groups.List returned no groups. There should be at least one group, Everyone, in all Okta accounts.")
	}

	filter := "type eq \"BUILT_IN\""
	groupList, _, err = client.Groups.ListFilter(context.Background(), filter)
	if err != nil {
		t.Fatalf("Groups.ListFilter(%q) returned error: %v", err, filter)
	}
	if len(groupList) == 0 {
		t.Fatalf("Groups.ListFilter(%q) returned no groups. There should be at least one group, Everyone, in all Okta accounts.", filter)
	}

	groupList, _, err = client.Groups.ListSearchByName(context.Background(), "Everyone")
	if err != nil {
		t.Fatalf("Groups.ListSearchByName(%q) returned error: %v", "Everyone", err)
	}

	if len(groupList) == 0 {
		t.Fatalf("Groups.ListSearchByName returned no groups when searching for Everyone. Everyone is a built in group in Okta and should exist.")
	}

	if groupList[0].Profile.Name != "Everyone" {
		t.Fatalf("Groups.List didn't return Everyone as it's first result. Everyone is a built in group in Okta and should be the first result for this query. It returned: %s", groupList[0].Profile.Name)
	}

	// groupList[0] should be the Everyone group
	groupMemberList, _, err := client.Groups.ListMembers(context.Background(), groupList[0].ID)
	if len(groupMemberList) == 0 {
		t.Fatalf("Groups.ListMembers (%q) returned no results when looking at the Everyone group. Everyone is a built in group in Okta and should exist and have members.", groupList[0].ID)
	}
}
