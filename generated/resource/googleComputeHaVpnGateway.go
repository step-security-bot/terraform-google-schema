package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeHaVpnGateway = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource.",
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network this VPN gateway is accepting traffic for.",
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
      "region": {
        "computed": true,
        "description": "The region this gateway should sit in.",
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
      },
      "vpn_interfaces": {
        "block": {
          "attributes": {
            "id": {
              "description": "The numeric ID of this VPN gateway interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interconnect_attachment": {
              "description": "URL of the interconnect attachment resource. When the value\nof this field is present, the VPN Gateway will be used for\nIPsec-encrypted Cloud Interconnect; all Egress or Ingress\ntraffic for this VPN Gateway interface will go through the\nspecified interconnect attachment resource.\n\nNot currently available publicly.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "computed": true,
              "description": "The external IP address for this VPN gateway interface.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "A list of interfaces on this VPN gateway.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeHaVpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeHaVpnGateway), &result)
	return &result
}
