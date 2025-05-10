# waroom-notification-action
 Action for Waroom incident notification

# waroom-notification-action - Notificate Waroom Incidents for GitHub Actions

## Introduction

This action references incident information in Waroom and notifies Slack of the acquired incident information.

> [!WARNING]
> Slack notifications only support incoming-webhook.

## Parameter Reference

### inputs

| Name          | Type     | Required | Default | Description                                                                      |
| ------------- | -------- | -------- | ------- | -------------------------------------------------------------------------------- |
| `api-key`     | `String` | `true`   |         | Waroom API key                                                                   |
| `webhook-url` | `String` | `true`   |         | Slack webhook URL                                                                |
| `template`    | `String` | `true`   |         | Go template for the message(You can also specify the template file path with @.) |

### Usage

```yaml
name: ci

on:
  push:
    branches:
      - 'main'

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Package List
        uses: chmikata/waroom-notification-action@v0.1.0
        with:
          api-key: xxxxxxxxx
          webhook-url: https://xxxxx
          template: @tpl/output.tpl
