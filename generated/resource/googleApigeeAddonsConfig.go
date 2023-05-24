package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeAddonsConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "org": {
        "description": "Name of the Apigee organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "addons_config": {
        "block": {
          "block_types": {
            "advanced_api_ops_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for the Monetization add-on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "api_security_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expires_at": {
                    "computed": true,
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Configuration for the Monetization add-on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "connectors_platform_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expires_at": {
                    "computed": true,
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Configuration for the Monetization add-on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "integration_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for the Monetization add-on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "monetization_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Flag that specifies whether the Advanced API Ops add-on is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for the Monetization add-on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Addon configurations of the Apigee organization.",
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

func GoogleApigeeAddonsConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeAddonsConfig), &result)
	return &result
}
