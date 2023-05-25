package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeOrganization = `{
  "block": {
    "attributes": {
      "analytics_region": {
        "description": "Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorized_network": {
        "description": "Compute Engine network used for Service Networking to be peered with Apigee runtime instances.\nSee [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).\nValid only when 'RuntimeType' is set to CLOUD. The value can be updated only when there are no runtime instances. For example: \"default\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_certificate": {
        "computed": true,
        "description": "Output only. Base64-encoded public certificate for the root CA of the Apigee organization.\nValid only when 'RuntimeType' is CLOUD. A base64-encoded string.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the Apigee organization.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the Apigee organization.",
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
        "description": "Output only. Name of the Apigee organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "description": "The project ID associated with the Apigee organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_database_encryption_key_name": {
        "description": "Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.\nUpdate is not allowed after the organization is created.\nIf not specified, a Google-Managed encryption key will be used.\nValid only when 'RuntimeType' is CLOUD. For example: 'projects/foo/locations/us/keyRings/bar/cryptoKeys/baz'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_type": {
        "description": "Runtime type of the Apigee organization based on the Apigee subscription purchased. Default value: \"CLOUD\" Possible values: [\"CLOUD\", \"HYBRID\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_type": {
        "computed": true,
        "description": "Output only. Subscription type of the Apigee organization.\nValid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).",
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

func GoogleApigeeOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeOrganization), &result)
	return &result
}
