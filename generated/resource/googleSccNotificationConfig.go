package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSccNotificationConfig = `{
  "block": {
    "attributes": {
      "config_id": {
        "description": "This must be unique within the organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the notification config (max of 1024 characters).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of this notification config, in the format\n'organizations/{{organization}}/notificationConfigs/{{config_id}}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "description": "The organization whose Cloud Security Command Center the Notification\nConfig lives in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pubsub_topic": {
        "description": "The Pub/Sub topic to send notifications to. Its format is\n\"projects/[project_id]/topics/[topic]\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "The service account that needs \"pubsub.topics.publish\" permission to\npublish to the Pub/Sub topic.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "streaming_config": {
        "block": {
          "attributes": {
            "filter": {
              "description": "Expression that defines the filter to apply across create/update\nevents of assets or findings as specified by the event type. The\nexpression is a list of zero or more restrictions combined via\nlogical operators AND and OR. Parentheses are supported, and OR\nhas higher precedence than AND.\n\nRestrictions have the form \u003cfield\u003e \u003coperator\u003e \u003cvalue\u003e and may have\na - character in front of them to indicate negation. The fields\nmap to those defined in the corresponding resource.\n\nThe supported operators are:\n\n* = for all value types.\n* \u003e, \u003c, \u003e=, \u003c= for integer values.\n* :, meaning substring matching, for strings.\n\nThe supported value types are:\n\n* string literals in quotes.\n* integer literals without quotes.\n* boolean literals true and false without quotes.\n\nSee\n[Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)\nfor information on how to write a filter.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The config for triggering streaming-based notifications.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSccNotificationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSccNotificationConfig), &result)
	return &result
}
