package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabaseInstances = `{
  "block": {
    "attributes": {
      "database_version": {
        "description": "To filter out the database instances which are of the specified database version.",
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_maintenance_versions": [
                "list",
                "string"
              ],
              "clone": [
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
              ],
              "connection_name": "string",
              "database_version": "string",
              "deletion_protection": "bool",
              "encryption_key_name": "string",
              "first_ip_address": "string",
              "instance_type": "string",
              "ip_address": [
                "list",
                [
                  "object",
                  {
                    "ip_address": "string",
                    "time_to_retire": "string",
                    "type": "string"
                  }
                ]
              ],
              "maintenance_version": "string",
              "master_instance_name": "string",
              "name": "string",
              "private_ip_address": "string",
              "project": "string",
              "public_ip_address": "string",
              "region": "string",
              "replica_configuration": [
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
              ],
              "restore_backup_context": [
                "list",
                [
                  "object",
                  {
                    "backup_run_id": "number",
                    "instance_id": "string",
                    "project": "string"
                  }
                ]
              ],
              "root_password": "string",
              "self_link": "string",
              "server_ca_cert": [
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
              ],
              "service_account_email_address": "string",
              "settings": [
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
          ]
        ]
      },
      "project": {
        "description": "Project ID of the project that contains the instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "To filter out the database instances which are located in this specified region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "description": "To filter out the database instances based on the current state of the database instance, valid values include : \"SQL_INSTANCE_STATE_UNSPECIFIED\", \"RUNNABLE\", \"SUSPENDED\", \"PENDING_DELETE\", \"PENDING_CREATE\", \"MAINTENANCE\" and \"FAILED\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tier": {
        "description": "To filter out the database instances based on the machine type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "description": "To filter out the database instances which are located in this specified zone.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSqlDatabaseInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabaseInstances), &result)
	return &result
}
