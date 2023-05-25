package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProjectOrganizationPolicy = `{
  "block": {
    "attributes": {
      "boolean_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enforced": "bool"
            }
          ]
        ]
      },
      "constraint": {
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
      "list_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow": [
                "list",
                [
                  "object",
                  {
                    "all": "bool",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "deny": [
                "list",
                [
                  "object",
                  {
                    "all": "bool",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "inherit_from_parent": "bool",
              "suggested_value": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "default": "bool"
            }
          ]
        ]
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleProjectOrganizationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProjectOrganizationPolicy), &result)
	return &result
}
