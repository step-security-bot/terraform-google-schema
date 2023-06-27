package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkloadIdentityPoolProvider = `{
  "block": {
    "attributes": {
      "attribute_condition": {
        "description": "[A Common Expression Language](https://opensource.google/projects/cel) expression, in\nplain text, to restrict what otherwise valid authentication credentials issued by the\nprovider should not be accepted.\n\nThe expression must output a boolean representing whether to allow the federation.\n\nThe following keywords may be referenced in the expressions:\n  * 'assertion': JSON representing the authentication credential issued by the provider.\n  * 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.\n  * 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.\n\nThe maximum length of the attribute condition expression is 4096 characters. If\nunspecified, all valid authentication credential are accepted.\n\nThe following example shows how to only allow credentials with a mapped 'google.groups'\nvalue of 'admins':\n'''\n\"'admins' in google.groups\"\n'''",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attribute_mapping": {
        "description": "Maps attributes from authentication credentials issued by an external identity provider\nto Google Cloud attributes, such as 'subject' and 'segment'.\n\nEach key must be a string specifying the Google Cloud IAM attribute to map to.\n\nThe following keys are supported:\n  * 'google.subject': The principal IAM is authenticating. You can reference this value\n    in IAM bindings. This is also the subject that appears in Cloud Logging logs.\n    Cannot exceed 127 characters.\n  * 'google.groups': Groups the external identity belongs to. You can grant groups\n    access to resources using an IAM 'principalSet' binding; access applies to all\n    members of the group.\n\nYou can also provide custom attributes by specifying 'attribute.{custom_attribute}',\nwhere '{custom_attribute}' is the name of the custom attribute to be mapped. You can\ndefine a maximum of 50 custom attributes. The maximum length of a mapped attribute key\nis 100 characters, and the key may only contain the characters [a-z0-9_].\n\nYou can reference these attributes in IAM policies to define fine-grained access for a\nworkload to Google Cloud resources. For example:\n  * 'google.subject':\n    'principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}'\n  * 'google.groups':\n    'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}'\n  * 'attribute.{custom_attribute}':\n    'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}'\n\nEach value must be a [Common Expression Language](https://opensource.google/projects/cel)\nfunction that maps an identity provider credential to the normalized attribute specified\nby the corresponding map key.\n\nYou can use the 'assertion' keyword in the expression to access a JSON representation of\nthe authentication credential issued by the provider.\n\nThe maximum length of an attribute mapping expression is 2048 characters. When evaluated,\nthe total size of all mapped attributes must not exceed 8KB.\n\nFor AWS providers, the following rules apply:\n  - If no attribute mapping is defined, the following default mapping applies:\n    '''\n    {\n      \"google.subject\":\"assertion.arn\",\n      \"attribute.aws_role\":\n        \"assertion.arn.contains('assumed-role')\"\n        \" ? assertion.arn.extract('{account_arn}assumed-role/')\"\n        \"   + 'assumed-role/'\"\n        \"   + assertion.arn.extract('assumed-role/{role_name}/')\"\n        \" : assertion.arn\",\n    }\n    '''\n  - If any custom attribute mappings are defined, they must include a mapping to the\n    'google.subject' attribute.\n\nFor OIDC providers, the following rules apply:\n  - Custom attribute mappings must be defined, and must include a mapping to the\n    'google.subject' attribute. For example, the following maps the 'sub' claim of the\n    incoming credential to the 'subject' attribute on a Google token.\n    '''\n    {\"google.subject\": \"assertion.sub\"}\n    '''",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "description": "A description for the provider. Cannot exceed 256 characters.",
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
        "description": "A display name for the provider. Cannot exceed 32 characters.",
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
      "name": {
        "computed": true,
        "description": "The resource name of the provider as\n'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the provider.\n* STATE_UNSPECIFIED: State unspecified.\n* ACTIVE: The provider is active, and may be used to validate authentication credentials.\n* DELETED: The provider is soft-deleted. Soft-deleted providers are permanently deleted\n  after approximately 30 days. You can restore a soft-deleted provider using\n  UndeleteWorkloadIdentityPoolProvider. You cannot reuse the ID of a soft-deleted provider\n  until it is permanently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_identity_pool_id": {
        "description": "The ID used for the pool, which is the final component of the pool resource name. This\nvalue should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix\n'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workload_identity_pool_provider_id": {
        "description": "The ID for the provider, which becomes the final component of the resource name. This\nvalue must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix\n'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "aws": {
        "block": {
          "attributes": {
            "account_id": {
              "description": "The AWS account ID.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An Amazon Web Services identity provider. Not compatible with the property oidc.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oidc": {
        "block": {
          "attributes": {
            "allowed_audiences": {
              "description": "Acceptable values for the 'aud' field (audience) in the OIDC token. Token exchange\nrequests are rejected if the token audience does not match one of the configured\nvalues. Each audience may be at most 256 characters. A maximum of 10 audiences may\nbe configured.\n\nIf this list is empty, the OIDC token audience must be equal to the full canonical\nresource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.\nFor example:\n'''\n//iam.googleapis.com/projects/\u003cproject-number\u003e/locations/\u003clocation\u003e/workloadIdentityPools/\u003cpool-id\u003e/providers/\u003cprovider-id\u003e\nhttps://iam.googleapis.com/projects/\u003cproject-number\u003e/locations/\u003clocation\u003e/workloadIdentityPools/\u003cpool-id\u003e/providers/\u003cprovider-id\u003e\n'''",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "issuer_uri": {
              "description": "The OIDC issuer URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "jwks_json": {
              "description": "OIDC JWKs in JSON String format. For details on definition of a\nJWK, see https:tools.ietf.org/html/rfc7517. If not set, then we\nuse the 'jwks_uri' from the discovery document fetched from the\n.well-known path for the 'issuer_uri'. Currently, RSA and EC asymmetric\nkeys are supported. The JWK must use following format and include only\nthe following fields:\n'''\n{\n  \"keys\": [\n    {\n          \"kty\": \"RSA/EC\",\n          \"alg\": \"\u003calgorithm\u003e\",\n          \"use\": \"sig\",\n          \"kid\": \"\u003ckey-id\u003e\",\n          \"n\": \"\",\n          \"e\": \"\",\n          \"x\": \"\",\n          \"y\": \"\",\n          \"crv\": \"\"\n    }\n  ]\n}\n'''",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "An OpenId Connect 1.0 identity provider. Not compatible with the property aws.",
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

func GoogleIamWorkloadIdentityPoolProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkloadIdentityPoolProvider), &result)
	return &result
}
