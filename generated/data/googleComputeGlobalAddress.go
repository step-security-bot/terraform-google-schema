package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeGlobalAddress = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "address_type": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_length": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purpose": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "users": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeGlobalAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeGlobalAddress), &result)
	return &result
}
