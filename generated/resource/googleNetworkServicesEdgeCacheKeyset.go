package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesEdgeCacheKeyset = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of label tags associated with the EdgeCache resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,\nand all following characters must be a dash, underscore, letter or digit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "public_key": {
        "block": {
          "attributes": {
            "id": {
              "description": "The ID of the public key. The ID must be 1-63 characters long, and comply with RFC1035.\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]*\nwhich means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "managed": {
              "description": "Set to true to have the CDN automatically manage this public key value.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "value": {
              "description": "The base64-encoded value of the Ed25519 public key. The base64 encoding can be padded (44 bytes) or unpadded (43 bytes).\nRepresentations or encodings of the public key other than this will be rejected with an error.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "An ordered list of Ed25519 public keys to use for validating signed requests.\nYou must specify 'public_keys' or 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.\nYou may specify no more than one Google-managed public key.\nIf you specify 'public_keys', you must specify at least one (1) key and may specify up to three (3) keys.\n\nEd25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.\nEnsure that the private key is kept secret, and that only authorized users can add public keys to a keyset.",
          "description_kind": "plain"
        },
        "max_items": 3,
        "nesting_mode": "list"
      },
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
      },
      "validation_shared_keys": {
        "block": {
          "attributes": {
            "secret_version": {
              "description": "The name of the secret version in Secret Manager.\n\nThe resource name of the secret version must be in the format 'projects/*/secrets/*/versions/*' where the '*' values are replaced by the secrets themselves.\nThe secrets must be at least 16 bytes large.  The recommended secret size depends on the signature algorithm you are using.\n* If you are using HMAC-SHA1, we suggest 20-byte secrets.\n* If you are using HMAC-SHA256, we suggest 32-byte secrets.\nSee RFC 2104, Section 3 for more details on these recommendations.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An ordered list of shared keys to use for validating signed requests.\nShared keys are secret.  Ensure that only authorized users can add 'validation_shared_keys' to a keyset.\nYou can rotate keys by appending (pushing) a new key to the list of 'validation_shared_keys' and removing any superseded keys.\nYou must specify 'public_keys' or 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.",
          "description_kind": "plain"
        },
        "max_items": 3,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleNetworkServicesEdgeCacheKeysetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesEdgeCacheKeyset), &result)
	return &result
}
