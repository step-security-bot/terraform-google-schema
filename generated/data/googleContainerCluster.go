package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerCluster = `{
  "block": {
    "attributes": {
      "additional_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "addons_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "horizontal_pod_autoscaling": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool"
                  }
                ]
              ],
              "http_load_balancing": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool"
                  }
                ]
              ],
              "kubernetes_dashboard": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool"
                  }
                ]
              ],
              "network_policy_config": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cluster_autoscaling": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "resource_limits": [
                "list",
                [
                  "object",
                  {
                    "maximum": "number",
                    "minimum": "number",
                    "resource_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cluster_ipv4_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_max_pods_per_node": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_binary_authorization": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_intranode_visibility": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_kubernetes_alpha": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_legacy_abac": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_tpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initial_node_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_group_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ip_allocation_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_ipv4_cidr_block": "string",
              "cluster_secondary_range_name": "string",
              "create_subnetwork": "bool",
              "node_ipv4_cidr_block": "string",
              "services_ipv4_cidr_block": "string",
              "services_secondary_range_name": "string",
              "subnetwork_name": "string",
              "use_ip_aliases": "bool"
            }
          ]
        ]
      },
      "location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_service": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "daily_maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "duration": "string",
                    "start_time": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "master_auth": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_certificate": "string",
              "client_certificate_config": [
                "list",
                [
                  "object",
                  {
                    "issue_client_certificate": "bool"
                  }
                ]
              ],
              "client_key": "string",
              "cluster_ca_certificate": "string",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "master_authorized_networks_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
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
              ]
            }
          ]
        ]
      },
      "master_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "min_master_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_service": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "provider": "string"
            }
          ]
        ]
      },
      "node_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_size_gb": "number",
              "disk_type": "string",
              "guest_accelerator": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "type": "string"
                  }
                ]
              ],
              "image_type": "string",
              "labels": [
                "map",
                "string"
              ],
              "local_ssd_count": "number",
              "machine_type": "string",
              "metadata": [
                "map",
                "string"
              ],
              "min_cpu_platform": "string",
              "oauth_scopes": [
                "set",
                "string"
              ],
              "preemptible": "bool",
              "sandbox_config": [
                "list",
                [
                  "object",
                  {
                    "sandbox_type": "string"
                  }
                ]
              ],
              "service_account": "string",
              "shielded_instance_config": [
                "list",
                [
                  "object",
                  {
                    "enable_integrity_monitoring": "bool",
                    "enable_secure_boot": "bool"
                  }
                ]
              ],
              "tags": [
                "list",
                "string"
              ],
              "taint": [
                "list",
                [
                  "object",
                  {
                    "effect": "string",
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "workload_metadata_config": [
                "list",
                [
                  "object",
                  {
                    "node_metadata": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "node_locations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "node_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "autoscaling": [
                "list",
                [
                  "object",
                  {
                    "max_node_count": "number",
                    "min_node_count": "number"
                  }
                ]
              ],
              "initial_node_count": "number",
              "instance_group_urls": [
                "list",
                "string"
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "auto_repair": "bool",
                    "auto_upgrade": "bool"
                  }
                ]
              ],
              "max_pods_per_node": "number",
              "name": "string",
              "name_prefix": "string",
              "node_config": [
                "list",
                [
                  "object",
                  {
                    "disk_size_gb": "number",
                    "disk_type": "string",
                    "guest_accelerator": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "type": "string"
                        }
                      ]
                    ],
                    "image_type": "string",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "local_ssd_count": "number",
                    "machine_type": "string",
                    "metadata": [
                      "map",
                      "string"
                    ],
                    "min_cpu_platform": "string",
                    "oauth_scopes": [
                      "set",
                      "string"
                    ],
                    "preemptible": "bool",
                    "sandbox_config": [
                      "list",
                      [
                        "object",
                        {
                          "sandbox_type": "string"
                        }
                      ]
                    ],
                    "service_account": "string",
                    "shielded_instance_config": [
                      "list",
                      [
                        "object",
                        {
                          "enable_integrity_monitoring": "bool",
                          "enable_secure_boot": "bool"
                        }
                      ]
                    ],
                    "tags": [
                      "list",
                      "string"
                    ],
                    "taint": [
                      "list",
                      [
                        "object",
                        {
                          "effect": "string",
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "workload_metadata_config": [
                      "list",
                      [
                        "object",
                        {
                          "node_metadata": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "node_count": "number",
              "version": "string"
            }
          ]
        ]
      },
      "node_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pod_security_policy_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "private_cluster_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_private_endpoint": "bool",
              "enable_private_nodes": "bool",
              "master_ipv4_cidr_block": "string",
              "private_endpoint": "string",
              "public_endpoint": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remove_default_node_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "services_ipv4_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleContainerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerCluster), &result)
	return &result
}
