package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstanceGroupManager = `{
  "block": {
    "attributes": {
      "base_instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_template": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_pools": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "update_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait_for_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_healing_policies": {
        "block": {
          "attributes": {
            "health_check": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "initial_delay_sec": {
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
      "named_port": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rolling_update_policy": {
        "block": {
          "attributes": {
            "max_surge_fixed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_surge_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_fixed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_ready_sec": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimal_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
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
      "update_policy": {
        "block": {
          "attributes": {
            "max_surge_fixed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_surge_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_fixed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_ready_sec": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimal_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "version": {
        "block": {
          "attributes": {
            "instance_template": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "target_size": {
              "block": {
                "attributes": {
                  "fixed": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeInstanceGroupManagerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstanceGroupManager), &result)
	return &result
}
