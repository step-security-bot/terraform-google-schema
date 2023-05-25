package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketObject = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_control": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_disposition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_encoding": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "crc32c": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detect_md5hash": {
        "computed": true,
        "description_kind": "plain",
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
        "optional": true,
        "type": "string"
      },
      "output_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "predefined_acl": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description_kind": "plain",
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
