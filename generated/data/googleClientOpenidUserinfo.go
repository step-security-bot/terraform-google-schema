package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClientOpenidUserinfo = `{
  "block": {
    "attributes": {
      "email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleClientOpenidUserinfoSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClientOpenidUserinfo), &result)
	return &result
}
