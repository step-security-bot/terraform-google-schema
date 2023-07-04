package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryReservation = `{
  "block": {
    "attributes": {
      "concurrency": {
        "description": "Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "edition": {
        "computed": true,
        "description": "The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS",
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
      "ignore_idle_slots": {
        "description": "If false, any query using this reservation will use idle slots from other reservations within\nthe same admin project. If true, a query using this reservation will execute with the slot\ncapacity specified above at most.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description": "The geographic location where the transfer config should reside.\nExamples: US, EU, asia-northeast1. The default value is US.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_region_auxiliary": {
        "description": "Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).\nIf set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The name of the reservation. This field must only contain alphanumeric characters or dash.",
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
      "slot_capacity": {
        "description": "Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the\nunit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "autoscale": {
        "block": {
          "attributes": {
            "current_slots": {
              "computed": true,
              "description": "The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].",
              "description_kind": "plain",
              "type": "number"
            },
            "max_slots": {
              "description": "Number of slots to be scaled when needed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "The configuration parameters for the auto scaling feature.",
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

func GoogleBigqueryReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryReservation), &result)
	return &result
}
