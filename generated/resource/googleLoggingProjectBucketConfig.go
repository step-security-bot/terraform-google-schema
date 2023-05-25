package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingProjectBucketConfig = `{
  "block": {
    "attributes": {
      "bucket_id": {
        "description": "The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description for this bucket.",
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
        "description": "The bucket's lifecycle such as active or deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the bucket",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "The parent project that contains the logging bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_days": {
        "description": "Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingProjectBucketConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingProjectBucketConfig), &result)
	return &result
}
