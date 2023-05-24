package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageTransferAgentPool = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Specifies the client-specified AgentPool description.",
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
      "name": {
        "description": "The ID of the agent pool to create.\n\nThe agentPoolId must meet the following requirements:\n* Length of 128 characters or less.\n* Not start with the string goog.\n* Start with a lowercase ASCII character, followed by:\n  * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).\n  * One or more numerals or lowercase ASCII characters.\n\nAs expressed by the regular expression: ^(?!goog)[a-z]([a-z0-9-._~]*[a-z0-9])?$.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Specifies the state of the AgentPool.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "bandwidth_limit": {
        "block": {
          "attributes": {
            "limit_mbps": {
              "description": "Bandwidth rate in megabytes per second, distributed across all the agents in the pool.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.",
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

func GoogleStorageTransferAgentPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageTransferAgentPool), &result)
	return &result
}
