package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubLiteSubscription = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the subscription.",
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
        "description": "The region of the pubsub lite topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic": {
        "description": "A reference to a Topic resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "description": "The zone of the pubsub lite topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "delivery_config": {
        "block": {
          "attributes": {
            "delivery_requirement": {
              "description": "When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: [\"DELIVER_IMMEDIATELY\", \"DELIVER_AFTER_STORED\", \"DELIVERY_REQUIREMENT_UNSPECIFIED\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The settings for this subscription's message delivery.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GooglePubsubLiteSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubLiteSubscription), &result)
	return &result
}
