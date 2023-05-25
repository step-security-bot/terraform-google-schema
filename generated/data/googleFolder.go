package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFolder = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder": {
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
      "lifecycle_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lookup_organization": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFolder), &result)
	return &result
}
