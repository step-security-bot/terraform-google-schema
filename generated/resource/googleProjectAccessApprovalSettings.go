package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectAccessApprovalSettings = `{
  "block": {
    "attributes": {
      "enrolled_ancestor": {
        "computed": true,
        "description": "If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Project.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the settings. Format is \"projects/{project_id}/accessApprovalSettings\"",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_emails": {
        "computed": true,
        "description": "A list of email addresses to which notifications relating to approval requests should be sent.\nNotifications relating to a resource will be sent to all emails in the settings of ancestor\nresources of that resource. A maximum of 50 email addresses are allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "deprecated": true,
        "description": "Deprecated in favor of 'project_id'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_id": {
        "description": "ID of the project of the access approval settings.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "enrolled_services": {
        "block": {
          "attributes": {
            "cloud_product": {
              "description": "The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):\n  all\n  appengine.googleapis.com\n  bigquery.googleapis.com\n  bigtable.googleapis.com\n  cloudkms.googleapis.com\n  compute.googleapis.com\n  dataflow.googleapis.com\n  iam.googleapis.com\n  pubsub.googleapis.com\n  storage.googleapis.com",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enrollment_level": {
              "description": "The enrollment level of the service. Default value: \"BLOCK_ALL\" Possible values: [\"BLOCK_ALL\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A list of Google Cloud Services for which the given resource has Access Approval enrolled.\nAccess requests for the resource given by name against any of these services contained here will be required\nto have explicit approval. Enrollment can only be done on an all or nothing basis.\n\nA maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func GoogleProjectAccessApprovalSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectAccessApprovalSettings), &result)
	return &result
}
