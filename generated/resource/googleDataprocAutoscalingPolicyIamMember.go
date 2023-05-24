package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocAutoscalingPolicyIamMember = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "member": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
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
      "role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "condition": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "title": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataprocAutoscalingPolicyIamMemberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocAutoscalingPolicyIamMember), &result)
	return &result
}
