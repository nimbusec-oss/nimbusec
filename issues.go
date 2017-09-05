package nimbusec

import (
	"context"
	"net/http"
)

type IssueService service

func (srv *IssueService) List(ctx context.Context) ([]Issue, error) {
	issues := []Issue{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/issues", nil, &issues)
	return issues, err
}

func (srv *IssueService) ListByDomain(ctx context.Context, id DomainID) ([]Issue, error) {
	issues := []Issue{}
	err := srv.client.Do(ctx, http.MethodGet, string(id)+"/issues", nil, &issues)
	return issues, err
}

func (srv *IssueService) Get(ctx context.Context, id IssueID) (Issue, error) {
	issue := Issue{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &issue)
	return issue, err
}

func (srv *IssueService) Update(ctx context.Context, id IssueID, update IssueUpdate) (Issue, error) {
	issue := Issue{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &issue)
	return issue, err
}
