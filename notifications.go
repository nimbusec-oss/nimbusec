package nimbusec

import (
	"context"
	"fmt"
	"net/http"
)

type NotificationService service

func (srv *NotificationService) List(ctx context.Context) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/notifications", nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) ListByDomain(ctx context.Context, id DomainID) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/domains/%d/notifications", id), nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) ListByUser(ctx context.Context, id UserID) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/users/%d/notifications", id), nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) Get(ctx context.Context, id NotificationID) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodGet, fmt.Sprintf("/v3/notifications/%d", id), nil, &notification)
	return notification, err
}

func (srv *NotificationService) Create(ctx context.Context, create Notification) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodPost, "/v3/notifications", create, &notification)
	return notification, err
}

func (srv *NotificationService) Update(ctx context.Context, id NotificationID, update NotificationUpdate) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodPut, fmt.Sprintf("/v3/notifications/%d", id), update, &notification)
	return notification, err
}

func (srv *NotificationService) Delete(ctx context.Context, id NotificationID) error {
	err := srv.client.Do(ctx, http.MethodDelete, fmt.Sprintf("/v3/notifications/%d", id), nil, nil)
	return err
}
