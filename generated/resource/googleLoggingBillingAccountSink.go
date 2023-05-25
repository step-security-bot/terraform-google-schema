package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingBillingAccountSink = `{
  "block": {
    "attributes": {
      "billing_account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "writer_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingBillingAccountSinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingBillingAccountSink), &result)
	return &result
}
