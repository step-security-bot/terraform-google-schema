package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamTestablePermissions = `{
  "block": {
    "attributes": {
      "custom_support_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "full_resource_name": {
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
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "api_disabled": "bool",
              "custom_support_level": "string",
              "name": "string",
              "stage": "string",
              "title": "string"
            }
          ]
        ]
      },
      "stages": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleIamTestablePermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamTestablePermissions), &result)
	return &result
}
