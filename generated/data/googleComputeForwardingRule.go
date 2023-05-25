package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRule = `{
  "block": {
    "attributes": {
      "backend_service": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancing_scheme": {
        "computed": true,
        "description_kind": "plain",
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
      "port_range": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ports": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeForwardingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeForwardingRule), &result)
	return &result
}
