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
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional textual description of the resource; provided by the\nclient when the resource is created.",
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
            "signed_url_cache_max_age_sec": {
              "description": "Maximum number of seconds the response to a signed URL request will\nbe considered fresh. Defaults to 1hr (3600s). After this time period,\nthe response will be revalidated before being served.\nWhen serving responses to signed URL requests,\nCloud CDN will internally behave as though\nall responses from this backend had a \"Cache-Control: public,\nmax-age=[TTL]\" header, regardless of any existing Cache-Control\nheader. The actual headers served in responses will not be altered.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
  "version": 0
}`

func GoogleComputeBackendBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendBucket), &result)
	return &result
}
