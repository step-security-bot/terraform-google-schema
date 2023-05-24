package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvgroupAttachment = `{
  "block": {
    "attributes": {
      "envgroup_id": {
        "description": "The Apigee environment group associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/envgroups/{{envgroup_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment": {
        "description": "The resource ID of the environment.",
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
      "name": {
        "computed": true,
        "description": "The name of the newly created  attachment (output parameter).",
        "description_kind": "plain",
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

func GoogleApigeeEnvgroupAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvgroupAttachment), &result)
	return &result
}
