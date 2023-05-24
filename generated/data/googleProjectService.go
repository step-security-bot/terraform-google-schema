package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectService = `{
  "block": {
    "attributes": {
      "disable_dependent_services": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "disable_on_destroy": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleProjectServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectService), &result)
	return &result
}
