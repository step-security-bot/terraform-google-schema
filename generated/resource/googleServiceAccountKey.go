package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_algorithm": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pgp_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "private_key_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_key_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "valid_after": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_before": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountKey), &result)
	return &result
}
