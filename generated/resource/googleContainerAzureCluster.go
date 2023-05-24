package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAzureCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "azure_region": {
        "description": "The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client": {
        "description": "Name of the AzureClient. The ` + "`" + `AzureClient` + "`" + ` resource must reside on the same GCP project and region as the ` + "`" + `AzureCluster` + "`" + `. ` + "`" + `AzureClient` + "`" + ` names are formatted as ` + "`" + `projects/\u003cproject-number\u003e/locations/\u003cregion\u003e/azureClients/\u003cclient-id\u003e` + "`" + `. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which this cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "Output only. The endpoint of the cluster's API server.",
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
        "description": "Output only. If set, there are currently changes in flight to the cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_id": {
        "description": "The ARM ID of the resource group where the cluster resources are deployed. For example: ` + "`" + `/subscriptions/*/resourceGroups/*` + "`" + `",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A globally unique identifier for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which this cluster was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_identity_config": {
        "computed": true,
        "description": "Output only. Workload Identity settings.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_provider": "string",
              "issuer_uri": "string",
              "workload_pool": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "authorization": {
        "block": {
          "block_types": {
            "admin_users": {
              "block": {
                "attributes": {
                  "username": {
                    "description": "The name of the user, e.g. ` + "`" + `my-gcp-id@gmail.com` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Users that can perform operations as a cluster admin. A new ClusterRoleBinding will be created to grant the cluster-admin ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration related to the cluster RBAC settings.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "azure_services_authentication": {
        "block": {
          "attributes": {
            "application_id": {
              "description": "The Azure Active Directory Application ID for Authentication configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tenant_id": {
              "description": "The Azure Active Directory Tenant ID for Authentication configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Azure authentication configuration for management of Azure resources",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "control_plane": {
        "block": {
          "attributes": {
            "subnet_id": {
              "description": "The ARM ID of the subnet where the control plane VMs are deployed. Example: ` + "`" + `/subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tags": {
              "description": "Optional. A set of tags to apply to all underlying control plane Azure resources.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "description": "The Kubernetes version to run on control plane replicas (e.g. ` + "`" + `1.19.10-gke.1000` + "`" + `). You can list all supported versions on a given Google Cloud region by calling GetAzureServerConfig.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_size": {
              "computed": true,
              "description": "Optional. The Azure VM size name. Example: ` + "`" + `Standard_DS2_v2` + "`" + `. For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions. When unspecified, it defaults to ` + "`" + `Standard_DS2_v2` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "database_encryption": {
              "block": {
                "attributes": {
                  "key_id": {
                    "description": "The ARM ID of the Azure Key Vault key to encrypt / decrypt data. For example: ` + "`" + `/subscriptions/\u003csubscription-id\u003e/resourceGroups/\u003cresource-group-id\u003e/providers/Microsoft.KeyVault/vaults/\u003ckey-vault-id\u003e/keys/\u003ckey-name\u003e` + "`" + ` Encryption will always take the latest version of the key and hence specific version is not supported.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Configuration related to application-layer secrets encryption.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "main_volume": {
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
                "description": "Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. When unspecified, it defaults to a 8-GiB Azure Disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
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
            "replica_placements": {
              "block": {
                "attributes": {
                  "azure_availability_zone": {
                    "description": "For a given replica, the Azure availability zone where to provision the control plane VM and the ETCD disk.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnet_id": {
                    "description": "For a given replica, the ARM ID of the subnet where the control plane VM is deployed. Make sure it's a subnet under the virtual network in the cluster configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for where to place the control plane replicas. Up to three replica placement instances can be specified. If replica_placements is set, the replica placement instances will be applied to the three control plane replicas as evenly as possible.",
                "description_kind": "plain"
              },
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
                "description": "Optional. Configuration related to the root volume provisioned for each control plane replica. When unspecified, it defaults to 32-GiB Azure Disk.",
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
                "description": "SSH configuration for how to access the underlying control plane machines.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration related to the cluster control plane.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "fleet": {
        "block": {
          "attributes": {
            "membership": {
              "computed": true,
              "description": "The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects/\u003cproject-number\u003e/locations/global/membership/\u003ccluster-id\u003e.",
              "description_kind": "plain",
              "type": "string"
            },
            "project": {
              "computed": true,
              "description": "The number of the Fleet host project where this cluster will be registered.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Fleet configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "networking": {
        "block": {
          "attributes": {
            "pod_address_cidr_blocks": {
              "description": "The IP address range of the pods in this cluster, in CIDR notation (e.g. ` + "`" + `10.96.0.0/14` + "`" + `). All pods in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "service_address_cidr_blocks": {
              "description": "The IP address range for services in this cluster, in CIDR notation (e.g. ` + "`" + `10.96.0.0/14` + "`" + `). All services in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creating a cluster.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "virtual_network_id": {
              "description": "The Azure Resource Manager (ARM) ID of the VNet associated with your cluster. All components in the cluster (i.e. control plane and node pools) run on a single VNet. Example: ` + "`" + `/subscriptions/*/resourceGroups/*/providers/Microsoft.Network/virtualNetworks/*` + "`" + ` This field cannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Cluster-wide networking configuration.",
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

func GoogleContainerAzureClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAzureCluster), &result)
	return &result
}
