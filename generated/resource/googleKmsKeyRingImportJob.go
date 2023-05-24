package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsKeyRingImportJob = `{
  "block": {
    "attributes": {
      "attestation": {
        "computed": true,
        "description": "Statement that was generated and signed by the key creator (for example, an HSM) at key creation time.\nUse this statement to verify attributes of the key as stored on the HSM, independently of Google.\nOnly present if the chosen ImportMethod is one with a protection level of HSM.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "content": "string",
              "format": "string"
            }
          ]
        ]
      },
      "expire_time": {
        "computed": true,
        "description": "The time at which this resource is scheduled for expiration and can no longer be used.\nThis is in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_job_id": {
        "description": "It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_method": {
        "description": "The wrapping method to be used for incoming key material. Possible values: [\"RSA_OAEP_3072_SHA1_AES_256\", \"RSA_OAEP_4096_SHA1_AES_256\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_ring": {
        "description": "The KeyRing that this import job belongs to.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.",
        "description_kind": "plain",
        "type": "string"
      },
      "protection_level": {
        "description": "The protection level of the ImportJob. This must match the protectionLevel of the\nversionTemplate on the CryptoKey you attempt to import into. Possible values: [\"SOFTWARE\", \"HSM\", \"EXTERNAL\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description": "The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pem": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "The current state of the ImportJob, indicating if it can be used.",
        "description_kind": "plain",
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

func GoogleKmsKeyRingImportJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsKeyRingImportJob), &result)
	return &result
}
