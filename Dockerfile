FROM golang:1.13

LABEL version="1.0.0"
LABEL maintainer="mattermost"
LABEL repository="http://github.com/mattermost/action-mattermost-notify"
LABEL homepage="http://github.com/mattermost/action-mattermost-notify"
LABEL "com.github.actions.name"="Mattermost"
LABEL "com.github.actions.description"="Send a Mattermost message"
LABEL "com.github.actions.icon"="send"
LABEL "com.github.actions.color"="blue"

RUN go get github.com/mattermost/action-mattermost-notify

ENTRYPOINT ["action-mattermost-notify"]
