package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerCertificateMap = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds with up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gclb_targets": {
        "computed": true,
        "description": "A list of target proxies that use this Certificate Map",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_configs": [
                "list",
                [
                  "object",
                  {
                    "ip_address": "string",
                    "ports": [
                      "list",
                      "number"
                    ]
                  }
                ]
              ],
              "target_https_proxy": "string",
              "target_ssl_proxy": "string"
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
      "labels": {
        "computed": true,
        "description": "Set of labels associated with a Certificate Map resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "A user-defined name of the Certificate Map. Certificate Map names must be unique\nglobally and match the pattern 'projects/*/locations/*/certificateMaps/*'.",
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
      "update_time": {
        "computed": true,
        "description": "Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds with up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
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

func GoogleCertificateManagerCertificateMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerCertificateMap), &result)
	return &result
}
