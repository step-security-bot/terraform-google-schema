package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrganizationIamCustomRole = `{
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
      "org_id": {
        "description_kind": "plain",
        "required": true,
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

func GoogleOrganizationIamCustomRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrganizationIamCustomRole), &result)
	return &result
}
