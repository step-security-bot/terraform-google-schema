package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKeyVersion = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key": {
        "description_kind": "plain",
        "required": true,
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
      "protection_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "pem": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsCryptoKeyVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKeyVersion), &result)
	return &result
}
