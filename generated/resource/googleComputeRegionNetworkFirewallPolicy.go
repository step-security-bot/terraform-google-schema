package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkFirewallPolicy = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of the resource. This field is used internally during updates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The location of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_network_firewall_policy_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_tuple_count": {
        "computed": true,
        "description": "Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.",
        "description_kind": "plain",
        "type": "number"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_id": {
        "computed": true,
        "description": "Server-defined URL for this resource with the resource id.",
        "description_kind": "plain",
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

func GoogleComputeRegionNetworkFirewallPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkFirewallPolicy), &result)
	return &result
}
