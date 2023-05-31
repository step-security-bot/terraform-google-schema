package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubTopicIamPolicy = `{
  "block": {
    "attributes": {
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GooglePubsubTopicIamPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubTopicIamPolicy), &result)
	return &result
}
