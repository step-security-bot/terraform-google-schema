package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerDatabase = `{
  "block": {
    "attributes": {
      "database_dialect": {
        "computed": true,
        "description": "The dialect of the Cloud Spanner Database.\nIf it is not provided, \"GOOGLE_STANDARD_SQL\" will be used. Possible values: [\"GOOGLE_STANDARD_SQL\", \"POSTGRESQL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
        "description": "Whether or not to allow Terraform to destroy the database. Defaults to true. Unless this field is set to false\nin Terraform state, a 'terraform destroy' or 'terraform apply' that would delete the database will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_drop_protection": {
        "description": "Whether drop protection is enabled for this database. Defaults to false.\nDrop protection is different from\nthe \"deletion_protection\" attribute in the following ways:\n(1) \"deletion_protection\" only protects the database from deletions in Terraform.\nwhereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.\n(2) Setting \"enableDropProtection\" to true also prevents the deletion of the parent instance containing the database.\n\"deletion_protection\" attribute does not provide protection against the deletion of the parent instance.",
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
      },
      "version_retention_period": {
        "computed": true,
        "description": "The retention period for the database. The retention period must be between 1 hour\nand 7 days, and can be specified in days, hours, minutes, or seconds. For example,\nthe values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.\nIf this property is used, you must avoid adding new DDL statements to 'ddl' that\nupdate the database's version_retention_period.",
        "description_kind": "plain",
        "optional": true,
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
