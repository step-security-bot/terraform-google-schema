package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlTiers = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "Project ID of the project for which to list tiers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tiers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_quota": "number",
              "ram": "number",
              "region": [
                "list",
                "string"
              ],
              "tier": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSqlTiersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlTiers), &result)
	return &result
}
