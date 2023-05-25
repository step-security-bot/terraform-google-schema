package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlSourceRepresentationInstance = `{
  "block": {
    "attributes": {
      "database_version": {
        "description": "The MySQL version running on your source database server. Possible values: [\"MYSQL_5_5\", \"MYSQL_5_6\", \"MYSQL_5_7\", \"MYSQL_8_0\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host": {
        "description": "The externally accessible IPv4 address for the source database server.",
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
      "name": {
        "description": "The name of the source representation instance. Use any valid Cloud SQL instance name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description": "The externally accessible port for the source database server.\nDefaults to 3306.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created instance should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSqlSourceRepresentationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlSourceRepresentationInstance), &result)
	return &result
}
