package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkPeeringRoutesConfig = `{
  "block": {
    "attributes": {
      "export_custom_routes": {
        "description": "Whether to export the custom routes to the peer network.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_custom_routes": {
        "description": "Whether to import the custom routes to the peer network.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "network": {
        "description": "The name of the primary network for the peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering": {
        "description": "Name of the peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func GoogleComputeNetworkPeeringRoutesConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkPeeringRoutesConfig), &result)
	return &result
}
