package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputePacketMirroring = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the rule.",
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
        "description": "The name of the packet mirroring rule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "Since only one rule can be active at a time, priority is\nused to break ties in the case of two rules that apply to\nthe same instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created address should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "collector_ilb": {
        "block": {
          "attributes": {
            "url": {
              "description": "The URL of the forwarding rule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)\nthat will be used as collector for mirrored traffic. The\nspecified forwarding rule must have is_mirroring_collector\nset to true.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "filter": {
        "block": {
          "attributes": {
            "cidr_ranges": {
              "description": "IP CIDR ranges that apply as a filter on the source (ingress) or\ndestination (egress) IP in the IP header. Only IPv4 is supported.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "direction": {
              "description": "Direction of traffic to mirror. Default value: \"BOTH\" Possible values: [\"INGRESS\", \"EGRESS\", \"BOTH\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_protocols": {
              "description": "Protocols that apply as a filter on mirrored traffic. Possible values: [\"tcp\", \"udp\", \"icmp\"]",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "A filter for mirrored traffic.  If unset, all traffic is mirrored.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mirrored_resources": {
        "block": {
          "attributes": {
            "tags": {
              "description": "All instances with these tags will be mirrored.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "instances": {
              "block": {
                "attributes": {
                  "url": {
                    "description": "The URL of the instances where this rule should be active.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "All the listed instances will be mirrored.  Specify at most 50.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "subnetworks": {
              "block": {
                "attributes": {
                  "url": {
                    "description": "The URL of the subnetwork where this rule should be active.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "All instances in one of these subnetworks will be mirrored.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A means of specifying which resources to mirror.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "network": {
        "block": {
          "attributes": {
            "url": {
              "description": "The full self_link URL of the network where this rule is active.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the mirrored VPC network. Only packets in this network\nwill be mirrored. All mirrored VMs should have a NIC in the given\nnetwork. All mirrored subnetworks should belong to the given network.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleComputePacketMirroringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputePacketMirroring), &result)
	return &result
}
