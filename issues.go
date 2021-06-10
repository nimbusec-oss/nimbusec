package nimbusec

import (
	"context"
	"fmt"
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

	issues := []Issue{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (srv *IssueService) ListByDomain(ctx context.Context, id DomainID, filter *IssueFilter) ([]Issue, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     fmt.Sprintf("/v3/domains/%d/issues", id),
		RawQuery: v.Encode(),
	}

	issues := []Issue{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (srv *IssueService) Get(ctx context.Context, id IssueID) (Issue, error) {
	issue := Issue{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/issues/%d", id), nil, &issue)
	if err != nil {
		return Issue{}, err
	}

	return issue, nil
}

func (srv *IssueService) Update(ctx context.Context, id IssueID, update IssueUpdate) (Issue, error) {
	issue := Issue{}
	err := srv.client.Do(ctx, http.MethodPut, fmt.Sprintf("/v3/issues/%d", id), update, &issue)
	if err != nil {
		return Issue{}, err
	}

	return issue, nil
}
