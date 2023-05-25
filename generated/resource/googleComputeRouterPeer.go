package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterPeer = `{
  "block": {
    "attributes": {
      "advertise_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "advertised_groups": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "advertised_route_priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interface": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_asn": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "peer_ip_address": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "advertised_ip_ranges": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRouterPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterPeer), &result)
	return &result
}
