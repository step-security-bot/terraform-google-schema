package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSharedflowDeployment = `{
  "block": {
    "attributes": {
      "environment": {
        "description": "The resource ID of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision": {
        "description": "Revision of the Sharedflow to be deployed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_account": {
        "description": "The service account represents the identity of the deployed proxy, and determines what permissions it has. The format must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sharedflow_id": {
        "description": "Id of the Sharedflow to be deployed.",
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

func GoogleApigeeSharedflowDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSharedflowDeployment), &result)
	return &result
}
