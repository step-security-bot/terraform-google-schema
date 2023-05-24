package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableAppProfile = `{
  "block": {
    "attributes": {
      "app_profile_id": {
        "description": "The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Long form description of the use case for this app profile.",
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
      "ignore_warnings": {
        "description": "If true, ignore safety checks when deleting/updating the app profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance": {
        "description": "The name of the instance to create the app profile within.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_cluster_routing_cluster_ids": {
        "description": "The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "multi_cluster_routing_use_any": {
        "description": "If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available\nin the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes\nconsistency to improve availability.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The unique name of the requested app profile. Values are of the form 'projects/\u003cproject\u003e/instances/\u003cinstance\u003e/appProfiles/\u003cappProfileId\u003e'.",
        "description_kind": "plain",
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
      "single_cluster_routing": {
        "block": {
          "attributes": {
            "allow_transactional_writes": {
              "description": "If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.\nIt is unsafe to send these requests to the same table/row/column in multiple clusters.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cluster_id": {
              "description": "The cluster to which read/write requests should be routed.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Use a single-cluster routing policy.",
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

func GoogleBigtableAppProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableAppProfile), &result)
	return &result
}
