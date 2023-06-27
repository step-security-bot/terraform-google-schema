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
      "apigee_project_id": {
        "computed": true,
        "description": "Output only. Project ID of the Apigee Tenant Project.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorized_network": {
        "description": "Compute Engine network used for Service Networking to be peered with Apigee runtime instances.\nSee [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).\nValid only when 'RuntimeType' is set to CLOUD. The value can be updated only when there are no runtime instances. For example: \"default\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "billing_type": {
        "computed": true,
        "description": "Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).",
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
      "retention": {
        "description": "Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType\nis not EVALUATION). It controls how long Organization data will be retained after the initial delete\noperation completes. During this period, the Organization may be restored to its last known state.\nAfter this period, the Organization will no longer be able to be restored. Default value: \"DELETION_RETENTION_UNSPECIFIED\" Possible values: [\"DELETION_RETENTION_UNSPECIFIED\", \"MINIMUM\"]",
        "description_kind": "plain",
        "optional": true,
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
      "properties": {
        "block": {
          "block_types": {
            "property": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value of the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of all properties in the object.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Properties defined in the Apigee organization profile.",
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

func GoogleApigeeOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeOrganization), &result)
	return &result
}
