package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFilestoreInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server-specified ETag for the instance resource to prevent\nsimultaneous updates from overwriting each other.",
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
        "description": "Resource labels to represent user-provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The resource name of the instance.",
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
      "tier": {
        "description": "The service tier of the instance. Possible values: [\"TIER_UNSPECIFIED\", \"STANDARD\", \"PREMIUM\", \"BASIC_HDD\", \"BASIC_SSD\", \"HIGH_SCALE_SSD\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "description": "The name of the Filestore zone of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "file_shares": {
        "block": {
          "attributes": {
            "capacity_gb": {
              "description": "File share capacity in GiB. This must be at least 1024 GiB\nfor the standard tier, or 2560 GiB for the premium tier.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "name": {
              "description": "The name of the fileshare (16 characters or less)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "File system shares on the instance. For this version, only a\nsingle file share is supported.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "networks": {
        "block": {
          "attributes": {
            "ip_addresses": {
              "computed": true,
              "description": "A list of IPv4 or IPv6 addresses.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "modes": {
              "description": "IP versions for which the instance has\nIP addresses assigned. Possible values: [\"ADDRESS_MODE_UNSPECIFIED\", \"MODE_IPV4\", \"MODE_IPV6\"]",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "network": {
              "description": "The name of the GCE VPC network to which the\ninstance is connected.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "reserved_ip_range": {
              "computed": true,
              "description": "A /29 CIDR block that identifies the range of IP\naddresses reserved for this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "VPC networks to which the instance is connected. For this version,\nonly a single network is supported.",
          "description_kind": "plain"
        },
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

func GoogleFilestoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFilestoreInstance), &result)
	return &result
}
