package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageNotification = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_attributes": {
        "description": " A set of key/value attribute pairs to attach to each Cloud Pub/Sub message published for this notification subscription",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "event_types": {
        "description": "List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: \"OBJECT_FINALIZE\", \"OBJECT_METADATA_UPDATE\", \"OBJECT_DELETE\", \"OBJECT_ARCHIVE\"",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_id": {
        "computed": true,
        "description": "The ID of the created notification.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_name_prefix": {
        "description": "Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payload_format": {
        "description": "The desired content of the Payload. One of \"JSON_API_V1\" or \"NONE\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "topic": {
        "description": "The Cloud Pub/Sub topic to which this subscription publishes. Expects either the  topic name, assumed to belong to the default GCP provider project, or the project-level name,  i.e. projects/my-gcp-project/topics/my-topic or my-topic. If the project is not set in the provider, you will need to use the project-level name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageNotification), &result)
	return &result
}
