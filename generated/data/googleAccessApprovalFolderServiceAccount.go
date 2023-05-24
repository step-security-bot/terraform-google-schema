package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessApprovalFolderServiceAccount = `{
  "block": {
    "attributes": {
      "account_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_id": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAccessApprovalFolderServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessApprovalFolderServiceAccount), &result)
	return &result
}
