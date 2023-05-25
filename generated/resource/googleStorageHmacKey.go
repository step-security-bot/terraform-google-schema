package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageHmacKey = `{
  "block": {
    "attributes": {
      "access_id": {
        "computed": true,
        "description": "The access ID of the HMAC Key.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret": {
        "computed": true,
        "description": "HMAC secret key material.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "service_account_email": {
        "description": "The email address of the key's associated service account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "description": "The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: \"ACTIVE\" Possible values: [\"ACTIVE\", \"INACTIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_created": {
        "computed": true,
        "description": "'The creation time of the HMAC key in RFC 3339 format. '",
        "description_kind": "plain",
        "type": "string"
      },
      "updated": {
        "computed": true,
        "description": "'The last modification time of the HMAC key metadata in RFC 3339 format.'",
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

func GoogleStorageHmacKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageHmacKey), &result)
	return &result
}
