package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubMembershipBinding = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the MembershipBinding was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Time the MembershipBinding was deleted in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_binding_id": {
        "description": "The client-provided identifier of the membership binding.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_id": {
        "description": "Id of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the membershipbinding itself",
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
        "description": "A Workspace resource name in the format\n'projects/*/locations/*/scopes/*'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the membership binding resource.",
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
        "description": "Time the MembershipBinding was updated in UTC.",
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

func GoogleGkeHubMembershipBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubMembershipBinding), &result)
	return &result
}
