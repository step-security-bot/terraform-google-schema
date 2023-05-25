package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceNetworkingConnection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reserved_peering_ranges": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceNetworkingConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceNetworkingConnection), &result)
	return &result
}
