package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkforcePoolProvider = `{
  "block": {
    "attributes": {
      "attribute_condition": {
        "description": "A [Common Expression Language](https://opensource.google/projects/cel) expression, in\nplain text, to restrict what otherwise valid authentication credentials issued by the\nprovider should not be accepted.\n\nThe expression must output a boolean representing whether to allow the federation.\n\nThe following keywords may be referenced in the expressions:\n  * 'assertion': JSON representing the authentication credential issued by the provider.\n  * 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.\n    'google.profile_photo' and 'google.display_name' are not supported.\n  * 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.\n\nThe maximum length of the attribute condition expression is 4096 characters.\nIf unspecified, all valid authentication credentials will be accepted.\n\nThe following example shows how to only allow credentials with a mapped 'google.groups' value of 'admins':\n'''\n\"'admins' in google.groups\"\n'''",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attribute_mapping": {
        "description": "Maps attributes from the authentication credentials issued by an external identity provider\nto Google Cloud attributes, such as 'subject' and 'segment'.\n\nEach key must be a string specifying the Google Cloud IAM attribute to map to.\n\nThe following keys are supported:\n  * 'google.subject': The principal IAM is authenticating. You can reference this value in IAM bindings.\n    This is also the subject that appears in Cloud Logging logs. This is a required field and\n    the mapped subject cannot exceed 127 bytes.\n  * 'google.groups': Groups the authenticating user belongs to. You can grant groups access to\n    resources using an IAM 'principalSet' binding; access applies to all members of the group.\n  * 'google.display_name': The name of the authenticated user. This is an optional field and\n    the mapped display name cannot exceed 100 bytes. If not set, 'google.subject' will be displayed instead.\n    This attribute cannot be referenced in IAM bindings.\n  * 'google.profile_photo': The URL that specifies the authenticated user's thumbnail photo.\n    This is an optional field. When set, the image will be visible as the user's profile picture.\n    If not set, a generic user icon will be displayed instead.\n    This attribute cannot be referenced in IAM bindings.\n\nYou can also provide custom attributes by specifying 'attribute.{custom_attribute}', where {custom_attribute}\nis the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.\nThe maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].\n\nYou can reference these attributes in IAM policies to define fine-grained access for a workforce pool\nto Google Cloud resources. For example:\n  * 'google.subject':\n    'principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}'\n  * 'google.groups':\n    'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}'\n  * 'attribute.{custom_attribute}':\n    'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}'\n\nEach value must be a [Common Expression Language](https://opensource.google/projects/cel)\nfunction that maps an identity provider credential to the normalized attribute specified\nby the corresponding map key.\n\nYou can use the 'assertion' keyword in the expression to access a JSON representation of\nthe authentication credential issued by the provider.\n\nThe maximum length of an attribute mapping expression is 2048 characters. When evaluated,\nthe total size of all mapped attributes must not exceed 8KB.\n\nFor OIDC providers, you must supply a custom mapping that includes the 'google.subject' attribute.\nFor example, the following maps the sub claim of the incoming credential to the 'subject' attribute\non a Google token:\n'''\n{\"google.subject\": \"assertion.sub\"}\n'''\n\nAn object containing a list of '\"key\": value' pairs.\nExample: '{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "description": "A user-specified description of the provider. Cannot exceed 256 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.\nHowever, existing tokens still grant access.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "A user-specified display name for the provider. Cannot exceed 32 characters.",
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
      "location": {
        "description": "The location for the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the provider.\nFormat: 'locations/{location}/workforcePools/{workforcePoolId}/providers/{providerId}'",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_id": {
        "description": "The ID for the provider, which becomes the final component of the resource name.\nThis value must be 4-32 characters, and may contain the characters [a-z0-9-].\nThe prefix 'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the provider.\n* STATE_UNSPECIFIED: State unspecified.\n* ACTIVE: The provider is active and may be used to validate authentication credentials.\n* DELETED: The provider is soft-deleted. Soft-deleted providers are permanently\n  deleted after approximately 30 days. You can restore a soft-deleted provider using\n  [providers.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools.providers/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePoolProvider).",
        "description_kind": "plain",
        "type": "string"
      },
      "workforce_pool_id": {
        "description": "The ID to use for the pool, which becomes the final component of the resource name.\nThe IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.\nIt must start with a letter, and cannot have a trailing hyphen.\nThe prefix 'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "oidc": {
        "block": {
          "attributes": {
            "client_id": {
              "description": "The client ID. Must match the audience claim of the JWT issued by the identity provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "issuer_uri": {
              "description": "The OIDC issuer URI. Must be a valid URI using the 'https' scheme.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "client_secret": {
              "block": {
                "block_types": {
                  "value": {
                    "block": {
                      "attributes": {
                        "plain_text": {
                          "description": "The plain text of the client secret value.",
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "thumbprint": {
                          "computed": true,
                          "description": "A thumbprint to represent the current client secret value.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "description": "The value of the client secret.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The optional client secret. Required to enable Authorization Code flow for web sign-in.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "web_sso_config": {
              "block": {
                "attributes": {
                  "additional_scopes": {
                    "description": "Additional scopes to request for in the OIDC authentication request on top of scopes requested by default. By default, the 'openid', 'profile' and 'email' scopes that are supported by the identity provider are requested.\nEach additional scope may be at most 256 characters. A maximum of 10 additional scopes may be configured.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "assertion_claims_behavior": {
                    "description": "The behavior for how OIDC Claims are included in the 'assertion' object used for attribute mapping and attribute condition.\n* MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS: Merge the UserInfo Endpoint Claims with ID Token Claims, preferring UserInfo Claim Values for the same Claim Name. This option is available only for the Authorization Code Flow.\n* ONLY_ID_TOKEN_CLAIMS: Only include ID Token Claims. Possible values: [\"MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS\", \"ONLY_ID_TOKEN_CLAIMS\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "response_type": {
                    "description": "The Response Type to request for in the OIDC Authorization Request for web sign-in.\n\nThe 'CODE' Response Type is recommended to avoid the Implicit Flow, for security reasons.\n* CODE: The 'response_type=code' selection uses the Authorization Code Flow for web sign-in. Requires a configured client secret.\n* ID_TOKEN: The 'response_type=id_token' selection uses the Implicit Flow for web sign-in. Possible values: [\"CODE\", \"ID_TOKEN\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for web single sign-on for the OIDC provider. Here, web sign-in refers to console sign-in and gcloud sign-in through the browser.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents an OpenId Connect 1.0 identity provider.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "saml": {
        "block": {
          "attributes": {
            "idp_metadata_xml": {
              "description": "SAML Identity provider configuration metadata xml doc.\nThe xml document should comply with [SAML 2.0 specification](https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf).\nThe max size of the acceptable xml document will be bounded to 128k characters.\n\nThe metadata xml document should satisfy the following constraints:\n1) Must contain an Identity Provider Entity ID.\n2) Must contain at least one non-expired signing key certificate.\n3) For each signing key:\n  a) Valid from should be no more than 7 days from now.\n  b) Valid to should be no more than 10 years in the future.\n4) Up to 3 IdP signing keys are allowed in the metadata xml.\n\nWhen updating the provider's metadata xml, at least one non-expired signing key\nmust overlap with the existing metadata. This requirement is skipped if there are\nno non-expired signing keys present in the existing metadata.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Represents a SAML identity provider.",
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

func GoogleIamWorkforcePoolProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkforcePoolProvider), &result)
	return &result
}
