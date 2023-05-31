package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEndpointsServiceConsumersIamPolicy = `{
  "block": {
    "attributes": {
      "consumer_project": {
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
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleEndpointsServiceConsumersIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEndpointsServiceConsumersIamPolicy), &result)
	return &result
}
