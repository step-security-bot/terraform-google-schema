package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbSupportedDatabaseFlags = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The canonical id for the location. For example: \"us-east1\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "Project ID of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supported_database_flags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accepts_multiple_values": "bool",
              "flag_name": "string",
              "integer_restrictions": [
                "list",
                [
                  "object",
                  {
                    "max_value": "string",
                    "min_value": "string"
                  }
                ]
              ],
              "name": "string",
              "requires_db_restart": "bool",
              "string_restrictions": [
                "list",
                [
                  "object",
                  {
                    "allowed_values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "supported_db_versions": [
                "list",
                "string"
              ],
              "value_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAlloydbSupportedDatabaseFlagsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbSupportedDatabaseFlags), &result)
	return &result
}
