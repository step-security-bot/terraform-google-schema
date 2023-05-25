package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocAutoscalingPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The  location where the autoscaling policy should reside.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The \"resource name\" of the autoscaling policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "description": "The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),\nand hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between\n3 and 50 characters.",
        "description_kind": "plain",
        "required": true,
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
      "basic_algorithm": {
        "block": {
          "attributes": {
            "cooldown_period": {
              "description": "Duration between scaling events. A scaling period starts after the\nupdate operation from the previous event has completed.\n\nBounds: [2m, 1d]. Default: 2m.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "yarn_config": {
              "block": {
                "attributes": {
                  "graceful_decommission_timeout": {
                    "description": "Timeout for YARN graceful decommissioning of Node Managers. Specifies the\nduration to wait for jobs to complete before forcefully removing workers\n(and potentially interrupting jobs). Only applicable to downscaling operations.\n\nBounds: [0s, 1d].",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "scale_down_factor": {
                    "description": "Fraction of average pending memory in the last cooldown period for which to\nremove workers. A scale-down factor of 1 will result in scaling down so that there\nis no available memory remaining after the update (more aggressive scaling).\nA scale-down factor of 0 disables removing workers, which can be beneficial for\nautoscaling a single job.\n\nBounds: [0.0, 1.0].",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "scale_down_min_worker_fraction": {
                    "description": "Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.\nFor example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must\nrecommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0\nmeans the autoscaler will scale down on any recommended change.\n\nBounds: [0.0, 1.0]. Default: 0.0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scale_up_factor": {
                    "description": "Fraction of average pending memory in the last cooldown period for which to\nadd workers. A scale-up factor of 1.0 will result in scaling up so that there\nis no pending memory remaining after the update (more aggressive scaling).\nA scale-up factor closer to 0 will result in a smaller magnitude of scaling up\n(less aggressive scaling).\n\nBounds: [0.0, 1.0].",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "scale_up_min_worker_fraction": {
                    "description": "Minimum scale-up threshold as a fraction of total cluster size before scaling\noccurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler\nmust recommend at least a 2-worker scale-up for the cluster to scale. A threshold of\n0 means the autoscaler will scale up on any recommended change.\n\nBounds: [0.0, 1.0]. Default: 0.0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "YARN autoscaling configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Basic algorithm for autoscaling.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "secondary_worker_config": {
        "block": {
          "attributes": {
            "max_instances": {
              "description": "Maximum number of instances for this group. Note that by default, clusters will not use\nsecondary workers. Required for secondary workers if the minimum secondary instances is set.\nBounds: [minInstances, ). Defaults to 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_instances": {
              "description": "Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weight": {
              "description": "Weight for the instance group, which is used to determine the fraction of total workers\nin the cluster from this instance group. For example, if primary workers have weight 2,\nand secondary workers have weight 1, the cluster will have approximately 2 primary workers\nfor each secondary worker.\n\nThe cluster may not reach the specified balance if constrained by min/max bounds or other\nautoscaling settings. For example, if maxInstances for secondary workers is 0, then only\nprimary workers will be added. The cluster can also be out of balance when created.\n\nIf weight is not set on any instance group, the cluster will default to equal weight for\nall groups: the cluster will attempt to maintain an equal number of workers in each group\nwithin the configured size bounds for each group. If weight is set for one group only,\nthe cluster will default to zero weight on the unset group. For example if weight is set\nonly on primary workers, the cluster will use primary workers only and no secondary workers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Describes how the autoscaler will operate for secondary workers.",
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
      "worker_config": {
        "block": {
          "attributes": {
            "max_instances": {
              "description": "Maximum number of instances for this group.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_instances": {
              "description": "Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weight": {
              "description": "Weight for the instance group, which is used to determine the fraction of total workers\nin the cluster from this instance group. For example, if primary workers have weight 2,\nand secondary workers have weight 1, the cluster will have approximately 2 primary workers\nfor each secondary worker.\n\nThe cluster may not reach the specified balance if constrained by min/max bounds or other\nautoscaling settings. For example, if maxInstances for secondary workers is 0, then only\nprimary workers will be added. The cluster can also be out of balance when created.\n\nIf weight is not set on any instance group, the cluster will default to equal weight for\nall groups: the cluster will attempt to maintain an equal number of workers in each group\nwithin the configured size bounds for each group. If weight is set for one group only,\nthe cluster will default to zero weight on the unset group. For example if weight is set\nonly on primary workers, the cluster will use primary workers only and no secondary workers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Describes how the autoscaler will operate for primary workers.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataprocAutoscalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocAutoscalingPolicy), &result)
	return &result
}
