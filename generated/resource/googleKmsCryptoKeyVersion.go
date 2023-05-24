package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKeyVersion = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description": "The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "attestation": {
        "computed": true,
        "description": "Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.\nOnly provided for key versions with protectionLevel HSM.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert_chains": [
                "list",
                [
                  "object",
                  {
                    "cavium_certs": "string",
                    "google_card_certs": "string",
                    "google_partition_certs": "string"
                  }
                ]
              ],
              "content": "string",
              "external_protection_level_options": [
                "list",
                [
                  "object",
                  {
                    "ekm_connection_key_path": "string",
                    "external_key_uri": "string"
                  }
                ]
              ],
              "format": "string"
            }
          ]
        ]
      },
      "crypto_key": {
        "description": "The name of the cryptoKey associated with the CryptoKeyVersions.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "generate_time": {
        "computed": true,
        "description": "The time this CryptoKeyVersion key material was generated",
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
        "description": "The resource name for this CryptoKeyVersion.",
        "description_kind": "plain",
        "type": "string"
      },
      "protection_level": {
        "computed": true,
        "description": "The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the CryptoKeyVersion. Possible values: [\"PENDING_GENERATION\", \"ENABLED\", \"DISABLED\", \"DESTROYED\", \"DESTROY_SCHEDULED\", \"PENDING_IMPORT\", \"IMPORT_FAILED\"]",
        "description_kind": "plain",
        "optional": true,
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
            },
            "update": {
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

func GoogleKmsCryptoKeyVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKeyVersion), &result)
	return &result
}
