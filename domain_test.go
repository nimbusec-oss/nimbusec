package nimbusec

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestDomainFilterToURL(t *testing.T) {
	expected := "/v3/domains?externalId=IMPORTANTE&name=nimbusec.com"

	v, err := query.Values(&DomainFilter{
		Name:       "nimbusec.com",
		ExternalID: "IMPORTANTE",
	})
	if err != nil {
		t.Error(err)
	}
	u := url.URL{
		Path:     "/v3/domains",
		RawQuery: v.Encode(),
	}
	result := u.String()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
