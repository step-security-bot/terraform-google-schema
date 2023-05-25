package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsSecret = `{
  "block": {
    "attributes": {
      "additional_authenticated_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ciphertext": {
        "description_kind": "plain",
        "required": true,
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
      "plaintext": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsSecret), &result)
	return &result
}
