package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabase = `{
  "block": {
    "attributes": {
      "charset": {
        "computed": true,
        "description": "The charset value. See MySQL's\n[Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)\nand Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)\nfor more details and supported values. Postgres databases only support\na value of 'UTF8' at creation time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "collation": {
        "computed": true,
        "description": "The collation value. See MySQL's\n[Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)\nand Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)\nfor more details and supported values. Postgres databases only support\na value of 'en_US.UTF8' at creation time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The name of the Cloud SQL instance. This does not include the project\nID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the database in the Cloud SQL instance.\nThis does not include the project ID or instance name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
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
            },
            "update": {
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

func GoogleSqlDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabase), &result)
	return &result
}
