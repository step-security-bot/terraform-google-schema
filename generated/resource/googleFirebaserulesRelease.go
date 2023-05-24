package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaserulesRelease = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Time the release was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description": "Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Format: ` + "`" + `projects/{project_id}/releases/{release_id}` + "`" + `\\Firestore Rules Releases will **always** have the name 'cloud.firestore'",
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
      "ruleset_name": {
        "description": "Name of the ` + "`" + `Ruleset` + "`" + ` referred to by this ` + "`" + `Release` + "`" + `. The ` + "`" + `Ruleset` + "`" + ` must exist for the ` + "`" + `Release` + "`" + ` to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Time the release was updated.",
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

func GoogleFirebaserulesReleaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaserulesRelease), &result)
	return &result
}
