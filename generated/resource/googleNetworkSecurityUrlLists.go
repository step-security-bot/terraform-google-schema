package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityUrlLists = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Time when the security policy was created.\nA timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.\nExamples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Free-text description of the resource.",
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
      "location": {
        "description": "The location of the url lists.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Short name of the UrlList resource to be created.\nThis value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. 'urlList'.",
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
        "description": "Output only. Time when the security policy was updated.\nA timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.\nExamples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "values": {
        "description": "FQDNs and URLs.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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

func GoogleNetworkSecurityUrlListsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityUrlLists), &result)
	return &result
}
