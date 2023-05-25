package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableGcPolicy = `{
  "block": {
    "attributes": {
      "column_family": {
        "description": "The name of the column family.",
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
      "instance_name": {
        "description": "The name of the Bigtable instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "description": "If multiple policies are set, you should choose between UNION OR INTERSECTION.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table": {
        "description": "The name of the table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "max_age": {
        "block": {
          "attributes": {
            "days": {
              "computed": true,
              "deprecated": true,
              "description": "Number of days before applying GC policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "duration": {
              "computed": true,
              "description": "Duration before applying GC policy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "GC policy that applies to all cells older than the given age.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "max_version": {
        "block": {
          "attributes": {
            "number": {
              "description": "Number of version before applying the GC policy.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "GC policy that applies to all versions of a cell except for the most recent.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigtableGcPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableGcPolicy), &result)
	return &result
}
