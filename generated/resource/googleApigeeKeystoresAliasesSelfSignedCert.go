package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeKeystoresAliasesSelfSignedCert = `{
  "block": {
    "attributes": {
      "alias": {
        "description": "Alias for the key/certificate pair. Values must match the regular expression [\\w\\s-.]{1,255}.\nThis must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either\nthis parameter or the JSON body.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cert_validity_in_days": {
        "description": "Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
        "description": "The Apigee environment name",
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
      "key_size": {
        "description": "Key size. Default and maximum value is 2048 bits.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keystore": {
        "description": "The Apigee keystore name associated in an Apigee environment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization name associated with the Apigee environment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sig_alg": {
        "description": "Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA",
        "description_kind": "plain",
        "required": true,
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
      "subject": {
        "block": {
          "attributes": {
            "common_name": {
              "description": "Common name of the organization. Maximum length is 64 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "description": "Two-letter country code. Example, IN for India, US for United States of America.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "description": "Email address. Max 255 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "locality": {
              "description": "City or town name. Maximum length is 128 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "org": {
              "description": "Organization name. Maximum length is 64 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "org_unit": {
              "description": "Organization team name. Maximum length is 64 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description": "State or district name. Maximum length is 128 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Subject details.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "subject_alternative_dns_names": {
        "block": {
          "attributes": {
            "subject_alternative_name": {
              "description": "Subject Alternative Name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "List of alternative host names. Maximum length is 255 characters for each value.",
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

func GoogleApigeeKeystoresAliasesSelfSignedCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeKeystoresAliasesSelfSignedCert), &result)
	return &result
}
