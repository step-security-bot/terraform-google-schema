package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleRuntimeconfigVariable = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the variable to manage. Note that variable names can be hierarchical using slashes (e.g. \"prod-variables/hostname\").",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The name of the RuntimeConfig resource containing this variable.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "text": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds, representing when the variable was last updated. Example: \"2016-10-09T12:33:37.578138407Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleRuntimeconfigVariableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleRuntimeconfigVariable), &result)
	return &result
}
