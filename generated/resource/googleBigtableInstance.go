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
        "description": "The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.",
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
              "description": "The ID of the Cloud Bigtable cluster.",
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
              "description": "The number of nodes in your Cloud Bigtable cluster. Required, with a minimum of 1 for a PRODUCTION instance. Must be left unset for a DEVELOPMENT instance.",
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
          "description": "A block of cluster configuration options. This can be specified at least once.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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
