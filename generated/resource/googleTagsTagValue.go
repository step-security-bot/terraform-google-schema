package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTagsTagValue = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Creation time.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-assigned description of the TagValue. Must not exceed 256 characters.",
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
      "name": {
        "computed": true,
        "description": "The generated numeric id for the TagValue.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespaced_name": {
        "computed": true,
        "description": "Output only. Namespaced name of the TagValue. Will be in the format {parentNamespace}/{tagKeyShortName}/{shortName}.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Input only. The resource name of the new TagValue's parent. Must be of the form tagKeys/{tag_key_id}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "short_name": {
        "description": "Input only. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey.\n\nThe short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleTagsTagValueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTagsTagValue), &result)
	return &result
}
