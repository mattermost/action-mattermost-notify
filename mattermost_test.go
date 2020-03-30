package main_test

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/stretchr/testify/assert"

	notify "github.com/mattermost/action-mattermost-notify"
)

func TestIsEmpty(t *testing.T) {
	for name, test := range map[string]struct {
		request *model.IncomingWebhookRequest
		result  bool
	}{
		"nil": {
			request: nil,
			result:  true,
		},
		"Text and Attachments empty": {
			request: &model.IncomingWebhookRequest{},
			result:  true,
		},
		"Text not empty": {
			request: &model.IncomingWebhookRequest{
				Text: "some Text",
			},
			result: false,
		},
		"Attachments not empty": {
			request: &model.IncomingWebhookRequest{
				Attachments: []*model.SlackAttachment{{
					Text: "some Text",
				}},
			},
			result: false,
		},
		"Attachments not empty, but no text": {
			request: &model.IncomingWebhookRequest{
				Attachments: []*model.SlackAttachment{},
			},
			result: false,
		},
		"Text and Attachments not empty": {
			request: &model.IncomingWebhookRequest{
				Text: "some Text",
				Attachments: []*model.SlackAttachment{{
					Text: "some Text",
				}},
			},
			result: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			r := notify.IsEmpty(test.request)
			assert.Equal(t, test.result, r)
		})
	}
}
