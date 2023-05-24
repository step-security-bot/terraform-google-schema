package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubMembership = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this membership.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "membership_id": {
        "description": "The client-provided identifier of the membership.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the membership.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "authority": {
        "block": {
          "attributes": {
            "issuer": {
              "description": "A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://' and // be a valid\nwith length \u003c2000 characters. For example: 'https://container.googleapis.com/v1/projects/my-project/locations/us-west1/clusters/my-cluster' (must be 'locations' rather than 'zones'). If the cluster is provisioned with Terraform, this is '\"https://container.googleapis.com/v1/${google_container_cluster.my-cluster.id}\"'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Authority encodes how Google will recognize identities from this Membership.\nSee the workload identity documentation for more details:\nhttps://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "endpoint": {
        "block": {
          "block_types": {
            "gke_cluster": {
              "block": {
                "attributes": {
                  "resource_link": {
                    "description": "Self-link of the GCP resource for the GKE cluster.\nFor example: '//container.googleapis.com/projects/my-project/zones/us-west1-a/clusters/my-cluster'.\nIt can be at the most 1000 characters in length. If the cluster is provisioned with Terraform,\nthis can be '\"//container.googleapis.com/${google_container_cluster.my-cluster.id}\"' or\n'google_container_cluster.my-cluster.id'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGkeHubMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubMembership), &result)
	return &result
}
