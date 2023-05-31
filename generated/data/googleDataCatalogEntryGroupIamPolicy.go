package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogEntryGroupIamPolicy = `{
  "block": {
    "attributes": {
      "entry_group": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
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

func GoogleDataCatalogEntryGroupIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogEntryGroupIamPolicy), &result)
	return &result
}
