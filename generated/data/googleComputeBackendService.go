package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendService = `{
  "block": {
    "attributes": {
      "affinity_cookie_ttl_sec": {
        "computed": true,
        "description": "Lifetime of cookies in seconds if session_affinity is\nGENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts\nonly until the end of the browser session (or equivalent). The\nmaximum allowed value for TTL is one day.\n\nWhen the load balancing scheme is INTERNAL, this field is not used.",
        "description_kind": "plain",
        "type": "number"
      },
      "backend": {
        "computed": true,
        "description": "The set of backends that serve this BackendService.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "balancing_mode": "string",
              "capacity_scaler": "number",
              "description": "string",
              "group": "string",
              "max_connections": "number",
              "max_connections_per_endpoint": "number",
              "max_connections_per_instance": "number",
              "max_rate": "number",
              "max_rate_per_endpoint": "number",
              "max_rate_per_instance": "number",
              "max_utilization": "number"
            }
          ]
        ]
      },
      "cdn_policy": {
        "computed": true,
        "description": "Cloud CDN configuration for this BackendService.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cache_key_policy": [
                "list",
                [
                  "object",
                  {
                    "include_host": "bool",
                    "include_protocol": "bool",
                    "include_query_string": "bool",
                    "query_string_blacklist": [
                      "set",
                      "string"
                    ],
                    "query_string_whitelist": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "signed_url_cache_max_age_sec": "number"
            }
          ]
        ]
      },
      "connection_draining_timeout_sec": {
        "computed": true,
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
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
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cdn": {
        "computed": true,
        "description": "If true, enable Cloud CDN for this BackendService.",
        "description_kind": "plain",
        "type": "bool"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this\nobject. This field is used in optimistic locking.",
        "description_kind": "plain",
        "type": "string"
      },
      "health_checks": {
        "computed": true,
        "description": "The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource\nfor health checking this BackendService. Currently at most one health\ncheck can be specified, and a health check is required.\n\nFor internal load balancing, a URL to a HealthCheck resource must be specified instead.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iap": {
        "computed": true,
        "description": "Settings for enabling Cloud Identity Aware Proxy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "oauth2_client_id": "string",
              "oauth2_client_secret": "string",
              "oauth2_client_secret_sha256": "string"
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
      "load_balancing_scheme": {
        "computed": true,
        "description": "Indicates whether the backend service will be used with internal or\nexternal load balancing. A backend service created for one type of\nload balancing cannot be used with the other. Must be 'EXTERNAL' or\n'INTERNAL_SELF_MANAGED' for a global backend service. Defaults to 'EXTERNAL'.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port_name": {
        "computed": true,
        "description": "Name of backend port. The same name should appear in the instance\ngroups referenced by this service. Required when the load balancing\nscheme is EXTERNAL.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol this BackendService uses to communicate with backends.\nPossible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is\nHTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer\ntypes and may result in errors if used with the GA API.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_policy": {
        "computed": true,
        "description": "The security policy associated with this backend service.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP.",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout_sec": {
        "computed": true,
        "description": "How many seconds to wait for the backend before considering it a\nfailed request. Default is 30 seconds. Valid range is [1, 86400].",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleComputeBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendService), &result)
	return &result
}
