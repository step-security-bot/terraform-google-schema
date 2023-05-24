package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSharedflow = `{
  "block": {
    "attributes": {
      "config_bundle": {
        "description": "Path to the config zip bundle",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "detect_md5hash": {
        "description": "A hash of local config bundle in string, user needs to use a Terraform Hash function of their choice. A change in hash will trigger an update.",
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
      "latest_revision_id": {
        "computed": true,
        "description": "The id of the most recently created revision for this shared flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "md5hash": {
        "computed": true,
        "description": "Base 64 MD5 hash of the uploaded config bundle.",
        "description_kind": "plain",
        "type": "string"
      },
      "meta_data": {
        "computed": true,
        "description": "Metadata describing the shared flow.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "last_modified_at": "string",
              "sub_type": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The ID of the shared flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization name associated with the Apigee instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision": {
        "computed": true,
        "description": "A list of revisions of this shared flow.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func GoogleApigeeSharedflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSharedflow), &result)
	return &result
}
