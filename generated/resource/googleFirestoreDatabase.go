package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreDatabase = `{
  "block": {
    "attributes": {
      "app_engine_integration_mode": {
        "computed": true,
        "description": "The App Engine integration mode to use for this database. Possible values: [\"ENABLED\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "concurrency_mode": {
        "computed": true,
        "description": "The concurrency control mode to use for this database. Possible values: [\"OPTIMISTIC\", \"PESSIMISTIC\", \"OPTIMISTIC_WITH_ENTITY_GROUPS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this database was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of other fields,\nand may be sent on update and delete requests to ensure the client has an\nup-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_prefix": {
        "computed": true,
        "description": "Output only. The keyPrefix for this database.\nThis keyPrefix is used, in combination with the project id (\"~\") to construct the application id\nthat is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.\nThis value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).",
        "description_kind": "plain",
        "type": "string"
      },
      "location_id": {
        "description": "The location of the database. Available databases are listed at\nhttps://cloud.google.com/firestore/docs/locations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID to use for the database, which will become the final\ncomponent of the database's resource name. This value should be 4-63\ncharacters. Valid characters are /[a-z][0-9]-/ with first character\na letter and the last a letter or a number. Must not be\nUUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.\n\"(default)\" database id is also valid.",
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
      "type": {
        "description": "The type of the database.\nSee https://cloud.google.com/datastore/docs/firestore-or-datastore\nfor information about how to choose. Possible values: [\"FIRESTORE_NATIVE\", \"DATASTORE_MODE\"]",
        "description_kind": "plain",
        "required": true,
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

func GoogleFirestoreDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreDatabase), &result)
	return &result
}
