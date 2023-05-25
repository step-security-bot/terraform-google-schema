package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabaseInstance = `{
  "block": {
    "attributes": {
      "connection_name": {
        "computed": true,
        "description": "The connection name of the instance to be used in connection strings. For example, when connecting with Cloud SQL Proxy.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_version": {
        "description": "The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database Version Policies includes an up-to-date reference of supported versions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "description": "Used to block Terraform from deleting a SQL Instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "first_ip_address": {
        "computed": true,
        "description": "The first IPv4 address of any type assigned. This is to support accessing the first address in the list in a terraform output when the resource is configured with a count.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "time_to_retire": "string",
              "type": "string"
            }
          ]
        ]
      },
      "master_instance_name": {
        "computed": true,
        "description": "The name of the instance that will act as the master in the replication setup. Note, this requires the master to have binary_log_enabled set, as well as existing backups.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "IPv4 address assigned. This is a workaround for an issue fixed in Terraform 0.12 but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description": "IPv4 address assigned. This is a workaround for an issue fixed in Terraform 0.12 but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region the instance will sit in. Note, Cloud SQL is not available in all regions. A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for instances if the provider region is not supported with Cloud SQL. If you choose not to provide the region argument for this resource, make sure you understand this.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "root_password": {
        "description": "Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_ca_cert": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert": "string",
              "common_name": "string",
              "create_time": "string",
              "expiration_time": "string",
              "sha1_fingerprint": "string"
            }
          ]
        ]
      },
      "service_account_email_address": {
        "computed": true,
        "description": "The service account email address assigned to the instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "clone": {
        "block": {
          "attributes": {
            "point_in_time": {
              "description": "The timestamp of the point in time that should be restored.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_instance_name": {
              "description": "The name of the instance from which the point in time should be restored.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for creating a new instance as a clone of another instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "replica_configuration": {
        "block": {
          "attributes": {
            "ca_certificate": {
              "description": "PEM representation of the trusted CA's x509 certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_certificate": {
              "description": "PEM representation of the replica's x509 certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_key": {
              "description": "PEM representation of the replica's private key. The corresponding public key in encoded in the client_certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connect_retry_interval": {
              "description": "The number of seconds between connect retries.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dump_file_path": {
              "description": "Path to a SQL file in Google Cloud Storage from which replica instances are created. Format is gs://bucket/filename.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failover_target": {
              "description": "Specifies if the replica is the failover target. If the field is set to true the replica will be designated as a failover replica. If the master instance fails, the replica instance will be promoted as the new master instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "master_heartbeat_period": {
              "description": "Time in ms between replication heartbeats.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "password": {
              "description": "Password for the replication connection.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "ssl_cipher": {
              "description": "Permissible ciphers for use in SSL encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description": "Username for replication connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_server_certificate": {
              "description": "True if the master's common name value is checked during the SSL handshake.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The configuration for replication.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_backup_context": {
        "block": {
          "attributes": {
            "backup_run_id": {
              "description": "The ID of the backup run to restore from.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_id": {
              "description": "The ID of the instance that the backup was taken from.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description": "The full project ID of the source instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "settings": {
        "block": {
          "attributes": {
            "activation_policy": {
              "computed": true,
              "description": "This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authorized_gae_applications": {
              "computed": true,
              "deprecated": true,
              "description": "This property is only applicable to First Generation instances. First Generation instances are now deprecated, see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances. A list of Google App Engine project names that are allowed to access this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "availability_type": {
              "computed": true,
              "description": "The availability type of the Cloud SQL instance, high availability\n(REGIONAL) or single zone (ZONAL). For MySQL instances, ensure that\nsettings.backup_configuration.enabled and\nsettings.backup_configuration.binary_log_enabled are both set to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "collation": {
              "description": "The name of server instance collation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "crash_safe_replication": {
              "computed": true,
              "deprecated": true,
              "description": "This property is only applicable to First Generation instances. First Generation instances are now deprecated, see here for information on how to upgrade to Second Generation instances. Specific to read instances, indicates when crash-safe replication flags are enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_autoresize": {
              "description": "Configuration to increase storage size automatically.  Note that future terraform apply calls will attempt to resize the disk to the value specified in disk_size - if this is set, do not set disk_size.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_autoresize_limit": {
              "description": "The maximum size, in GB, to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_size": {
              "computed": true,
              "description": "The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description": "The type of data disk: PD_SSD or PD_HDD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pricing_plan": {
              "description": "Pricing plan for this instance, can only be PER_USE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "replication_type": {
              "computed": true,
              "deprecated": true,
              "description": "This property is only applicable to First Generation instances. First Generation instances are now deprecated, see here for information on how to upgrade to Second Generation instances. Replication type for this instance, can be one of ASYNCHRONOUS or SYNCHRONOUS.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tier": {
              "description": "The machine type to use. See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_labels": {
              "computed": true,
              "description": "A set of key/value user label pairs to assign to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "computed": true,
              "description": "Used to make sure changes to the settings block are atomic.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "block_types": {
            "backup_configuration": {
              "block": {
                "attributes": {
                  "binary_log_enabled": {
                    "description": "True if binary logging is enabled. If settings.backup_configuration.enabled is false, this must be as well. Cannot be used with Postgres.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description": "True if backup configuration is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "location": {
                    "description": "Location of the backup configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "point_in_time_recovery_enabled": {
                    "description": "True if Point-in-time recovery is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "start_time": {
                    "computed": true,
                    "description": "HH:MM format time indicating when backup configuration starts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transaction_log_retention_days": {
                    "computed": true,
                    "description": "The number of days of transaction logs we retain for point in time restore, from 1-7.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "backup_retention_settings": {
                    "block": {
                      "attributes": {
                        "retained_backups": {
                          "description": "Number of backups to retain.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "retention_unit": {
                          "description": "The unit that 'retainedBackups' represents. Defaults to COUNT",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "database_flags": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the flag.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value of the flag.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "insights_config": {
              "block": {
                "attributes": {
                  "query_insights_enabled": {
                    "description": "True if Query Insights feature is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "query_string_length": {
                    "description": "Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "record_application_tags": {
                    "description": "True if Query Insights will record application tags from query when enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "record_client_address": {
                    "description": "True if Query Insights will record client address when enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration of Query Insights.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ip_configuration": {
              "block": {
                "attributes": {
                  "ipv4_enabled": {
                    "description": "Whether this Cloud SQL instance should be assigned a public IPV4 address. At least ipv4_enabled must be enabled or a private_network must be configured.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "private_network": {
                    "description": "The VPC network from which the Cloud SQL instance is accessible for private IP. For example, projects/myProject/global/networks/default. Specifying a network enables private IP. At least ipv4_enabled must be enabled or a private_network must be configured. This setting can be updated, but it cannot be removed after it is set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_ssl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "authorized_networks": {
                    "block": {
                      "attributes": {
                        "expiration_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "location_preference": {
              "block": {
                "attributes": {
                  "follow_gae_application": {
                    "description": "A Google App Engine application whose zone to remain in. Must be in the same region as this instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zone": {
                    "description": "The preferred compute engine zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of week (1-7), starting on Monday",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "hour": {
                    "description": "Hour of day (0-23), ignored if day not set",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "update_track": {
                    "description": "Receive updates earlier (canary) or later (stable)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Declares a one-hour maintenance window when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The settings to use for the database. The configuration is detailed below.",
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

func GoogleSqlDatabaseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabaseInstance), &result)
	return &result
}
