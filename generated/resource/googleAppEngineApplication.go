package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineApplication = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gcr_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serving_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "url_dispatch_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain": "string",
              "path": "string",
              "service": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "feature_settings": {
        "block": {
          "attributes": {
            "split_health_checks": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAppEngineApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineApplication), &result)
	return &result
}
