package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSslCertificate = `{
  "block": {
    "attributes": {
      "certificate": {
        "computed": true,
        "description": "The certificate in PEM format.\nThe certificate chain must be no greater than 5 certs long.\nThe chain must include at least one intermediate cert.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "An optional description of this resource.",
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.\n\n\nThese are in the same namespace as the managed SSL certificates.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description": "The write-only private key in PEM format.",
        "description_kind": "plain",
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

func GoogleComputeSslCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSslCertificate), &result)
	return &result
}
