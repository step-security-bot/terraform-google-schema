package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSnapshot = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_size_gb": {
        "computed": true,
        "description": "Size of the snapshot, specified in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource. Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this Snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "licenses": {
        "computed": true,
        "description": "A list of public visible licenses that apply to this snapshot. This\ncan be because the original image had licenses attached (such as a\nWindows image).  snapshotEncryptionKey nested object Encrypts the\nsnapshot using a customer-supplied encryption key.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
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
      "snapshot_encryption_key_raw": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "snapshot_encryption_key_sha256": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_disk": {
        "description": "A reference to the disk used to create this snapshot.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_disk_encryption_key_raw": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "source_disk_encryption_key_sha256": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_bytes": {
        "computed": true,
        "description": "A size of the the storage used by the snapshot. As snapshots share\nstorage, this number is expected to change with snapshot\ncreation/deletion.",
        "description_kind": "plain",
        "type": "number"
      },
      "zone": {
        "computed": true,
        "description": "A reference to the zone where the disk is hosted.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "snapshot_encryption_key": {
        "block": {
          "attributes": {
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied\nencryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_disk_encryption_key": {
        "block": {
          "attributes": {
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleComputeSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSnapshot), &result)
	return &result
}
