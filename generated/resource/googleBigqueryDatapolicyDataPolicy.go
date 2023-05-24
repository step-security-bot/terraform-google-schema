package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDatapolicyDataPolicy = `{
  "block": {
    "attributes": {
      "data_policy_id": {
        "description": "User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {dataPolicyId} in part of the resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_policy_type": {
        "description": "The enrollment level of the service. Possible values: [\"COLUMN_LEVEL_SECURITY_POLICY\", \"DATA_MASKING_POLICY\"]",
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
      "location": {
        "description": "The name of the location of the data policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Resource name of this data policy, in the format of projects/{project_number}/locations/{locationId}/dataPolicies/{dataPolicyId}.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_tag": {
        "description": "Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_masking_policy": {
        "block": {
          "attributes": {
            "predefined_expression": {
              "description": "The available masking rules. Learn more here: https://cloud.google.com/bigquery/docs/column-data-masking-intro#masking_options. Possible values: [\"SHA256\", \"ALWAYS_NULL\", \"DEFAULT_MASKING_VALUE\", \"LAST_FOUR_CHARACTERS\", \"FIRST_FOUR_CHARACTERS\", \"EMAIL_MASK\", \"DATE_YEAR_MASK\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The data masking policy that specifies the data masking rule to use.",
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

func GoogleBigqueryDatapolicyDataPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDatapolicyDataPolicy), &result)
	return &result
}
