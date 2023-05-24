package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerCertificateMapEntry = `{
  "block": {
    "attributes": {
      "certificates": {
        "description": "A set of Certificates defines for the given hostname.\nThere can be defined up to fifteen certificates in each Certificate Map Entry.\nEach certificate must match pattern projects/*/locations/*/certificates/*.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC \"Zulu\" format,\nwith nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hostname": {
        "description": "A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)\nfor a set of hostnames with common suffix. Used as Server Name Indication (SNI) for\nselecting a proper certificate.",
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
        "computed": true,
        "description": "Set of labels associated with a Certificate Map Entry.\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "map": {
        "description": "A map entry that is inputted into the cetrificate map",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "matcher": {
        "description": "A predefined matcher for particular cases, other than SNI selection",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the Certificate Map Entry. Certificate Map Entry\nnames must be unique globally and match pattern\n'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'",
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
      "state": {
        "computed": true,
        "description": "A serving state of this Certificate Map Entry.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC \"Zulu\" format,\nwith nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleCertificateManagerCertificateMapEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerCertificateMapEntry), &result)
	return &result
}
