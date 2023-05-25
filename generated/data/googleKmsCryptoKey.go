package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_ring": {
        "description": "The KeyRing that this key belongs to.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels with user-defined metadata to apply to this resource.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The immutable purpose of this CryptoKey. See the\n[purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)\nfor possible inputs.",
        "description_kind": "plain",
        "type": "string"
      },
      "rotation_period": {
        "computed": true,
        "description": "Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.\nThe first rotation will take place after the specified period. The rotation period has\nthe format of a decimal number with up to 9 fractional digits, followed by the\nletter 's' (seconds). It must be greater than a day (ie, 86400).",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version_template": {
        "computed": true,
        "description": "A template describing settings for new crypto key versions.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "protection_level": "string"
            }
          ]
        ]
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
