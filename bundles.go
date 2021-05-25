package nimbusec

import (
	"context"
	"fmt"
	"net/http"
)

type BundleService service

func (srv *BundleService) List(ctx context.Context) ([]Bundle, error) {
	bundles := []Bundle{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/bundles", nil, &bundles)
	return bundles, err
}

func (srv *BundleService) Get(ctx context.Context, id BundleID) (Bundle, error) {
	bundle := Bundle{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/bundles/%s", id), nil, &bundle)
	return bundle, err
}
