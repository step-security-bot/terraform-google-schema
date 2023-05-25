package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstanceFromTemplate = `{
  "block": {
    "attributes": {
      "allow_stopping_for_update": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "attached_disk": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "kms_key_self_link": "string",
              "mode": "string",
              "source": "string"
            }
          ]
        ]
      },
      "can_ip_forward": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cpu_platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_display": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "hostname": {
        "computed": true,
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
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description_kind": "plain",
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
      "metadata_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_startup_script": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_cpu_platform": {
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
      "scratch_disk": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "interface": "string"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "email": "string",
              "scopes": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "source_instance_template": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags_fingerprint": {
        "computed": true,
        "description_kind": "plain",
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
      "boot_disk": {
        "block": {
          "attributes": {
            "auto_delete": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_key_raw": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "disk_encryption_key_sha256": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_self_link": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "initialize_params": {
              "block": {
                "attributes": {
                  "image": {
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
                  "size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
      "network_interface": {
        "block": {
          "attributes": {
            "access_config": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "assigned_nat_ip": "string",
                    "nat_ip": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ]
            },
            "address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "alias_ip_range": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "ip_cidr_range": "string",
                    "subnetwork_range_name": "string"
                  }
                ]
              ]
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "network": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork_project": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "scheduling": {
        "block": {
          "attributes": {
            "automatic_restart": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "on_host_maintenance": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preemptible": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "node_affinities": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_secure_boot": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_vtpm": {
              "computed": true,
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

func GoogleComputeInstanceFromTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstanceFromTemplate), &result)
	return &result
}
