package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionDisk = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this disk.  A list of key-\u003evalue pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_attach_timestamp": {
        "computed": true,
        "description": "Last attach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_detach_timestamp": {
        "computed": true,
        "description": "Last detach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_block_size_bytes": {
        "computed": true,
        "description": "Physical block size of the persistent disk, in bytes. If not present\nin a request, a default value is used. Currently supported sizes\nare 4096 and 16384, other sizes may be added in the future.\nIf an unsupported value is requested, the error message will list\nthe supported values for the caller's project.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "A reference to the region where the disk resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_zones": {
        "description": "URLs of the zones where the disk should be replicated to.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "Size of the persistent disk, specified in GB. You can specify this\nfield when creating a persistent disk using the sourceImage or\nsourceSnapshot parameter, or specify it alone to create an empty\npersistent disk.\n\nIf you specify this field along with sourceImage or sourceSnapshot,\nthe value of sizeGb must not be less than the size of the sourceImage\nor the size of the snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot": {
        "description": "The source snapshot used to create this disk. You can provide this as\na partial or full URL to the resource. For example, the following are\nvalid values:\n\n* 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot'\n* 'projects/project/global/snapshots/snapshot'\n* 'global/snapshots/snapshot'\n* 'snapshot'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_snapshot_id": {
        "computed": true,
        "description": "The unique ID of the snapshot used to create this disk. This value\nidentifies the exact snapshot that was used to create this persistent\ndisk. For example, if you created the persistent disk from a snapshot\nthat was later deleted and recreated under the same name, the source\nsnapshot ID would identify the exact version of the snapshot that was\nused.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "URL of the disk type resource describing which disk type to use to\ncreate the disk. Provide this when creating the disk.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "Links to the users of the disk (attached instances) in form:\nproject/zones/zone/instances/instance",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "disk_encryption_key": {
        "block": {
          "attributes": {
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
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
      "source_snapshot_encryption_key": {
        "block": {
          "attributes": {
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
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

func GoogleComputeRegionDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionDisk), &result)
	return &result
}
