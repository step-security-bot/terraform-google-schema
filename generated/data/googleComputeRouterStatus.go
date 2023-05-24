package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterStatus = `{
  "block": {
    "attributes": {
      "best_routes": {
        "computed": true,
        "description": "Best routes for this router's network.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "dest_range": "string",
              "name": "string",
              "network": "string",
              "next_hop_gateway": "string",
              "next_hop_ilb": "string",
              "next_hop_instance": "string",
              "next_hop_instance_zone": "string",
              "next_hop_ip": "string",
              "next_hop_network": "string",
              "next_hop_vpn_tunnel": "string",
              "priority": "number",
              "project": "string",
              "self_link": "string",
              "tags": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "best_routes_for_router": {
        "computed": true,
        "description": "Best routes learned by this router.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "dest_range": "string",
              "name": "string",
              "network": "string",
              "next_hop_gateway": "string",
              "next_hop_ilb": "string",
              "next_hop_instance": "string",
              "next_hop_instance_zone": "string",
              "next_hop_ip": "string",
              "next_hop_network": "string",
              "next_hop_vpn_tunnel": "string",
              "priority": "number",
              "project": "string",
              "self_link": "string",
              "tags": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the router to query.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "URI of the network to which this router belongs.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "Project ID of the target router.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region of the target router.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRouterStatusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterStatus), &result)
	return &result
}
