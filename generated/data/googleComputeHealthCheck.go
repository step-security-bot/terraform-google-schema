package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeHealthCheck = `{
  "block": {
    "attributes": {
      "check_interval_sec": {
        "computed": true,
        "description": "How often (in seconds) to send a health check. The default value is 5\nseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "grpc_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "grpc_service_name": "string",
              "port": "number",
              "port_name": "string",
              "port_specification": "string"
            }
          ]
        ]
      },
      "healthy_threshold": {
        "computed": true,
        "description": "A so-far unhealthy instance will be marked healthy after this many\nconsecutive successes. The default value is 2.",
        "description_kind": "plain",
        "type": "number"
      },
      "http2_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "port": "number",
              "port_name": "string",
              "port_specification": "string",
              "proxy_header": "string",
              "request_path": "string",
              "response": "string"
            }
          ]
        ]
      },
      "http_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "port": "number",
              "port_name": "string",
              "port_specification": "string",
              "proxy_header": "string",
              "request_path": "string",
              "response": "string"
            }
          ]
        ]
      },
      "https_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "port": "number",
              "port_name": "string",
              "port_specification": "string",
              "proxy_header": "string",
              "request_path": "string",
              "response": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_config": {
        "computed": true,
        "description": "Configure logging on this health check.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable": "bool"
            }
          ]
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "port": "number",
              "port_name": "string",
              "port_specification": "string",
              "proxy_header": "string",
              "request": "string",
              "response": "string"
            }
          ]
        ]
      },
      "tcp_health_check": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "port": "number",
              "port_name": "string",
              "port_specification": "string",
              "proxy_header": "string",
              "request": "string",
              "response": "string"
            }
          ]
        ]
      },
      "timeout_sec": {
        "computed": true,
        "description": "How long (in seconds) to wait before claiming failure.\nThe default value is 5 seconds.  It is invalid for timeoutSec to have\ngreater value than checkIntervalSec.",
        "description_kind": "plain",
        "type": "number"
      },
      "type": {
        "computed": true,
        "description": "The type of the health check. One of HTTP, HTTPS, TCP, or SSL.",
        "description_kind": "plain",
        "type": "string"
      },
      "unhealthy_threshold": {
        "computed": true,
        "description": "A so-far healthy instance will be marked unhealthy after this many\nconsecutive failures. The default value is 2.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeHealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeHealthCheck), &result)
	return &result
}
