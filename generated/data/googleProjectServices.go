package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectServices = `{
  "block": {
    "attributes": {
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
      "services": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleProjectServicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectServices), &result)
	return &result
}
