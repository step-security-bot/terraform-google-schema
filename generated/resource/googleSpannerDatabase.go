package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerDatabase = `{
  "block": {
    "attributes": {
      "ddl": {
        "description": "An optional list of DDL statements to run inside the newly created\ndatabase. Statements can create tables, indexes, etc. These statements\nexecute atomically with the creation of the database: if there is an\nerror in any statement, the database is not created.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The instance to create the database on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A unique identifier for the database, which cannot be changed after\nthe instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].",
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
      "state": {
        "computed": true,
        "description": "An explanation of the status of the database.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Fully qualified name of the KMS key to use to encrypt this database. This key must exist\nin the same location as the Spanner Database.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Encryption configuration for the database",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func GoogleSpannerDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerDatabase), &result)
	return &result
}
