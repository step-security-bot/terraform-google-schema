package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryBiReservation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "LOCATION_DESCRIPTION",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the singleton BI reservation. Reservation names have the form 'projects/{projectId}/locations/{locationId}/biReservation'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size": {
        "description": "Size of a reservation, in bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a reservation.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "preferred_tables": {
        "block": {
          "attributes": {
            "dataset_id": {
              "description": "The ID of the dataset in the above project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_id": {
              "description": "The assigned project ID of the project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "table_id": {
              "description": "The ID of the table in the above dataset.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Preferred tables to use BI capacity for.",
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

func GoogleBigqueryBiReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryBiReservation), &result)
	return &result
}
