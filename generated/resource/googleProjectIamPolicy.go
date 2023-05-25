package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectIamPolicy = `{
  "block": {
    "attributes": {
      "authoritative": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_project": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
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
      "policy_data": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleProjectIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectIamPolicy), &result)
	return &result
}
