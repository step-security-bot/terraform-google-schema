package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeTargetPool = `{
  "block": {
    "attributes": {
      "backup_pool": {
        "description": "URL to the backup target pool. Must also set failover_ratio.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Textual description field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover_ratio": {
        "description": "Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_checks": {
        "description": "List of zero or one health check name or self_link. Only legacy google_compute_http_health_check is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description": "List of instances in the pool. They can be given as URLs, or in the form of \"zone/name\". Note that the instances need not exist at the time of target pool creation, so there is no need to use the Terraform interpolators to create a dependency on the instances from the target pool.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "description": "A unique name for the resource, required by GCE. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Where the target pool resides. Defaults to project region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "description": "How to distribute load. Options are \"NONE\" (no affinity). \"CLIENT_IP\" (hash of the source/dest addresses / ports), and \"CLIENT_IP_PROTO\" also includes the protocol (default \"NONE\").",
        "description_kind": "plain",
        "optional": true,
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

func GoogleComputeTargetPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeTargetPool), &result)
	return &result
}
