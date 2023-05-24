package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProject = `{
  "block": {
    "attributes": {
      "auto_create_network": {
        "description": "Create the 'default' network automatically.  Default true. If set to false, the default network will be deleted.  Note that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even if you set auto_create_network to false, since the network will exist momentarily.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "billing_account": {
        "description": "The alphanumeric ID of the billing account this project belongs to. The user or service account performing this operation with Terraform must have Billing Account Administrator privileges (roles/billing.admin) in the organization. See Google Cloud Billing API Access Control for more details.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "folder_id": {
        "description": "The numeric ID of the folder this project should be created under. Only one of org_id or folder_id may be specified. If the folder_id is specified, then the project is created under the specified folder. Changing this forces the project to be migrated to the newly specified folder.",
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
      "labels": {
        "description": "A set of key/value label pairs to assign to the project.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The display name of the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number": {
        "computed": true,
        "description": "The numeric identifier of the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The numeric ID of the organization this project belongs to. Changing this forces a new project to be created.  Only one of org_id or folder_id may be specified. If the org_id is specified then the project is created at the top level. Changing this forces the project to be migrated to the newly specified organization.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_id": {
        "description": "The project ID. Changing this forces a new project to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "skip_delete": {
        "computed": true,
        "description": "If true, the Terraform resource can be deleted without deleting the Project via the Google API.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
            "read": {
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
  "version": 1
}`

func GoogleProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProject), &result)
	return &result
}
