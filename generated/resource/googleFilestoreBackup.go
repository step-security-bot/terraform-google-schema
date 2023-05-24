package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFilestoreBackup = `{
  "block": {
    "attributes": {
      "capacity_gb": {
        "computed": true,
        "description": "The amount of bytes needed to allocate a full copy of the snapshot content.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the snapshot was created in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "download_bytes": {
        "computed": true,
        "description": "Amount of bytes that will be downloaded if the backup is restored.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "KMS key name used for data encryption.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Resource labels to represent user-provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location of the instance. This can be a region for ENTERPRISE tier instances.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the backup. The name must be unique within the specified instance.\n\nThe name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "source_file_share": {
        "description": "Name of the file share in the source Cloud Filestore instance that the backup is created from.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_instance": {
        "description": "The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_instance_tier": {
        "computed": true,
        "description": "The service tier of the source Cloud Filestore instance that this backup is created from.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The backup state.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_bytes": {
        "computed": true,
        "description": "The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.",
        "description_kind": "plain",
        "type": "string"
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

func GoogleFilestoreBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFilestoreBackup), &result)
	return &result
}
