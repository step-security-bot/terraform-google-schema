package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreField = `{
  "block": {
    "attributes": {
      "collection": {
        "description": "The id of the collection group to configure.",
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
      "field": {
        "description": "The id of the field to configure.",
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
        "description": "The name of this field. Format:\n'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/fields/{{field}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "index_config": {
        "block": {
          "block_types": {
            "indexes": {
              "block": {
                "attributes": {
                  "array_config": {
                    "description": "Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can\nbe specified. Possible values: [\"CONTAINS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "order": {
                    "description": "Indicates that this field supports ordering by the specified order or comparing using =, \u003c, \u003c=, \u003e, \u003e=, !=.\nOnly one of 'order' and 'arrayConfig' can be specified. Possible values: [\"ASCENDING\", \"DESCENDING\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query_scope": {
                    "description": "The scope at which a query is run. Collection scoped queries require you specify\nthe collection at query time. Collection group scope allows queries across all\ncollections with the same id. Default value: \"COLLECTION\" Possible values: [\"COLLECTION\", \"COLLECTION_GROUP\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The indexes to configure on the field. Order or array contains must be specified.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The single field index configuration for this field.\nCreating an index configuration for this field will override any inherited configuration with the\nindexes specified. Configuring the index configuration with an empty block disables all indexes on\nthe field.",
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
      },
      "ttl_config": {
        "block": {
          "attributes": {
            "state": {
              "computed": true,
              "description": "The state of the TTL configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "If set, this field is configured for TTL deletion.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleFirestoreFieldSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreField), &result)
	return &result
}
