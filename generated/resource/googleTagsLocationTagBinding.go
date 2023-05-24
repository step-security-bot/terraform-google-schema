package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTagsLocationTagBinding = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The geographic location where the transfer config should reside.\nExamples: US, EU, asia-northeast1. The default value is US.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_value": {
        "description": "The TagValue of the TagBinding. Must be of the form tagValues/456.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func GoogleTagsLocationTagBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTagsLocationTagBinding), &result)
	return &result
}
