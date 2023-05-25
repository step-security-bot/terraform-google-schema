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
