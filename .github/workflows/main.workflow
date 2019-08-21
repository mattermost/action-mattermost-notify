workflow "Build and Deploy" {
  on = "pull_request"
  resolves = ["Generate", "Send Message"]
}

action "Generate" {
  runs = ["sh", "-c", "echo '{\"text\":\"testing github action :tada:\"}' > mattermost.json"]
}

action "Send Message" {
  needs = ["Generate"]
  uses = "cpanato/action-mattermost-notify@master"
  secrets = ["MATTERMOST_WEBHOOK_URL"]
}
