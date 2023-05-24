package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsKeys = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "DNS keys identifier",
        "description_kind": "markdown",
        "type": "string"
      },
      "key_signing_keys": {
        "computed": true,
        "description": "A list of Key-signing key (KSK) records.",
        "description_kind": "markdown",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "creation_time": "string",
              "description": "string",
              "digests": [
                "list",
                [
                  "object",
                  {
                    "digest": "string",
                    "type": "string"
                  }
                ]
              ],
              "ds_record": "string",
              "id": "string",
              "is_active": "bool",
              "key_length": "number",
              "key_tag": "number",
              "public_key": "string"
            }
          ]
        ]
      },
      "managed_zone": {
        "description": "The Name of the zone.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project for the Google Cloud.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "zone_signing_keys": {
        "computed": true,
        "description": "A list of Zone-signing key (ZSK) records.",
        "description_kind": "markdown",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "creation_time": "string",
              "description": "string",
              "digests": [
                "list",
                [
                  "object",
                  {
                    "digest": "string",
                    "type": "string"
                  }
                ]
              ],
              "id": "string",
              "is_active": "bool",
              "key_length": "number",
              "key_tag": "number",
              "public_key": "string"
            }
          ]
        ]
      }
    },
    "description": "Get the DNSKEY and DS records of DNSSEC-signed managed zones",
    "description_kind": "markdown"
  },
  "version": 0
}`

func GoogleDnsKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsKeys), &result)
	return &result
}
