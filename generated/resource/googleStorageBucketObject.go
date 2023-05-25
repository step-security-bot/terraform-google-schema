package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketObject = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cache_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "content_disposition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "crc32c": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detect_md5hash": {
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
      "md5hash": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "predefined_acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageBucketObjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucketObject), &result)
	return &result
}
