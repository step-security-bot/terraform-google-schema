package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterNat = `{
  "block": {
    "attributes": {
      "drain_nat_ips": {
        "description": "A list of URLs of the IP resources to be drained. These IPs must be\nvalid static external IPs that have been assigned to the NAT.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "enable_endpoint_independent_mapping": {
        "description": "Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information\nsee the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "icmp_idle_timeout_sec": {
        "description": "Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_ports_per_vm": {
        "description": "Minimum number of ports allocated to a VM from this NAT.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of the NAT service. The name must be 1-63 characters long and\ncomply with RFC1035.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_allocate_option": {
        "description": "How external IPs should be allocated for this NAT. Valid values are\n'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud\nPlatform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: [\"MANUAL_ONLY\", \"AUTO_ONLY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ips": {
        "description": "Self-links of NAT IPs. Only valid if natIpAllocateOption\nis set to MANUAL_ONLY.",
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
      "source_subnetwork_ip_ranges_to_nat": {
        "description": "How NAT should be configured per Subnetwork.\nIf 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the\nIP ranges in every Subnetwork are allowed to Nat.\nIf 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP\nranges in every Subnetwork are allowed to Nat.\n'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat\n(specified in the field subnetwork below). Note that if this field\ncontains ALL_SUBNETWORKS_ALL_IP_RANGES or\nALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any\nother RouterNat section in any Router for this network in this region. Possible values: [\"ALL_SUBNETWORKS_ALL_IP_RANGES\", \"ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES\", \"LIST_OF_SUBNETWORKS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tcp_established_idle_timeout_sec": {
        "description": "Timeout (in seconds) for TCP established connections.\nDefaults to 1200s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tcp_transitory_idle_timeout_sec": {
        "description": "Timeout (in seconds) for TCP transitory connections.\nDefaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "udp_idle_timeout_sec": {
        "description": "Timeout (in seconds) for UDP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "enable": {
              "description": "Indicates whether or not to export logs.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "filter": {
              "description": "Specifies the desired filtering of logs on this NAT. Possible values: [\"ERRORS_ONLY\", \"TRANSLATIONS_ONLY\", \"ALL\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for logging on NAT",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "subnetwork": {
        "block": {
          "attributes": {
            "name": {
              "description": "Self-link of subnetwork to NAT",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secondary_ip_range_names": {
              "description": "List of the secondary ranges of the subnetwork that are allowed\nto use NAT. This can be populated only if\n'LIST_OF_SECONDARY_IP_RANGES' is one of the values in\nsourceIpRangesToNat",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "source_ip_ranges_to_nat": {
              "description": "List of options for which source IPs in the subnetwork\nshould have NAT enabled. Supported values include:\n'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES',\n'PRIMARY_IP_RANGE'.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "One or more subnetwork NAT configurations. Only used if\n'source_subnetwork_ip_ranges_to_nat' is set to 'LIST_OF_SUBNETWORKS'",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleComputeRouterNatSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterNat), &result)
	return &result
}
