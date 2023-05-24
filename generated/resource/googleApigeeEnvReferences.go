package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvReferences = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Optional. A human-readable description of this reference.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "description": "Required. The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "refers": {
        "description": "Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "description": "The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.",
        "description_kind": "plain",
        "required": true,
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

func GoogleApigeeEnvReferencesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvReferences), &result)
	return &result
}
