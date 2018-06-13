package nimbusec

import (
	"context"
	"net/http"
)

type DomainMetadataService service

func (srv *DomainMetadataService) Get(ctx context.Context, id DomainID) (DomainMetadata, error) {
	domainMetadata := DomainMetadata{}
	err := srv.client.Do(ctx, http.MethodGet, string(id)+"/metadata", nil, &domainMetadata)
	return domainMetadata, err
}
