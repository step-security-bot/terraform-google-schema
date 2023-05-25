package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeHttpsHealthCheck = `{
  "block": {
    "attributes": {
      "check_interval_sec": {
        "description": "How often (in seconds) to send a health check. The default value is 5\nseconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "healthy_threshold": {
        "description": "A so-far unhealthy instance will be marked healthy after this many\nconsecutive successes. The default value is 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "host": {
        "description": "The value of the host header in the HTTPS health check request. If\nleft empty (default value), the public IP on behalf of which this\nhealth check is performed will be used.",
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
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description": "The TCP port number for the HTTPS health check request.\nThe default value is 80.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_path": {
        "description": "The request path of the HTTPS health check request.\nThe default value is /.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "timeout_sec": {
        "description": "How long (in seconds) to wait before claiming failure.\nThe default value is 5 seconds.  It is invalid for timeoutSec to have\ngreater value than checkIntervalSec.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "unhealthy_threshold": {
        "description": "A so-far healthy instance will be marked unhealthy after this many\nconsecutive failures. The default value is 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func GoogleComputeHttpsHealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeHttpsHealthCheck), &result)
	return &result
}
