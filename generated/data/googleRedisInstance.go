package data

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
        "type": "string"
      },
      "auth_enabled": {
        "computed": true,
        "description": "Optional. Indicates whether OSS Redis AUTH is enabled for the\ninstance. If set to \"true\" AUTH is enabled on the instance.\nDefault value is \"false\" meaning AUTH is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "auth_string": {
        "computed": true,
        "description": "AUTH String set on the instance. This field will only be populated if auth_enabled is true.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorized_network": {
        "computed": true,
        "description": "The full name of the Google Compute Engine network to which the\ninstance is connected. If left unspecified, the default network\nwill be used.",
        "description_kind": "plain",
        "type": "string"
      },
      "connect_mode": {
        "computed": true,
        "description": "The connection mode of the Redis instance. Default value: \"DIRECT_PEERING\" Possible values: [\"DIRECT_PEERING\", \"PRIVATE_SERVICE_ACCESS\"]",
        "description_kind": "plain",
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
        "computed": true,
        "description": "An arbitrary and optional user-provided name for the instance.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location_id": {
        "computed": true,
        "description": "The zone where the instance will be provisioned. If not provided,\nthe service will choose a zone for the instance. For STANDARD_HA tier,\ninstances will be created across two zones for protection against\nzonal failures. If [alternativeLocationId] is also provided, it must\nbe different from [locationId].",
        "description_kind": "plain",
        "type": "string"
      },
      "memory_size_gb": {
        "computed": true,
        "description": "Redis memory size in GiB.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The ID of the instance or a fully qualified identifier for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redis_configs": {
        "computed": true,
        "description": "Redis configuration parameters, according to http://redis.io/topics/config.\nPlease check Memorystore documentation for the list of supported parameters:\nhttps://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "redis_version": {
        "computed": true,
        "description": "The version of Redis software. If not provided, latest supported\nversion will be used. Please check the API documentation linked \nat the top for the latest valid values.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "description": "The name of the Redis region of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_ip_range": {
        "computed": true,
        "description": "The CIDR range of internal addresses that are reserved for this\ninstance. If not provided, the service will choose an unused /29\nblock, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be\nunique and non-overlapping with existing subnets in an authorized\nnetwork.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The service tier of the instance. Must be one of these values:\n\n- BASIC: standalone instance\n- STANDARD_HA: highly available primary/replica instances Default value: \"BASIC\" Possible values: [\"BASIC\", \"STANDARD_HA\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_encryption_mode": {
        "computed": true,
        "description": "The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.\n\n- SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentcation Default value: \"DISABLED\" Possible values: [\"SERVER_AUTHENTICATION\", \"DISABLED\"]",
        "description_kind": "plain",
        "type": "string"
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
