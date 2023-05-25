package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketAccessControl = `{
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
        "description": "The entity holding the permission, in one of the following forms:\n  user-userId\n  user-email\n  group-groupId\n  group-email\n  domain-domain\n  project-team-projectId\n  allUsers\n  allAuthenticatedUsers\nExamples:\n  The user liz@example.com would be user-liz@example.com.\n  The group example@googlegroups.com would be\n  group-example@googlegroups.com.\n  To refer to all members of the Google Apps for Business domain\n  example.com, the entity would be domain-example.com.",
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
      "role": {
        "description": "The access permission for the entity.",
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

func GoogleStorageBucketAccessControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucketAccessControl), &result)
	return &result
}
