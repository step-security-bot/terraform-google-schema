package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMemcacheInstance = `{
  "block": {
    "attributes": {
      "authorized_network": {
        "computed": true,
        "description": "The full name of the GCE network to connect the instance to.  If not provided,\n'default' will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_endpoint": {
        "computed": true,
        "description": "Endpoint for Discovery API",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "A user-visible name for the instance.",
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
      "labels": {
        "description": "Resource labels to represent user-provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "maintenance_schedule": {
        "computed": true,
        "description": "Output only. Published maintenance schedule.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_time": "string",
              "schedule_deadline_time": "string",
              "start_time": "string"
            }
          ]
        ]
      },
      "memcache_full_version": {
        "computed": true,
        "description": "The full version of memcached server running on this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "memcache_nodes": {
        "computed": true,
        "description": "Additional information about the instance state, if available.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "node_id": "string",
              "port": "number",
              "state": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "memcache_version": {
        "description": "The major version of Memcached software. If not provided, latest supported version will be used.\nCurrently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically\ndetermined by our system based on the latest supported minor version. Default value: \"MEMCACHE_1_5\" Possible values: [\"MEMCACHE_1_5\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "description": "Number of nodes in the memcache instance.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The region of the Memcache instance. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zones": {
        "computed": true,
        "description": "Zones where memcache nodes should be provisioned.  If not\nprovided, all zones will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "maintenance_policy": {
        "block": {
          "attributes": {
            "create_time": {
              "computed": true,
              "description": "Output only. The time when the policy was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "description": "Optional. Description of what this policy is for.\nCreate/Update methods return INVALID_ARGUMENT if the\nlength is greater than 512.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_time": {
              "computed": true,
              "description": "Output only. The time when the policy was updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "weekly_maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Required. The day of week that maintenance updates occur.\n- DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.\n- MONDAY: Monday\n- TUESDAY: Tuesday\n- WEDNESDAY: Wednesday\n- THURSDAY: Thursday\n- FRIDAY: Friday\n- SATURDAY: Saturday\n- SUNDAY: Sunday Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
                    "description": "Required. The length of the maintenance window, ranging from 3 hours to 8 hours.\nA duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "start_time": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of day in 24 hour format. Should be from 0 to 23.\nAn API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "description": "Minutes of hour of day. Must be from 0 to 59.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "nanos": {
                          "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Seconds of minutes of the time. Must normally be from 0 to 59.\nAn API may allow the value 60 if it allows leap-seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Required. Start time of the window in UTC time.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Required. Maintenance window that is applied to resources covered by this policy.\nMinimum 1. For the current version, the maximum number of weekly_maintenance_windows\nis expected to be one.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Maintenance policy for an instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "memcache_parameters": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "This is a unique ID associated with this set of parameters.",
              "description_kind": "plain",
              "type": "string"
            },
            "params": {
              "description": "User-defined set of parameters to use in the memcache process.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "User-specified parameters for this memcache instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "cpu_count": {
              "description": "Number of CPUs per node.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "memory_size_mb": {
              "description": "Memory size in Mebibytes for each memcache node.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Configuration for memcache nodes.",
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

func GoogleMemcacheInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMemcacheInstance), &result)
	return &result
}
