package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableInstance = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "description": "Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The human-readable display name of the Bigtable instance. Defaults to the instance name.",
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
      "instance_type": {
        "deprecated": true,
        "description": "The instance type to create. One of \"DEVELOPMENT\" or \"PRODUCTION\". Defaults to \"PRODUCTION\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A mapping of labels to assign to the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance. Must be 6-33 characters and must only contain hyphens, lowercase letters and numbers.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cluster": {
        "block": {
          "attributes": {
            "cluster_id": {
              "description": "The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key_name": {
              "computed": true,
              "description": "Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster. 3) All clusters within an instance must use the same CMEK key. Values are of the form projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "num_nodes": {
              "computed": true,
              "description": "The number of nodes in the cluster. If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "storage_type": {
              "description": "The storage type to use. One of \"SSD\" or \"HDD\". Defaults to \"SSD\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zone": {
              "computed": true,
              "description": "The zone to create the Cloud Bigtable cluster in. Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the Cloud Bigtable locations page.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "autoscaling_config": {
              "block": {
                "attributes": {
                  "cpu_target": {
                    "description": "The target CPU utilization for autoscaling. Value must be between 10 and 80.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "max_nodes": {
                    "description": "The maximum number of nodes for autoscaling.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "min_nodes": {
                    "description": "The minimum number of nodes for autoscaling.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "storage_target": {
                    "computed": true,
                    "description": "The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "A list of Autoscaling configurations. Only one element is used and allowed.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A block of cluster configuration options. This can be specified at least once.",
          "description_kind": "plain"
        },
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
  "version": 1
}`

func GoogleBigtableInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableInstance), &result)
	return &result
}
