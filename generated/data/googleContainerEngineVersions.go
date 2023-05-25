package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerEngineVersions = `{
  "block": {
    "attributes": {
      "default_cluster_version": {
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
      "latest_master_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest_node_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "valid_master_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "valid_node_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "version_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleContainerEngineVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerEngineVersions), &result)
	return &result
}
