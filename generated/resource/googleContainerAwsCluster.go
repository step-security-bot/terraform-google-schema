package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAwsCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "aws_region": {
        "description": "The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.",
        "description_kind": "plain",
        "required": true,
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
                "description": "Users to perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the ` + "`" + `cluster-admin` + "`" + ` ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles",
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
      "control_plane": {
        "block": {
          "attributes": {
            "iam_instance_profile": {
              "description": "The name of the AWS IAM instance pofile to assign to each control plane replica.",
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
            "security_group_ids": {
              "description": "Optional. The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "description": "The list of subnets where control plane replicas will run. A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tags": {
              "description": "Optional. A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "description": "The Kubernetes version to run on control plane replicas (e.g. ` + "`" + `1.19.10-gke.1000` + "`" + `). You can list all supported versions on a given Google Cloud region by calling .",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "aws_services_authentication": {
              "block": {
                "attributes": {
                  "role_arn": {
                    "description": "The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_session_name": {
                    "computed": true,
                    "description": "Optional. An identifier for the assumed role session. When unspecified, it defaults to ` + "`" + `multicloud-service-agent` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Authentication configuration for management of AWS resources.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "config_encryption": {
              "block": {
                "attributes": {
                  "kms_key_arn": {
                    "description": "The ARN of the AWS KMS key used to encrypt cluster configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The ARN of the AWS KMS key used to encrypt cluster configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "database_encryption": {
              "block": {
                "attributes": {
                  "kms_key_arn": {
                    "description": "The ARN of the AWS KMS key used to encrypt cluster secrets.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The ARN of the AWS KMS key used to encrypt cluster secrets.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "main_volume": {
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
                  "throughput": {
                    "computed": true,
                    "description": "Optional. The throughput to provision for the volume, in MiB/s. Only valid if the volume type is GP3.",
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
                "description": "Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 8 GiB with the GP2 volume type.",
                "description_kind": "plain"
              },
              "max_items": 1,
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
                  "throughput": {
                    "computed": true,
                    "description": "Optional. The throughput to provision for the volume, in MiB/s. Only valid if the volume type is GP3.",
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
                "description": "Optional. Configuration related to the root volume provisioned for each control plane replica. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.",
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
                "description": "Optional. SSH configuration for how to access the underlying control plane machines.",
                "description_kind": "plain"
              },
              "max_items": 1,
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
            "per_node_pool_sg_rules_disabled": {
              "description": "Disable the per node pool subnet security group rules on the control plane security group. When set to true, you must also provide one or more security groups that ensure node pools are able to send requests to the control plane on TCP/443 and TCP/8132. Failure to do so may result in unavailable node pools.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pod_address_cidr_blocks": {
              "description": "All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "service_address_cidr_blocks": {
              "description": "All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "description": "The VPC associated with the cluster. All component clusters (i.e. control plane and node pools) run on a single VPC. This field cannot be changed after creation.",
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

func GoogleContainerAwsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAwsCluster), &result)
	return &result
}
