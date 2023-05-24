package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbLocations = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "locations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "display_name": "string",
              "labels": [
                "map",
                "string"
              ],
              "location_id": "string",
              "metadata": [
                "map",
                "string"
              ],
              "name": "string"
            }
          ]
        ]
      },
      "project": {
        "description": "Project ID of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAlloydbLocationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbLocations), &result)
	return &result
}
