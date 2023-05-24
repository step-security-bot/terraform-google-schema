package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeBackupBackupPlan = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The source cluster from which Backups will be created via this BackupPlan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deactivated": {
        "computed": true,
        "description": "This flag indicates whether this BackupPlan has been deactivated.\nSetting this field to True locks the BackupPlan such that no further updates will be allowed\n(except deletes), including the deactivated field itself. It also prevents any new Backups\nfrom being created via this BackupPlan (including scheduled Backups).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "User specified descriptive string for this BackupPlan.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "etag is used for optimistic concurrency control as a way to help prevent simultaneous\nupdates of a backup plan from overwriting each other. It is strongly suggested that\nsystems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates\nin order to avoid race conditions: An etag is returned in the response to backupPlans.get,\nand systems are expected to put that etag in the request to backupPlans.patch or\nbackupPlans.delete to ensure that their change will be applied to the same version of the resource.",
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
        "description": "Description: A set of custom labels supplied by the user.\nA list of key-\u003evalue pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The region of the Backup Plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full name of the BackupPlan Resource.",
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
      "protected_pod_count": {
        "computed": true,
        "description": "The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.",
        "description_kind": "plain",
        "type": "number"
      },
      "uid": {
        "computed": true,
        "description": "Server generated, unique identifier of UUID format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "backup_config": {
        "block": {
          "attributes": {
            "all_namespaces": {
              "description": "If True, include all namespaced resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_secrets": {
              "computed": true,
              "description": "This flag specifies whether Kubernetes Secret resources should be included\nwhen they fall into the scope of Backups.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_volume_data": {
              "computed": true,
              "description": "This flag specifies whether volume data should be backed up when PVCs are\nincluded in the scope of a Backup.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "encryption_key": {
              "block": {
                "attributes": {
                  "gcp_kms_encryption_key": {
                    "description": "Google Cloud KMS encryption key. Format: projects/*/locations/*/keyRings/*/cryptoKeys/*",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "This defines a customer managed encryption key that will be used to encrypt the \"config\"\nportion (the Kubernetes resources) of Backups created via this plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_applications": {
              "block": {
                "block_types": {
                  "namespaced_names": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "namespace": {
                          "description": "The namespace of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of namespaced Kubernetes resources.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of namespaced Kubernetes Resources.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_namespaces": {
              "block": {
                "attributes": {
                  "namespaces": {
                    "description": "A list of Kubernetes Namespaces.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "If set, include just the resources in the listed namespaces.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the configuration of Backups created via this BackupPlan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "backup_schedule": {
        "block": {
          "attributes": {
            "cron_schedule": {
              "description": "A standard cron string that defines a repeating schedule for\ncreating Backups via this BackupPlan.\nIf this is defined, then backupRetainDays must also be defined.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "paused": {
              "computed": true,
              "description": "This flag denotes whether automatic Backup creation is paused for this BackupPlan.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Defines a schedule for automatic Backup creation via this BackupPlan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "backup_delete_lock_days": {
              "computed": true,
              "description": "Minimum age for a Backup created via this BackupPlan (in days).\nMust be an integer value between 0-90 (inclusive).\nA Backup created under this BackupPlan will not be deletable\nuntil it reaches Backup's (create time + backup_delete_lock_days).\nUpdating this field of a BackupPlan does not affect existing Backups.\nBackups created after a successful update will inherit this new value.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "backup_retain_days": {
              "computed": true,
              "description": "The default maximum age of a Backup created via this BackupPlan.\nThis field MUST be an integer value \u003e= 0 and \u003c= 365. If specified,\na Backup created under this BackupPlan will be automatically deleted\nafter its age reaches (createTime + backupRetainDays).\nIf not specified, Backups created under this BackupPlan will NOT be\nsubject to automatic deletion. Updating this field does NOT affect\nexisting Backups under it. Backups created AFTER a successful update\nwill automatically pick up the new value.\nNOTE: backupRetainDays must be \u003e= backupDeleteLockDays.\nIf cronSchedule is defined, then this must be \u003c= 360 * the creation interval.]",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "locked": {
              "computed": true,
              "description": "This flag denotes whether the retention policy of this BackupPlan is locked.\nIf set to True, no further update is allowed on this policy, including\nthe locked field itself.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "RetentionPolicy governs lifecycle of Backups created under this plan.",
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

func GoogleGkeBackupBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeBackupBackupPlan), &result)
	return &result
}
