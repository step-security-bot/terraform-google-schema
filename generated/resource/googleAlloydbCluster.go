package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbCluster = `{
  "block": {
    "attributes": {
      "backup_source": {
        "computed": true,
        "description": "Cluster created from backup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_name": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "description": "The ID of the alloydb cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "continuous_backup_info": {
        "computed": true,
        "description": "ContinuousBackupInfo describes the continuous backup properties of a cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "earliest_restorable_time": "string",
              "enabled_time": "string",
              "encryption_info": [
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
              ],
              "schedule": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "database_version": {
        "computed": true,
        "description": "The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "User-settable and human-readable display name for the Cluster.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the alloydb cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the alloydb cluster should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "migration_source": {
        "computed": true,
        "description": "Cluster created via DMS migration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host_port": "string",
              "reference_id": "string",
              "source_type": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the cluster resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "description": "The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:\n\n\"projects/{projectNumber}/global/networks/{network_id}\".",
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
      "uid": {
        "computed": true,
        "description": "The system-generated UID of the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "automated_backup_policy": {
        "block": {
          "attributes": {
            "backup_window": {
              "computed": true,
              "description": "The length of the time window during which a backup can be taken. If a backup does not succeed within this time window, it will be canceled and considered failed.\n\nThe backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description": "Whether automated backups are enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "labels": {
              "description": "Labels to apply to backups created using this configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "location": {
              "computed": true,
              "description": "The location where the backup will be stored. Currently, the only supported option is to store the backup in the same region as the cluster.",
              "description_kind": "plain",
              "optional": true,
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
            "quantity_based_retention": {
              "block": {
                "attributes": {
                  "count": {
                    "description": "The number of backups to retain.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Quantity-based Backup retention policy to retain recent backups. Conflicts with 'time_based_retention', both can't be set together.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "time_based_retention": {
              "block": {
                "attributes": {
                  "retention_period": {
                    "description": "The retention period.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Time-based Backup retention policy. Conflicts with 'quantity_based_retention', both can't be set together.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "weekly_schedule": {
              "block": {
                "attributes": {
                  "days_of_week": {
                    "description": "The days of the week to perform a backup. At least one day of the week must be provided. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "start_times": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "description": "Minutes of hour of day. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "nanos": {
                          "description": "Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Seconds of minutes of the time. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Weekly schedule for the Backup.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "continuous_backup_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether continuous backup recovery is enabled. If not set, defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "recovery_window_days": {
              "computed": true,
              "description": "The numbers of days that are eligible to restore from using PITR. To support the entire recovery window, backups and logs are retained for one day more than the recovery window.\n\nIf not set, defaults to 14 days.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
            }
          },
          "description": "The continuous backup config for this cluster.\n\nIf no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
      "initial_user": {
        "block": {
          "attributes": {
            "password": {
              "description": "The initial password for the user.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "user": {
              "description": "The database username.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Initial user to setup during cluster creation.",
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

func GoogleAlloydbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbCluster), &result)
	return &result
}
