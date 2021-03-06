package nimbusec

import (
	"context"
	"fmt"
	"net/http"
)

type DomainMetadataService service

func (srv *DomainMetadataService) List(ctx context.Context) ([]DomainMetadata, error) {
	domainMetadatas := []DomainMetadata{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/domains/metadata", nil, &domainMetadatas)
	return domainMetadatas, err
}

func (srv *DomainMetadataService) Get(ctx context.Context, id DomainID) (DomainMetadata, error) {
	domainMetadata := DomainMetadata{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/domains/%d/metadata", id), nil, &domainMetadata)
	return domainMetadata, err
}
