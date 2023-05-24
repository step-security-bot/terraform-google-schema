package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerInstance = `{
  "block": {
    "attributes": {
      "config": {
        "description": "The name of the instance's configuration (similar but not\nquite the same as a region) which defines the geographic placement and\nreplication of your databases in this instance. It determines where your data\nis stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.\nIn order to obtain a valid list please consult the\n[Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The descriptive name for this instance as it appears in UIs. Must be\nunique per project and between 4 and 30 characters in length.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_destroy": {
        "computed": true,
        "description": "When deleting a spanner instance, this boolean option will delete all backups of this instance.\nThis must be set to true if you created a backup manually in the console.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "An object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "A unique identifier for the instance, which cannot be changed after\nthe instance is created. The name must be between 6 and 30 characters\nin length.\n\n\nIf not provided, a random string starting with 'tf-' will be selected.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "num_nodes": {
        "computed": true,
        "description": "The number of nodes allocated to this instance. Exactly one of either node_count or processing_units\nmust be present in terraform.",
        "description_kind": "plain",
        "type": "number"
      },
      "processing_units": {
        "computed": true,
        "description": "The number of processing units allocated to this instance. Exactly one of processing_units\nor node_count must be present in terraform.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSpannerInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerInstance), &result)
	return &result
}
