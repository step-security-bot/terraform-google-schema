package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkPeering = `{
  "block": {
    "attributes": {
      "export_custom_routes": {
        "description": "Whether to export the custom routes to the peer network. Defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "export_subnet_routes_with_public_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_custom_routes": {
        "description": "Whether to export the custom routes from the peer network. Defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "import_subnet_routes_with_public_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "Name of the peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The primary network of the peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_network": {
        "description": "The peer network in the peering. The peer network may belong to a different project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State for the peering, either ACTIVE or INACTIVE. The peering is ACTIVE when there's a matching configuration in the peer network.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "Details about the current state of the peering.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeNetworkPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkPeering), &result)
	return &result
}
