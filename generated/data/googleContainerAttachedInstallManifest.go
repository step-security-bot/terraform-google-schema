package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAttachedInstallManifest = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manifest": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleContainerAttachedInstallManifestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAttachedInstallManifest), &result)
	return &result
}
