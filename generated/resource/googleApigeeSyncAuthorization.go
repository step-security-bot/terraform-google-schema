package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSyncAuthorization = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description": "Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting each other.\nUsed internally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identities": {
        "description": "Array of service accounts to grant access to control plane resources, each specified using the following format: 'serviceAccount:service-account-name'.\n\nThe 'service-account-name' is formatted like an email address. For example: my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com\n\nYou might specify multiple service accounts, for example, if you have multiple environments and wish to assign a unique service account to each one.\n\nThe service accounts must have **Apigee Synchronizer Manager** role. See also [Create service accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the Apigee organization.",
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

func GoogleApigeeSyncAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSyncAuthorization), &result)
	return &result
}
