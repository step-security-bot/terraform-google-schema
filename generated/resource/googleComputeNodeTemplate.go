package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNodeTemplate = `{
  "block": {
    "attributes": {
      "cpu_overcommit_type": {
        "description": "CPU overcommit. Default value: \"NONE\" Possible values: [\"ENABLED\", \"NONE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "name": {
        "description": "Name of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_affinity_labels": {
        "description": "Labels to use for node affinity, which will be used in\ninstance scheduling.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "node_type": {
        "description": "Node type to use for nodes group that are created from this template.\nOnly one of nodeTypeFlexibility and nodeType can be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where nodes using the node template will be created.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "node_type_flexibility": {
        "block": {
          "attributes": {
            "cpus": {
              "description": "Number of virtual CPUs to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_ssd": {
              "computed": true,
              "description": "Use local SSD",
              "description_kind": "plain",
              "type": "string"
            },
            "memory": {
              "description": "Physical memory available to the node, defined in MB.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Flexible properties for the desired node type. Node groups that\nuse this node template will create nodes of a type that matches\nthese properties. Only one of nodeTypeFlexibility and nodeType can\nbe specified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "server_binding": {
        "block": {
          "attributes": {
            "type": {
              "description": "Type of server binding policy. If 'RESTART_NODE_ON_ANY_SERVER',\nnodes using this template will restart on any physical server\nfollowing a maintenance event.\n\nIf 'RESTART_NODE_ON_MINIMAL_SERVER', nodes using this template\nwill restart on the same physical server following a maintenance\nevent, instead of being live migrated to or restarted on a new\nphysical server. This option may be useful if you are using\nsoftware licenses tied to the underlying server characteristics\nsuch as physical sockets or cores, to avoid the need for\nadditional licenses when maintenance occurs. However, VMs on such\nnodes will experience outages while maintenance is applied. Possible values: [\"RESTART_NODE_ON_ANY_SERVER\", \"RESTART_NODE_ON_MINIMAL_SERVERS\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The server binding policy for nodes using this template. Determines\nwhere the nodes should restart following a maintenance event.",
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

func GoogleComputeNodeTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNodeTemplate), &result)
	return &result
}
