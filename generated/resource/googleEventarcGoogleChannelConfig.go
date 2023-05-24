package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcGoogleChannelConfig = `{
  "block": {
    "attributes": {
      "crypto_key_name": {
        "description": "Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern ` + "`" + `projects/*/locations/*/keyRings/*/cryptoKeys/*` + "`" + `.",
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
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Required. The resource name of the config. Must be in the format of, ` + "`" + `projects/{project}/locations/{location}/googleChannelConfig` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last-modified time.",
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

func GoogleEventarcGoogleChannelConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcGoogleChannelConfig), &result)
	return &result
}
