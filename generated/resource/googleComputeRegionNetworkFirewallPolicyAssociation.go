package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkFirewallPolicyAssociation = `{
  "block": {
    "attributes": {
      "attachment_target": {
        "description": "The target that the firewall policy is attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_policy": {
        "description": "The firewall policy ID of the association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name for an association.",
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
      "short_name": {
        "computed": true,
        "description": "The short name of the firewall policy of the association.",
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

func GoogleComputeRegionNetworkFirewallPolicyAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkFirewallPolicyAssociation), &result)
	return &result
}
