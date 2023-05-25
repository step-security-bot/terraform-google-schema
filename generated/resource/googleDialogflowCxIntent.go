package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxIntent = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the intent, unique within the agent.",
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
      "is_fallback": {
        "description": "Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. \nAdding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "labels": {
        "description": "The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.\nPrefix \"sys-\" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. \"sys-head\" means the intent is a head intent. \"sys.contextual\" means the intent is a contextual intent.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "language_code": {
        "description": "The language of the following fields in intent:\nIntent.training_phrases.parts.text\nIf not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the intent.  \nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/intents/\u003cIntent ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create an intent for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "The priority of this intent. Higher numbers represent higher priorities.\nIf the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.\nIf the supplied value is negative, the intent is ignored in runtime detect intent requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "parameters": {
        "block": {
          "attributes": {
            "entity_type": {
              "description": "The entity type of the parameter. \nFormat: projects/-/locations/-/agents/-/entityTypes/\u003cSystem Entity Type ID\u003e for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/entityTypes/\u003cEntity Type ID\u003e for developer entity types.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "description": "The unique identifier of the parameter. This field is used by training phrases to annotate their parts.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "is_list": {
              "description": "Indicates whether the parameter represents a list of values.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "redact": {
              "description": "Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging. \nNote: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The collection of parameters associated with the intent.",
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
      },
      "training_phrases": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The unique identifier of the training phrase.",
              "description_kind": "plain",
              "type": "string"
            },
            "repeat_count": {
              "description": "Indicates how many times this example was added to the intent.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "parts": {
              "block": {
                "attributes": {
                  "parameter_id": {
                    "description": "The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "text": {
                    "description": "The text for this part.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The ordered list of training phrase parts. The parts are concatenated in order to form the training phrase.\nNote: The API does not automatically annotate training phrases like the Dialogflow Console does.\nNote: Do not forget to include whitespace at part boundaries, so the training phrase is well formatted when the parts are concatenated.\nIf the training phrase does not need to be annotated with parameters, you just need a single part with only the Part.text field set.\nIf you want to annotate the training phrase, you must create multiple parts, where the fields of each part are populated in one of two ways:\nPart.text is set to a part of the phrase that has no parameters.\nPart.text is set to a part of the phrase that you want to annotate, and the parameterId field is set.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The collection of training phrases the agent is trained on to identify the intent.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowCxIntentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxIntent), &result)
	return &result
}
