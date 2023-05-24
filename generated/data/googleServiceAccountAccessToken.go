package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountAccessToken = `{
  "block": {
    "attributes": {
      "access_token": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
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
      "lifetime": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scopes": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_service_account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountAccessTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountAccessToken), &result)
	return &result
}
