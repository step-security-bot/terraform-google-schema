package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRule = `{
  "block": {
    "attributes": {
      "all_ports": {
        "description": "For internal TCP/UDP load balancing (i.e. load balancing scheme is\nINTERNAL and protocol is TCP/UDP), set this to true to allow packets\naddressed to any ports to be forwarded to the backends configured\nwith this forwarding rule. Used with backend service. Cannot be set\nif port or portRange are set.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backend_service": {
        "description": "A BackendService to receive the matched traffic. This is used only\nfor INTERNAL load balancing.",
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
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
      "ip_address": {
        "computed": true,
        "description": "The IP address that this forwarding rule is serving on behalf of.\n\nAddresses are restricted based on the forwarding rule's load balancing\nscheme (EXTERNAL or INTERNAL) and scope (global or regional).\n\nWhen the load balancing scheme is EXTERNAL, for global forwarding\nrules, the address must be a global IP, and for regional forwarding\nrules, the address must live in the same region as the forwarding\nrule. If this field is empty, an ephemeral IPv4 address from the same\nscope (global or regional) will be assigned. A regional forwarding\nrule supports IPv4 only. A global forwarding rule supports either IPv4\nor IPv6.\n\nWhen the load balancing scheme is INTERNAL, this can only be an RFC\n1918 IP address belonging to the network/subnet configured for the\nforwarding rule. By default, if this field is empty, an ephemeral\ninternal IP address will be automatically allocated from the IP range\nof the subnet or network configured for this forwarding rule.\n\n~\u003e **NOTE** The address should be specified as a literal IP address,\ne.g. '100.1.2.3' to avoid a permanent diff, as the server returns the\nIP address regardless of the input value.\n\nThe server accepts a literal IP address or a URL reference to an existing\nAddress resource. The following examples are all valid but only the first\nwill prevent a permadiff. If you are using 'google_compute_address' or\nsimilar, interpolate using '.address' instead of '.self_link' or similar\nto prevent a diff on re-apply.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol to which this rule applies. Valid options are TCP,\nUDP, ESP, AH, SCTP or ICMP.\n\nWhen the load balancing scheme is INTERNAL, only TCP and UDP are\nvalid.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "description": "ipVersion is not a valid field for regional forwarding rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_scheme": {
        "description": "This signifies what the ForwardingRule will be used for and can be\nEXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic\nCloud VPN gateways, protocol forwarding to VMs from an external IP address,\nand HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.\nINTERNAL is used for protocol forwarding to VMs from an internal IP address,\nand internal TCP/UDP load balancers.\nINTERNAL_MANAGED is used for internal HTTP(S) load balancers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "For internal load balancing, this field identifies the network that\nthe load balanced IP should belong to for this Forwarding Rule. If\nthis field is not specified, the default network will be used.\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "The networking tier used for configuring this address. This field can\ntake the following values: PREMIUM or STANDARD. If this field is not\nspecified, it is assumed to be PREMIUM.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port_range": {
        "description": "This field is used along with the target field for TargetHttpProxy,\nTargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,\nTargetPool, TargetInstance.\n\nApplicable only when IPProtocol is TCP, UDP, or SCTP, only packets\naddressed to ports in the specified range will be forwarded to target.\nForwarding rules with the same [IPAddress, IPProtocol] pair must have\ndisjoint port ranges.\n\nSome types of forwarding target have constraints on the acceptable\nports:\n\n* TargetHttpProxy: 80, 8080\n* TargetHttpsProxy: 443\n* TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,\n                  1883, 5222\n* TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,\n                  1883, 5222\n* TargetVpnGateway: 500, 4500",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ports": {
        "description": "This field is used along with the backend_service field for internal\nload balancing.\n\nWhen the load balancing scheme is INTERNAL, a single port or a comma\nseparated list of ports can be configured. Only packets addressed to\nthese ports will be forwarded to the backends configured with this\nforwarding rule.\n\nYou may specify a maximum of up to 5 ports.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "A reference to the region where the regional forwarding rule resides.\nThis field is not applicable to global forwarding rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_label": {
        "description": "An optional prefix to the service name for this Forwarding Rule.\nIf specified, will be the first label of the fully qualified service\nname.\n\nThe label must be 1-63 characters long, and comply with RFC1035.\nSpecifically, the label must be 1-63 characters long and match the\nregular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters\nmust be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.\n\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description": "The internal fully qualified service name for this Forwarding Rule.\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "The subnetwork that the load balanced IP should belong to for this\nForwarding Rule.  This field is only used for INTERNAL load balancing.\n\nIf the network specified is in auto subnet mode, this field is\noptional. However, if the network is in custom subnet mode, a\nsubnetwork must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target": {
        "description": "This field is only used for EXTERNAL load balancing.\nA reference to a TargetPool resource to receive the matched traffic.\nThis target must live in the same region as the forwarding rule.\nThe forwarded traffic must be of a type appropriate to the target\nobject.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleComputeForwardingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeForwardingRule), &result)
	return &result
}
