package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrganizationIamAuditConfig = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description": "The etag of iam policy",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "org_id": {
        "description": "The numeric ID of the organization in which you want to manage the audit logging config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service": {
        "description": "Service which will be enabled for audit logging. The special value allServices covers all services.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "audit_log_config": {
        "block": {
          "attributes": {
            "exempted_members": {
              "description": "Identities that do not cause logging for this type of permission. Each entry can have one of the following values:user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com. serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "log_type": {
              "description": "Permission type for which logging is to be configured. Must be one of DATA_READ, DATA_WRITE, or ADMIN_READ.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The configuration for logging of each type of permission. This can be specified multiple times.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOrganizationIamAuditConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrganizationIamAuditConfig), &result)
	return &result
}
