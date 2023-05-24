package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogTaxonomy = `{
  "block": {
    "attributes": {
      "activated_policy_types": {
        "description": "A list of policy types that are activated for this taxonomy. If not set,\ndefaults to an empty list. Possible values: [\"POLICY_TYPE_UNSPECIFIED\", \"FINE_GRAINED_ACCESS_CONTROL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "Description of this taxonomy. It must: contain only unicode characters,\ntabs, newlines, carriage returns and page breaks; and be at most 2000 bytes\nlong when encoded in UTF-8. If not set, defaults to an empty description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User defined name of this taxonomy.\nIt must: contain only unicode letters, numbers, underscores, dashes\nand spaces; not start or end with spaces; and be at most 200 bytes\nlong when encoded in UTF-8.",
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
        "computed": true,
        "description": "Resource name of this taxonomy, whose format is:\n\"projects/{project}/locations/{region}/taxonomies/{taxonomy}\".",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Taxonomy location region.",
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

func GoogleDataCatalogTaxonomySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogTaxonomy), &result)
	return &result
}
