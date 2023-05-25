package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreIndex = `{
  "block": {
    "attributes": {
      "collection": {
        "description": "The collection being indexed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database": {
        "description": "The Firestore database id. Defaults to '\"(default)\"'.",
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
      "name": {
        "computed": true,
        "description": "A server defined name for this index. Format:\n'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_scope": {
        "description": "The scope at which a query is run. Default value: \"COLLECTION\" Possible values: [\"COLLECTION\", \"COLLECTION_GROUP\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "fields": {
        "block": {
          "attributes": {
            "array_config": {
              "description": "Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can\nbe specified. Possible values: [\"CONTAINS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_path": {
              "description": "Name of the field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order": {
              "description": "Indicates that this field supports ordering by the specified order or comparing using =, \u003c, \u003c=, \u003e, \u003e=.\nOnly one of 'order' and 'arrayConfig' can be specified. Possible values: [\"ASCENDING\", \"DESCENDING\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The fields supported by this index. The last field entry is always for\nthe field path '__name__'. If, on creation, '__name__' was not\nspecified as the last field, it will be added automatically with the\nsame direction as that of the last field defined. If the final field\nin a composite index is not directional, the '__name__' will be\nordered '\"ASCENDING\"' (unless explicitly specified otherwise).",
          "description_kind": "plain"
        },
        "min_items": 2,
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

func GoogleFirestoreIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreIndex), &result)
	return &result
}
