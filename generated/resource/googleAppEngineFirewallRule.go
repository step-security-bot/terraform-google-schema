package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineFirewallRule = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The action to take if this rule matches. Possible values: [\"UNSPECIFIED_ACTION\", \"ALLOW\", \"DENY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "An optional string description of this rule.",
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
      "priority": {
        "description": "A positive integer that defines the order of rule evaluation.\nRules with the lowest priority are evaluated first.\n\nA default rule at priority Int32.MaxValue matches all IPv4 and\nIPv6 traffic when no previous rule matches. Only the action of\nthis rule can be modified by the user.",
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
      "source_range": {
        "description": "IP address or range, defined using CIDR notation, of requests that this rule applies to.",
        "description_kind": "plain",
        "required": true,
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

func GoogleAppEngineFirewallRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineFirewallRule), &result)
	return &result
}
