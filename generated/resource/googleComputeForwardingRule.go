package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRule = `{
  "block": {
    "attributes": {
      "all_ports": {
        "description": "This field can only be used:\n* If 'IPProtocol' is one of TCP, UDP, or SCTP.\n* By internal TCP/UDP load balancers, backend service-based network load\nbalancers, and internal and external protocol forwarding.\n\n\nSet this field to true to allow packets addressed to any port or packets\nlacking destination port information (for example, UDP fragments after the\nfirst fragment) to be forwarded to the backends configured with this\nforwarding rule.\n\nThe 'ports', 'port_range', and\n'allPorts' fields are mutually exclusive.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allow_global_access": {
        "description": "This field is used along with the 'backend_service' field for\ninternal load balancing or with the 'target' field for internal\nTargetInstance.\n\nIf the field is set to 'TRUE', clients can access ILB from all\nregions.\n\nOtherwise only allows access from clients in the same region as the\ninternal load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backend_service": {
        "description": "Identifies the backend service to which the forwarding rule sends traffic.\n\nRequired for Internal TCP/UDP Load Balancing and Network Load Balancing;\nmust be omitted for all other load balancer types.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_forwarding_rule": {
        "computed": true,
        "description": "[Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.",
        "description_kind": "plain",
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
        "description": "IP address for which this forwarding rule accepts traffic. When a client\nsends traffic to this IP address, the forwarding rule directs the traffic\nto the referenced 'target' or 'backendService'.\n\nWhile creating a forwarding rule, specifying an 'IPAddress' is\nrequired under the following circumstances:\n\n* When the 'target' is set to 'targetGrpcProxy' and\n'validateForProxyless' is set to 'true', the\n'IPAddress' should be set to '0.0.0.0'.\n* When the 'target' is a Private Service Connect Google APIs\nbundle, you must specify an 'IPAddress'.\n\n\nOtherwise, you can optionally specify an IP address that references an\nexisting static (reserved) IP address resource. When omitted, Google Cloud\nassigns an ephemeral IP address.\n\nUse one of the following formats to specify an IP address while creating a\nforwarding rule:\n\n* IP address number, as in '100.1.2.3'\n* IPv6 address range, as in '2600:1234::/96'\n* Full resource URL, as in\n'https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name'\n* Partial URL or by name, as in:\n  * 'projects/project_id/regions/region/addresses/address-name'\n  * 'regions/region/addresses/address-name'\n  * 'global/addresses/address-name'\n  * 'address-name'\n\n\nThe forwarding rule's 'target' or 'backendService',\nand in most cases, also the 'loadBalancingScheme', determine the\ntype of IP address that you can use. For detailed information, see\n[IP address\nspecifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).\n\nWhen reading an 'IPAddress', the API always returns the IP\naddress number.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol to which this rule applies.\n\nFor protocol forwarding, valid\noptions are 'TCP', 'UDP', 'ESP',\n'AH', 'SCTP', 'ICMP' and\n'L3_DEFAULT'.\n\nThe valid IP protocols are different for different load balancing products\nas described in [Load balancing\nfeatures](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends). Possible values: [\"TCP\", \"UDP\", \"ESP\", \"AH\", \"SCTP\", \"ICMP\", \"L3_DEFAULT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_mirroring_collector": {
        "description": "Indicates whether or not this load balancer can be used as a collector for\npacket mirroring. To prevent mirroring loops, instances behind this\nload balancer will not have their traffic mirrored even if a\n'PacketMirroring' rule applies to them.\n\nThis can only be set to true for load balancers that have their\n'loadBalancingScheme' set to 'INTERNAL'.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this forwarding rule.  A list of key-\u003evalue pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "description": "Specifies the forwarding rule type.\n\nFor more information about forwarding rules, refer to\n[Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts). Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"EXTERNAL_MANAGED\", \"INTERNAL\", \"INTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-63 characters long, and comply with\n[RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n\nSpecifically, the name must be 1-63 characters long and match the regular\nexpression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters must\nbe a dash, lowercase letter, or digit, except the last character, which\ncannot be a dash.\n\nFor Private Service Connect forwarding rules that forward traffic to Google\nAPIs, the forwarding rule name must be a 1-20 characters string with\nlowercase letters and numbers and must start with a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "This field is not used for external load balancing.\n\nFor Internal TCP/UDP Load Balancing, this field identifies the network that\nthe load balanced IP should belong to for this Forwarding Rule.\nIf the subnetwork is specified, the network of the subnetwork will be used.\nIf neither subnetwork nor this field is specified, the default network will\nbe used.\n\nFor Private Service Connect forwarding rules that forward traffic to Google\nAPIs, a network must be provided.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "This signifies the networking tier used for configuring\nthis load balancer and can only take the following values:\n'PREMIUM', 'STANDARD'.\n\nFor regional ForwardingRule, the valid values are 'PREMIUM' and\n'STANDARD'. For GlobalForwardingRule, the valid value is\n'PREMIUM'.\n\nIf this field is not specified, it is assumed to be 'PREMIUM'.\nIf 'IPAddress' is specified, this value must be equal to the\nnetworkTier of the Address. Possible values: [\"PREMIUM\", \"STANDARD\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port_range": {
        "computed": true,
        "description": "This field can only be used:\n\n* If 'IPProtocol' is one of TCP, UDP, or SCTP.\n* By backend service-based network load balancers, target pool-based\nnetwork load balancers, internal proxy load balancers, external proxy load\nbalancers, Traffic Director, external protocol forwarding, and Classic VPN.\nSome products have restrictions on what ports can be used. See\n[port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)\nfor details.\n\n\nOnly packets addressed to ports in the specified range will be forwarded to\nthe backends configured with this forwarding rule.\n\nThe 'ports' and 'port_range' fields are mutually exclusive.\n\nFor external forwarding rules, two or more forwarding rules cannot use the\nsame '[IPAddress, IPProtocol]' pair, and cannot have\noverlapping 'portRange's.\n\nFor internal forwarding rules within the same VPC network, two or more\nforwarding rules cannot use the same '[IPAddress, IPProtocol]'\npair, and cannot have overlapping 'portRange's.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ports": {
        "description": "This field can only be used:\n\n* If 'IPProtocol' is one of TCP, UDP, or SCTP.\n* By internal TCP/UDP load balancers, backend service-based network load\nbalancers, and internal protocol forwarding.\n\n\nYou can specify a list of up to five ports by number, separated by commas.\nThe ports can be contiguous or discontiguous. Only packets addressed to\nthese ports will be forwarded to the backends configured with this\nforwarding rule.\n\nFor external forwarding rules, two or more forwarding rules cannot use the\nsame '[IPAddress, IPProtocol]' pair, and cannot share any values\ndefined in 'ports'.\n\nFor internal forwarding rules within the same VPC network, two or more\nforwarding rules cannot use the same '[IPAddress, IPProtocol]'\npair, and cannot share any values defined in 'ports'.\n\nThe 'ports' and 'port_range' fields are mutually exclusive.",
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
      "psc_connection_id": {
        "computed": true,
        "description": "The PSC connection id of the PSC Forwarding Rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "psc_connection_status": {
        "computed": true,
        "description": "The PSC connection status of the PSC Forwarding Rule. Possible values: 'STATUS_UNSPECIFIED', 'PENDING', 'ACCEPTED', 'REJECTED', 'CLOSED'",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "A reference to the region where the regional forwarding rule resides.\n\nThis field is not applicable to global forwarding rules.",
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
        "description": "The internal fully qualified service name for this Forwarding Rule.\n\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_ip_ranges": {
        "description": "If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnetwork": {
        "computed": true,
        "description": "This field identifies the subnetwork that the load balanced IP should\nbelong to for this Forwarding Rule, used in internal load balancing and\nnetwork load balancing with IPv6.\n\nIf the network specified is in auto subnet mode, this field is optional.\nHowever, a subnetwork must be specified if the network is in custom subnet\nmode or when creating external forwarding rule with IPv6.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target": {
        "description": "The URL of the target resource to receive the matched traffic.  For\nregional forwarding rules, this target must be in the same region as the\nforwarding rule. For global forwarding rules, this target must be a global\nload balancing resource.\n\nThe forwarded traffic must be of a type appropriate to the target object.\n*  For load balancers, see the \"Target\" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).\n*  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:\n  *  'vpc-sc' - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).\n  *  'all-apis' - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).\n\n\nFor Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "service_directory_registrations": {
        "block": {
          "attributes": {
            "namespace": {
              "computed": true,
              "description": "Service Directory namespace to register the forwarding rule under.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service": {
              "description": "Service Directory service to register the forwarding rule under.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Service Directory resources to register this forwarding rule with.\n\nCurrently, only supports a single Service Directory resource.",
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

func GoogleComputeForwardingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeForwardingRule), &result)
	return &result
}
