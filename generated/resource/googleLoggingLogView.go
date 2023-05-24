package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingLogView = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The bucket of the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The creation timestamp of the view.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Describes this view.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter": {
        "description": "Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE(\"projects/myproject\") AND resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")",
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
      "location": {
        "computed": true,
        "description": "The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the view. For example: \\'projects/my-project/locations/global/buckets/my-bucket/views/my-view\\'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description": "The parent of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last update timestamp of the view.",
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

func GoogleLoggingLogViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingLogView), &result)
	return &result
}
