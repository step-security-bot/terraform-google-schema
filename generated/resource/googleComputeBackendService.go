package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendService = `{
  "block": {
    "attributes": {
      "affinity_cookie_ttl_sec": {
        "description": "Lifetime of cookies in seconds if session_affinity is\nGENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts\nonly until the end of the browser session (or equivalent). The\nmaximum allowed value for TTL is one day.\n\nWhen the load balancing scheme is INTERNAL, this field is not used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "connection_draining_timeout_sec": {
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
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
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_cdn": {
        "description": "If true, enable Cloud CDN for this BackendService.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this\nobject. This field is used in optimistic locking.",
        "description_kind": "plain",
        "type": "string"
      },
      "health_checks": {
        "description": "The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource\nfor health checking this BackendService. Currently at most one health\ncheck can be specified, and a health check is required.\n\nFor internal load balancing, a URL to a HealthCheck resource must be specified instead.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_scheme": {
        "description": "Indicates whether the backend service will be used with internal or\nexternal load balancing. A backend service created for one type of\nload balancing cannot be used with the other. Must be 'EXTERNAL' or\n'INTERNAL_SELF_MANAGED' for a global backend service. Defaults to 'EXTERNAL'.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol this BackendService uses to communicate with backends.\nPossible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is\nHTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer\ntypes and may result in errors if used with the GA API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_policy": {
        "description": "The security policy associated with this backend service.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "timeout_sec": {
        "computed": true,
        "description": "How many seconds to wait for the backend before considering it a\nfailed request. Default is 30 seconds. Valid range is [1, 86400].",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "backend": {
        "block": {
          "attributes": {
            "balancing_mode": {
              "description": "Specifies the balancing mode for this backend.\n\nFor global HTTP(S) or TCP/SSL load balancing, the default is\nUTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S))\nand CONNECTION (for TCP/SSL).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "capacity_scaler": {
              "description": "A multiplier applied to the group's maximum servicing capacity\n(based on UTILIZATION, RATE or CONNECTION).\n\nDefault value is 1, which means the group will serve up to 100%\nof its configured capacity (depending on balancingMode). A\nsetting of 0 means the group is completely drained, offering\n0% of its available Capacity. Valid range is [0.0,1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "description": {
              "description": "An optional description of this resource.\nProvide this property when you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group": {
              "description": "The fully-qualified URL of an Instance Group or Network Endpoint\nGroup resource. In case of instance group this defines the list\nof instances that serve traffic. Member virtual machine\ninstances from each instance group must live in the same zone as\nthe instance group itself. No two backends in a backend service\nare allowed to use same Instance Group resource.\n\nFor Network Endpoint Groups this defines list of endpoints. All\nendpoints of Network Endpoint Group must be hosted on instances\nlocated in the same zone as the Network Endpoint Group.\n\nBackend services cannot mix Instance Group and\nNetwork Endpoint Group backends.\n\nNote that you must specify an Instance Group or Network Endpoint\nGroup resource using the fully-qualified URL, rather than a\npartial URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_connections": {
              "description": "The max number of simultaneous connections for the group. Can\nbe used with either CONNECTION or UTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or one\nof maxConnectionsPerInstance or maxConnectionsPerEndpoint,\nas appropriate for group type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_endpoint": {
              "description": "The max number of simultaneous connections that a single backend\nnetwork endpoint can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either\nmaxConnections or maxConnectionsPerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_instance": {
              "description": "The max number of simultaneous connections that a single\nbackend instance can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or\nmaxConnectionsPerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate": {
              "description": "The max requests per second (RPS) of the group.\n\nCan be used with either RATE or UTILIZATION balancing modes,\nbut required if RATE mode. For RATE mode, either maxRate or one\nof maxRatePerInstance or maxRatePerEndpoint, as appropriate for\ngroup type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_endpoint": {
              "description": "The max requests per second (RPS) that a single backend network\nendpoint can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_instance": {
              "description": "The max requests per second (RPS) that a single backend\ninstance can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_utilization": {
              "description": "Used when balancingMode is UTILIZATION. This ratio defines the\nCPU utilization target for the group. The default is 0.8. Valid\nrange is [0.0, 1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "cdn_policy": {
        "block": {
          "attributes": {
            "signed_url_cache_max_age_sec": {
              "description": "Maximum number of seconds the response to a signed URL request\nwill be considered fresh, defaults to 1hr (3600s). After this\ntime period, the response will be revalidated before\nbeing served.\n\nWhen serving responses to signed URL requests, Cloud CDN will\ninternally behave as though all responses from this backend had a\n\"Cache-Control: public, max-age=[TTL]\" header, regardless of any\nexisting Cache-Control header. The actual headers served in\nresponses will not be altered.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "cache_key_policy": {
              "block": {
                "attributes": {
                  "include_host": {
                    "description": "If true requests to different hosts will be cached separately.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "include_protocol": {
                    "description": "If true, http and https requests will be cached separately.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "include_query_string": {
                    "description": "If true, include query string parameters in the cache key\naccording to query_string_whitelist and\nquery_string_blacklist. If neither is set, the entire query\nstring will be included.\n\nIf false, the query string will be excluded from the cache\nkey entirely.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "query_string_blacklist": {
                    "description": "Names of query string parameters to exclude in cache keys.\n\nAll other parameters will be included. Either specify\nquery_string_whitelist or query_string_blacklist, not both.\n'\u0026' and '=' will be percent encoded and not treated as\ndelimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "query_string_whitelist": {
                    "description": "Names of query string parameters to include in cache keys.\n\nAll other parameters will be excluded. Either specify\nquery_string_whitelist or query_string_blacklist, not both.\n'\u0026' and '=' will be percent encoded and not treated as\ndelimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "iap": {
        "block": {
          "attributes": {
            "oauth2_client_id": {
              "description": "OAuth2 Client ID for IAP",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "oauth2_client_secret": {
              "description": "OAuth2 Client Secret for IAP",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "oauth2_client_secret_sha256": {
              "computed": true,
              "description": "OAuth2 Client Secret SHA-256 for IAP",
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            }
          },
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
  "version": 1
}`

func GoogleComputeBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendService), &result)
	return &result
}
