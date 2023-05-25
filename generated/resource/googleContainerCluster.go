package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerCluster = `{
  "block": {
    "attributes": {
      "additional_zones": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cluster_autoscaling": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "resource_limits": [
                "list",
                [
                  "object",
                  {
                    "maximum": "number",
                    "minimum": "number",
                    "resource_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cluster_ipv4_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_max_pods_per_node": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_binary_authorization": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_intranode_visibility": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_kubernetes_alpha": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_legacy_abac": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_tpu": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
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
      "initial_node_count": {
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
      "ip_allocation_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "cluster_ipv4_cidr_block": "string",
              "cluster_secondary_range_name": "string",
              "create_subnetwork": "bool",
              "node_ipv4_cidr_block": "string",
              "services_ipv4_cidr_block": "string",
              "services_secondary_range_name": "string",
              "subnetwork_name": "string",
              "use_ip_aliases": "bool"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_service": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "min_master_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitoring_service": {
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
      "network": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_locations": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "node_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remove_default_node_pool": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_labels": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "services_ipv4_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
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
      "addons_config": {
        "block": {
          "block_types": {
            "horizontal_pod_autoscaling": {
              "block": {
                "attributes": {
                  "disabled": {
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
            "http_load_balancing": {
              "block": {
                "attributes": {
                  "disabled": {
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
            "kubernetes_dashboard": {
              "block": {
                "attributes": {
                  "disabled": {
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
            "network_policy_config": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "maintenance_policy": {
        "block": {
          "block_types": {
            "daily_maintenance_window": {
              "block": {
                "attributes": {
                  "duration": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "master_auth": {
        "block": {
          "attributes": {
            "client_certificate": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "client_key": {
              "computed": true,
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            },
            "cluster_ca_certificate": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "client_certificate_config": {
              "block": {
                "attributes": {
                  "issue_client_certificate": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
      "master_authorized_networks_config": {
        "block": {
          "block_types": {
            "cidr_blocks": {
              "block": {
                "attributes": {
                  "cidr_block": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "display_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "network_policy": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "provider": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "node_pool": {
        "block": {
          "attributes": {
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
            "version": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "pod_security_policy_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_cluster_config": {
        "block": {
          "attributes": {
            "enable_private_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_private_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "master_ipv4_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_endpoint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "public_endpoint": {
              "computed": true,
              "description_kind": "plain",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleContainerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerCluster), &result)
	return &result
}
