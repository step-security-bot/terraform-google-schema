package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryAnalyticsHubListing = `{
  "block": {
    "attributes": {
      "categories": {
        "description": "Categories of the listing. Up to two categories are allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "data_exchange_id": {
        "description": "The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (\u0026) and can't start or end with spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "documentation": {
        "description": "Documentation describing the listing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "icon": {
        "description": "Base64 encoded image representing the listing.",
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
      "listing_id": {
        "description": "The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The name of the location this data exchange listing.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the listing. e.g. \"projects/myproject/locations/US/dataExchanges/123/listings/456\"",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_contact": {
        "description": "Email or URL of the primary point of contact of the listing.",
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
      "request_access": {
        "description": "Email or URL of the request access of the listing. Subscribers can use this reference to request access.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "bigquery_dataset": {
        "block": {
          "attributes": {
            "dataset": {
              "description": "Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Shared dataset i.e. BigQuery dataset source.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "data_provider": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the data provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_contact": {
              "description": "Email or URL of the data provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Details of the data provider who owns the source data.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "publisher": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the listing publisher.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_contact": {
              "description": "Email or URL of the listing publisher.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Details of the publisher who owns the listing and who can share the source data.",
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

func GoogleBigqueryAnalyticsHubListingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryAnalyticsHubListing), &result)
	return &result
}
