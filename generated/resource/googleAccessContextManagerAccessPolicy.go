package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerAccessPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the AccessPolicy was created in UTC.",
        "description_kind": "plain",
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
        "description": "Resource name of the AccessPolicy. Format: {policy_id}",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of this AccessPolicy in the Cloud Resource Hierarchy.\nFormat: organizations/{organization_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scopes": {
        "description": "Folder or project on which this policy is applicable.\nFormat: folders/{{folder_id}} or projects/{{project_id}}",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "title": {
        "description": "Human readable title. Does not affect behavior.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the AccessPolicy was updated in UTC.",
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

func GoogleAccessContextManagerAccessPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerAccessPolicy), &result)
	return &result
}
