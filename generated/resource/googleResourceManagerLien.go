package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleResourceManagerLien = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time of creation",
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
        "description": "A system-generated unique identifier for this Lien.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin": {
        "description": "A stable, user-visible/meaningful string identifying the origin\nof the Lien, intended to be inspected programmatically. Maximum length of\n200 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "A reference to the resource this Lien is attached to.\nThe server will validate the parent against those for which Liens are supported.\nSince a variety of objects can have Liens against them, you must provide the type\nprefix (e.g. \"projects/my-project-name\").",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reason": {
        "description": "Concise user-visible strings indicating why an action cannot be performed\non a resource. Maximum length of 200 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restrictions": {
        "description": "The types of operations which should be blocked as a result of this Lien.\nEach value should correspond to an IAM permission. The server will validate\nthe permissions against those for which Liens are supported.  An empty\nlist is meaningless and will be rejected.\ne.g. ['resourcemanager.projects.delete']",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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

func GoogleResourceManagerLienSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleResourceManagerLien), &result)
	return &result
}
