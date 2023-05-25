package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatastoreIndex = `{
  "block": {
    "attributes": {
      "ancestor": {
        "description": "Policy for including ancestors in the index. Default value: \"NONE\" Possible values: [\"NONE\", \"ALL_ANCESTORS\"]",
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
      "index_id": {
        "computed": true,
        "description": "The index id.",
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "description": "The entity kind which the index applies to.",
        "description_kind": "plain",
        "required": true,
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
      "properties": {
        "block": {
          "attributes": {
            "direction": {
              "description": "The direction the index should optimize for sorting. Possible values: [\"ASCENDING\", \"DESCENDING\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "The property name to index.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An ordered list of properties to index on.",
          "description_kind": "plain"
        },
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

func GoogleDatastoreIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatastoreIndex), &result)
	return &result
}
