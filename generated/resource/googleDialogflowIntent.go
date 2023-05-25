package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowIntent = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The name of the action associated with the intent.\nNote: The action name must not contain whitespaces.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_response_platforms": {
        "description": "The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED\n(i.e. default platform). Possible values: [\"FACEBOOK\", \"SLACK\", \"TELEGRAM\", \"KIK\", \"SKYPE\", \"LINE\", \"VIBER\", \"ACTIONS_ON_GOOGLE\", \"GOOGLE_HANGOUTS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "The name of this intent to be displayed on the console.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "events": {
        "description": "The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of\nthe contexts must be present in the active user session for an event to trigger this intent. See the \n[events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "followup_intent_info": {
        "computed": true,
        "description": "Information about all followup intents that have this intent as a direct or indirect parent. We populate this field\nonly in the output.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "followup_intent_name": "string",
              "parent_followup_intent_name": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_context_names": {
        "description": "The list of context names required for this intent to be triggered.\nFormat: projects/\u003cProject ID\u003e/agent/sessions/-/contexts/\u003cContext ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "is_fallback": {
        "computed": true,
        "description": "Indicates whether this is a fallback intent.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ml_disabled": {
        "computed": true,
        "description": "Indicates whether Machine Learning is disabled for the intent.\nNote: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML\nONLY match mode. Also, auto-markup in the UI is turned off.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of this intent. \nFormat: projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_followup_intent_name": {
        "computed": true,
        "description": "The unique identifier of the parent intent in the chain of followup intents.\nFormat: projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "The priority of this intent. Higher numbers represent higher priorities.\n  - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds\n  to the Normal priority in the console.\n  - If the supplied value is negative, the intent is ignored in runtime detect intent requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reset_contexts": {
        "computed": true,
        "description": "Indicates whether to delete all contexts in the current session when this intent is matched.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "root_followup_intent_name": {
        "computed": true,
        "description": "The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup\nintents chain for this intent.\nFormat: projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "webhook_state": {
        "computed": true,
        "description": "Indicates whether webhooks are enabled for the intent.\n* WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.\n* WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot\nfilling prompt is forwarded to the webhook. Possible values: [\"WEBHOOK_STATE_ENABLED\", \"WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING\"]",
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
  "version": 0
}`

func GoogleDialogflowIntentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowIntent), &result)
	return &result
}
