package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerCluster = `{
  "block": {
    "attributes": {
      "addons_config": {
        "computed": true,
        "description": "The configuration for addons supported by GKE.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloudrun_config": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool",
                    "load_balancer_type": "string"
                  }
                ]
              ],
              "config_connector_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "dns_cache_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "gce_persistent_disk_csi_driver_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "gcp_filestore_csi_driver_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "gke_backup_agent_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
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
      "authenticator_groups_config": {
        "computed": true,
        "description": "Configuration for the Google Groups for GKE feature.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "security_group": "string"
            }
          ]
        ]
      },
      "binary_authorization": {
        "computed": true,
        "description": "Configuration options for the Binary Authorization feature.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "evaluation_mode": "string"
            }
          ]
        ]
      },
      "cluster_autoscaling": {
        "computed": true,
        "description": "Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_provisioning_defaults": [
                "list",
                [
                  "object",
                  {
                    "boot_disk_kms_key": "string",
                    "disk_size": "number",
                    "disk_type": "string",
                    "image_type": "string",
                    "management": [
                      "list",
                      [
                        "object",
                        {
                          "auto_repair": "bool",
                          "auto_upgrade": "bool",
                          "upgrade_options": [
                            "list",
                            [
                              "object",
                              {
                                "auto_upgrade_start_time": "string",
                                "description": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "min_cpu_platform": "string",
                    "oauth_scopes": [
                      "list",
                      "string"
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
                    "upgrade_settings": [
                      "list",
                      [
                        "object",
                        {
                          "blue_green_settings": [
                            "list",
                            [
                              "object",
                              {
                                "node_pool_soak_duration": "string",
                                "standard_rollout_policy": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "batch_node_count": "number",
                                      "batch_percentage": "number",
                                      "batch_soak_duration": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "max_surge": "number",
                          "max_unavailable": "number",
                          "strategy": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
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
        "description": "The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined.",
        "description_kind": "plain",
        "type": "string"
      },
      "confidential_nodes": {
        "computed": true,
        "description": "Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster.",
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
      "cost_management_config": {
        "computed": true,
        "description": "Cost management configuration for the cluster.",
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
      "database_encryption": {
        "computed": true,
        "description": "Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: \"ENCRYPTED\"; \"DECRYPTED\". key_name is the name of a CloudKMS key.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key_name": "string",
              "state": "string"
            }
          ]
        ]
      },
      "datapath_provider": {
        "computed": true,
        "description": "The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_max_pods_per_node": {
        "computed": true,
        "description": "The default maximum number of pods per node in this cluster. This doesn't work on \"routes-based\" clusters, clusters that don't have IP Aliasing enabled.",
        "description_kind": "plain",
        "type": "number"
      },
      "default_snat_status": {
        "computed": true,
        "description": "Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disabled": "bool"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": " Description of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_config": {
        "computed": true,
        "description": "Configuration for Cloud DNS for Kubernetes Engine.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_dns": "string",
              "cluster_dns_domain": "string",
              "cluster_dns_scope": "string"
            }
          ]
        ]
      },
      "enable_autopilot": {
        "computed": true,
        "description": "Enable Autopilot for this cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_binary_authorization": {
        "computed": true,
        "description": "Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_intranode_visibility": {
        "computed": true,
        "description": "Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_kubernetes_alpha": {
        "computed": true,
        "description": "Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_l4_ilb_subsetting": {
        "computed": true,
        "description": "Whether L4ILB Subsetting is enabled for this cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_legacy_abac": {
        "computed": true,
        "description": "Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_shielded_nodes": {
        "computed": true,
        "description": "Enable Shielded Nodes features on all nodes in this cluster. Defaults to true.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_tpu": {
        "computed": true,
        "description": "Whether to enable Cloud TPU resources in this cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The IP address of this cluster's Kubernetes master.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_api_config": {
        "computed": true,
        "description": "Configuration for GKE Gateway API controller.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "channel": "string"
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
      "initial_node_count": {
        "computed": true,
        "description": "The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true.",
        "description_kind": "plain",
        "type": "number"
      },
      "ip_allocation_policy": {
        "computed": true,
        "description": "Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_ipv4_cidr_block": "string",
              "cluster_secondary_range_name": "string",
              "pod_cidr_overprovision_config": [
                "list",
                [
                  "object",
                  {
                    "disabled": "bool"
                  }
                ]
              ],
              "services_ipv4_cidr_block": "string",
              "services_secondary_range_name": "string",
              "stack_type": "string"
            }
          ]
        ]
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint of the set of labels for this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_config": {
        "computed": true,
        "description": "Logging configuration for the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_components": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "logging_service": {
        "computed": true,
        "description": "The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_policy": {
        "computed": true,
        "description": "The maintenance policy to use for the cluster.",
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
              ],
              "maintenance_exclusion": [
                "set",
                [
                  "object",
                  {
                    "end_time": "string",
                    "exclusion_name": "string",
                    "exclusion_options": [
                      "list",
                      [
                        "object",
                        {
                          "scope": "string"
                        }
                      ]
                    ],
                    "start_time": "string"
                  }
                ]
              ],
              "recurring_window": [
                "list",
                [
                  "object",
                  {
                    "end_time": "string",
                    "recurrence": "string",
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
        "description": "The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials permission.",
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
              "cluster_ca_certificate": "string"
            }
          ]
        ]
      },
      "master_authorized_networks_config": {
        "computed": true,
        "description": "The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists).",
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
              ],
              "gcp_public_cidrs_access_enabled": "bool"
            }
          ]
        ]
      },
      "master_version": {
        "computed": true,
        "description": "The current version of the master in the cluster. This may be different than the min_master_version set in the config if the master has been updated by GKE.",
        "description_kind": "plain",
        "type": "string"
      },
      "mesh_certificates": {
        "computed": true,
        "description": "If set, and enable_certificates=true, the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_certificates": "bool"
            }
          ]
        ]
      },
      "min_master_version": {
        "computed": true,
        "description": "The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version).",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_config": {
        "computed": true,
        "description": "Monitoring configuration for the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_components": [
                "list",
                "string"
              ],
              "managed_prometheus": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "monitoring_service": {
        "computed": true,
        "description": "The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the cluster, unique within the project and location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The name or self_link of the Google Compute Engine network to which the cluster is connected. For Shared VPC, set this to the self link of the shared network.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_policy": {
        "computed": true,
        "description": "Configuration options for the NetworkPolicy feature.",
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
      "networking_mode": {
        "computed": true,
        "description": "Determines whether alias IPs or routes will be used for pod IPs in the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_config": {
        "computed": true,
        "description": "The configuration of the nodepool",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "advanced_machine_features": [
                "list",
                [
                  "object",
                  {
                    "threads_per_core": "number"
                  }
                ]
              ],
              "boot_disk_kms_key": "string",
              "disk_size_gb": "number",
              "disk_type": "string",
              "ephemeral_storage_local_ssd_config": [
                "list",
                [
                  "object",
                  {
                    "local_ssd_count": "number"
                  }
                ]
              ],
              "gcfs_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "guest_accelerator": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "gpu_partition_size": "string",
                    "gpu_sharing_config": [
                      "list",
                      [
                        "object",
                        {
                          "gpu_sharing_strategy": "string",
                          "max_shared_clients_per_gpu": "number"
                        }
                      ]
                    ],
                    "type": "string"
                  }
                ]
              ],
              "gvnic": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "image_type": "string",
              "kubelet_config": [
                "list",
                [
                  "object",
                  {
                    "cpu_cfs_quota": "bool",
                    "cpu_cfs_quota_period": "string",
                    "cpu_manager_policy": "string",
                    "pod_pids_limit": "number"
                  }
                ]
              ],
              "labels": [
                "map",
                "string"
              ],
              "linux_node_config": [
                "list",
                [
                  "object",
                  {
                    "sysctls": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "local_nvme_ssd_block_config": [
                "list",
                [
                  "object",
                  {
                    "local_ssd_count": "number"
                  }
                ]
              ],
              "local_ssd_count": "number",
              "logging_variant": "string",
              "machine_type": "string",
              "metadata": [
                "map",
                "string"
              ],
              "min_cpu_platform": "string",
              "node_group": "string",
              "oauth_scopes": [
                "set",
                "string"
              ],
              "preemptible": "bool",
              "reservation_affinity": [
                "list",
                [
                  "object",
                  {
                    "consume_reservation_type": "string",
                    "key": "string",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "resource_labels": [
                "map",
                "string"
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
              "spot": "bool",
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
                    "mode": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "node_locations": {
        "computed": true,
        "description": "The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "node_pool": {
        "computed": true,
        "description": "List of node pools associated with this cluster. See google_container_node_pool for schema. Warning: node pools defined inside a cluster can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability to say \"these are the only node pools associated with this cluster\", use the google_container_node_pool resource instead of this property.",
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
                    "location_policy": "string",
                    "max_node_count": "number",
                    "min_node_count": "number",
                    "total_max_node_count": "number",
                    "total_min_node_count": "number"
                  }
                ]
              ],
              "initial_node_count": "number",
              "instance_group_urls": [
                "list",
                "string"
              ],
              "managed_instance_group_urls": [
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
              "network_config": [
                "list",
                [
                  "object",
                  {
                    "create_pod_range": "bool",
                    "enable_private_nodes": "bool",
                    "pod_cidr_overprovision_config": [
                      "list",
                      [
                        "object",
                        {
                          "disabled": "bool"
                        }
                      ]
                    ],
                    "pod_ipv4_cidr_block": "string",
                    "pod_range": "string"
                  }
                ]
              ],
              "node_config": [
                "list",
                [
                  "object",
                  {
                    "advanced_machine_features": [
                      "list",
                      [
                        "object",
                        {
                          "threads_per_core": "number"
                        }
                      ]
                    ],
                    "boot_disk_kms_key": "string",
                    "disk_size_gb": "number",
                    "disk_type": "string",
                    "ephemeral_storage_local_ssd_config": [
                      "list",
                      [
                        "object",
                        {
                          "local_ssd_count": "number"
                        }
                      ]
                    ],
                    "gcfs_config": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool"
                        }
                      ]
                    ],
                    "guest_accelerator": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "gpu_partition_size": "string",
                          "gpu_sharing_config": [
                            "list",
                            [
                              "object",
                              {
                                "gpu_sharing_strategy": "string",
                                "max_shared_clients_per_gpu": "number"
                              }
                            ]
                          ],
                          "type": "string"
                        }
                      ]
                    ],
                    "gvnic": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool"
                        }
                      ]
                    ],
                    "image_type": "string",
                    "kubelet_config": [
                      "list",
                      [
                        "object",
                        {
                          "cpu_cfs_quota": "bool",
                          "cpu_cfs_quota_period": "string",
                          "cpu_manager_policy": "string",
                          "pod_pids_limit": "number"
                        }
                      ]
                    ],
                    "labels": [
                      "map",
                      "string"
                    ],
                    "linux_node_config": [
                      "list",
                      [
                        "object",
                        {
                          "sysctls": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "local_nvme_ssd_block_config": [
                      "list",
                      [
                        "object",
                        {
                          "local_ssd_count": "number"
                        }
                      ]
                    ],
                    "local_ssd_count": "number",
                    "logging_variant": "string",
                    "machine_type": "string",
                    "metadata": [
                      "map",
                      "string"
                    ],
                    "min_cpu_platform": "string",
                    "node_group": "string",
                    "oauth_scopes": [
                      "set",
                      "string"
                    ],
                    "preemptible": "bool",
                    "reservation_affinity": [
                      "list",
                      [
                        "object",
                        {
                          "consume_reservation_type": "string",
                          "key": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "resource_labels": [
                      "map",
                      "string"
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
                    "spot": "bool",
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
                          "mode": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "node_count": "number",
              "node_locations": [
                "set",
                "string"
              ],
              "placement_policy": [
                "list",
                [
                  "object",
                  {
                    "type": "string"
                  }
                ]
              ],
              "upgrade_settings": [
                "list",
                [
                  "object",
                  {
                    "blue_green_settings": [
                      "list",
                      [
                        "object",
                        {
                          "node_pool_soak_duration": "string",
                          "standard_rollout_policy": [
                            "list",
                            [
                              "object",
                              {
                                "batch_node_count": "number",
                                "batch_percentage": "number",
                                "batch_soak_duration": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "max_surge": "number",
                    "max_unavailable": "number",
                    "strategy": "string"
                  }
                ]
              ],
              "version": "string"
            }
          ]
        ]
      },
      "node_pool_defaults": {
        "computed": true,
        "description": "The default nodel pool settings for the entire cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "node_config_defaults": [
                "list",
                [
                  "object",
                  {
                    "logging_variant": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "node_version": {
        "computed": true,
        "description": "The Kubernetes version on the nodes. Must either be unset or set to the same value as min_master_version on create. Defaults to the default version set by GKE which is not necessarily the latest version. This only affects nodes in the default node pool. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way. To update nodes in other node pools, use the version attribute on the node pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_config": {
        "computed": true,
        "description": "The notification config for sending cluster upgrade notifications",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pubsub": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "filter": [
                      "list",
                      [
                        "object",
                        {
                          "event_type": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "topic": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_cluster_config": {
        "computed": true,
        "description": "Configuration for private clusters, clusters with private nodes.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_private_endpoint": "bool",
              "enable_private_nodes": "bool",
              "master_global_access_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "master_ipv4_cidr_block": "string",
              "peering_name": "string",
              "private_endpoint": "string",
              "private_endpoint_subnetwork": "string",
              "public_endpoint": "string"
            }
          ]
        ]
      },
      "private_ipv6_google_access": {
        "computed": true,
        "description": "The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4).",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_channel": {
        "computed": true,
        "description": "Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the \"UNSPECIFIED\" channel.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "channel": "string"
            }
          ]
        ]
      },
      "remove_default_node_pool": {
        "computed": true,
        "description": "If true, deletes the default node pool upon cluster creation. If you're using google_container_node_pool resources with no default node pool, this should be set to true, alongside setting initial_node_count to at least 1.",
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_labels": {
        "computed": true,
        "description": "The GCE resource labels (a map of key/value pairs) to be applied to the cluster.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "resource_usage_export_config": {
        "computed": true,
        "description": "Configuration for the ResourceUsageExportConfig feature.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bigquery_destination": [
                "list",
                [
                  "object",
                  {
                    "dataset_id": "string"
                  }
                ]
              ],
              "enable_network_egress_metering": "bool",
              "enable_resource_consumption_metering": "bool"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_external_ips_config": {
        "computed": true,
        "description": "If set, and enabled=true, services with external ips field will not be blocked",
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
      "services_ipv4_cidr": {
        "computed": true,
        "description": "The IP address range of the Kubernetes services in this cluster, in CIDR notation (e.g. 1.2.3.4/29). Service addresses are typically put in the last /16 from the container CIDR.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "The name or self_link of the Google Compute Engine subnetwork in which the cluster's instances are launched.",
        "description_kind": "plain",
        "type": "string"
      },
      "tpu_ipv4_cidr_block": {
        "computed": true,
        "description": "The IP address range of the Cloud TPUs in this cluster, in CIDR notation (e.g. 1.2.3.4/29).",
        "description_kind": "plain",
        "type": "string"
      },
      "vertical_pod_autoscaling": {
        "computed": true,
        "description": "Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.",
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
      "workload_identity_config": {
        "computed": true,
        "description": "Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "workload_pool": "string"
            }
          ]
        ]
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
