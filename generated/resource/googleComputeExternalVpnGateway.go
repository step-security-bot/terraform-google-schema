package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeExternalVpnGateway = `{
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
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels for the external VPN gateway resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "redundancy_type": {
        "description": "Indicates the redundancy type of this external VPN gateway Possible values: [\"FOUR_IPS_REDUNDANCY\", \"SINGLE_IP_INTERNALLY_REDUNDANT\", \"TWO_IPS_REDUNDANCY\"]",
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
      "interface": {
        "block": {
          "attributes": {
            "id": {
              "description": "The numeric ID for this interface. Allowed values are based on the redundancy type\nof this external VPN gateway\n* '0 - SINGLE_IP_INTERNALLY_REDUNDANT'\n* '0, 1 - TWO_IPS_REDUNDANCY'\n* '0, 1, 2, 3 - FOUR_IPS_REDUNDANCY'",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ip_address": {
              "description": "IP address of the interface in the external VPN gateway.\nOnly IPv4 is supported. This IP address can be either from\nyour on-premise gateway or another Cloud provider's VPN gateway,\nit cannot be an IP address from Google Compute Engine.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A list of interfaces on this external VPN gateway.",
          "description_kind": "plain"
        },
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

func GoogleComputeExternalVpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeExternalVpnGateway), &result)
	return &result
}
