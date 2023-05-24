package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerDnsAuthorization = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_resource_record": {
        "computed": true,
        "description": "The structure describing the DNS Resource Record that needs to be added\nto DNS configuration for the authorization to be usable by\ncertificate.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data": "string",
              "name": "string",
              "type": "string"
            }
          ]
        ]
      },
      "domain": {
        "description": "A domain which is being authorized. A DnsAuthorization resource covers a\nsingle domain and its wildcard, e.g. authorization for \"example.com\" can\nbe used to issue certificates for \"example.com\" and \"*.example.com\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of label tags associated with the DNS Authorization resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,\nand all following characters must be a dash, underscore, letter or digit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func GoogleCertificateManagerDnsAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerDnsAuthorization), &result)
	return &result
}
