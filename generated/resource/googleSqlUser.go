package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlUser = `{
  "block": {
    "attributes": {
      "deletion_policy": {
        "description": "The deletion policy for the user. Setting ABANDON allows the resource\n\t\t\t\tto be abandoned rather than deleted. This is useful for Postgres, where users cannot be deleted from the API if they\n\t\t\t\thave been granted SQL roles. Possible values are: \"ABANDON\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "computed": true,
        "description": "The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL instances. Can be an IP address. Changing this forces a new resource to be created.",
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
        "description": "The name of the Cloud SQL instance. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the user. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description": "The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to\n                either CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The user type. It determines the method to authenticate the user during login.\n                The default is the database's built-in user type. Flags include \"BUILT_IN\", \"CLOUD_IAM_USER\", or \"CLOUD_IAM_SERVICE_ACCOUNT\".",
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
  "version": 1
}`

func GoogleSqlUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlUser), &result)
	return &result
}
