package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeNatAddress = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description": "The Apigee instance associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/instances/{{instance_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The allocated NAT IP address.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Resource ID of the NAT address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the NAT IP address.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleApigeeNatAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeNatAddress), &result)
	return &result
}
