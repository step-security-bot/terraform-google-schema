package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendBucket = `{
  "block": {
    "attributes": {
      "bucket_name": {
        "computed": true,
        "description": "Cloud Storage bucket name.",
        "description_kind": "plain",
        "type": "string"
      },
      "cdn_policy": {
        "computed": true,
        "description": "Cloud CDN configuration for this Backend Bucket.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cache_mode": "string",
              "client_ttl": "number",
              "default_ttl": "number",
              "max_ttl": "number",
              "negative_caching": "bool",
              "negative_caching_policy": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "ttl": "number"
                  }
                ]
              ],
              "serve_while_stale": "number",
              "signed_url_cache_max_age_sec": "number"
            }
          ]
        ]
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_response_headers": {
        "computed": true,
        "description": "Headers that the HTTP/S load balancer should add to proxied responses.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "An optional textual description of the resource; provided by the\nclient when the resource is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cdn": {
        "computed": true,
        "description": "If true, enable Cloud CDN for this BackendBucket.",
        "description_kind": "plain",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeBackendBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendBucket), &result)
	return &result
}
