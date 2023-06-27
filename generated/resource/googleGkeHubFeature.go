package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubFeature = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. When the Feature resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. When the Feature resource was deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "GCP labels for this Feature.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full, unique name of this Feature resource",
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
      "resource_state": {
        "computed": true,
        "description": "State of the Feature resource itself.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "has_resources": "bool",
              "state": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. The Hub-wide Feature state",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "state": [
                "list",
                [
                  "object",
                  {
                    "code": "string",
                    "description": "string",
                    "update_time": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Output only. When the Feature resource was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "block_types": {
            "multiclusteringress": {
              "block": {
                "attributes": {
                  "config_membership": {
                    "description": "Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Multicluster Ingress-specific spec.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
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

func GoogleGkeHubFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFeature), &result)
	return &result
}
