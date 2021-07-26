FROM docker.io/golang:1.16 as builder

ENV GOPATH=/go/
USER root
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o action-mattermost-notify .

FROM gcr.io/distroless/base

LABEL version="1.1.0"
LABEL maintainer="mattermost"
LABEL repository="http://github.com/mattermost/action-mattermost-notify"
LABEL homepage="http://github.com/mattermost/action-mattermost-notify"
LABEL "com.github.actions.name"="Mattermost"
LABEL "com.github.actions.description"="Send a Mattermost message"
LABEL "com.github.actions.icon"="send"
LABEL "com.github.actions.color"="blue"

COPY --from=builder /app/action-mattermost-notify /bin/action-mattermost-notify
ENTRYPOINT ["action-mattermost-notify"]
