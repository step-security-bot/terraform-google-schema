package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageObjectSignedUrl = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_md5": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extension_headers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "http_method": {
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
      "path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signed_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageObjectSignedUrlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageObjectSignedUrl), &result)
	return &result
}
