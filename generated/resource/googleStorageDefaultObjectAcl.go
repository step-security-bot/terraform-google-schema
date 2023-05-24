package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageDefaultObjectAcl = `{
  "block": {
    "attributes": {
      "bucket": {
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
      "role_entity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func GoogleStorageDefaultObjectAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageDefaultObjectAcl), &result)
	return &result
}
