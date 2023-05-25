package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeGlobalForwardingRule = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
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
        "description": "IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule. If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address: * IPv4 dotted decimal, as in ` + "`" + `100.1.2.3` + "`" + ` * Full URL, as in ` + "`" + `https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name` + "`" + ` * Partial URL or by name, as in: * ` + "`" + `projects/project_id/regions/region/addresses/address-name` + "`" + ` * ` + "`" + `regions/region/addresses/address-name` + "`" + ` * ` + "`" + `global/addresses/address-name` + "`" + ` * ` + "`" + `address-name` + "`" + ` The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol to which this rule applies. For protocol forwarding, valid options are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ESP` + "`" + `, ` + "`" + `AH` + "`" + `, ` + "`" + `SCTP` + "`" + ` or ` + "`" + `ICMP` + "`" + `. For Internal TCP/UDP Load Balancing, the load balancing scheme is ` + "`" + `INTERNAL` + "`" + `, and one of ` + "`" + `TCP` + "`" + ` or ` + "`" + `UDP` + "`" + ` are valid. For Traffic Director, the load balancing scheme is ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + `, and only ` + "`" + `TCP` + "`" + `is valid. For Internal HTTP(S) Load Balancing, the load balancing scheme is ` + "`" + `INTERNAL_MANAGED` + "`" + `, and only ` + "`" + `TCP` + "`" + ` is valid. For HTTP(S), SSL Proxy, and TCP Proxy Load Balancing, the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + ` and only ` + "`" + `TCP` + "`" + ` is valid. For Network TCP/UDP Load Balancing, the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + `, and one of ` + "`" + `TCP` + "`" + ` or ` + "`" + `UDP` + "`" + ` is valid.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "description": "The IP Version that will be used by this forwarding rule. Valid options are ` + "`" + `IPV4` + "`" + ` or ` + "`" + `IPV6` + "`" + `. This can only be specified for an external global forwarding rule. Possible values: UNSPECIFIED_VERSION, IPV4, IPV6",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "Used internally during label updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this rule.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "description": "Specifies the forwarding rule type.\n\n*   ` + "`" + `EXTERNAL` + "`" + ` is used for:\n    *   Classic Cloud VPN gateways\n    *   Protocol forwarding to VMs from an external IP address\n    *   The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP\n*   ` + "`" + `INTERNAL` + "`" + ` is used for:\n    *   Protocol forwarding to VMs from an internal IP address\n    *   Internal TCP/UDP load balancers\n*   ` + "`" + `INTERNAL_MANAGED` + "`" + ` is used for:\n    *   Internal HTTP(S) load balancers\n*   ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` is used for:\n    *   Traffic Director\n\nFor more information about forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values: INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]*[a-z0-9])?` + "`" + ` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "This field is not used for external load balancing. For ` + "`" + `INTERNAL` + "`" + ` and ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` load balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port_range": {
        "description": "When the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + `, ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` and ` + "`" + `INTERNAL_MANAGED` + "`" + `, you can specify a ` + "`" + `port_range` + "`" + `. Use with a forwarding rule that points to a target proxy or a target pool. Do not use with a forwarding rule that points to a backend service. This field is used along with the ` + "`" + `target` + "`" + ` field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when ` + "`" + `IPProtocol` + "`" + ` is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, or ` + "`" + `SCTP` + "`" + `, only packets addressed to ports in the specified range will be forwarded to ` + "`" + `target` + "`" + `. Forwarding rules with the same ` + "`" + `[IPAddress, IPProtocol]` + "`" + ` pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:\n\n*   TargetHttpProxy: 80, 8080\n*   TargetHttpsProxy: 443\n*   TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n*   TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n*   TargetVpnGateway: 500, 4500\n\n@pattern: d+(?:-d+)?",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project this resource belongs in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "[Output Only] Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "description": "The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. For ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` load balancing, only ` + "`" + `targetHttpProxy` + "`" + ` is valid, not ` + "`" + `targetHttpsProxy` + "`" + `.",
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
              "description": "Specifies how individual ` + "`" + `filterLabel` + "`" + ` matches within the list of ` + "`" + `filterLabel` + "`" + `s contribute towards the overall ` + "`" + `metadataFilter` + "`" + ` match.\n\nSupported values are:\n\n*   MATCH_ANY: At least one of the ` + "`" + `filterLabels` + "`" + ` must have a matching label in the provided metadata.\n*   MATCH_ALL: All ` + "`" + `filterLabels` + "`" + ` must have matching labels in the provided metadata. Possible values: NOT_SET, MATCH_ALL, MATCH_ANY",
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
                    "description": "Name of metadata label.\n\nThe name can have a maximum length of 1024 characters and must be at least 1 character long.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value of the label must match the specified value.\n\nvalue can have a maximum length of 1024 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of label value pairs that must match labels in the provided metadata based on ` + "`" + `filterMatchCriteria` + "`" + `\n\nThis list must not be empty and can have at the most 64 entries.",
                "description_kind": "plain"
              },
              "max_items": 64,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of [xDS](https://github.com/envoyproxy/data-plane-api/blob/master/XDS_PROTOCOL.md) compliant clients. In their xDS requests to Loadbalancer, xDS clients present [node metadata](https://github.com/envoyproxy/data-plane-api/search?q=%22message+Node%22+in%3A%2Fenvoy%2Fapi%2Fv2%2Fcore%2Fbase.proto\u0026). If a match takes place, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. ` + "`" + `TargetHttpProxy` + "`" + `, ` + "`" + `UrlMap` + "`" + `) referenced by the ` + "`" + `ForwardingRule` + "`" + ` will not be visible to those proxies.\n\nFor each ` + "`" + `metadataFilter` + "`" + ` in this list, if its ` + "`" + `filterMatchCriteria` + "`" + ` is set to MATCH_ANY, at least one of the ` + "`" + `filterLabel` + "`" + `s must match the corresponding label provided in the metadata. If its ` + "`" + `filterMatchCriteria` + "`" + ` is set to MATCH_ALL, then all of its ` + "`" + `filterLabel` + "`" + `s must match with corresponding labels provided in the metadata.\n\n` + "`" + `metadataFilters` + "`" + ` specified here will be applifed before those specified in the ` + "`" + `UrlMap` + "`" + ` that this ` + "`" + `ForwardingRule` + "`" + ` references.\n\n` + "`" + `metadataFilters` + "`" + ` only applies to Loadbalancers that have their loadBalancingScheme set to ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + `.",
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
