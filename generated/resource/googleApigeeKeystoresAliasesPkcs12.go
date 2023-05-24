package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeKeystoresAliasesPkcs12 = `{
  "block": {
    "attributes": {
      "alias": {
        "description": "Alias Name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certs_info": {
        "computed": true,
        "description": "Chain of certificates under this alias.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert_info": [
                "list",
                [
                  "object",
                  {
                    "basic_constraints": "string",
                    "expiry_date": "string",
                    "is_valid": "string",
                    "issuer": "string",
                    "public_key": "string",
                    "serial_number": "string",
                    "sig_alg_name": "string",
                    "subject": "string",
                    "subject_alternative_names": [
                      "list",
                      "string"
                    ],
                    "valid_from": "string",
                    "version": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "environment": {
        "description": "Environment associated with the alias",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file": {
        "description": "Cert content",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filehash": {
        "description": "Hash of the pkcs file",
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
      "keystore": {
        "description": "Keystore Name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "Organization ID associated with the alias",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "Password for the Private Key if it's encrypted",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Optional.Type of Alias",
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

func GoogleApigeeKeystoresAliasesPkcs12Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeKeystoresAliasesPkcs12), &result)
	return &result
}
