package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRoute = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource. Provide this property\nwhen you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_range": {
        "description": "The destination range of outgoing packets that this route applies to.\nOnly IPv4 is supported.",
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network that this route applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_hop_gateway": {
        "description": "URL to a gateway that should handle matching packets.\nCurrently, you can only specify the internet gateway, using a full or\npartial valid URL:\n* 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'\n* 'projects/project/global/gateways/default-internet-gateway'\n* 'global/gateways/default-internet-gateway'\n* The string 'default-internet-gateway'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_instance": {
        "description": "URL to an instance that should handle matching packets.\nYou can specify this as a full or partial URL. For example:\n* 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'\n* 'projects/project/zones/zone/instances/instance'\n* 'zones/zone/instances/instance'\n* Just the instance name, with the zone in 'next_hop_instance_zone'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_instance_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_ip": {
        "computed": true,
        "description": "Network IP address of an instance that should handle matching packets.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_network": {
        "computed": true,
        "description": "URL to a Network that should handle matching packets.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_vpn_tunnel": {
        "description": "URL to a VpnTunnel that should handle matching packets.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "The priority of this route. Priority is used to break ties in cases\nwhere there is more than one matching route of equal prefix length.\n\nIn the case of two routes with equal prefix length, the one with the\nlowest-numbered priority value wins.\n\nDefault value is 1000. Valid range is 0 through 65535.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "tags": {
        "description": "A list of instance tags to which this route applies.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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

func GoogleComputeRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRoute), &result)
	return &result
}
