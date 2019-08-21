package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/model"
)

// ReadMessage reads a message from the given file.
func ReadMessage(path string) (*model.IncomingWebhookRequest, error) {
	bytesMsg, err := ioutil.ReadFile(path)
	if err != nil {
		return &model.IncomingWebhookRequest{}, err
	}

	iwr, _ := model.IncomingWebhookRequestFromJson(bytes.NewReader(bytesMsg))
	if iwr == nil {
		return &model.IncomingWebhookRequest{}, errors.Errorf("error parsing the message")
	}

	return iwr, nil
}

// Send a message to the given webhook url.
func Send(url string, msg *model.IncomingWebhookRequest) error {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return errors.Wrap(err, "marshaling")
	}

	res, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return errors.Wrap(err, "requesting")
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return errors.Errorf("%s response", res.Status)
	}

	return nil
}
