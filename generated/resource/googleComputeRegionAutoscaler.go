package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionAutoscaler = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "region": {
        "computed": true,
        "description": "URL of the region where the instance group resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "description": "URL of the managed instance group that this autoscaler will scale.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling_policy": {
        "block": {
          "attributes": {
            "cooldown_period": {
              "description": "The number of seconds that the autoscaler should wait before it\nstarts collecting information from a new instance. This prevents\nthe autoscaler from collecting information when the instance is\ninitializing, during which the collected usage would not be\nreliable. The default time autoscaler waits is 60 seconds.\n\nVirtual machine initialization times might vary because of\nnumerous factors. We recommend that you test how long an\ninstance may take to initialize. To do this, create an instance\nand time the startup process.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_replicas": {
              "description": "The maximum number of instances that the autoscaler can scale up\nto. This is required when creating or updating an autoscaler. The\nmaximum number of replicas should not be lower than minimal number\nof replicas.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_replicas": {
              "description": "The minimum number of replicas that the autoscaler can scale down\nto. This cannot be less than 0. If not provided, autoscaler will\nchoose a default value depending on maximum number of instances\nallowed.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "cpu_utilization": {
              "block": {
                "attributes": {
                  "target": {
                    "description": "The target CPU utilization that the autoscaler should maintain.\nMust be a float value in the range (0, 1]. If not specified, the\ndefault is 0.6.\n\nIf the CPU level is below the target utilization, the autoscaler\nscales down the number of instances until it reaches the minimum\nnumber of instances you specified or until the average CPU of\nyour instances reaches the target utilization.\n\nIf the average CPU is above the target utilization, the autoscaler\nscales up until it reaches the maximum number of instances you\nspecified or until the average utilization reaches the target\nutilization.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "load_balancing_utilization": {
              "block": {
                "attributes": {
                  "target": {
                    "description": "Fraction of backend capacity utilization (set in HTTP(s) load\nbalancing configuration) that autoscaler should maintain. Must\nbe a positive float value. If not defined, the default is 0.8.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metric": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The identifier (type) of the Stackdriver Monitoring metric.\nThe metric cannot have negative values.\n\nThe metric must have a value type of INT64 or DOUBLE.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "target": {
                    "description": "The target value of the metric that autoscaler should\nmaintain. This must be a positive value. A utilization\nmetric scales number of virtual machines handling requests\nto increase or decrease proportionally to the metric.\n\nFor example, a good metric to use as a utilizationTarget is\nwww.googleapis.com/compute/instance/network/received_bytes_count.\nThe autoscaler will work to keep this value constant for each\nof the instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description": "Defines how target utilization value is expressed for a\nStackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,\nor DELTA_PER_MINUTE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
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

func GoogleComputeRegionAutoscalerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionAutoscaler), &result)
	return &result
}
