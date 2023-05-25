package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeGlobalForwardingRule = `{
  "block": {
    "attributes": {
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
        "description": "The IP protocol to which this rule applies. Valid options are TCP,\nUDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is\nINTERNAL_SELF_MANAGED, only TCP is valid.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "description": "The IP Version that will be used by this global forwarding rule.\nValid options are IPV4 or IPV6.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_scheme": {
        "description": "This signifies what the GlobalForwardingRule will be used for.\nThe value of INTERNAL_SELF_MANAGED means that this will be used for\nInternal Global HTTP(S) LB. The value of EXTERNAL means that this\nwill be used for External Global Load Balancing (HTTP(S) LB,\nExternal TCP/UDP LB, SSL Proxy)\n\nNOTE: Currently global forwarding rules cannot be used for INTERNAL\nload balancing.",
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
      "port_range": {
        "description": "This field is used along with the target field for TargetHttpProxy,\nTargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,\nTargetPool, TargetInstance.\n\nApplicable only when IPProtocol is TCP, UDP, or SCTP, only packets\naddressed to ports in the specified range will be forwarded to target.\nForwarding rules with the same [IPAddress, IPProtocol] pair must have\ndisjoint port ranges.\n\nSome types of forwarding target have constraints on the acceptable\nports:\n\n* TargetHttpProxy: 80, 8080\n* TargetHttpsProxy: 443\n* TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,\n                  1883, 5222\n* TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,\n                  1883, 5222\n* TargetVpnGateway: 500, 4500",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "description": "The URL of the target resource to receive the matched traffic.\nThe forwarded traffic must be of a type appropriate to the target object.\nFor INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets\nare valid.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "metadata_filters": {
        "block": {
          "attributes": {
            "filter_match_criteria": {
              "description": "Specifies how individual filterLabel matches within the list of\nfilterLabels contribute towards the overall metadataFilter match.\n\nMATCH_ANY - At least one of the filterLabels must have a matching\nlabel in the provided metadata.\nMATCH_ALL - All filterLabels must have matching labels in the\nprovided metadata.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "filter_labels": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the metadata label. The length must be between\n1 and 1024 characters, inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value that the label must match. The value has a maximum\nlength of 1024 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 64,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
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

func GoogleComputeGlobalForwardingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeGlobalForwardingRule), &result)
	return &result
}
