package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstanceGroupNamedPort = `{
  "block": {
    "attributes": {
      "group": {
        "description": "The name of the instance group.",
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
        "description": "The name for this named port. The name must be 1-63 characters\nlong, and comply with RFC1035.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description": "The port number, which can be a value between 1 and 65535.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "The zone of the instance group.",
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

func GoogleComputeInstanceGroupNamedPortSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstanceGroupNamedPort), &result)
	return &result
}
