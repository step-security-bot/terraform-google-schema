package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendBucket = `{
  "block": {
    "attributes": {
      "bucket_name": {
        "description": "Cloud Storage bucket name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compression_mode": {
        "description": "Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: [\"AUTOMATIC\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_response_headers": {
        "description": "Headers that the HTTP/S load balancer should add to proxied responses.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "An optional textual description of the resource; provided by the\nclient when the resource is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_security_policy": {
        "description": "The security policy associated with this backend bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_cdn": {
        "description": "If true, enable Cloud CDN for this BackendBucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cdn_policy": {
        "block": {
          "attributes": {
            "cache_mode": {
              "computed": true,
              "description": "Specifies the cache setting for all responses from this backend.\nThe possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: [\"USE_ORIGIN_HEADERS\", \"FORCE_CACHE_ALL\", \"CACHE_ALL_STATIC\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_ttl": {
              "computed": true,
              "description": "Specifies the maximum allowed TTL for cached content served by this origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "default_ttl": {
              "computed": true,
              "description": "Specifies the default TTL for cached content served by this origin for responses\nthat do not have an existing valid TTL (max-age or s-max-age).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_ttl": {
              "computed": true,
              "description": "Specifies the maximum allowed TTL for cached content served by this origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "negative_caching": {
              "computed": true,
              "description": "Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "request_coalescing": {
              "description": "If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "serve_while_stale": {
              "computed": true,
              "description": "Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "signed_url_cache_max_age_sec": {
              "description": "Maximum number of seconds the response to a signed URL request will\nbe considered fresh. After this time period,\nthe response will be revalidated before being served.\nWhen serving responses to signed URL requests,\nCloud CDN will internally behave as though\nall responses from this backend had a \"Cache-Control: public,\nmax-age=[TTL]\" header, regardless of any existing Cache-Control\nheader. The actual headers served in responses will not be altered.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "bypass_cache_on_request_headers": {
              "block": {
                "attributes": {
                  "header_name": {
                    "description": "The header field name to match on when bypassing cache. Values are case-insensitive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.",
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "cache_key_policy": {
              "block": {
                "attributes": {
                  "include_http_headers": {
                    "description": "Allows HTTP request headers (by name) to be used in the\ncache key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_whitelist": {
                    "description": "Names of query string parameters to include in cache keys.\nDefault parameters are always included. '\u0026' and '=' will\nbe percent encoded and not treated as delimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The CacheKeyPolicy for this CdnPolicy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "negative_caching_policy": {
              "block": {
                "attributes": {
                  "code": {
                    "description": "The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501\ncan be specified as values, and you cannot specify a status code more than once.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ttl": {
                    "description": "The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s\n(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.\nOmitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Cloud CDN configuration for this Backend Bucket.",
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

func GoogleComputeBackendBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendBucket), &result)
	return &result
}
