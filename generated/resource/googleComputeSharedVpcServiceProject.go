package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSharedVpcServiceProject = `{
  "block": {
    "attributes": {
      "deletion_policy": {
        "description": "The deletion policy for the shared VPC service. Setting ABANDON allows the resource\n\t\t\t\tto be abandoned rather than deleted. Possible values are: \"ABANDON\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_project": {
        "description": "The ID of a host project to associate.",
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
      "service_project": {
        "description": "The ID of the project that will serve as a Shared VPC service project.",
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

func GoogleComputeSharedVpcServiceProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSharedVpcServiceProject), &result)
	return &result
}
