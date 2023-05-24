package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectDefaultServiceAccounts = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE.\n\t\t\t\tNote that DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.",
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
      "project": {
        "description": "The project ID where service accounts are created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_policy": {
        "description": "The action to be performed in the default service accounts on the resource destroy.\n\t\t\t\tValid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_accounts": {
        "computed": true,
        "description": "The Service Accounts changed by this resource. It is used for revert the action on the destroy.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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

func GoogleProjectDefaultServiceAccountsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectDefaultServiceAccounts), &result)
	return &result
}
