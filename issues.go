package nimbusec

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type IssueService service

func (srv *IssueService) List(ctx context.Context, filter *IssueFilter) ([]Issue, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/v3/issues",
		RawQuery: v.Encode(),
	}

	issues := []issueDTO{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &issues)
	if err != nil {
		return nil, err
	}

	converted := make([]Issue, len(issues))
	for i, issue := range issues {
		converted[i], err = translateIssue(issue)
		if err != nil {
			return nil, err
		}
	}

	return converted, nil
}

func (srv *IssueService) ListByDomain(ctx context.Context, id DomainID, filter *IssueFilter) ([]Issue, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     string(id) + "/issues",
		RawQuery: v.Encode(),
	}

	issues := []issueDTO{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &issues)
	if err != nil {
		return nil, err
	}

	converted := make([]Issue, len(issues))
	for i, issue := range issues {
		converted[i], err = translateIssue(issue)
		if err != nil {
			return nil, err
		}
	}

	return converted, nil
}

func (srv *IssueService) Get(ctx context.Context, id IssueID) (Issue, error) {
	issue := issueDTO{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &issue)
	if err != nil {
		return Issue{}, err
	}

	return translateIssue(issue)
}

func (srv *IssueService) Update(ctx context.Context, id IssueID, update IssueUpdate) (Issue, error) {
	issue := issueDTO{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &issue)
	if err != nil {
		return Issue{}, err
	}

	return translateIssue(issue)
}

// issueDTO is a temporay struct to read json issues until they are fully parsed
// and translated into the exported Issue type.
type issueDTO struct {
	Issue
	Details json.RawMessage `json:"details,omitempty"`
}

func translateIssue(in issueDTO) (Issue, error) {
	var err error
	var out = in.Issue
	var details interface{}

	switch in.Event {
	case IssueEventBlacklist:
		details = BlacklistDetails{}

	case IssueEventChangedFile:
		details = ChangedFileDetails{}

	case IssueEventCMSVersion:
		details = ApplicationOutdatedDetails{}

	case IssueEventCMSVulnerable:
		details = ApplicationVulnerableDetails{}

	case IssueEventDefacement:
		details = DefacementDetails{}

	case IssueEventMalware:
		details = MalwareDetails{}

	case IssueEventSuspiciousLink:
		details = SuspiciousLinkDetails{}

	case IssueEventSuspciousRequest:
		details = SuspiciousRequestDetails{}

	case IssueEventTLSCipherSuite, IssueEventTLSProtocol:
		details = TLSConfigurationDetails{}

	case IssueEventTLSExpires, IssueEventTLSHostname, IssueEventTLSNoTrust, IssueEventTLSSigAlg:
		details = TLSCertificateDetails{}

	case IssueEventWebshell:
		details = WebshellDetails{}
	default:
		return out, nil
	}
	err = json.Unmarshal(in.Details, &details)
	out.Details = details

	return out, err
}
