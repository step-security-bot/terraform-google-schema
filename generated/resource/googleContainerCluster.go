package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerCluster = `{
  "block": {
    "attributes": {
      "allow_net_admin": {
        "description": "Enable NET_ADMIN for this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_ipv4_cidr": {
        "computed": true,
        "description": "The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "datapath_provider": {
        "computed": true,
        "description": "The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_max_pods_per_node": {
        "computed": true,
        "description": "The default maximum number of pods per node in this cluster. This doesn't work on \"routes-based\" clusters, clusters that don't have IP Aliasing enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description": " Description of the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_autopilot": {
        "description": "Enable Autopilot for this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_binary_authorization": {
        "deprecated": true,
        "description": "Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_intranode_visibility": {
        "computed": true,
        "description": "Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_kubernetes_alpha": {
        "description": "Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_l4_ilb_subsetting": {
        "description": "Whether L4ILB Subsetting is enabled for this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_legacy_abac": {
        "description": "Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_shielded_nodes": {
        "description": "Enable Shielded Nodes features on all nodes in this cluster. Defaults to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_tpu": {
        "description": "Whether to enable Cloud TPU resources in this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The IP address of this cluster's Kubernetes master.",
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
        "description": "The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint of the set of labels for this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_service": {
        "computed": true,
        "description": "The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_version": {
        "computed": true,
        "description": "The current version of the master in the cluster. This may be different than the min_master_version set in the config if the master has been updated by GKE.",
        "description_kind": "plain",
        "type": "string"
      },
      "min_master_version": {
        "description": "The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitoring_service": {
        "computed": true,
        "description": "The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the cluster, unique within the project and location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The name or self_link of the Google Compute Engine network to which the cluster is connected. For Shared VPC, set this to the self link of the shared network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "networking_mode": {
        "computed": true,
        "description": "Determines whether alias IPs or routes will be used for pod IPs in the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_locations": {
        "computed": true,
        "description": "The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "node_version": {
        "computed": true,
        "description": "The Kubernetes version on the nodes. Must either be unset or set to the same value as min_master_version on create. Defaults to the default version set by GKE which is not necessarily the latest version. This only affects nodes in the default node pool. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way. To update nodes in other node pools, use the version attribute on the node pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ipv6_google_access": {
        "computed": true,
        "description": "The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remove_default_node_pool": {
        "description": "If true, deletes the default node pool upon cluster creation. If you're using google_container_node_pool resources with no default node pool, this should be set to true, alongside setting initial_node_count to at least 1.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_labels": {
        "description": "The GCE resource labels (a map of key/value pairs) to be applied to the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
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
        "optional": true,
        "type": "string"
      },
      "tpu_ipv4_cidr_block": {
        "computed": true,
        "description": "The IP address range of the Cloud TPUs in this cluster, in CIDR notation (e.g. 1.2.3.4/29).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "addons_config": {
        "block": {
          "block_types": {
            "cloudrun_config": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "load_balancer_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The status of the CloudRun addon. It is disabled by default. Set disabled = false to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "config_connector_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The of the Config Connector addon.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "dns_cache_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the NodeLocal DNSCache addon. It is disabled by default. Set enabled = true to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gce_persistent_disk_csi_driver_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to enabled; set disabled = true to disable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcp_filestore_csi_driver_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the Filestore CSI driver addon, which allows the usage of filestore instance as volumes. Defaults to disabled; set enabled = true to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcs_fuse_csi_driver_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the GCS Fuse CSI driver addon, which allows the usage of gcs bucket as volumes. Defaults to disabled; set enabled = true to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gke_backup_agent_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the Backup for GKE Agent addon. It is disabled by default. Set enabled = true to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "horizontal_pod_autoscaling": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set disabled = true to disable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_load_balancing": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set disabled = true to disable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_policy_config": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a network_policy block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set disabled = false to enable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration for addons supported by GKE.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "authenticator_groups_config": {
        "block": {
          "attributes": {
            "security_group": {
              "description": "The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for the Google Groups for GKE feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "binary_authorization": {
        "block": {
          "attributes": {
            "enabled": {
              "deprecated": true,
              "description": "Enable Binary Authorization for this cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "evaluation_mode": {
              "description": "Mode of operation for Binary Authorization policy evaluation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration options for the Binary Authorization feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cluster_autoscaling": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "auto_provisioning_defaults": {
              "block": {
                "attributes": {
                  "boot_disk_kms_key": {
                    "description": "The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size": {
                    "description": "Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "disk_type": {
                    "description": "Type of the disk attached to each node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_type": {
                    "description": "The default image type used by NAP once a new node pool is being created.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "computed": true,
                    "description": "Scopes that are used by NAP when creating node pools.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service_account": {
                    "description": "The Google Cloud Platform Service Account to be used by the node VMs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "management": {
                    "block": {
                      "attributes": {
                        "auto_repair": {
                          "computed": true,
                          "description": "Specifies whether the node auto-repair is enabled for the node pool. If enabled, the nodes in this node pool will be monitored and, if they fail health checks too many times, an automatic repair action will be triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "auto_upgrade": {
                          "computed": true,
                          "description": "Specifies whether node auto-upgrade is enabled for the node pool. If enabled, node auto-upgrade helps keep the nodes in your node pool up to date with the latest release version of Kubernetes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "upgrade_options": {
                          "computed": true,
                          "description": "Specifies the Auto Upgrade knobs for the node pool.",
                          "description_kind": "plain",
                          "type": [
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
                      },
                      "description": "NodeManagement configuration for this NodePool.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "shielded_instance_config": {
                    "block": {
                      "attributes": {
                        "enable_integrity_monitoring": {
                          "description": "Defines whether the instance has integrity monitoring enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_secure_boot": {
                          "description": "Defines whether the instance has Secure Boot enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Shielded Instance options.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "upgrade_settings": {
                    "block": {
                      "attributes": {
                        "max_surge": {
                          "description": "The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_unavailable": {
                          "description": "The maximum number of nodes that can be simultaneously unavailable during the upgrade process.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "strategy": {
                          "computed": true,
                          "description": "Update strategy of the node pool.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "blue_green_settings": {
                          "block": {
                            "attributes": {
                              "node_pool_soak_duration": {
                                "computed": true,
                                "description": "Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.\n\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "standard_rollout_policy": {
                                "block": {
                                  "attributes": {
                                    "batch_node_count": {
                                      "computed": true,
                                      "description": "Number of blue nodes to drain in a batch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "batch_percentage": {
                                      "computed": true,
                                      "description": "Percentage of the bool pool nodes to drain in a batch. The range of this field should be (0.0, 1.0].",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "batch_soak_duration": {
                                      "description": "Soak time after each batch gets drained.\n\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Standard policy for the blue-green upgrade.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Settings for blue-green upgrade strategy.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies the upgrade settings for NAP created node pools",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Contains defaults for a node pool created by NAP.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "resource_limits": {
              "block": {
                "attributes": {
                  "maximum": {
                    "description": "Maximum amount of the resource in the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minimum": {
                    "description": "Minimum amount of the resource in the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_type": {
                    "description": "The type of the resource. For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Global constraints for machine resources in the cluster. Configuring the cpu and memory types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "confidential_nodes": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether Confidential Nodes feature is enabled for all nodes in this cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cost_management_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether to enable GKE cost allocation. When you enable GKE cost allocation, the cluster name and namespace of your GKE workloads appear in the labels field of the billing export to BigQuery. Defaults to false.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Cost management configuration for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "database_encryption": {
        "block": {
          "attributes": {
            "key_name": {
              "description": "The key to use to encrypt/decrypt secrets.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description": "ENCRYPTED or DECRYPTED.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: \"ENCRYPTED\"; \"DECRYPTED\". key_name is the name of a CloudKMS key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "default_snat_status": {
        "block": {
          "attributes": {
            "disabled": {
              "description": "When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "dns_config": {
        "block": {
          "attributes": {
            "cluster_dns": {
              "description": "Which in-cluster DNS provider should be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_dns_domain": {
              "description": "The suffix used for all cluster service records.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_dns_scope": {
              "description": "The scope of access to cluster DNS records.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration for Cloud DNS for Kubernetes Engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "enable_k8s_beta_apis": {
        "block": {
          "attributes": {
            "enabled_apis": {
              "description": "Enabled Kubernetes Beta APIs.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Configuration for Kubernetes Beta APIs.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gateway_api_config": {
        "block": {
          "attributes": {
            "channel": {
              "description": "The Gateway API release channel to use for Gateway API.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for GKE Gateway API controller.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ip_allocation_policy": {
        "block": {
          "attributes": {
            "cluster_ipv4_cidr_block": {
              "computed": true,
              "description": "The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_secondary_range_name": {
              "computed": true,
              "description": "The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "services_ipv4_cidr_block": {
              "computed": true,
              "description": "The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "services_secondary_range_name": {
              "computed": true,
              "description": "The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stack_type": {
              "description": "The IP Stack type of the cluster. Choose between IPV4 and IPV4_IPV6. Default type is IPV4 Only if not set",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "additional_pod_ranges_config": {
              "block": {
                "attributes": {
                  "pod_range_names": {
                    "description": "Name for pod secondary ipv4 range which has the actual range defined ahead.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "AdditionalPodRangesConfig is the configuration for additional pod secondary ranges supporting the ClusterUpdate message.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pod_cidr_overprovision_config": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for cluster level pod cidr overprovision. Default is disabled=false.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "logging_config": {
        "block": {
          "attributes": {
            "enable_components": {
              "description": "GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Logging configuration for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_policy": {
        "block": {
          "block_types": {
            "daily_maintenance_window": {
              "block": {
                "attributes": {
                  "duration": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Time window specified for daily maintenance operations. Specify start_time in RFC3339 format \"HH:MM”, where HH : [00-23] and MM : [00-59] GMT.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "maintenance_exclusion": {
              "block": {
                "attributes": {
                  "end_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "exclusion_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "exclusion_options": {
                    "block": {
                      "attributes": {
                        "scope": {
                          "description": "The scope of automatic upgrades to restrict in the exclusion window.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Maintenance exclusion related options.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Exceptions to maintenance window. Non-emergency maintenance should not occur in these windows.",
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "recurring_window": {
              "block": {
                "attributes": {
                  "end_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "recurrence": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Time window for recurring maintenance operations.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The maintenance policy to use for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "master_auth": {
        "block": {
          "attributes": {
            "client_certificate": {
              "computed": true,
              "description": "Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_key": {
              "computed": true,
              "description": "Base64 encoded private key used by clients to authenticate to the cluster endpoint.",
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            },
            "cluster_ca_certificate": {
              "computed": true,
              "description": "Base64 encoded public certificate that is the root of trust for the cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "client_certificate_config": {
              "block": {
                "attributes": {
                  "issue_client_certificate": {
                    "description": "Whether client certificate authorization is enabled for this cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Whether client certificate authorization is enabled for this cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials permission.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "master_authorized_networks_config": {
        "block": {
          "attributes": {
            "gcp_public_cidrs_access_enabled": {
              "computed": true,
              "description": "Whether master is accessbile via Google Compute Engine Public IP addresses.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "cidr_blocks": {
              "block": {
                "attributes": {
                  "cidr_block": {
                    "description": "External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "display_name": {
                    "description": "Field for users to identify CIDR blocks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "External networks that can access the Kubernetes cluster master through HTTPS.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists).",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mesh_certificates": {
        "block": {
          "attributes": {
            "enable_certificates": {
              "description": "When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "If set, and enable_certificates=true, the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitoring_config": {
        "block": {
          "attributes": {
            "enable_components": {
              "computed": true,
              "description": "GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, and SCHEDULER.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "advanced_datapath_observability_config": {
              "block": {
                "attributes": {
                  "enable_metrics": {
                    "description": "Whether or not the advanced datapath metrics are enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "relay_mode": {
                    "computed": true,
                    "description": "Mode used to make Relay available.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration of Advanced Datapath Observability features.",
                "description_kind": "plain"
              },
              "max_items": 2,
              "nesting_mode": "list"
            },
            "managed_prometheus": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not the managed collection is enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for Google Cloud Managed Services for Prometheus.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Monitoring configuration for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_policy": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether network policy is enabled on the cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "provider": {
              "description": "The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration options for the NetworkPolicy feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "boot_disk_kms_key": {
              "description": "The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
              "description": "Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description": "Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "guest_accelerator": {
              "computed": true,
              "description": "List of the type and count of accelerator cards attached to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "gpu_driver_installation_config": [
                      "list",
                      [
                        "object",
                        {
                          "gpu_driver_version": "string"
                        }
                      ]
                    ],
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
              ]
            },
            "image_type": {
              "computed": true,
              "description": "The image type to use for this node. Note that for a given image type, the latest version of it will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description": "The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "local_ssd_count": {
              "computed": true,
              "description": "The number of local SSD disks to be attached to the node.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "logging_variant": {
              "description": "Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machine_type": {
              "computed": true,
              "description": "The name of a Google Compute Engine machine type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "The metadata key/value pairs assigned to instances in the cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "min_cpu_platform": {
              "computed": true,
              "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group": {
              "description": "Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oauth_scopes": {
              "computed": true,
              "description": "The set of Google API scopes to be made available on all of the node VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "preemptible": {
              "description": "Whether the nodes are created as preemptible VM instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "resource_labels": {
              "description": "The GCE resource labels (a map of key/value pairs) to be applied to the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "service_account": {
              "computed": true,
              "description": "The Google Cloud Platform Service Account to be used by the node VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot": {
              "description": "Whether the nodes are created as spot VM instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tags": {
              "description": "The list of instance tags applied to all nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "taint": {
              "computed": true,
              "description": "List of Kubernetes taints to be applied to each node.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "effect": "string",
                    "key": "string",
                    "value": "string"
                  }
                ]
              ]
            }
          },
          "block_types": {
            "advanced_machine_features": {
              "block": {
                "attributes": {
                  "threads_per_core": {
                    "description": "The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies options for controlling advanced machine features.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ephemeral_storage_local_ssd_config": {
              "block": {
                "attributes": {
                  "local_ssd_count": {
                    "description": "Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcfs_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not GCFS is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "GCFS configuration for this node.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gvnic": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not gvnic is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Enable or disable gvnic in the node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "host_maintenance_policy": {
              "block": {
                "attributes": {
                  "maintenance_interval": {
                    "description": ".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The maintenance policy for the hosts on which the GKE VMs run on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kubelet_config": {
              "block": {
                "attributes": {
                  "cpu_cfs_quota": {
                    "description": "Enable CPU CFS quota enforcement for containers that specify CPU limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cpu_cfs_quota_period": {
                    "description": "Set the CPU CFS quota period value 'cpu.cfs_period_us'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_manager_policy": {
                    "description": "Control the CPU management policy on the node.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "pod_pids_limit": {
                    "description": "Controls the maximum number of processes allowed to run in a pod.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Node kubelet configs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "linux_node_config": {
              "block": {
                "attributes": {
                  "sysctls": {
                    "description": "The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Parameters that can be configured on Linux nodes.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "local_nvme_ssd_block_config": {
              "block": {
                "attributes": {
                  "local_ssd_count": {
                    "description": "Number of raw-block local NVMe SSD disks to be attached to the node. Each local SSD is 375 GB in size.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Parameters for raw-block local NVMe SSDs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "reservation_affinity": {
              "block": {
                "attributes": {
                  "consume_reservation_type": {
                    "description": "Corresponds to the type of reservation consumption.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "The label key of a reservation resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "The label values of the reservation resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "The reservation affinity configuration for the node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "shielded_instance_config": {
              "block": {
                "attributes": {
                  "enable_integrity_monitoring": {
                    "description": "Defines whether the instance has integrity monitoring enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_secure_boot": {
                    "description": "Defines whether the instance has Secure Boot enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Shielded Instance options.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sole_tenant_config": {
              "block": {
                "block_types": {
                  "node_affinity": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": ".",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description": "Node affinity options for sole tenant node pools.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "workload_metadata_config": {
              "block": {
                "attributes": {
                  "mode": {
                    "description": "Mode is the configuration for how to expose metadata to workloads running on the node.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The workload metadata configuration for this node.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration of the nodepool",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_pool": {
        "block": {
          "attributes": {
            "initial_node_count": {
              "computed": true,
              "description": "The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_group_urls": {
              "computed": true,
              "description": "The resource URLs of the managed instance groups associated with this node pool.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "managed_instance_group_urls": {
              "computed": true,
              "description": "List of instance group URLs which have been assigned to this node pool.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "max_pods_per_node": {
              "computed": true,
              "description": "The maximum number of pods per node in this node pool. Note that this does not work on node pools which are \"route-based\" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "The name of the node pool. If left blank, Terraform will auto-generate a unique name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name_prefix": {
              "computed": true,
              "description": "Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_count": {
              "computed": true,
              "description": "The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "node_locations": {
              "computed": true,
              "description": "The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "version": {
              "computed": true,
              "description": "The Kubernetes version for the nodes in this pool. Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "autoscaling": {
              "block": {
                "attributes": {
                  "location_policy": {
                    "computed": true,
                    "description": "Location policy specifies the algorithm used when scaling-up the node pool. \"BALANCED\" - Is a best effort policy that aims to balance the sizes of available zones. \"ANY\" - Instructs the cluster autoscaler to prioritize utilization of unused reservations, and reduces preemption risk for Spot VMs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_node_count": {
                    "description": "Maximum number of nodes per zone in the node pool. Must be \u003e= min_node_count. Cannot be used with total limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_node_count": {
                    "description": "Minimum number of nodes per zone in the node pool. Must be \u003e=0 and \u003c= max_node_count. Cannot be used with total limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "total_max_node_count": {
                    "description": "Maximum number of all nodes in the node pool. Must be \u003e= total_min_node_count. Cannot be used with per zone limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "total_min_node_count": {
                    "description": "Minimum number of all nodes in the node pool. Must be \u003e=0 and \u003c= total_max_node_count. Cannot be used with per zone limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "management": {
              "block": {
                "attributes": {
                  "auto_repair": {
                    "description": "Whether the nodes will be automatically repaired.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "auto_upgrade": {
                    "description": "Whether the nodes will be automatically upgraded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Node management configuration, wherein auto-repair and auto-upgrade is configured.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_config": {
              "block": {
                "attributes": {
                  "create_pod_range": {
                    "description": "Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_private_nodes": {
                    "computed": true,
                    "description": "Whether nodes have internal IP addresses only.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "pod_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The IP address range for pod IPs in this node pool. Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pod_range": {
                    "computed": true,
                    "description": "The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "pod_cidr_overprovision_config": {
                    "block": {
                      "attributes": {
                        "disabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configuration for node-pool level pod cidr overprovision. If not set, the cluster level setting will be inherited",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "node_config": {
              "block": {
                "attributes": {
                  "boot_disk_kms_key": {
                    "description": "The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size_gb": {
                    "computed": true,
                    "description": "Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "disk_type": {
                    "computed": true,
                    "description": "Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "guest_accelerator": {
                    "computed": true,
                    "description": "List of the type and count of accelerator cards attached to the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "gpu_driver_installation_config": [
                            "list",
                            [
                              "object",
                              {
                                "gpu_driver_version": "string"
                              }
                            ]
                          ],
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
                    ]
                  },
                  "image_type": {
                    "computed": true,
                    "description": "The image type to use for this node. Note that for a given image type, the latest version of it will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "labels": {
                    "computed": true,
                    "description": "The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "local_ssd_count": {
                    "computed": true,
                    "description": "The number of local SSD disks to be attached to the node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "logging_variant": {
                    "description": "Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The name of a Google Compute Engine machine type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metadata": {
                    "computed": true,
                    "description": "The metadata key/value pairs assigned to instances in the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_group": {
                    "description": "Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "computed": true,
                    "description": "The set of Google API scopes to be made available on all of the node VMs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "preemptible": {
                    "description": "Whether the nodes are created as preemptible VM instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "resource_labels": {
                    "description": "The GCE resource labels (a map of key/value pairs) to be applied to the node pool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "service_account": {
                    "computed": true,
                    "description": "The Google Cloud Platform Service Account to be used by the node VMs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot": {
                    "description": "Whether the nodes are created as spot VM instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tags": {
                    "description": "The list of instance tags applied to all nodes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "taint": {
                    "computed": true,
                    "description": "List of Kubernetes taints to be applied to each node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "effect": "string",
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "advanced_machine_features": {
                    "block": {
                      "attributes": {
                        "threads_per_core": {
                          "description": "The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Specifies options for controlling advanced machine features.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ephemeral_storage_local_ssd_config": {
                    "block": {
                      "attributes": {
                        "local_ssd_count": {
                          "description": "Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gcfs_config": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description": "Whether or not GCFS is enabled",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "description": "GCFS configuration for this node.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gvnic": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description": "Whether or not gvnic is enabled",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "description": "Enable or disable gvnic in the node pool.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "host_maintenance_policy": {
                    "block": {
                      "attributes": {
                        "maintenance_interval": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The maintenance policy for the hosts on which the GKE VMs run on.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "kubelet_config": {
                    "block": {
                      "attributes": {
                        "cpu_cfs_quota": {
                          "description": "Enable CPU CFS quota enforcement for containers that specify CPU limits.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "cpu_cfs_quota_period": {
                          "description": "Set the CPU CFS quota period value 'cpu.cfs_period_us'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cpu_manager_policy": {
                          "description": "Control the CPU management policy on the node.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "pod_pids_limit": {
                          "description": "Controls the maximum number of processes allowed to run in a pod.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Node kubelet configs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "linux_node_config": {
                    "block": {
                      "attributes": {
                        "sysctls": {
                          "description": "The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Parameters that can be configured on Linux nodes.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "local_nvme_ssd_block_config": {
                    "block": {
                      "attributes": {
                        "local_ssd_count": {
                          "description": "Number of raw-block local NVMe SSD disks to be attached to the node. Each local SSD is 375 GB in size.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Parameters for raw-block local NVMe SSDs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "reservation_affinity": {
                    "block": {
                      "attributes": {
                        "consume_reservation_type": {
                          "description": "Corresponds to the type of reservation consumption.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description": "The label key of a reservation resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "values": {
                          "description": "The label values of the reservation resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "The reservation affinity configuration for the node pool.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "shielded_instance_config": {
                    "block": {
                      "attributes": {
                        "enable_integrity_monitoring": {
                          "description": "Defines whether the instance has integrity monitoring enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_secure_boot": {
                          "description": "Defines whether the instance has Secure Boot enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Shielded Instance options.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sole_tenant_config": {
                    "block": {
                      "block_types": {
                        "node_affinity": {
                          "block": {
                            "attributes": {
                              "key": {
                                "description": ".",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "operator": {
                                "description": ".",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "values": {
                                "description": ".",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": ".",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description": "Node affinity options for sole tenant node pools.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "workload_metadata_config": {
                    "block": {
                      "attributes": {
                        "mode": {
                          "description": "Mode is the configuration for how to expose metadata to workloads running on the node.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The workload metadata configuration for this node.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration of the nodepool",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "placement_policy": {
              "block": {
                "attributes": {
                  "policy_name": {
                    "description": "If set, refers to the name of a custom resource policy supplied by the user. The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Type defines the type of placement policy",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the node placement policy",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "upgrade_settings": {
              "block": {
                "attributes": {
                  "max_surge": {
                    "computed": true,
                    "description": "The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_unavailable": {
                    "computed": true,
                    "description": "The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "strategy": {
                    "description": "Update strategy for the given nodepool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "blue_green_settings": {
                    "block": {
                      "attributes": {
                        "node_pool_soak_duration": {
                          "computed": true,
                          "description": "Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "standard_rollout_policy": {
                          "block": {
                            "attributes": {
                              "batch_node_count": {
                                "computed": true,
                                "description": "Number of blue nodes to drain in a batch.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "batch_percentage": {
                                "computed": true,
                                "description": "Percentage of the blue pool nodes to drain in a batch.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "batch_soak_duration": {
                                "computed": true,
                                "description": "Soak time after each batch gets drained.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Standard rollout policy is the default policy for blue-green.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Settings for BlueGreen node pool upgrade.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of node pools associated with this cluster. See google_container_node_pool for schema. Warning: node pools defined inside a cluster can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability to say \"these are the only node pools associated with this cluster\", use the google_container_node_pool resource instead of this property.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "node_pool_defaults": {
        "block": {
          "block_types": {
            "node_config_defaults": {
              "block": {
                "attributes": {
                  "logging_variant": {
                    "description": "Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Subset of NodeConfig message that has defaults.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The default nodel pool settings for the entire cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_config": {
        "block": {
          "block_types": {
            "pubsub": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not the notification config is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "topic": {
                    "description": "The pubsub topic to push upgrade notifications to. Must be in the same project as the cluster. Must be in the format: projects/{project}/topics/{topic}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "filter": {
                    "block": {
                      "attributes": {
                        "event_type": {
                          "description": "Can be used to filter what notifications are sent. Valid values include include UPGRADE_AVAILABLE_EVENT, UPGRADE_EVENT and SECURITY_BULLETIN_EVENT",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Allows filtering to one or more specific event types. If event types are present, those and only those event types will be transmitted to the cluster. Other types will be skipped. If no filter is specified, or no event types are present, all event types will be sent",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Notification config for Cloud Pub/Sub",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The notification config for sending cluster upgrade notifications",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_cluster_config": {
        "block": {
          "attributes": {
            "enable_private_endpoint": {
              "description": "When true, the cluster's private endpoint is used as the cluster endpoint and access through the public endpoint is disabled. When false, either endpoint can be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_private_nodes": {
              "description": "Enables the private cluster feature, creating a private endpoint on the cluster. In a private cluster, nodes only have RFC 1918 private addresses and communicate with the master's private endpoint via private networking.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "master_ipv4_cidr_block": {
              "computed": true,
              "description": "The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning private IP addresses to the cluster master(s) and the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network, and it must be a /28 subnet. See Private Cluster Limitations for more details. This field only applies to private clusters, when enable_private_nodes is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "peering_name": {
              "computed": true,
              "description": "The name of the peering between this cluster and the Google owned VPC.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_endpoint": {
              "computed": true,
              "description": "The internal IP address of this cluster's master endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_endpoint_subnetwork": {
              "description": "Subnetwork in cluster's network where master's endpoint will be provisioned.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_endpoint": {
              "computed": true,
              "description": "The external IP address of this cluster's master endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "master_global_access_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether the cluster master is accessible globally or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Controls cluster master global access settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for private clusters, clusters with private nodes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "release_channel": {
        "block": {
          "attributes": {
            "channel": {
              "description": "The selected release channel. Accepted values are:\n* UNSPECIFIED: Not set.\n* RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.\n* REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.\n* STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the \"UNSPECIFIED\" channel.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resource_usage_export_config": {
        "block": {
          "attributes": {
            "enable_network_egress_metering": {
              "description": "Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_resource_consumption_metering": {
              "description": "Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "bigquery_destination": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The ID of a BigQuery Dataset.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Parameters for using BigQuery as the destination of resource usage export.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for the ResourceUsageExportConfig feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "security_posture_config": {
        "block": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include DISABLED and BASIC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vulnerability_mode": {
              "computed": true,
              "description": "Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include VULNERABILITY_DISABLED and VULNERABILITY_BASIC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Defines the config needed to enable/disable features for the Security Posture API",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_external_ips_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "When enabled, services with exterenal ips specified will be allowed.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "If set, and enabled=true, services with external ips field will not be blocked",
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
            "read": {
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
      },
      "vertical_pod_autoscaling": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Enables vertical pod autoscaling.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "workload_identity_config": {
        "block": {
          "attributes": {
            "workload_pool": {
              "description": "The workload pool to attach all Kubernetes service accounts to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
