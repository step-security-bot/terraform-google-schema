package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbBackup = `{
  "block": {
    "attributes": {
      "backup_id": {
        "description": "The ID of the alloydb backup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_name": {
        "description": "The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the Backup was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description of the backup.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_info": {
        "computed": true,
        "description": "EncryptionInfo describes the encryption information of a cluster or a backup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "encryption_type": "string",
              "kms_key_versions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "etag": {
        "computed": true,
        "description": "A hash of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the alloydb backup.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the alloydb backup should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the Backup was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).",
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

func GoogleAlloydbBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbBackup), &result)
	return &result
}
