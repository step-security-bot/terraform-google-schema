package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjects = `{
  "block": {
    "attributes": {
      "filter": {
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
      "projects": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "labels": [
                "map",
                "string"
              ],
              "lifecycle_state": "string",
              "name": "string",
              "number": "string",
              "parent": [
                "map",
                "string"
              ],
              "project_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleProjectsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjects), &result)
	return &result
}
