package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineApplicationUrlDispatchRules = `{
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
      }
    },
    "block_types": {
      "dispatch_rules": {
        "block": {
          "attributes": {
            "domain": {
              "description": "Domain name to match against. The wildcard \"*\" is supported if specified before a period: \"*.\".\nDefaults to matching all domains: \"*\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description": "Pathname within the host. Must start with a \"/\". A single \"*\" can be included at the end of the path.\nThe sum of the lengths of the domain and path may not exceed 100 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service": {
              "description": "Pathname within the host. Must start with a \"/\". A single \"*\" can be included at the end of the path.\nThe sum of the lengths of the domain and path may not exceed 100 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Rules to match an HTTP request and dispatch that request to a service.",
          "description_kind": "plain"
        },
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

func GoogleAppEngineApplicationUrlDispatchRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineApplicationUrlDispatchRules), &result)
	return &result
}
