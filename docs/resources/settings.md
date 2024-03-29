---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awsteam_settings Resource - terraform-provider-awsteam"
subcategory: ""
description: |-
  Allows configuration of the settings within an AWS TEAM deployment.
  Important: By default the settings resource will already exist on a fresh deployment of AWS TEAM and will cause a create to fail. Use the import block example below when deploying to a fresh instance of AWS TEAM to import the existing settings
---

# awsteam_settings (Resource)

Allows configuration of the settings within an AWS TEAM deployment.

> **Important:** By default the `settings` resource will already exist on a fresh deployment of AWS TEAM and will cause a create to fail. Use the import block example below when deploying to a fresh instance of AWS TEAM to import the existing `settings`

## Example Usage

```terraform
resource "awsteam_settings" "example" {
  duration           = 5
  expiry             = 3
  team_admin_group   = "My-Team-Admin-Group"
  team_auditor_group = "My-Team-Auditor-Group"
}

// Import the existing settings on a fresh install of AWS TEAM
import {
  to = awsteam_settings.example
  id = "settings"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `duration` (Number) Default maximum request duration in hours.
- `expiry` (Number) Number of time in hours before an unapproved TEAM request expires.
- `team_admin_group` (String) Group of users responsible for managing TEAM administrative configurations
- `team_auditor_group` (String) Group of users allowed to audit TEAM elevated access requests

### Optional

- `approval` (Boolean) If disabled, approval will not be required for all elevated access requests. If enabled, approval requirement is managed in eligibility policy configuration.
- `comments` (Boolean) Determines if comment field is mandatory for all elevated access requests.
- `ses_notifications_enabled` (Boolean) Enable sending notifications via Amazon SES.
- `ses_source_arn` (String) ARN of a verified SES identity in another AWS account. Must be configured to authorize sending mail from the TEAM account.
- `ses_source_email` (String) Email address to send notifications from. Must be verified in SES.
- `slack_notifications_enabled` (Boolean) Enable to send notifications directly to users in Slack via a Slack bot app.
- `slack_token` (String, Sensitive) Slack OAuth token associated with the installed app.
- `sns_notifications_enabled` (Boolean) Send notifications via Amazon SNS. Once enabled, create a subscription to the SNS topic (TeamNotifications-main) in the TEAM account.
- `ticket_no` (Boolean) Determines if ticket number field is mandatory for elevated access requests

### Read-Only

- `created_at` (String) The date and time that the item was created
- `id` (String) The settings identifier
- `modified_by` (String) The user to last modify the item
- `updated_at` (String) The date and time of the last time the item was updated

## Import

Import is supported using the following syntax:

```shell
# Import using settings ID, this is always "settings"
terraform import awsteam_settings.example settings
```
