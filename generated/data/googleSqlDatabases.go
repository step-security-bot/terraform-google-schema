package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabases = `{
  "block": {
    "attributes": {
      "databases": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "charset": "string",
              "collation": "string",
              "deletion_policy": "string",
              "instance": "string",
              "name": "string",
              "project": "string",
              "self_link": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The name of the Cloud SQL database instance in which the database belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "Project ID of the project that contains the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSqlDatabasesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabases), &result)
	return &result
}
