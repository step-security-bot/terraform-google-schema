package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFolders = `{
  "block": {
    "attributes": {
      "folders": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "delete_time": "string",
              "display_name": "string",
              "etag": "string",
              "name": "string",
              "parent": "string",
              "state": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleFoldersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFolders), &result)
	return &result
}
