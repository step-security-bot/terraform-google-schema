package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowFulfillment = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The human-readable name of the fulfillment, unique within the agent.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description": "Whether fulfillment is enabled.",
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
        "description": "The unique identifier of the fulfillment. \nFormat: projects/\u003cProject ID\u003e/agent/fulfillment - projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent/fulfillment",
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
      "features": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of the feature that enabled for fulfillment.\n* SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: [\"SMALLTALK\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The field defines whether the fulfillment is enabled for certain features.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "generic_web_service": {
        "block": {
          "attributes": {
            "password": {
              "description": "The password for HTTP Basic authentication.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_headers": {
              "description": "The HTTP request headers to send together with fulfillment requests.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "uri": {
              "description": "The fulfillment URI for receiving POST requests. It must use https protocol.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
              "description": "The user name for HTTP Basic authentication.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.",
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

func GoogleDialogflowFulfillmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowFulfillment), &result)
	return &result
}
