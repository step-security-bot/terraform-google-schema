package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocCluster = `{
  "block": {
    "attributes": {
      "id": {
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
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cluster_config": {
        "block": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "delete_autogen_bucket": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "staging_bucket": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
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
            "gce_cluster_config": {
              "block": {
                "attributes": {
                  "internal_ip_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "metadata": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "network": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_scopes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnetwork": {
                    "description_kind": "plain",
                    "optional": true,
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
                  "zone": {
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
            },
            "initialization_action": {
              "block": {
                "attributes": {
                  "script": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout_sec": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "master_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_instances": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "preemptible_worker_config": {
              "block": {
                "attributes": {
                  "instance_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "num_instances": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "software_config": {
              "block": {
                "attributes": {
                  "image_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "optional_components": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "override_properties": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "worker_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_instances": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
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
  "version": 0
}`

func GoogleDataprocClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocCluster), &result)
	return &result
}
