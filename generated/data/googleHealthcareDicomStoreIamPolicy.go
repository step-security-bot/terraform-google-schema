package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcareDicomStoreIamPolicy = `{
  "block": {
    "attributes": {
      "dicom_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleHealthcareDicomStoreIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcareDicomStoreIamPolicy), &result)
	return &result
}
