package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageObjectAccessControl = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain associated with the entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
        "computed": true,
        "description": "The email address associated with the entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "entity": {
        "description": "The entity holding the permission, in one of the following forms:\n  * user-{{userId}}\n  * user-{{email}} (such as \"user-liz@example.com\")\n  * group-{{groupId}}\n  * group-{{email}} (such as \"group-example@googlegroups.com\")\n  * domain-{{domain}} (such as \"domain-example.com\")\n  * project-team-{{projectId}}\n  * allUsers\n  * allAuthenticatedUsers",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_id": {
        "computed": true,
        "description": "The ID for the entity",
        "description_kind": "plain",
        "type": "string"
      },
      "generation": {
        "computed": true,
        "description": "The content generation of the object, if applied to an object.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object": {
        "description": "The name of the object to apply the access control to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_team": {
        "computed": true,
        "description": "The project team associated with the entity",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "project_number": "string",
              "team": "string"
            }
          ]
        ]
      },
      "role": {
        "description": "The access permission for the entity.",
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

func GoogleStorageObjectAccessControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageObjectAccessControl), &result)
	return &result
}
