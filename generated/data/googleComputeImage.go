package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeImage = `{
  "block": {
    "attributes": {
      "archive_size_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "creation_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "family": {
        "computed": true,
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
      "image_encryption_key_sha256": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "licenses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk_encryption_key_sha256": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeImage), &result)
	return &result
}
