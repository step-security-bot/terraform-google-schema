package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerSecretVersion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the Secret was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "destroy_time": {
        "computed": true,
        "description": "The time at which the Secret was destroyed. Only present if state is DESTROYED.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description": "The current state of the SecretVersion.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the SecretVersion. Format:\n'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "secret": {
        "description": "Secret Manager secret resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_data": {
        "description": "The secret data. Must be no larger than 64KiB.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the Secret.",
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

func GoogleSecretManagerSecretVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerSecretVersion), &result)
	return &result
}
