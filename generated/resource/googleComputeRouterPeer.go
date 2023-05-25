package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterPeer = `{
  "block": {
    "attributes": {
      "advertise_mode": {
        "description": "User-specified flag to indicate which mode to use for advertisement.\nValid values of this enum field are: 'DEFAULT', 'CUSTOM' Default value: \"DEFAULT\" Possible values: [\"DEFAULT\", \"CUSTOM\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "advertised_groups": {
        "description": "User-specified list of prefix groups to advertise in custom\nmode, which can take one of the following options:\n\n* 'ALL_SUBNETS': Advertises all available subnets, including peer VPC subnets.\n* 'ALL_VPC_SUBNETS': Advertises the router's own VPC subnets.\n* 'ALL_PEER_VPC_SUBNETS': Advertises peer subnets of the router's VPC network.\n\n\nNote that this field can only be populated if advertiseMode is 'CUSTOM'\nand overrides the list defined for the router (in the \"bgp\" message).\nThese groups are advertised in addition to any specified prefixes.\nLeave this field blank to advertise no custom groups.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "advertised_route_priority": {
        "description": "The priority of routes advertised to this BGP peer.\nWhere there is more than one matching route of maximum\nlength, the routes with the lowest priority value win.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable": {
        "description": "The status of the BGP peer connection. If set to false, any active session\nwith the peer is terminated and all associated routing information is removed.\nIf set to true, the peer connection can be established with routing information.\nThe default is true.",
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
      "interface": {
        "description": "Name of the interface the BGP peer is associated with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "IP address of the interface inside Google Cloud Platform.\nOnly IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "management_type": {
        "computed": true,
        "description": "The resource that configures and manages this BGP peer.\n\n* 'MANAGED_BY_USER' is the default value and can be managed by\nyou or other users\n* 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and\nmanaged by Cloud Interconnect, specifically by an\nInterconnectAttachment of type PARTNER. Google automatically\ncreates, updates, and deletes this type of BGP peer when the\nPARTNER InterconnectAttachment is created, updated,\nor deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of this BGP peer. The name must be 1-63 characters long,\nand comply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which\nmeans the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_asn": {
        "description": "Peer BGP Autonomous System Number (ASN).\nEach BGP interface may use a different value.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "peer_ip_address": {
        "description": "IP address of the BGP interface outside Google Cloud Platform.\nOnly IPv4 is supported.",
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
        "description": "Region where the router and BgpPeer reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "The name of the Cloud Router in which this BgpPeer will be configured.",
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
              "description": "User-specified description for the IP range.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range": {
              "description": "The IP range to advertise. The value must be a\nCIDR-formatted string.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "User-specified list of individual IP ranges to advertise in\ncustom mode. This field can only be populated if advertiseMode\nis 'CUSTOM' and is advertised to all peers of the router. These IP\nranges will be advertised in addition to any specified groups.\nLeave this field blank to advertise no custom IP ranges.",
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

func GoogleComputeRouterPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterPeer), &result)
	return &result
}
