package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreDocument = `{
  "block": {
    "attributes": {
      "collection": {
        "description": "The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "database": {
        "description": "The Firestore database id. Defaults to '\"(default)\"'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "document_id": {
        "description": "The client-assigned document ID to use for this document during creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fields": {
        "description": "The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.",
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
        "computed": true,
        "description": "A server defined name for this index. Format:\n'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "A relative path to the collection this document exists within",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Last update timestamp in RFC3339 format.",
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

func GoogleFirestoreDocumentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreDocument), &result)
	return &result
}
