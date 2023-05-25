package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNodeGroup = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional textual description of the resource.",
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
      "initial_size": {
        "description": "The initial number of nodes in the node group. One of 'initial_size' or 'size' must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maintenance_policy": {
        "description": "Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_template": {
        "description": "The URL of the node template to which this node group belongs.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "The total number of nodes in the node group. One of 'initial_size' or 'size' must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "zone": {
        "computed": true,
        "description": "Zone where this node group is located",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling_policy": {
        "block": {
          "attributes": {
            "max_nodes": {
              "computed": true,
              "description": "Maximum size of the node group. Set to a value less than or equal\nto 100 and greater than or equal to min-nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_nodes": {
              "computed": true,
              "description": "Minimum size of the node group. Must be less\nthan or equal to max-nodes. The default value is 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
              "computed": true,
              "description": "The autoscaling mode. Set to one of the following:\n  - OFF: Disables the autoscaler.\n  - ON: Enables scaling in and scaling out.\n  - ONLY_SCALE_OUT: Enables only scaling out.\n  You must use this mode if your node groups are configured to\n  restart their hosted VMs on minimal servers. Possible values: [\"OFF\", \"ON\", \"ONLY_SCALE_OUT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "If you use sole-tenant nodes for your workloads, you can use the node\ngroup autoscaler to automatically manage the sizes of your node groups.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_window": {
        "block": {
          "attributes": {
            "start_time": {
              "description": "instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "contains properties for the timeframe of maintenance",
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

func GoogleComputeNodeGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNodeGroup), &result)
	return &result
}
