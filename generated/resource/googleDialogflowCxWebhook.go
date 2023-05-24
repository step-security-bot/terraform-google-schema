package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxWebhook = `{
  "block": {
    "attributes": {
      "disabled": {
        "description": "Indicates whether the webhook is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The human-readable name of the webhook, unique within the agent.",
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
      "name": {
        "computed": true,
        "description": "The unique identifier of the webhook.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a webhook for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
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
      "timeout": {
        "description": "Webhook execution timeout.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "generic_web_service": {
        "block": {
          "attributes": {
            "allowed_ca_certs": {
              "description": "Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "request_headers": {
              "description": "The HTTP request headers to send together with webhook requests.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "uri": {
              "description": "Whether to use speech adaptation for speech recognition.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for a generic web service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_directory": {
        "block": {
          "attributes": {
            "service": {
              "description": "The name of Service Directory service.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "generic_web_service": {
              "block": {
                "attributes": {
                  "allowed_ca_certs": {
                    "description": "Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "request_headers": {
                    "description": "The HTTP request headers to send together with webhook requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "uri": {
                    "description": "Whether to use speech adaptation for speech recognition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The name of Service Directory service.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for a Service Directory service.",
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

func GoogleDialogflowCxWebhookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxWebhook), &result)
	return &result
}
