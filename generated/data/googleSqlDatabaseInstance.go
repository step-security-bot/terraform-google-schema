package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabaseInstance = `{
  "block": {
    "attributes": {
      "available_maintenance_versions": {
        "computed": true,
        "description": "Available Maintenance versions.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "clone": {
        "computed": true,
        "description": "Configuration for creating a new instance as a clone of another instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocated_ip_range": "string",
              "database_names": [
                "list",
                "string"
              ],
              "point_in_time": "string",
              "source_instance_name": "string"
            }
          ]
        ]
      },
      "connection_name": {
        "computed": true,
        "description": "The connection name of the instance to be used in connection strings. For example, when connecting with Cloud SQL Proxy.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_version": {
        "computed": true,
        "description": "The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database Version Policies includes an up-to-date reference of supported versions.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Used to block Terraform from deleting a SQL Instance. Defaults to true.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption_key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "instance_type": {
        "computed": true,
        "description": "The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.",
        "description_kind": "plain",
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
      "maintenance_version": {
        "computed": true,
        "description": "Maintenance version.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_instance_name": {
        "computed": true,
        "description": "The name of the instance that will act as the master in the replication setup. Note, this requires the master to have binary_log_enabled set, as well as existing backups.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "IPv4 address assigned. This is a workaround for an issue fixed in Terraform 0.12 but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
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
        "type": "string"
      },
      "replica_configuration": {
        "computed": true,
        "description": "The configuration for replication.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ca_certificate": "string",
              "client_certificate": "string",
              "client_key": "string",
              "connect_retry_interval": "number",
              "dump_file_path": "string",
              "failover_target": "bool",
              "master_heartbeat_period": "number",
              "password": "string",
              "ssl_cipher": "string",
              "username": "string",
              "verify_server_certificate": "bool"
            }
          ]
        ]
      },
      "restore_backup_context": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_run_id": "number",
              "instance_id": "string",
              "project": "string"
            }
          ]
        ]
      },
      "root_password": {
        "computed": true,
        "description": "Initial root password. Required for MS SQL Server.",
        "description_kind": "plain",
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
      },
      "settings": {
        "computed": true,
        "description": "The settings to use for the database. The configuration is detailed below.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "activation_policy": "string",
              "active_directory_config": [
                "list",
                [
                  "object",
                  {
                    "domain": "string"
                  }
                ]
              ],
              "advanced_machine_features": [
                "list",
                [
                  "object",
                  {
                    "threads_per_core": "number"
                  }
                ]
              ],
              "availability_type": "string",
              "backup_configuration": [
                "list",
                [
                  "object",
                  {
                    "backup_retention_settings": [
                      "list",
                      [
                        "object",
                        {
                          "retained_backups": "number",
                          "retention_unit": "string"
                        }
                      ]
                    ],
                    "binary_log_enabled": "bool",
                    "enabled": "bool",
                    "location": "string",
                    "point_in_time_recovery_enabled": "bool",
                    "start_time": "string",
                    "transaction_log_retention_days": "number"
                  }
                ]
              ],
              "collation": "string",
              "connector_enforcement": "string",
              "data_cache_config": [
                "list",
                [
                  "object",
                  {
                    "data_cache_enabled": "bool"
                  }
                ]
              ],
              "database_flags": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "deletion_protection_enabled": "bool",
              "deny_maintenance_period": [
                "list",
                [
                  "object",
                  {
                    "end_date": "string",
                    "start_date": "string",
                    "time": "string"
                  }
                ]
              ],
              "disk_autoresize": "bool",
              "disk_autoresize_limit": "number",
              "disk_size": "number",
              "disk_type": "string",
              "edition": "string",
              "insights_config": [
                "list",
                [
                  "object",
                  {
                    "query_insights_enabled": "bool",
                    "query_plans_per_minute": "number",
                    "query_string_length": "number",
                    "record_application_tags": "bool",
                    "record_client_address": "bool"
                  }
                ]
              ],
              "ip_configuration": [
                "list",
                [
                  "object",
                  {
                    "allocated_ip_range": "string",
                    "authorized_networks": [
                      "set",
                      [
                        "object",
                        {
                          "expiration_time": "string",
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "enable_private_path_for_google_cloud_services": "bool",
                    "ipv4_enabled": "bool",
                    "private_network": "string",
                    "require_ssl": "bool"
                  }
                ]
              ],
              "location_preference": [
                "list",
                [
                  "object",
                  {
                    "follow_gae_application": "string",
                    "secondary_zone": "string",
                    "zone": "string"
                  }
                ]
              ],
              "maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "day": "number",
                    "hour": "number",
                    "update_track": "string"
                  }
                ]
              ],
              "password_validation_policy": [
                "list",
                [
                  "object",
                  {
                    "complexity": "string",
                    "disallow_username_substring": "bool",
                    "enable_password_policy": "bool",
                    "min_length": "number",
                    "password_change_interval": "string",
                    "reuse_interval": "number"
                  }
                ]
              ],
              "pricing_plan": "string",
              "sql_server_audit_config": [
                "list",
                [
                  "object",
                  {
                    "bucket": "string",
                    "retention_interval": "string",
                    "upload_interval": "string"
                  }
                ]
              ],
              "tier": "string",
              "time_zone": "string",
              "user_labels": [
                "map",
                "string"
              ],
              "version": "number"
            }
          ]
        ]
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
