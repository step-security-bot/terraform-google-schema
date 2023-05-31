package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrganizationIamPolicy = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "org_id": {
        "description": "The numeric ID of the organization in which you want to manage the audit logging config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOrganizationIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrganizationIamPolicy), &result)
	return &result
}
