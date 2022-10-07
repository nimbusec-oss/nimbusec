package nimbusec

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type ScreenshotService service

func (srv *ScreenshotService) GetScreenshots(ctx context.Context, id DomainID, filter *ScreenshotFilter) (Screenshot, error) {
	v, err := query.Values(filter)
	if err != nil {
		return Screenshot{}, err
	}
	u := url.URL{
		Path:     fmt.Sprintf("/v3/domains/%d/screenshots", id),
		RawQuery: v.Encode(),
	}

	screenshot := Screenshot{}
	err = srv.client.Do(ctx, http.MethodGet, u.String(), nil, &screenshot)
	return screenshot, err
}

func (srv *ScreenshotService) GetImage(ctx context.Context, id DomainID, typez ScreenshotType) ([]byte, error) {
	image := []byte{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/domains/%d/screenshots/%s.jpg", id, typez), nil, &image)
	return image, err
}
