package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterNat = `{
  "block": {
    "attributes": {
      "drain_nat_ips": {
        "computed": true,
        "description": "A list of URLs of the IP resources to be drained. These IPs must be\nvalid static external IPs that have been assigned to the NAT.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "enable_dynamic_port_allocation": {
        "computed": true,
        "description": "Enable Dynamic Port Allocation.\nIf minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.\nIf minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.\nIf maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.\nIf maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.\n\nMutually exclusive with enableEndpointIndependentMapping.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_endpoint_independent_mapping": {
        "computed": true,
        "description": "Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information\nsee the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).",
        "description_kind": "plain",
        "type": "bool"
      },
      "icmp_idle_timeout_sec": {
        "computed": true,
        "description": "Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_config": {
        "computed": true,
        "description": "Configuration for logging on NAT",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable": "bool",
              "filter": "string"
            }
          ]
        ]
      },
      "max_ports_per_vm": {
        "computed": true,
        "description": "Maximum number of ports allocated to a VM from this NAT.\nThis field can only be set when enableDynamicPortAllocation is enabled.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_ports_per_vm": {
        "computed": true,
        "description": "Minimum number of ports allocated to a VM from this NAT.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "Name of the NAT service. The name must be 1-63 characters long and\ncomply with RFC1035.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_allocate_option": {
        "computed": true,
        "description": "How external IPs should be allocated for this NAT. Valid values are\n'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud\nPlatform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: [\"MANUAL_ONLY\", \"AUTO_ONLY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "nat_ips": {
        "computed": true,
        "description": "Self-links of NAT IPs. Only valid if natIpAllocateOption\nis set to MANUAL_ONLY.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "Region where the router and NAT reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "The name of the Cloud Router in which this NAT will be configured.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "A list of rules associated with this NAT.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "action": [
                "list",
                [
                  "object",
                  {
                    "source_nat_active_ips": [
                      "set",
                      "string"
                    ],
                    "source_nat_drain_ips": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "description": "string",
              "match": "string",
              "rule_number": "number"
            }
          ]
        ]
      },
      "source_subnetwork_ip_ranges_to_nat": {
        "computed": true,
        "description": "How NAT should be configured per Subnetwork.\nIf 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the\nIP ranges in every Subnetwork are allowed to Nat.\nIf 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP\nranges in every Subnetwork are allowed to Nat.\n'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat\n(specified in the field subnetwork below). Note that if this field\ncontains ALL_SUBNETWORKS_ALL_IP_RANGES or\nALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any\nother RouterNat section in any Router for this network in this region. Possible values: [\"ALL_SUBNETWORKS_ALL_IP_RANGES\", \"ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES\", \"LIST_OF_SUBNETWORKS\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "One or more subnetwork NAT configurations. Only used if\n'source_subnetwork_ip_ranges_to_nat' is set to 'LIST_OF_SUBNETWORKS'",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "secondary_ip_range_names": [
                "set",
                "string"
              ],
              "source_ip_ranges_to_nat": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "tcp_established_idle_timeout_sec": {
        "computed": true,
        "description": "Timeout (in seconds) for TCP established connections.\nDefaults to 1200s if not set.",
        "description_kind": "plain",
        "type": "number"
      },
      "tcp_time_wait_timeout_sec": {
        "computed": true,
        "description": "Timeout (in seconds) for TCP connections that are in TIME_WAIT state.\nDefaults to 120s if not set.",
        "description_kind": "plain",
        "type": "number"
      },
      "tcp_transitory_idle_timeout_sec": {
        "computed": true,
        "description": "Timeout (in seconds) for TCP transitory connections.\nDefaults to 30s if not set.",
        "description_kind": "plain",
        "type": "number"
      },
      "udp_idle_timeout_sec": {
        "computed": true,
        "description": "Timeout (in seconds) for UDP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRouterNatSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterNat), &result)
	return &result
}
