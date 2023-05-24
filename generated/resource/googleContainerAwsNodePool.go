package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAwsNodePool = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "cluster": {
        "description": "The awsCluster for the resource",
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
        "description": "Output only. If set, there are currently changes in flight to the node pool.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "Output only. The lifecycle state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "description": "The subnet where the node pool node run.",
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
        "description": "The Kubernetes version to run on this node pool (e.g. ` + "`" + `1.19.10-gke.1000` + "`" + `). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.",
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
              "description": "Maximum number of nodes in the NodePool. Must be \u003e= min_node_count.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
              "description": "Minimum number of nodes in the NodePool. Must be \u003e= 1 and \u003c= max_node_count.",
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
            "iam_instance_profile": {
              "description": "The name of the AWS IAM role assigned to nodes in the pool.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_type": {
              "computed": true,
              "description": "Optional. The AWS instance type. When unspecified, it defaults to ` + "`" + `m5.large` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description": "Optional. The initial labels assigned to nodes of this node pool. An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "security_group_ids": {
              "description": "Optional. The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tags": {
              "description": "Optional. Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "autoscaling_metrics_collection": {
              "block": {
                "attributes": {
                  "granularity": {
                    "description": "The frequency at which EC2 Auto Scaling sends aggregated data to AWS CloudWatch. The only valid value is \"1Minute\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metrics": {
                    "description": "The metrics to enable. For a list of valid metrics, see https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html. If you specify granularity and don't specify any metrics, all metrics are enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Optional. Configuration related to CloudWatch metrics collection on the Auto Scaling group of the node pool. When unspecified, metrics collection is disabled.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "config_encryption": {
              "block": {
                "attributes": {
                  "kms_key_arn": {
                    "description": "The ARN of the AWS KMS key used to encrypt node pool configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The ARN of the AWS KMS key used to encrypt node pool configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "proxy_config": {
              "block": {
                "attributes": {
                  "secret_arn": {
                    "description": "The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_version": {
                    "description": "The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.",
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
                  "iops": {
                    "computed": true,
                    "description": "Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_arn": {
                    "description": "Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size_gib": {
                    "computed": true,
                    "description": "Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "volume_type": {
                    "computed": true,
                    "description": "Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Template for the root volume provisioned for node pool nodes. Volumes will be provisioned in the availability zone assigned to the node pool subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ssh_config": {
              "block": {
                "attributes": {
                  "ec2_key_pair": {
                    "description": "The name of the EC2 key pair used to login into cluster machines.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. The SSH configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "taints": {
              "block": {
                "attributes": {
                  "effect": {
                    "description": "The taint effect. Possible values: EFFECT_UNSPECIFIED, NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Key for the taint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value for the taint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. The initial taints assigned to nodes of this node pool.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The configuration of the node pool.",
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

func GoogleContainerAwsNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAwsNodePool), &result)
	return &result
}
