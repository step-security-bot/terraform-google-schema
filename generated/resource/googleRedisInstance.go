package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleRedisInstance = `{
  "block": {
    "attributes": {
      "alternative_location_id": {
        "computed": true,
        "description": "Only applicable to STANDARD_HA tier which protects the instance\nagainst zonal failures by provisioning it across two zones.\nIf provided, it must be a different zone from the one provided in\n[locationId].",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auth_enabled": {
        "description": "Optional. Indicates whether OSS Redis AUTH is enabled for the\ninstance. If set to \"true\" AUTH is enabled on the instance.\nDefault value is \"false\" meaning AUTH is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auth_string": {
        "computed": true,
        "description": "AUTH String set on the instance. This field will only be populated if auth_enabled is true.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "authorized_network": {
        "computed": true,
        "description": "The full name of the Google Compute Engine network to which the\ninstance is connected. If left unspecified, the default network\nwill be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connect_mode": {
        "description": "The connection mode of the Redis instance. Default value: \"DIRECT_PEERING\" Possible values: [\"DIRECT_PEERING\", \"PRIVATE_SERVICE_ACCESS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the instance was created in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_location_id": {
        "computed": true,
        "description": "The current zone where the Redis endpoint is placed.\nFor Basic Tier instances, this will always be the same as the\n[locationId] provided by the user at creation time. For Standard Tier\ninstances, this can be either [locationId] or [alternativeLocationId]\nand can change after a failover event.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "description": "Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis\ninstance. If this is provided, CMEK is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "An arbitrary and optional user-provided name for the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "computed": true,
        "description": "Hostname or IP address of the exposed Redis endpoint used by clients\nto connect to the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location_id": {
        "computed": true,
        "description": "The zone where the instance will be provisioned. If not provided,\nthe service will choose a zone for the instance. For STANDARD_HA tier,\ninstances will be created across two zones for protection against\nzonal failures. If [alternativeLocationId] is also provided, it must\nbe different from [locationId].",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_schedule": {
        "computed": true,
        "description": "Upcoming maintenance schedule.",
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
      "memory_size_gb": {
        "description": "Redis memory size in GiB.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "name": {
        "description": "The ID of the instance or a fully qualified identifier for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nodes": {
        "computed": true,
        "description": "Output only. Info per node.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "persistence_iam_identity": {
        "computed": true,
        "description": "Output only. Cloud IAM identity used by import / export operations\nto transfer data to/from Cloud Storage. Format is \"serviceAccount:\".\nThe value may change over time for a given instance so should be\nchecked before each import/export operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number of the exposed Redis endpoint.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_endpoint": {
        "computed": true,
        "description": "Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.\nTargets all healthy replica nodes in instance. Replication is asynchronous and replica nodes\nwill exhibit some lag behind the primary. Write requests must target 'host'.",
        "description_kind": "plain",
        "type": "string"
      },
      "read_endpoint_port": {
        "computed": true,
        "description": "Output only. The port number of the exposed readonly redis endpoint. Standard tier only.\nWrite requests should target 'port'.",
        "description_kind": "plain",
        "type": "number"
      },
      "read_replicas_mode": {
        "computed": true,
        "description": "Optional. Read replica mode. Can only be specified when trying to create the instance.\nIf not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.\n- READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the\ninstance cannot scale up or down the number of replicas.\n- READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance\ncan scale up and down the number of replicas. Possible values: [\"READ_REPLICAS_DISABLED\", \"READ_REPLICAS_ENABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redis_configs": {
        "description": "Redis configuration parameters, according to http://redis.io/topics/config.\nPlease check Memorystore documentation for the list of supported parameters:\nhttps://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "redis_version": {
        "computed": true,
        "description": "The version of Redis software. If not provided, latest supported\nversion will be used. Please check the API documentation linked\nat the top for the latest valid values.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The name of the Redis region of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_count": {
        "computed": true,
        "description": "Optional. The number of replica nodes. The valid range for the Standard Tier with\nread replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled\nfor a Standard Tier instance, the only valid value is 1 and the default is 1.\nThe valid value for basic tier is 0 and the default is also 0.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "reserved_ip_range": {
        "computed": true,
        "description": "The CIDR range of internal addresses that are reserved for this\ninstance. If not provided, the service will choose an unused /29\nblock, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be\nunique and non-overlapping with existing subnets in an authorized\nnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_ip_range": {
        "computed": true,
        "description": "Optional. Additional IP range for node placement. Required when enabling read replicas on\nan existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or\n\"auto\". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address\nrange associated with the private service access connection, or \"auto\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_ca_certs": {
        "computed": true,
        "description": "List of server CA certificates for the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert": "string",
              "create_time": "string",
              "expire_time": "string",
              "serial_number": "string",
              "sha1_fingerprint": "string"
            }
          ]
        ]
      },
      "tier": {
        "description": "The service tier of the instance. Must be one of these values:\n\n- BASIC: standalone instance\n- STANDARD_HA: highly available primary/replica instances Default value: \"BASIC\" Possible values: [\"BASIC\", \"STANDARD_HA\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_encryption_mode": {
        "description": "The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.\n\n- SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication Default value: \"DISABLED\" Possible values: [\"SERVER_AUTHENTICATION\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "maintenance_policy": {
        "block": {
          "attributes": {
            "create_time": {
              "computed": true,
              "description": "Output only. The time when the policy was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
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
              "description": "Output only. The time when the policy was last updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "weekly_maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Required. The day of week that maintenance updates occur.\n\n- DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.\n- MONDAY: Monday\n- TUESDAY: Tuesday\n- WEDNESDAY: Wednesday\n- THURSDAY: Thursday\n- FRIDAY: Friday\n- SATURDAY: Saturday\n- SUNDAY: Sunday Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
                    "computed": true,
                    "description": "Output only. Duration of the maintenance window.\nThe current window is fixed at 1 hour.\nA duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
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
                "description": "Optional. Maintenance window that is applied to resources covered by this policy.\nMinimum 1. For the current version, the maximum number\nof weekly_window is expected to be one.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Maintenance policy for an instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "persistence_config": {
        "block": {
          "attributes": {
            "persistence_mode": {
              "computed": true,
              "description": "Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.\n\n- DISABLED: \tPersistence is disabled for the instance, and any existing snapshots are deleted.\n- RDB: RDB based Persistence is enabled. Possible values: [\"DISABLED\", \"RDB\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rdb_next_snapshot_time": {
              "computed": true,
              "description": "Output only. The next time that a snapshot attempt is scheduled to occur.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up\nto nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "type": "string"
            },
            "rdb_snapshot_period": {
              "description": "Optional. Available snapshot periods for scheduling.\n\n- ONE_HOUR:\tSnapshot every 1 hour.\n- SIX_HOURS:\tSnapshot every 6 hours.\n- TWELVE_HOURS:\tSnapshot every 12 hours.\n- TWENTY_FOUR_HOURS:\tSnapshot every 24 hours. Possible values: [\"ONE_HOUR\", \"SIX_HOURS\", \"TWELVE_HOURS\", \"TWENTY_FOUR_HOURS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rdb_snapshot_start_time": {
              "computed": true,
              "description": "Optional. Date and time that the first snapshot was/will be attempted,\nand to which future snapshots will be aligned. If not provided,\nthe current time will be used.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution\nand up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Persistence configuration for an instance.",
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

func GoogleRedisInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleRedisInstance), &result)
	return &result
}
