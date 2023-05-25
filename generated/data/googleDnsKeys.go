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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_signing_keys": {
        "computed": true,
        "description_kind": "plain",
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
      "zone_signing_keys": {
        "computed": true,
        "description_kind": "plain",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDnsKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsKeys), &result)
	return &result
}
