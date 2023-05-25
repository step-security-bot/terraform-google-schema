package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerNodePool = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The cluster to create the node pool for. Cluster must be present in location provided for zonal clusters.",
        "description_kind": "plain",
        "required": true,
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
      "location": {
        "computed": true,
        "description": "The location (region or zone) of the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which to create the node pool. If blank, the provider-configured project will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
            "max_node_count": {
              "description": "Maximum number of nodes in the NodePool. Must be \u003e= min_node_count.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
              "description": "Minimum number of nodes in the NodePool. Must be \u003e=0 and \u003c= max_node_count.",
              "description_kind": "plain",
              "required": true,
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
      "node_config": {
        "block": {
          "attributes": {
            "disk_size_gb": {
              "computed": true,
              "description": "Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description": "Type of the disk attached to each node.",
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
                    "gpu_partition_size": "string",
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
              "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.",
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
            "service_account": {
              "computed": true,
              "description": "The Google Cloud Platform Service Account to be used by the node VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "workload_metadata_config": {
              "block": {
                "attributes": {
                  "mode": {
                    "computed": true,
                    "description": "Mode is the configuration for how to expose metadata to workloads running on the node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_metadata": {
                    "computed": true,
                    "deprecated": true,
                    "description": "NodeMetadata is the configuration for how to expose metadata to the workloads running on the node.",
                    "description_kind": "plain",
                    "optional": true,
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
      },
      "upgrade_settings": {
        "block": {
          "attributes": {
            "max_surge": {
              "description": "The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "max_unavailable": {
              "description": "The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.",
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

func GoogleContainerNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerNodePool), &result)
	return &result
}
