package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAzureNodePool = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "azure_availability_zone": {
        "computed": true,
        "description": "Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to ` + "`" + `1` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster": {
        "description": "The azureCluster for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which this node pool was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of this resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. If set, there are currently pending changes to the node pool.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "description": "The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A globally unique identifier for the node pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which this node pool was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description": "The Kubernetes version (e.g. ` + "`" + `1.19.10-gke.1000` + "`" + `) running on this node pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling": {
        "block": {
          "attributes": {
            "max_node_count": {
              "description": "Maximum number of nodes in the node pool. Must be \u003e= min_node_count.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
              "description": "Minimum number of nodes in the node pool. Must be \u003e= 1 and \u003c= max_node_count.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Autoscaler configuration for this node pool.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "config": {
        "block": {
          "attributes": {
            "tags": {
              "description": "Optional. A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "vm_size": {
              "computed": true,
              "description": "Optional. The Azure VM size name. Example: ` + "`" + `Standard_DS2_v2` + "`" + `. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to ` + "`" + `Standard_DS2_v2` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "proxy_config": {
              "block": {
                "attributes": {
                  "resource_group_id": {
                    "description": "The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as ` + "`" + `/subscriptions/\u003csubscription-id\u003e/resourceGroups/\u003cresource-group-name\u003e` + "`" + `",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_id": {
                    "description": "The URL the of the proxy setting secret with its version. Secret ids are formatted as ` + "`" + `https:\u003ckey-vault-name\u003e.vault.azure.net/secrets/\u003csecret-name\u003e/\u003csecret-version\u003e` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Proxy configuration for outbound HTTP(S) traffic.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "root_volume": {
              "block": {
                "attributes": {
                  "size_gib": {
                    "computed": true,
                    "description": "Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Optional. Configuration related to the root volume provisioned for each node pool machine. When unspecified, it defaults to a 32-GiB Azure Disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ssh_config": {
              "block": {
                "attributes": {
                  "authorized_key": {
                    "description": "The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "SSH configuration for how to access the node pool machines.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The node configuration of the node pool.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "max_pods_constraint": {
        "block": {
          "attributes": {
            "max_pods_per_node": {
              "description": "The maximum number of pods to schedule on a single node.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleContainerAzureNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAzureNodePool), &result)
	return &result
}
