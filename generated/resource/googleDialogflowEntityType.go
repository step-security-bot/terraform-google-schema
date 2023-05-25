package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowEntityType = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The name of this entity type to be displayed on the console.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_fuzzy_extraction": {
        "description": "Enables fuzzy entity extraction during classification.",
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
      "kind": {
        "description": "Indicates the kind of entity type.\n* KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.\n* KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity\ntypes can contain references to other entity types (with or without aliases).\n* KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: [\"KIND_MAP\", \"KIND_LIST\", \"KIND_REGEXP\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the entity type. \nFormat: projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity type ID\u003e.",
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
      "entities": {
        "block": {
          "attributes": {
            "synonyms": {
              "description": "A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym\ncould be green onions.\nFor KIND_LIST entity types:\n* This collection must contain exactly one synonym equal to value.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "value": {
              "description": "The primary value associated with this entity entry. For example, if the entity type is vegetable, the value\ncould be scallions.\nFor KIND_MAP entity types:\n* A reference value to be used in place of synonyms.\nFor KIND_LIST entity types:\n* A string that can contain references to other entity types (with or without aliases).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The collection of entity entries associated with the entity type.",
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

func GoogleDialogflowEntityTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowEntityType), &result)
	return &result
}
