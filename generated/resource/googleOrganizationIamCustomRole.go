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
        "description": "The current deleted state of the role.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "description": "A human-readable description for the role.",
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
      "name": {
        "computed": true,
        "description": "The name of the role in the format organizations/{{org_id}}/roles/{{role_id}}. Like id, this field can be used as a reference in other resources such as IAM role bindings.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The numeric ID of the organization in which you want to create a custom role.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions": {
        "description": "The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "role_id": {
        "description": "The role id to use for this role.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "description": "The current launch stage of the role. Defaults to GA.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "description": "A human-readable title for the role.",
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
