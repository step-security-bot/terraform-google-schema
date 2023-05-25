package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesEdgeCacheOrigin = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover_origin": {
        "description": "The Origin resource to try when the current origin cannot be reached.\nAfter maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.\n\nThe value of timeout.maxAttemptsTimeout dictates the timeout across all origins.\nA reference to a Topic resource.",
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
        "description": "Set of label tags associated with the EdgeCache resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "max_attempts": {
        "description": "The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.\n\nOnce maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,\nretryConditions and failoverOrigin to control its own cache fill failures.\n\nThe total number of allowed attempts to cache fill across this and failover origins is limited to four.\nThe total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.\n\nThe last valid response from an origin will be returned to the client.\nIf no origin returns a valid response, an HTTP 503 will be returned to the client.\n\nDefaults to 1. Must be a value greater than 0 and less than 4.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,\nand all following characters must be a dash, underscore, letter or digit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_address": {
        "description": "A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.\n\nThis address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname\n\nWhen providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.\nIf a Cloud Storage bucket is provided, it must be in the canonical \"gs://bucketname\" format. Other forms, such as \"storage.googleapis.com\", will be rejected.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port to connect to the origin on.\nDefaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.",
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
      "protocol": {
        "computed": true,
        "description": "The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security \u0026 performance.\n\nWhen using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: [\"HTTP2\", \"HTTPS\", \"HTTP\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retry_conditions": {
        "computed": true,
        "description": "Specifies one or more retry conditions for the configured origin.\n\nIf the failure mode during a connection attempt to the origin matches the configured retryCondition(s),\nthe origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.\n\nThe default retryCondition is \"CONNECT_FAILURE\".\n\nretryConditions apply to this origin, and not subsequent failoverOrigin(s),\nwhich may specify their own retryConditions and maxAttempts.\n\nValid values are:\n\n- CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.\n- HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.\n- GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.\n- RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)\n- NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet. Possible values: [\"CONNECT_FAILURE\", \"HTTP_5XX\", \"GATEWAY_ERROR\", \"RETRIABLE_4XX\", \"NOT_FOUND\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "timeout": {
        "block": {
          "attributes": {
            "connect_timeout": {
              "description": "The maximum duration to wait for the origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.\n\nDefaults to 5 seconds. The timeout must be a value between 1s and 15s.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_attempts_timeout": {
              "description": "The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 503 will be returned if the timeout is reached before a response is returned.\n\nDefaults to 5 seconds. The timeout must be a value between 1s and 15s.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_timeout": {
              "description": "The maximum duration to wait for data to arrive when reading from the HTTP connection/stream.\n\nDefaults to 5 seconds. The timeout must be a value between 1s and 30s.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The connection and HTTP timeout configuration for this origin.",
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

func GoogleNetworkServicesEdgeCacheOriginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesEdgeCacheOrigin), &result)
	return &result
}
