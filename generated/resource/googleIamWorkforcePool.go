package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkforcePool = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A user-specified description of the pool. Cannot exceed 256 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,\nor use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.",
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
        "description": "The location for the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the pool.\nFormat: 'locations/{location}/workforcePools/{workforcePoolId}'",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Immutable. The resource name of the parent. Format: 'organizations/{org-id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_duration": {
        "description": "Duration that the Google Cloud access tokens, console sign-in sessions,\nand 'gcloud' sign-in sessions from this pool are valid.\nMust be greater than 15 minutes (900s) and less than 12 hours (43200s).\nIf 'sessionDuration' is not configured, minted credentials have a default duration of one hour (3600s).\nA duration in seconds with up to nine fractional digits, ending with ''s''. Example: \"'3.5s'\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of the pool.\n * STATE_UNSPECIFIED: State unspecified.\n * ACTIVE: The pool is active, and may be used in Google Cloud policies.\n * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted\n   after approximately 30 days. You can restore a soft-deleted pool using\n   [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).\n   You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.\n   While a pool is deleted, you cannot use it to exchange tokens, or use\n   existing tokens to access resources. If the pool is undeleted, existing\n   tokens grant access again.",
        "description_kind": "plain",
        "type": "string"
      },
      "workforce_pool_id": {
        "description": "The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,\ndigits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.\nThe prefix 'gcp-' is reserved for use by Google, and may not be specified.",
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

func GoogleIamWorkforcePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkforcePool), &result)
	return &result
}
