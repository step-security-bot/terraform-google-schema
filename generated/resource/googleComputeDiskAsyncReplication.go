package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeDiskAsyncReplication = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_disk": {
        "description": "Primary disk for asynchronous replication.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "secondary_disk": {
        "block": {
          "attributes": {
            "disk": {
              "description": "Secondary disk for asynchronous replication.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "Output-only. Status of replication on the secondary disk.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Secondary disk for asynchronous replication.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
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

func GoogleComputeDiskAsyncReplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeDiskAsyncReplication), &result)
	return &result
}
