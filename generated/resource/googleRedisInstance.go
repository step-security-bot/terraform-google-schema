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
      "authorized_network": {
        "computed": true,
        "description": "The full name of the Google Compute Engine network to which the\ninstance is connected. If left unspecified, the default network\nwill be used.",
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
        "description": "The version of Redis software. If not provided, latest supported\nversion will be used. Currently, the supported values are:\n\n- REDIS_4_0 for Redis 4.0 compatibility\n- REDIS_3_2 for Redis 3.2 compatibility",
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
      "reserved_ip_range": {
        "computed": true,
        "description": "The CIDR range of internal addresses that are reserved for this\ninstance. If not provided, the service will choose an unused /29\nblock, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be\nunique and non-overlapping with existing subnets in an authorized\nnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tier": {
        "description": "The service tier of the instance. Must be one of these values:\n\n- BASIC: standalone instance\n- STANDARD_HA: highly available primary/replica instances",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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
