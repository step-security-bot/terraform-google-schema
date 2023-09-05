package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBiglakeCatalog = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The creation time of the catalog. A timestamp in RFC3339 UTC\n\"Zulu\" format, with nanosecond resolution and up to nine fractional\ndigits.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. The deletion time of the catalog. Only set after the catalog\nis deleted. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "Output only. The time when this catalog is considered expired. Only set\nafter the catalog is deleted. Only set after the catalog is deleted.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The geographic location where the Catalog should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the Catalog. Format:\nprojects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last modification time of the catalog. A timestamp in\nRFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits.",
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

func GoogleBiglakeCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBiglakeCatalog), &result)
	return &result
}
