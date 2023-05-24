package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxEntityType = `{
  "block": {
    "attributes": {
      "auto_expansion_mode": {
        "description": "Represents kinds of entities.\n* AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.\n* AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity. Possible values: [\"AUTO_EXPANSION_MODE_DEFAULT\", \"AUTO_EXPANSION_MODE_UNSPECIFIED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the entity type, unique within the agent.",
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
        "description": "Indicates whether the entity type can be automatically expanded.\n* KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.\n* KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).\n* KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: [\"KIND_MAP\", \"KIND_LIST\", \"KIND_REGEXP\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language_code": {
        "description": "The language of the following fields in entityType:\nEntityType.entities.value\nEntityType.entities.synonyms\nEntityType.excluded_phrases.value\nIf not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the entity type.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/entityTypes/\u003cEntity Type ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a entity type for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redact": {
        "description": "Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "entities": {
        "block": {
          "attributes": {
            "synonyms": {
              "description": "A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.\nFor KIND_LIST entity types: This collection must contain exactly one synonym equal to value.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "value": {
              "description": "The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.\nFor KIND_MAP entity types: A canonical value to be used in place of synonyms.\nFor KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The collection of entity entries associated with the entity type.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "excluded_phrases": {
        "block": {
          "attributes": {
            "value": {
              "description": "The word or phrase to be excluded.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.\nIf the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.",
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

func GoogleDialogflowCxEntityTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxEntityType), &result)
	return &result
}
