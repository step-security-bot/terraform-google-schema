package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAzureVersions = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supported_regions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "valid_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleContainerAzureVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAzureVersions), &result)
	return &result
}
