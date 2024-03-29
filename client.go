package nimbusec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type service struct {
	client *Client
}

type Client struct {
	http *http.Client
	base string

	common         service // Reuse a single struct instead of allocating one for each service on the heap.
	Bundles        *BundleService
	Domains        *DomainService
	DomainMetadata *DomainMetadataService
	DomainStats    *DomainStatsService
	Issues         *IssueService
	Notifications  *NotificationService
	Screenshots    *ScreenshotService
}

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d| %s", e.StatusCode, e.Message)
}

type Config struct {
	ClientID     string
	ClientSecret string
	Endpoint     Endpoint
}

type Endpoint struct {
	AuthURL   string
	ServerURL string
}

var DefaultEndpoint = Endpoint{
	AuthURL:   "https://auth.nimbusec.com/oauth2/token",
	ServerURL: "https://api.nimbusec.com",
}

func NewClient(ctx context.Context, config Config) *Client {
	oauth := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.AuthURL,
		Scopes:       []string{"onehorn:admin"},
	}

	c := &Client{http: oauth.Client(ctx), base: config.Endpoint.ServerURL}
	c.common = service{client: c}
	c.Bundles = (*BundleService)(&c.common)
	c.Domains = (*DomainService)(&c.common)
	c.DomainMetadata = (*DomainMetadataService)(&c.common)
	c.DomainStats = (*DomainStatsService)(&c.common)
	c.Issues = (*IssueService)(&c.common)
	c.Notifications = (*NotificationService)(&c.common)
	c.Screenshots = (*ScreenshotService)(&c.common)
	return c
}

func (client *Client) Ping(ctx context.Context) error {
	err := client.Do(ctx, http.MethodGet, "/v3/ping", nil, nil)
	return err
}

func (client *Client) Do(ctx context.Context, method, url string, in, out interface{}) error {
	var req *http.Request
	var err error

	// encode `in` object as json body if provided
	if in != nil {
		body := &bytes.Buffer{}
		err := json.NewEncoder(body).Encode(in)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, client.base+url, body)
	} else {
		req, err = http.NewRequest(method, client.base+url, nil)
	}

	if err != nil {
		return err
	}

	// perform request
	req = req.WithContext(ctx)
	resp, err := client.http.Do(req)
	if err != nil {
		return err
	}

	// decode response body into `out` object if provided
	defer resp.Body.Close()

	// check for api errors, try to decode in error object
	if resp.StatusCode >= 300 {
		msg := Error{StatusCode: resp.StatusCode}
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			return err
		}

		return msg
	}

	if out != nil {
		switch e := out.(type) {
		case *[]byte:
			read, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			*e = read
		default:
			err = json.NewDecoder(resp.Body).Decode(&out)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
