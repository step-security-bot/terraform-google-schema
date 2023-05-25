package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerInstance = `{
  "block": {
    "attributes": {
      "config": {
        "description": "The name of the instance's configuration (similar but not\nquite the same as a region) which defines defines the geographic placement and\nreplication of your databases in this instance. It determines where your data\nis stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.\nIn order to obtain a valid list please consult the\n[Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The descriptive name for this instance as it appears in UIs. Must be\nunique per project and between 4 and 30 characters in length.",
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
      "labels": {
        "description": "An object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for the instance, which cannot be changed after\nthe instance is created. The name must be between 6 and 30 characters\nin length.\n\n\nIf not provided, a random string starting with 'tf-' will be selected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "num_nodes": {
        "description": "The number of nodes allocated to this instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Instance status: 'CREATING' or 'READY'.",
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

func GoogleSpannerInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerInstance), &result)
	return &result
}
