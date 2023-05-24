package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIdentityPlatformTenantInboundSamlConfig = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Human friendly display name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description": "If this config allows users to sign in with the provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,\nhyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an\nalphanumeric character, and have at least 2 characters.",
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
      "tenant": {
        "description": "The name of the tenant where this inbound SAML config resource exists",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "idp_config": {
        "block": {
          "attributes": {
            "idp_entity_id": {
              "description": "Unique identifier for all SAML entities",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sign_request": {
              "description": "Indicates if outbounding SAMLRequest should be signed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sso_url": {
              "description": "URL to send Authentication request to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "idp_certificates": {
              "block": {
                "attributes": {
                  "x509_certificate": {
                    "description": "The x509 certificate",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "SAML IdP configuration when the project acts as the relying party",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "sp_config": {
        "block": {
          "attributes": {
            "callback_uri": {
              "description": "Callback URI where responses from IDP are handled. Must start with 'https://'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sp_certificates": {
              "computed": true,
              "description": "The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "x509_certificate": "string"
                  }
                ]
              ]
            },
            "sp_entity_id": {
              "description": "Unique identifier for all SAML entities.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "SAML SP (Service Provider) configuration when the project acts as the relying party to receive\nand accept an authentication assertion issued by a SAML identity provider.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleIdentityPlatformTenantInboundSamlConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIdentityPlatformTenantInboundSamlConfig), &result)
	return &result
}
