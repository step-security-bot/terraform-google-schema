package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeFlowhook = `{
  "block": {
    "attributes": {
      "continue_on_error": {
        "description": "Flag that specifies whether execution should continue if the flow hook throws an exception. Set to true to continue execution. Set to false to stop execution if the flow hook throws an exception. Defaults to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "Description of the flow hook.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment": {
        "description": "The resource ID of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "flow_hook_point": {
        "description": "Where in the API call flow the flow hook is invoked. Must be one of PreProxyFlowHook, PostProxyFlowHook, PreTargetFlowHook, or PostTargetFlowHook.",
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
      "org_id": {
        "description": "The Apigee Organization associated with the environment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sharedflow": {
        "description": "Id of the Sharedflow attaching to a flowhook point.",
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

func GoogleApigeeFlowhookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeFlowhook), &result)
	return &result
}
