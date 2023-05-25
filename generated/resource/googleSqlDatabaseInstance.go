package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabaseInstance = `{
  "block": {
    "attributes": {
      "connection_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "first_ip_address": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "time_to_retire": "string",
              "type": "string"
            }
          ]
        ]
      },
      "master_instance_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
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
      "server_ca_cert": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert": "string",
              "common_name": "string",
              "create_time": "string",
              "expiration_time": "string",
              "sha1_fingerprint": "string"
            }
          ]
        ]
      },
      "service_account_email_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "replica_configuration": {
        "block": {
          "attributes": {
            "ca_certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connect_retry_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dump_file_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failover_target": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "master_heartbeat_period": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "ssl_cipher": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_server_certificate": {
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
      "settings": {
        "block": {
          "attributes": {
            "activation_policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authorized_gae_applications": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "availability_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "crash_safe_replication": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_autoresize": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_size": {
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
            "pricing_plan": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "replication_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_labels": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "block_types": {
            "backup_configuration": {
              "block": {
                "attributes": {
                  "binary_log_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start_time": {
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
            "database_flags": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ip_configuration": {
              "block": {
                "attributes": {
                  "ipv4_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "private_network": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_ssl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "authorized_networks": {
                    "block": {
                      "attributes": {
                        "expiration_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
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
            "location_preference": {
              "block": {
                "attributes": {
                  "follow_gae_application": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zone": {
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
            "maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "hour": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "update_track": {
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

func GoogleSqlDatabaseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabaseInstance), &result)
	return &result
}
