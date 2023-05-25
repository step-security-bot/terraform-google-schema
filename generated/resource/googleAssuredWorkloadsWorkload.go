package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAssuredWorkloadsWorkload = `{
  "block": {
    "attributes": {
      "billing_account": {
        "description": "Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form ` + "`" + `billingAccounts/{billing_account_id}` + "`" + `. For example, 'billingAccounts/012345-567890-ABCDEF` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compliance_regime": {
        "description": "Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Immutable. The Workload creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload",
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
      "labels": {
        "description": "Optional. Labels applied to the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "description": "The organization for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioned_resources_parent": {
        "description": "Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description": "Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_id": "number",
              "resource_type": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "kms_settings": {
        "block": {
          "attributes": {
            "next_rotation_time": {
              "description": "Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rotation_period": {
              "description": "Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resource_settings": {
        "block": {
          "attributes": {
            "resource_id": {
              "description": "Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_type": {
              "description": "Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.",
          "description_kind": "plain"
        },
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

func GoogleAssuredWorkloadsWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAssuredWorkloadsWorkload), &result)
	return &result
}
