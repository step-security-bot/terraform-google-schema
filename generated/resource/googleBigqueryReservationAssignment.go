package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryReservationAssignment = `{
  "block": {
    "attributes": {
      "assignee": {
        "description": "The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.",
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
      "job_type": {
        "description": "Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The location for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the assignment.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reservation": {
        "description": "The reservation for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE",
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

func GoogleBigqueryReservationAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryReservationAssignment), &result)
	return &result
}
