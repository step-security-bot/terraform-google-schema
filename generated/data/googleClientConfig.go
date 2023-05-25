package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClientConfig = `{
  "block": {
    "attributes": {
      "access_token": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleClientConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClientConfig), &result)
	return &result
}
