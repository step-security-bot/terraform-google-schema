package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubNamespace = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the Namespace was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Time the Namespace was deleted in UTC.",
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
        "description": "The resource name for the namespace",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "The name of the Scope instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope_id": {
        "description": "Id of the scope",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope_namespace_id": {
        "description": "The client-provided identifier of the namespace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the namespace resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "Google-generated UUID for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the Namespace was updated in UTC.",
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

func GoogleGkeHubNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubNamespace), &result)
	return &result
}
