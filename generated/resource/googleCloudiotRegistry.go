package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudiotRegistry = `{
  "block": {
    "attributes": {
      "http_config": {
        "computed": true,
        "description": "Activate or deactivate HTTP.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_level": {
        "description": "The default logging verbosity for activity from devices in this\nregistry. Specifies which events should be written to logs. For\nexample, if the LogLevel is ERROR, only events that terminate in\nerrors will be logged. LogLevel is inclusive; enabling INFO logging\nwill also enable ERROR logging. Default value: \"NONE\" Possible values: [\"NONE\", \"ERROR\", \"INFO\", \"DEBUG\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mqtt_config": {
        "computed": true,
        "description": "Activate or deactivate MQTT.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "A unique name for the resource, required by device registry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region in which the created registry should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state_notification_config": {
        "description": "A PubSub topic to publish device state updates.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "credentials": {
        "block": {
          "attributes": {
            "public_key_certificate": {
              "description": "A public key certificate format and data.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "List of public key certificates to authenticate devices.",
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "list"
      },
      "event_notification_configs": {
        "block": {
          "attributes": {
            "pubsub_topic_name": {
              "description": "PubSub topic name to publish device events.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subfolder_matches": {
              "description": "If the subfolder name matches this string exactly, this\nconfiguration will be used. The string must not include the\nleading '/' character. If empty, all strings are matched. Empty\nvalue can only be used for the last 'event_notification_configs'\nitem.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "List of configurations for event notifications, such as PubSub topics\nto publish device events to.",
          "description_kind": "plain"
        },
        "max_items": 10,
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

func GoogleCloudiotRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudiotRegistry), &result)
	return &result
}
