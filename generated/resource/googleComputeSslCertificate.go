package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSslCertificate = `{
  "block": {
    "attributes": {
      "certificate": {
        "description": "The certificate in PEM format.\nThe certificate chain must be no greater than 5 certs long.\nThe chain must include at least one intermediate cert.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "certificate_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
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
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "Expire time of the certificate in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.\n\n\nThese are in the same namespace as the managed SSL certificates.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description": "Creates a unique name beginning with the specified prefix. Conflicts with name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_key": {
        "description": "The write-only private key in PEM format.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
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

func GoogleComputeSslCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSslCertificate), &result)
	return &result
}
