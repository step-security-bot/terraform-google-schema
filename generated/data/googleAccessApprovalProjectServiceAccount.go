package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessApprovalProjectServiceAccount = `{
  "block": {
    "attributes": {
      "account_email": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAccessApprovalProjectServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessApprovalProjectServiceAccount), &result)
	return &result
}
