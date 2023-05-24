package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubLiteReservation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the reservation.",
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
        "description": "The region of the pubsub lite reservation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "throughput_capacity": {
        "description": "The reserved throughput capacity. Every unit of throughput capacity is\nequivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed\nmessages.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
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

func GooglePubsubLiteReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubLiteReservation), &result)
	return &result
}
