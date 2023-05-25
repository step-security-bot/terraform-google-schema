package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamRole = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "included_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleIamRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamRole), &result)
	return &result
}
