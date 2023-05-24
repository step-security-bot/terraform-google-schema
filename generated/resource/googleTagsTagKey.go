package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTagsTagKey = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Creation time.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-assigned description of the TagKey. Must not exceed 256 characters.",
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
        "description": "The generated numeric id for the TagKey.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespaced_name": {
        "computed": true,
        "description": "Output only. Namespaced name of the TagKey.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id} or projects/{project_id_or_number}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "purpose": {
        "description": "Optional. A purpose cannot be changed once set.\n\nA purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. Possible values: [\"GCE_FIREWALL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purpose_data": {
        "description": "Optional. Purpose data cannot be changed once set.\n\nPurpose data corresponds to the policy system that the tag is intended for. For example, the GCE_FIREWALL purpose expects data in the following format: 'network = \"\u003cproject-name\u003e/\u003cvpc-name\u003e\"'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "short_name": {
        "description": "Input only. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.\n\nThe short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleTagsTagKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTagsTagKey), &result)
	return &result
}
