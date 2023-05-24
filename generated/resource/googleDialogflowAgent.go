package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowAgent = `{
  "block": {
    "attributes": {
      "api_version": {
        "computed": true,
        "description": "API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query\ndifferent service endpoints for different API versions. However, bots connectors and webhook calls will follow\nthe specified API version.\n* API_VERSION_V1: Legacy V1 API.\n* API_VERSION_V2: V2 API.\n* API_VERSION_V2_BETA_1: V2beta1 API. Possible values: [\"API_VERSION_V1\", \"API_VERSION_V2\", \"API_VERSION_V2_BETA_1\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "avatar_uri": {
        "description": "The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered\ninto this field, the Dialogflow will save the image in the backend. The address of the backend image returned\nfrom the API will be shown in the [avatarUriBackend] field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "avatar_uri_backend": {
        "computed": true,
        "description": "The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar,\nthe [avatarUri] field can be used.",
        "description_kind": "plain",
        "type": "string"
      },
      "classification_threshold": {
        "description": "To filter out false positive results and still get variety in matched natural language inputs for your agent,\nyou can tune the machine learning classification threshold. If the returned score value is less than the threshold\nvalue, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be\ntriggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the\ndefault of 0.3 is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_language_code": {
        "description": "The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)\nfor a list of the currently supported language codes. This field cannot be updated after creation.",
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
        "description": "The name of this agent.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_logging": {
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
      "match_mode": {
        "computed": true,
        "description": "Determines how intents are detected from user queries.\n* MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates\nsyntax and composite entities.\n* MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones\nusing @sys.any or very large developer entities. Possible values: [\"MATCH_MODE_HYBRID\", \"MATCH_MODE_ML_ONLY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supported_language_codes": {
        "description": "The list of all languages supported by this agent (except for the defaultLanguageCode).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tier": {
        "description": "The agent tier. If not specified, TIER_STANDARD is assumed.\n* TIER_STANDARD: Standard tier.\n* TIER_ENTERPRISE: Enterprise tier (Essentials).\n* TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).\nNOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between\nthe Terraform state and Dialogflow if the agent tier is changed outside of Terraform. Possible values: [\"TIER_STANDARD\", \"TIER_ENTERPRISE\", \"TIER_ENTERPRISE_PLUS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "description": "The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,\nEurope/Paris.",
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

func GoogleDialogflowAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowAgent), &result)
	return &result
}
