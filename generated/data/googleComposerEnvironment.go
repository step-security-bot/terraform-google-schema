package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComposerEnvironment = `{
  "block": {
    "attributes": {
      "config": {
        "computed": true,
        "description": "Configuration parameters for this environment.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "airflow_uri": "string",
              "dag_gcs_prefix": "string",
              "database_config": [
                "list",
                [
                  "object",
                  {
                    "machine_type": "string"
                  }
                ]
              ],
              "encryption_config": [
                "list",
                [
                  "object",
                  {
                    "kms_key_name": "string"
                  }
                ]
              ],
              "environment_size": "string",
              "gke_cluster": "string",
              "maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "end_time": "string",
                    "recurrence": "string",
                    "start_time": "string"
                  }
                ]
              ],
              "master_authorized_networks_config": [
                "list",
                [
                  "object",
                  {
                    "cidr_blocks": [
                      "set",
                      [
                        "object",
                        {
                          "cidr_block": "string",
                          "display_name": "string"
                        }
                      ]
                    ],
                    "enabled": "bool"
                  }
                ]
              ],
              "node_config": [
                "list",
                [
                  "object",
                  {
                    "disk_size_gb": "number",
                    "enable_ip_masq_agent": "bool",
                    "ip_allocation_policy": [
                      "list",
                      [
                        "object",
                        {
                          "cluster_ipv4_cidr_block": "string",
                          "cluster_secondary_range_name": "string",
                          "services_ipv4_cidr_block": "string",
                          "services_secondary_range_name": "string",
                          "use_ip_aliases": "bool"
                        }
                      ]
                    ],
                    "machine_type": "string",
                    "network": "string",
                    "oauth_scopes": [
                      "set",
                      "string"
                    ],
                    "service_account": "string",
                    "subnetwork": "string",
                    "tags": [
                      "set",
                      "string"
                    ],
                    "zone": "string"
                  }
                ]
              ],
              "node_count": "number",
              "private_environment_config": [
                "list",
                [
                  "object",
                  {
                    "cloud_composer_connection_subnetwork": "string",
                    "cloud_composer_network_ipv4_cidr_block": "string",
                    "cloud_sql_ipv4_cidr_block": "string",
                    "connection_type": "string",
                    "enable_private_endpoint": "bool",
                    "enable_privately_used_public_ips": "bool",
                    "master_ipv4_cidr_block": "string",
                    "web_server_ipv4_cidr_block": "string"
                  }
                ]
              ],
              "recovery_config": [
                "list",
                [
                  "object",
                  {
                    "scheduled_snapshots_config": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "snapshot_creation_schedule": "string",
                          "snapshot_location": "string",
                          "time_zone": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "resilience_mode": "string",
              "software_config": [
                "list",
                [
                  "object",
                  {
                    "airflow_config_overrides": [
                      "map",
                      "string"
                    ],
                    "env_variables": [
                      "map",
                      "string"
                    ],
                    "image_version": "string",
                    "pypi_packages": [
                      "map",
                      "string"
                    ],
                    "python_version": "string",
                    "scheduler_count": "number"
                  }
                ]
              ],
              "web_server_config": [
                "list",
                [
                  "object",
                  {
                    "machine_type": "string"
                  }
                ]
              ],
              "web_server_network_access_control": [
                "list",
                [
                  "object",
                  {
                    "allowed_ip_range": [
                      "set",
                      [
                        "object",
                        {
                          "description": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "workloads_config": [
                "list",
                [
                  "object",
                  {
                    "scheduler": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "cpu": "number",
                          "memory_gb": "number",
                          "storage_gb": "number"
                        }
                      ]
                    ],
                    "web_server": [
                      "list",
                      [
                        "object",
                        {
                          "cpu": "number",
                          "memory_gb": "number",
                          "storage_gb": "number"
                        }
                      ]
                    ],
                    "worker": [
                      "list",
                      [
                        "object",
                        {
                          "cpu": "number",
                          "max_count": "number",
                          "memory_gb": "number",
                          "min_count": "number",
                          "storage_gb": "number"
                        }
                      ]
                    ]
                  }
                ]
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
        "computed": true,
        "description": "User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63 characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be associated with a given environment. Both keys and values must be \u003c= 128 bytes in size.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The location or Compute Engine region for the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComposerEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComposerEnvironment), &result)
	return &result
}
