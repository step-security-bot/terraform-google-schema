package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvironmentIamPolicy = `{
  "block": {
    "attributes": {
      "env_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "org_id": {
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

func GoogleApigeeEnvironmentIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvironmentIamPolicy), &result)
	return &result
}
