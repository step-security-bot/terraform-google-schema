package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkloadIdentityPool = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of the pool. Cannot exceed 256 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use\nexisting tokens to access resources. If the pool is re-enabled, existing tokens grant\naccess again.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "A display name for the pool. Cannot exceed 32 characters.",
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
        "description": "The resource name of the pool as\n'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the pool.\n* STATE_UNSPECIFIED: State unspecified.\n* ACTIVE: The pool is active, and may be used in Google Cloud policies.\n* DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after\n  approximately 30 days. You can restore a soft-deleted pool using\n  UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is\n  permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or\n  use existing tokens to access resources. If the pool is undeleted, existing tokens grant\n  access again.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_identity_pool_id": {
        "description": "The ID to use for the pool, which becomes the final component of the resource name. This\nvalue should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix\n'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
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

func GoogleIamWorkloadIdentityPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkloadIdentityPool), &result)
	return &result
}
