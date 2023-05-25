package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsSecretCiphertext = `{
  "block": {
    "attributes": {
      "ciphertext": {
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
      "plaintext": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsSecretCiphertextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsSecretCiphertext), &result)
	return &result
}
