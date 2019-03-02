package nimbusec

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestIssueFilterToURL(t *testing.T) {
	expected := "/v3/issues?category=foo&event=bar&externalId=DESC1&limit=1&sort=firstSeen&status=pending"

	v, err := query.Values(&IssueFilter{
		Category: "foo",
		Event:    "bar",
		Limit:    1,
		// Severity: 2,
		Sort:       "firstSeen",
		Status:     IssueStatusPending,
		ExternalID: "DESC1",
	})
	if err != nil {
		t.Error(err)
	}
	u := url.URL{
		Path:     "/v3/issues",
		RawQuery: v.Encode(),
	}
	result := u.String()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
