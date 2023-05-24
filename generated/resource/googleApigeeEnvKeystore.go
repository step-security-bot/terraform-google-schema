package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvKeystore = `{
  "block": {
    "attributes": {
      "aliases": {
        "computed": true,
        "description": "Aliases in this keystore.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "env_id": {
        "description": "The Apigee environment group associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/environments/{{env_name}}'.",
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
      "name": {
        "description": "The name of the newly created keystore.",
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

func GoogleApigeeEnvKeystoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvKeystore), &result)
	return &result
}
