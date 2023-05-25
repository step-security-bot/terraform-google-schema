package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstance = `{
  "block": {
    "attributes": {
      "allow_stopping_for_update": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "can_ip_forward": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_display": {
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_cpu_platform": {
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
      "tags": {
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
      "attached_disk": {
        "block": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_key_raw": {
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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "boot_disk": {
        "block": {
          "attributes": {
            "auto_delete": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "disk": {
        "block": {
          "attributes": {
            "auto_delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_key_raw": {
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
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scratch": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
          "block_types": {
            "access_config": {
              "block": {
                "attributes": {
                  "assigned_nat_ip": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "nat_ip": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_tier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "public_ptr_domain_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "alias_ip_range": {
              "block": {
                "attributes": {
                  "ip_cidr_range": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnetwork_range_name": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "scheduling": {
        "block": {
          "attributes": {
            "automatic_restart": {
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
      "scratch_disk": {
        "block": {
          "attributes": {
            "interface": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "service_account": {
        "block": {
          "attributes": {
            "email": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scopes": {
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
            },
            "enable_vtpm": {
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
  "version": 6
}`

func GoogleComputeInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstance), &result)
	return &result
}
