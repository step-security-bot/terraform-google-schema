package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerNodePool = `{
  "block": {
    "attributes": {
      "cluster": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initial_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instance_group_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_pods_per_node": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling": {
        "block": {
          "attributes": {
            "max_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
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
      "management": {
        "block": {
          "attributes": {
            "auto_repair": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_upgrade": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "disk_size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "guest_accelerator": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "type": "string"
                  }
                ]
              ]
            },
            "image_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "local_ssd_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "machine_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "min_cpu_platform": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oauth_scopes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "preemptible": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_account": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "sandbox_config": {
              "block": {
                "attributes": {
                  "sandbox_type": {
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
            "shielded_instance_config": {
              "block": {
                "attributes": {
                  "enable_integrity_monitoring": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_secure_boot": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "taint": {
              "block": {
                "attributes": {
                  "effect": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "workload_metadata_config": {
              "block": {
                "attributes": {
                  "node_metadata": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
  "version": 1
}`

func GoogleContainerNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerNodePool), &result)
	return &result
}
