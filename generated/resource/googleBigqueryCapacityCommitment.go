package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryCapacityCommitment = `{
  "block": {
    "attributes": {
      "capacity_commitment_id": {
        "description": "The optional capacity commitment ID. Capacity commitment name will be generated automatically if this field is\nempty. This field must only contain lower case alphanumeric characters or dashes. The first and last character\ncannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split\nor merged.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "commitment_end_time": {
        "computed": true,
        "description": "The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.",
        "description_kind": "plain",
        "type": "string"
      },
      "commitment_start_time": {
        "computed": true,
        "description": "The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.",
        "description_kind": "plain",
        "type": "string"
      },
      "edition": {
        "description": "The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enforce_single_admin_project_per_org": {
        "description": "If true, fail the request if another project in the organization has a capacity commitment.",
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
        "description": "The geographic location where the transfer config should reside.\nExamples: US, EU, asia-northeast1. The default value is US.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the capacity commitment, e.g., projects/myproject/locations/US/capacityCommitments/123",
        "description_kind": "plain",
        "type": "string"
      },
      "plan": {
        "description": "Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan",
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
      "renewal_plan": {
        "description": "The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable some commitment plans.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slot_count": {
        "description": "Number of slots in this commitment.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "State of the commitment",
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

func GoogleBigqueryCapacityCommitmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryCapacityCommitment), &result)
	return &result
}
