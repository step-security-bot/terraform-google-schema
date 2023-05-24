package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineServiceNetworkSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "The name of the service these settings apply to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "network_settings": {
        "block": {
          "attributes": {
            "ingress_traffic_allowed": {
              "description": "The ingress settings for version or service. Default value: \"INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED\" Possible values: [\"INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED\", \"INGRESS_TRAFFIC_ALLOWED_ALL\", \"INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY\", \"INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Ingress settings for this service. Will apply to all versions.",
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

func GoogleAppEngineServiceNetworkSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineServiceNetworkSettings), &result)
	return &result
}
