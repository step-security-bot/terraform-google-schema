package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringNotificationChannel = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.",
        "description_kind": "plain",
        "type": "bool"
      },
      "force_delete": {
        "computed": true,
        "description": "If true, the notification channel will be deleted regardless\nof its use in alert policies (the policies will be updated\nto remove the channel). If false, channels that are still\nreferenced by an existing alerting policy will fail to be\ndeleted in a delete operation.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Configuration fields that define the channel and its behavior. The\npermissible and required labels are specified in the\nNotificationChannelDescriptor corresponding to the type field.\n\nLabels with sensitive data are obfuscated by the API and therefore Terraform cannot\ndetermine if there are upstream changes to these fields. They can also be configured via\nthe sensitive_labels block, but cannot be configured in both places.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The full REST resource name for this channel. The syntax is:\nprojects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]\nThe [CHANNEL_ID] is automatically assigned by the server on creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sensitive_labels": {
        "computed": true,
        "description": "Different notification type behaviors are configured primarily using the the 'labels' field on this\nresource. This block contains the labels which contain secrets or passwords so that they can be marked\nsensitive and hidden from plan output. The name of the field, eg: password, will be the key\nin the 'labels' map in the api request.\n\nCredentials may not be specified in both locations and will cause an error. Changing from one location\nto a different credential configuration in the config will require an apply to update state.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auth_token": "string",
              "password": "string",
              "service_key": "string"
            }
          ]
        ]
      },
      "type": {
        "description": "The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as \"email\", \"slack\", etc...",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_labels": {
        "description": "User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "verification_status": {
        "computed": true,
        "description": "Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMonitoringNotificationChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringNotificationChannel), &result)
	return &result
}
