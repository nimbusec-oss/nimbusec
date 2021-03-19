package nimbusec

import (
	"context"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type DomainService service

func (srv *DomainService) List(ctx context.Context, filter *DomainFilter) ([]Domain, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/v3/domains",
		RawQuery: v.Encode(),
	}

	domains := []Domain{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &domains)
	return domains, err
}

func (srv *DomainService) Get(ctx context.Context, id DomainID) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &domain)
	return domain, err
}

func (srv *DomainService) Create(ctx context.Context, create Domain) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodPost, "/v3/domains", create, &domain)
	return domain, err
}

func (srv *DomainService) Update(ctx context.Context, id DomainID, update Domain) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &domain)
	return domain, err
}

func (srv *DomainService) Disable(ctx context.Context, id DomainID) error {
	err := srv.client.Do(ctx, http.MethodPatch, string(id)+"/disable", nil, nil)
	return err
}

func (srv *DomainService) Delete(ctx context.Context, id DomainID) error {
	err := srv.client.Do(ctx, http.MethodDelete, string(id), nil, nil)
	return err
}
