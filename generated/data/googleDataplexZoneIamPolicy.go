package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexZoneIamPolicy = `{
  "block": {
    "attributes": {
      "dataplex_zone": {
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
      "lake": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataplexZoneIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexZoneIamPolicy), &result)
	return &result
}
