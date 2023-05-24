package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingLinkedDataset = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The bucket to which the linked dataset is attached.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The creation timestamp of the link. A timestamp in RFC3339 UTC \"Zulu\" format,\nwith nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\"\nand \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Describes this link. The maximum length of the description is 8000 characters.",
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
      "lifecycle_state": {
        "computed": true,
        "description": "Output only. The linked dataset lifecycle state.",
        "description_kind": "plain",
        "type": "string"
      },
      "link_id": {
        "description": "The id of the linked dataset.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The location of the linked dataset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the linked dataset. The name can have up to 100 characters. A valid link id\n(at the end of the link name) must only have alphanumeric characters and underscores within it.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description": "The parent of the linked dataset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "bigquery_dataset": {
        "block": {
          "attributes": {
            "dataset_id": {
              "computed": true,
              "description": "Output only. The full resource name of the BigQuery dataset. The DATASET_ID will match the ID\nof the link, so the link must match the naming restrictions of BigQuery datasets\n(alphanumeric characters and underscores only). The dataset will have a resource path of\n\"bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET_ID]\"",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along\nwith it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery\nViews corresponding to the LogViews in the bucket.",
          "description_kind": "plain"
        },
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

func GoogleLoggingLinkedDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingLinkedDataset), &result)
	return &result
}
