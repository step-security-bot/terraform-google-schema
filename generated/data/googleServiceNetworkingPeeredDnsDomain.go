package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceNetworkingPeeredDnsDomain = `{
  "block": {
    "attributes": {
      "dns_suffix": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func GoogleServiceNetworkingPeeredDnsDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceNetworkingPeeredDnsDomain), &result)
	return &result
}
