package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccount = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
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
      "member": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unique_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccount), &result)
	return &result
}
