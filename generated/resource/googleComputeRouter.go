package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouter = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
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
        "description": "Name of the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'\nwhich means the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "A reference to the network to which this router belongs.",
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
        "description": "Region where the router resides.",
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
      "bgp": {
        "block": {
          "attributes": {
            "advertise_mode": {
              "description": "User-specified flag to indicate which mode to use for advertisement.\n\nValid values of this enum field are: DEFAULT, CUSTOM",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "advertised_groups": {
              "description": "User-specified list of prefix groups to advertise in custom mode.\nThis field can only be populated if advertiseMode is CUSTOM and\nis advertised to all peers of the router. These groups will be\nadvertised in addition to any specified prefixes. Leave this field\nblank to advertise no custom groups.\n\nThis enum field has the one valid value: ALL_SUBNETS",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "asn": {
              "description": "Local BGP Autonomous System Number (ASN). Must be an RFC6996\nprivate ASN, either 16-bit or 32-bit. The value will be fixed for\nthis router resource. All VPN tunnels that link to this router\nwill have the same local ASN.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "advertised_ip_ranges": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "User-specified description for the IP range.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "range": {
                    "description": "The IP range to advertise. The value must be a\nCIDR-formatted string.",
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

func GoogleComputeRouterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouter), &result)
	return &result
}
