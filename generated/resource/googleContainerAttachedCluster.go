package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAttachedCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Annotations on the cluster. This field has the same\nrestrictions as Kubernetes annotations. The total size of all keys and\nvalues combined is limited to 256k. Key can have 2 segments: prefix (optional)\nand name (required), separated by a slash (/). Prefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "cluster_region": {
        "computed": true,
        "description": "Output only. The region where this cluster runs.\n\nFor EKS clusters, this is an AWS region. For AKS clusters,\nthis is an Azure region.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which this cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "description": "Policy to determine what flags to send on delete.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A human readable description of this attached cluster. Cannot be longer\nthan 255 UTF-8 encoded bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distribution": {
        "description": "The Kubernetes distribution of the underlying attached cluster. Supported values:\n\"eks\", \"aks\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "errors": {
        "computed": true,
        "description": "A set of errors found in the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "message": "string"
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
      "kubernetes_version": {
        "computed": true,
        "description": "The Kubernetes version of the cluster.",
        "description_kind": "plain",
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
      "platform_version": {
        "description": "The platform version for the cluster (e.g. '1.23.0-gke.1').",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "If set, there are currently changes in flight to the cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the cluster. Possible values:\nSTATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,\nDEGRADED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "A globally unique identifier for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which this cluster was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_identity_config": {
        "computed": true,
        "description": "Workload Identity settings.",
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
          "attributes": {
            "admin_users": {
              "description": "Users that can perform operations as a cluster admin. A managed\nClusterRoleBinding will be created to grant the 'cluster-admin' ClusterRole\nto the users. Up to ten admin users can be provided.\n\nFor more info on RBAC, see\nhttps://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Configuration related to the cluster RBAC settings.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "fleet": {
        "block": {
          "attributes": {
            "membership": {
              "computed": true,
              "description": "The name of the managed Hub Membership resource associated to this\ncluster. Membership names are formatted as\nprojects/\u003cproject-number\u003e/locations/global/membership/\u003ccluster-id\u003e.",
              "description_kind": "plain",
              "type": "string"
            },
            "project": {
              "description": "The number of the Fleet host project where this cluster will be registered.",
              "description_kind": "plain",
              "required": true,
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
      "logging_config": {
        "block": {
          "block_types": {
            "component_config": {
              "block": {
                "attributes": {
                  "enable_components": {
                    "description": "The components to be enabled. Possible values: [\"SYSTEM_COMPONENTS\", \"WORKLOADS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The configuration of the logging components",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Logging configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitoring_config": {
        "block": {
          "block_types": {
            "managed_prometheus_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Enable Managed Collection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Enable Google Cloud Managed Service for Prometheus in the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Monitoring configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oidc_config": {
        "block": {
          "attributes": {
            "issuer_url": {
              "description": "A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "jwks": {
              "description": "OIDC verification keys in JWKS format (RFC 7517).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "OIDC discovery information of the target cluster.\n\nKubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster\nAPI server. This fields indicates how GCP services\nvalidate KSA tokens in order to allow system workloads (such as GKE Connect\nand telemetry agents) to authenticate back to GCP.\n\nBoth clusters with public and private issuer URLs are supported.\nClusters with public issuers only need to specify the 'issuer_url' field\nwhile clusters with private issuers need to provide both\n'issuer_url' and 'jwks'.",
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

func GoogleContainerAttachedClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAttachedCluster), &result)
	return &result
}
