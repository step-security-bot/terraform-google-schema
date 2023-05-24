package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcareHl7V2StoreIamBinding = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hl7_v2_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "members": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "condition": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "title": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleHealthcareHl7V2StoreIamBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcareHl7V2StoreIamBinding), &result)
	return &result
}
