package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountIamPolicy = `{
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
      "policy_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountIamPolicy), &result)
	return &result
}
