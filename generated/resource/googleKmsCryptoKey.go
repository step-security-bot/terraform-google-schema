package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKey = `{
  "block": {
    "attributes": {
      "destroy_scheduled_duration": {
        "computed": true,
        "description": "The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.\nIf not specified at creation time, the default duration is 24 hours.",
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
      "import_only": {
        "computed": true,
        "description": "Whether this key may contain imported versions only.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "key_ring": {
        "description": "The KeyRing that this key belongs to.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels with user-defined metadata to apply to this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The resource name for the CryptoKey.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "purpose": {
        "description": "The immutable purpose of this CryptoKey. See the\n[purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)\nfor possible inputs. Default value: \"ENCRYPT_DECRYPT\" Possible values: [\"ENCRYPT_DECRYPT\", \"ASYMMETRIC_SIGN\", \"ASYMMETRIC_DECRYPT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rotation_period": {
        "description": "Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.\nThe first rotation will take place after the specified period. The rotation period has\nthe format of a decimal number with up to 9 fractional digits, followed by the\nletter 's' (seconds). It must be greater than a day (ie, 86400).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "deprecated": true,
        "description": "The self link of the created KeyRing in the format projects/{project}/locations/{location}/keyRings/{name}.",
        "description_kind": "plain",
        "type": "string"
      },
      "skip_initial_version_creation": {
        "description": "If set to true, the request will create a CryptoKey without any CryptoKeyVersions. \nYou must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      },
      "version_template": {
        "block": {
          "attributes": {
            "algorithm": {
              "description": "The algorithm to use when creating a version based on this template.\nSee the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protection_level": {
              "description": "The protection level to use when creating a version based on this template. Default value: \"SOFTWARE\" Possible values: [\"SOFTWARE\", \"HSM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A template describing settings for new crypto key versions.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleKmsCryptoKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKey), &result)
	return &result
}
