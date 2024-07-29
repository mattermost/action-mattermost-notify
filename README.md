## Github Action Mattermost notify

[Mattermost](https://mattermost.com) — Send Mattermost messages. Fork/inspired
from https://github.com/apex/actions/tree/master/slack

For the message format you can check the
[Mattermost documentation](https://developers.mattermost.com/integrate/webhooks/incoming/)

## Inputs

| Input                  | Description                                        | Default         | Required |
| ---------------------- | -------------------------------------------------- | --------------- | -------- |
| MATTERMOST_WEBHOOK_URL | The Mattermost Incoming Webhook URL                |                 | true     |
| MATTERMOST_CHANNEL     | The Mattermost channel to sent the message         |                 | false    |
| MATTERMOST_USERNAME    | The Mattermost username shown in the webhook       |                 | false    |
| MATTERMOST_ICON_URL    | The Mattermost Icon URL for the webhook            |                 | false    |
| TEXT                   | The text for the webhook message                   |                 | false    |
| PAYLOAD                | The payload for the webhook                        |                 | false    |
| PAYLOAD_FILENAME       | The payload file name for the webhook (deprecated) | mattermost.json | false    |

## Usage

You can use it as below:

```yaml
on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: mattermost/action-mattermost-notify@master
        with:
          MATTERMOST_WEBHOOK_URL: ${{ secrets.MM_WEBHOOK_URL }}
          MATTERMOST_CHANNEL: the-best-channel
          TEXT: |
            This is a message from ${{ github.repository }}.
            [Pipeline](${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}) was completed :white_check_mark:
          MATTERMOST_USERNAME: ${{ github.triggering_actor }}
          MATTERMOST_ICON_URL: https://cdn3.iconfinder.com/data/icons/system-basic-vol-4-1/20/icon-note-attention-alt3-512.png
```

or alternatively you can use the `PAYLOAD` input:

```yaml
on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: mattermost/action-mattermost-notify@master
        with:
          MATTERMOST_WEBHOOK_URL: ${{ secrets.MM_WEBHOOK_URL }}
          PAYLOAD: |-
            {
              "text": "What a message from ${{ github.repository }}",
              "channel": "the-best-channel",
              "username": "${{ github.triggering_actor }}",
              "icon": "https://cdn3.iconfinder.com/data/icons/system-basic-vol-4-1/20/icon-note-attention-alt3-512.png"
            }
```

## Update the Action Metadata

The [`action.yml`](action.yml) file defines metadata about your action, such as
input(s) and output(s). For details about this file, see
[Metadata syntax for GitHub Actions](https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions).

When you copy this repository, update `action.yml` with the name, description,
inputs, and outputs for your action.

## Update the Action Code

The [`src/`](./src/) directory is the heart of your action! This contains the
source code that will be run when your action is invoked. You can replace the
contents of this directory with your own code.

There are a few things to keep in mind when writing your action code:

- Most GitHub Actions toolkit and CI/CD operations are processed asynchronously.
  In `main.js`, you will see that the action is run in an `async` function.

  ```javascript
  const core = require('@actions/core')
  //...

  async function run() {
    try {
      //...
    } catch (error) {
      core.setFailed(error.message)
    }
  }
  ```

  For more information about the GitHub Actions toolkit, see the
  [documentation](https://github.com/actions/toolkit/blob/master/README.md).

So, what are you waiting for? Go ahead and start customizing your action!

1. Create a new branch

   ```bash
   git checkout -b releases/v1
   ```

1. Replace the contents of `src/` with your action code
1. Add tests to `__tests__/` for your source code
1. Format, test, and build the action

   ```bash
   npm run all
   ```

1. Commit your changes

   ```bash
   git add .
   git commit -m "My first action is ready!"
   ```

1. Push them to your repository

   ```bash
   git push -u origin releases/v1
   ```

1. Create a pull request and get feedback on your action
1. Merge the pull request into the `main` branch

Your action is now published! :rocket:

For information about versioning your action, see
[Versioning](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md)
in the GitHub Actions toolkit.
