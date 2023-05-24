package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountIdToken = `{
  "block": {
    "attributes": {
      "delegates": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id_token": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "include_email": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_audience": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_service_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountIdTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountIdToken), &result)
	return &result
}
