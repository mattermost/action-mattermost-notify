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
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return &model.IncomingWebhookRequest{}, err
	}

	incomingMessage, err := model.IncomingWebhookRequestFromJson(bytes.NewReader(b))
	if err != nil {
		return &model.IncomingWebhookRequest{}, errors.Wrap(err, "unmarshaling")
	}

	return incomingMessage, nil
}

// // WriteMessage writes a message to the given file.
// func WriteMessage(path string, msg *Message) error {
// 	b, err := json.MarshalIndent(msg, "", "  ")
// 	if err != nil {
// 		return errors.Wrap(err, "marshaling")
// 	}

// 	err = ioutil.WriteFile(path, b, 0755)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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
