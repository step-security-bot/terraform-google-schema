package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageTransferProjectServiceAccount = `{
  "block": {
    "attributes": {
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageTransferProjectServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageTransferProjectServiceAccount), &result)
	return &result
}
