package nimbusec

import (
	"context"
	"net/http"
)

type DomainStatsService service

func (srv *DomainStatsService) Get(ctx context.Context) (DomainStats, error) {
	domainStats := DomainStats{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/domains/stats", nil, &domainStats)

	return domainStats, err
}
