package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeKeystoresAliasesKeyCertFile = `{
  "block": {
    "attributes": {
      "alias": {
        "description": "Alias Name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cert": {
        "description": "Cert content",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment": {
        "description": "Environment associated with the alias",
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
      "key": {
        "description": "Private Key content, omit if uploading to truststore",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
        "description": "Password for the Private Key if it's encrypted",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "certs_info": {
        "block": {
          "block_types": {
            "cert_info": {
              "block": {
                "attributes": {
                  "basic_constraints": {
                    "computed": true,
                    "description": "X.509 basic constraints extension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expiry_date": {
                    "computed": true,
                    "description": "X.509 notAfter validity period in milliseconds since epoch.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_valid": {
                    "computed": true,
                    "description": "Flag that specifies whether the certificate is valid. \nFlag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "issuer": {
                    "computed": true,
                    "description": "X.509 issuer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "public_key": {
                    "computed": true,
                    "description": "Public key component of the X.509 subject public key info.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "serial_number": {
                    "computed": true,
                    "description": "X.509 serial number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sig_alg_name": {
                    "computed": true,
                    "description": "X.509 signatureAlgorithm.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description": "X.509 subject.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject_alternative_names": {
                    "computed": true,
                    "description": "X.509 subject alternative names (SANs) extension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "valid_from": {
                    "computed": true,
                    "description": "X.509 notBefore validity period in milliseconds since epoch.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "X.509 version.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "List of all properties in the object.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Chain of certificates under this alias.",
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
            "read": {
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

func GoogleApigeeKeystoresAliasesKeyCertFileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeKeystoresAliasesKeyCertFile), &result)
	return &result
}
