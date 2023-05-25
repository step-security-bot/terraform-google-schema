package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGameServicesGameServerCluster = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description": "Required. The resource name of the game server cluster",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Human readable description of the cluster.",
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
      "labels": {
        "description": "The labels associated with this game server cluster. Each label is a\nkey-value pair.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the Cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource id of the game server cluster, eg:\n\n'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'.\nFor example,\n\n'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "realm_id": {
        "description": "The realm id of the game server realm.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "connection_info": {
        "block": {
          "attributes": {
            "namespace": {
              "description": "Namespace designated on the game server cluster where the game server\ninstances will be created. The namespace existence will be validated\nduring creation.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "gke_cluster_reference": {
              "block": {
                "attributes": {
                  "cluster": {
                    "description": "The full or partial name of a GKE cluster, using one of the following\nforms:\n\n* 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'\n* 'locations/{location}/clusters/{cluster_id}'\n* '{cluster_id}'\n\nIf project and location are not specified, the project and location of the\nGameServerCluster resource are used to generate the full name of the\nGKE cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Reference of the GKE cluster where the game servers are installed.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Game server cluster connection information. This information is used to\nmanage game server clusters.",
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

func GoogleGameServicesGameServerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGameServicesGameServerCluster), &result)
	return &result
}
