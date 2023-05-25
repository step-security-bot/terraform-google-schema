package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRule = `{
  "block": {
    "attributes": {
      "all_ports": {
        "computed": true,
        "description": "This field is used along with the ` + "`" + `backend_service` + "`" + ` field for internal load balancing or with the ` + "`" + `target` + "`" + ` field for internal TargetInstance. This field cannot be used with ` + "`" + `port` + "`" + ` or ` + "`" + `portRange` + "`" + ` fields. When the load balancing scheme is ` + "`" + `INTERNAL` + "`" + ` and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.",
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_global_access": {
        "computed": true,
        "description": "This field is used along with the ` + "`" + `backend_service` + "`" + ` field for internal load balancing or with the ` + "`" + `target` + "`" + ` field for internal TargetInstance. If the field is set to ` + "`" + `TRUE` + "`" + `, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "backend_service": {
        "computed": true,
        "description": "This field is only used for ` + "`" + `INTERNAL` + "`" + ` load balancing. For internal load balancing, this field identifies the BackendService resource to receive the matched traffic.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "[Output Only] Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol to which this rule applies. For protocol forwarding, valid options are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ESP` + "`" + `, ` + "`" + `AH` + "`" + `, ` + "`" + `SCTP` + "`" + ` or ` + "`" + `ICMP` + "`" + `. For Internal TCP/UDP Load Balancing, the load balancing scheme is ` + "`" + `INTERNAL` + "`" + `, and one of ` + "`" + `TCP` + "`" + ` or ` + "`" + `UDP` + "`" + ` are valid. For Traffic Director, the load balancing scheme is ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + `, and only ` + "`" + `TCP` + "`" + `is valid. For Internal HTTP(S) Load Balancing, the load balancing scheme is ` + "`" + `INTERNAL_MANAGED` + "`" + `, and only ` + "`" + `TCP` + "`" + ` is valid. For HTTP(S), SSL Proxy, and TCP Proxy Load Balancing, the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + ` and only ` + "`" + `TCP` + "`" + ` is valid. For Network TCP/UDP Load Balancing, the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + `, and one of ` + "`" + `TCP` + "`" + ` or ` + "`" + `UDP` + "`" + ` is valid.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_mirroring_collector": {
        "computed": true,
        "description": "Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a ` + "`" + `PacketMirroring` + "`" + ` rule applies to them. This can only be set to true for load balancers that have their ` + "`" + `loadBalancingScheme` + "`" + ` set to ` + "`" + `INTERNAL` + "`" + `.",
        "description_kind": "plain",
        "type": "bool"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "Used internally during label updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels to apply to this rule.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "computed": true,
        "description": "Specifies the forwarding rule type.\n\n*   ` + "`" + `EXTERNAL` + "`" + ` is used for:\n    *   Classic Cloud VPN gateways\n    *   Protocol forwarding to VMs from an external IP address\n    *   The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP\n*   ` + "`" + `INTERNAL` + "`" + ` is used for:\n    *   Protocol forwarding to VMs from an internal IP address\n    *   Internal TCP/UDP load balancers\n*   ` + "`" + `INTERNAL_MANAGED` + "`" + ` is used for:\n    *   Internal HTTP(S) load balancers\n*   ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` is used for:\n    *   Traffic Director\n\nFor more information about forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values: INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL",
        "description_kind": "plain",
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
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "This signifies the networking tier used for configuring this load balancer and can only take the following values: ` + "`" + `PREMIUM` + "`" + `, ` + "`" + `STANDARD` + "`" + `. For regional ForwardingRule, the valid values are ` + "`" + `PREMIUM` + "`" + ` and ` + "`" + `STANDARD` + "`" + `. For GlobalForwardingRule, the valid value is ` + "`" + `PREMIUM` + "`" + `. If this field is not specified, it is assumed to be ` + "`" + `PREMIUM` + "`" + `. If ` + "`" + `IPAddress` + "`" + ` is specified, this value must be equal to the networkTier of the Address.",
        "description_kind": "plain",
        "type": "string"
      },
      "port_range": {
        "computed": true,
        "description": "When the load balancing scheme is ` + "`" + `EXTERNAL` + "`" + `, ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` and ` + "`" + `INTERNAL_MANAGED` + "`" + `, you can specify a ` + "`" + `port_range` + "`" + `. Use with a forwarding rule that points to a target proxy or a target pool. Do not use with a forwarding rule that points to a backend service. This field is used along with the ` + "`" + `target` + "`" + ` field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when ` + "`" + `IPProtocol` + "`" + ` is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, or ` + "`" + `SCTP` + "`" + `, only packets addressed to ports in the specified range will be forwarded to ` + "`" + `target` + "`" + `. Forwarding rules with the same ` + "`" + `[IPAddress, IPProtocol]` + "`" + ` pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:\n\n*   TargetHttpProxy: 80, 8080\n*   TargetHttpsProxy: 443\n*   TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n*   TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n*   TargetVpnGateway: 500, 4500\n\n@pattern: d+(?:-d+)?",
        "description_kind": "plain",
        "type": "string"
      },
      "ports": {
        "computed": true,
        "description": "This field is used along with the ` + "`" + `backend_service` + "`" + ` field for internal load balancing. When the load balancing scheme is ` + "`" + `INTERNAL` + "`" + `, a list of ports can be configured, for example, ['80'], ['8000','9000']. Only packets addressed to these ports are forwarded to the backends configured with the forwarding rule. If the forwarding rule's loadBalancingScheme is INTERNAL, you can specify ports in one of the following ways: * A list of up to five ports, which can be non-contiguous * Keyword ` + "`" + `ALL` + "`" + `, which causes the forwarding rule to forward traffic on any port of the forwarding rule's protocol. @pattern: d+(?:-d+)? For more information, refer to [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "description": "The project this resource belongs in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The location of this resource.",
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
      "service_label": {
        "computed": true,
        "description": "An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the label must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]*[a-z0-9])?` + "`" + ` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description": "[Output Only] The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "This field is only used for ` + "`" + `INTERNAL` + "`" + ` load balancing. For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule. If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "computed": true,
        "description": "The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. For ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` load balancing, only ` + "`" + `targetHttpProxy` + "`" + ` is valid, not ` + "`" + `targetHttpsProxy` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
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
