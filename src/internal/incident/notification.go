package incident

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type SlackRequest struct {
	*HttpRequest
}

func (sr *SlackRequest) ExecHttpReq(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-type", "application/json")
	resp, err := sr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf("HTTP status code: %d, error: %s", resp.StatusCode, body)
	}
	return body, nil
}

var _ = Requester(&SlackRequest{})

type Notificaton struct {
	webhookUrl string
	request    Requester
}

type NotificatorOption func(*Notificaton)

func WithRequest(request Requester) NotificatorOption {
	return func(i *Notificaton) {
		i.request = request
	}
}

func NewNotificator(webhookUrl string, options ...NotificatorOption) *Notificaton {
	notification := &Notificaton{
		webhookUrl: webhookUrl,
		request: &SlackRequest{
			HttpRequest: NewHttpRequest(),
		},
	}
	for _, option := range options {
		option(notification)
	}
	return notification
}

func (n *Notificaton) SendNotification(content string) error {
	req, err := http.NewRequest(
		"POST",
		n.webhookUrl,
		bytes.NewBuffer([]byte(content)),
	)
	if err != nil {
		return err
	}

	_, err = n.request.ExecHttpReq(req)
	if err != nil {
		return err
	}
	return nil
}
