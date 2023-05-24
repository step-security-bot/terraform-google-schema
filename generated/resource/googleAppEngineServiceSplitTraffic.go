package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineServiceSplitTraffic = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migrate_traffic": {
        "description": "If set to true traffic will be migrated to this version.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "The name of the service these settings apply to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "split": {
        "block": {
          "attributes": {
            "allocations": {
              "description": "Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "shard_by": {
              "description": "Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed. Possible values: [\"UNSPECIFIED\", \"COOKIE\", \"IP\", \"RANDOM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Mapping that defines fractional HTTP traffic diversion to different versions within the service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleAppEngineServiceSplitTrafficSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineServiceSplitTraffic), &result)
	return &result
}
