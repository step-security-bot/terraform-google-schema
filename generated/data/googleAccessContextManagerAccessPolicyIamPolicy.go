package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerAccessPolicyIamPolicy = `{
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
      "name": {
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

func GoogleAccessContextManagerAccessPolicyIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerAccessPolicyIamPolicy), &result)
	return &result
}
