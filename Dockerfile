FROM golang:1.12

LABEL version="1.0.0"
LABEL maintainer="cpanato"
LABEL repository="http://github.com/cpanato/action-mattermost-notify"
LABEL homepage="http://github.com/cpanato/action-mattermost-notify"
LABEL "com.github.actions.name"="Mattermost"
LABEL "com.github.actions.description"="Send a Mattermost message"
LABEL "com.github.actions.icon"="message-circle"
LABEL "com.github.actions.color"="white"

RUN go get github.com/cpanato/action-mattermost-notify

ENTRYPOINT ["mattermost"]
