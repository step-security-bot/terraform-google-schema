package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIapWebBackendServiceIamPolicy = `{
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "web_backend_service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleIapWebBackendServiceIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIapWebBackendServiceIamPolicy), &result)
	return &result
}
