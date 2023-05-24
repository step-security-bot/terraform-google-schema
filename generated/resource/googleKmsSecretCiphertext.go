package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsSecretCiphertext = `{
  "block": {
    "attributes": {
      "additional_authenticated_data": {
        "description": "The additional authenticated data used for integrity checks during encryption and decryption.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "ciphertext": {
        "computed": true,
        "description": "Contains the result of encrypting the provided plaintext, encoded in base64.",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key": {
        "description": "The full name of the CryptoKey that will be used to encrypt the provided plaintext.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''",
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
        "description": "The plaintext to be encrypted.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
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
