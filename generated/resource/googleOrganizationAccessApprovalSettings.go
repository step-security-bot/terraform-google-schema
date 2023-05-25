package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrganizationAccessApprovalSettings = `{
  "block": {
    "attributes": {
      "enrolled_ancestor": {
        "computed": true,
        "description": "This field will always be unset for the organization since organizations do not have ancestors.",
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
        "description": "The resource name of the settings. Format is \"organizations/{organization_id}/accessApprovalSettings\"",
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
      "organization_id": {
        "description": "ID of the organization of the access approval settings.",
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
          "description": "A list of Google Cloud Services for which the given resource has Access Approval enrolled.\nAccess requests for the resource given by name against any of these services contained here will be required\nto have explicit approval. Enrollment can be done for individual services.\n\nA maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.",
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

func GoogleOrganizationAccessApprovalSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrganizationAccessApprovalSettings), &result)
	return &result
}
