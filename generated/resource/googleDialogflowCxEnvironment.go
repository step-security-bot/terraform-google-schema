package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxEnvironment = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the environment (unique in an agent). Limit of 64 characters.",
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
      "name": {
        "computed": true,
        "description": "The name of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The Agent to create an Environment for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Update time of this environment. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
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
      },
      "version_configs": {
        "block": {
          "attributes": {
            "version": {
              "description": "Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowCxEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxEnvironment), &result)
	return &result
}
