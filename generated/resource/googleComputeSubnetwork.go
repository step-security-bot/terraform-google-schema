package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSubnetwork = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource. This field can be set only at resource\ncreation time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_flow_logs": {
        "computed": true,
        "description": "Whether to enable flow logging for this subnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. This field is used internally during\nupdates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_address": {
        "computed": true,
        "description": "The gateway address for default routes to reach destination addresses\noutside this subnetwork.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_cidr_range": {
        "description": "The range of internal addresses that are owned by this subnetwork.\nProvide this property when you create the subnetwork. For example,\n10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and\nnon-overlapping within a network. Only IPv4 is supported.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the resource, provided by the client when initially\ncreating the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which\nmeans the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network this subnet belongs to.\nOnly networks that are in the distributed mode can have subnetworks.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_ip_google_access": {
        "description": "When enabled, VMs in this subnetwork without external IP addresses can\naccess Google APIs and services by using Private Google Access.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "URL of the GCP region for this subnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_ip_range": {
        "computed": true,
        "description": "An array of configurations for secondary IP ranges for VM instances\ncontained in this subnetwork. The primary IP of such VM must belong\nto the primary ipCidrRange of the subnetwork. The alias IPs may belong\nto either primary or secondary ranges.\nThis field uses attr-as-block mode to avoid breaking\nusers during the 0.12 upgrade. See [the Attr-as-Block page](https://www.terraform.io/docs/configuration/attr-as-blocks.html)\nfor more details.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "ip_cidr_range": "string",
              "range_name": "string"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "aggregation_interval": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nToggles the aggregation interval for collecting flow logs. Increasing the\ninterval time will reduce the amount of generated flow logs for long\nlasting connections. Default is an interval of 5 seconds per connection.\nPossible values are INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN,\nINTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flow_sampling": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nThe value of the field must be in [0, 1]. Set the sampling rate of VPC\nflow logs within the subnetwork where 1.0 means all collected logs are\nreported and 0.0 means no logs are reported. Default is 0.5 which means\nhalf of all collected logs are reported.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metadata": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nConfigures whether metadata fields should be added to the reported VPC\nflow logs. Default is 'INCLUDE_ALL_METADATA'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func GoogleComputeSubnetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSubnetwork), &result)
	return &result
}
