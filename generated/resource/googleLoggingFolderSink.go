package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingFolderSink = `{
  "block": {
    "attributes": {
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter": {
        "description_kind": "plain",
        "optional": true,
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
      "include_children": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "writer_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingFolderSinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingFolderSink), &result)
	return &result
}
