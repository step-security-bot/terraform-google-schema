package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkEndpoints = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_endpoint_group": {
        "description": "The network endpoint group these endpoints are part of.",
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
      "zone": {
        "computed": true,
        "description": "Zone where the containing network endpoint group is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "network_endpoints": {
        "block": {
          "attributes": {
            "instance": {
              "description": "The name for a specific VM instance that the IP address belongs to.\nThis is required for network endpoints of type GCE_VM_IP_PORT.\nThe instance must be in the same zone as the network endpoint group.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "description": "IPv4 address of network endpoint. The IP address must belong\nto a VM in GCE (either the primary IP or as part of an aliased IP\nrange).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "Port number of network endpoint.\n**Note** 'port' is required unless the Network Endpoint Group is created\nwith the type of 'GCE_VM_IP'",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "The network endpoints to be added to the enclosing network endpoint group\n(NEG). Each endpoint specifies an IP address and port, along with\nadditional information depending on the NEG type.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleComputeNetworkEndpointsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkEndpoints), &result)
	return &result
}
