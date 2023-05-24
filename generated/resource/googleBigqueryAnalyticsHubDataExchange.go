package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryAnalyticsHubDataExchange = `{
  "block": {
    "attributes": {
      "data_exchange_id": {
        "description": "The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the data exchange.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "documentation": {
        "description": "Documentation describing the data exchange.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "icon": {
        "description": "Base64 encoded image representing the data exchange.",
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
      "listing_count": {
        "computed": true,
        "description": "Number of listings contained in the data exchange.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "description": "The name of the location this data exchange.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the data exchange, for example:\n\"projects/myproject/locations/US/dataExchanges/123\"",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_contact": {
        "description": "Email or URL of the primary point of contact of the data exchange.",
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

func GoogleBigqueryAnalyticsHubDataExchangeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryAnalyticsHubDataExchange), &result)
	return &result
}
