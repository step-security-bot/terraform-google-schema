package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectIamCustomRole = `{
  "block": {
    "attributes": {
      "deleted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
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
      "permissions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "description_kind": "plain",
        "optional": true,
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
  "version": 0
}`

func GoogleProjectIamCustomRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectIamCustomRole), &result)
	return &result
}
