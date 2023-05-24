package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerSecretVersionAccess = `{
  "block": {
    "attributes": {
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_data": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "version": {
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

func GoogleSecretManagerSecretVersionAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerSecretVersionAccess), &result)
	return &result
}
