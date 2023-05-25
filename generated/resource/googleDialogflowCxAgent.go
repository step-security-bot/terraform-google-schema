package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxAgent = `{
  "block": {
    "attributes": {
      "avatar_uri": {
        "description": "The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_language_code": {
        "description": "The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) \nfor a list of the currently supported language codes. This field cannot be updated after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the agent, unique within the location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_spell_correction": {
        "description": "Indicates if automatic spell correction is enabled in detect intent requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_stackdriver_logging": {
        "description": "Determines whether this agent should log conversation queries.",
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
      "location": {
        "description": "The name of the location this agent is located in.\n\n~\u003e **Note:** The first time you are deploying an Agent in your project you must configure location settings.\n This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.\n Another options is to use global location so you don't need to manually configure location settings.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_settings": {
        "description": "Name of the SecuritySettings reference for the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/securitySettings/\u003cSecurity Settings ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_flow": {
        "computed": true,
        "description": "Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "supported_language_codes": {
        "description": "The list of all languages supported by this agent (except for the default_language_code).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "time_zone": {
        "description": "The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,\nEurope/Paris.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "speech_to_text_settings": {
        "block": {
          "attributes": {
            "enable_speech_adaptation": {
              "description": "Whether to use speech adaptation for speech recognition.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings related to speech recognition.",
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

func GoogleDialogflowCxAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxAgent), &result)
	return &result
}
