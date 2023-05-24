package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFolderOrganizationPolicy = `{
  "block": {
    "attributes": {
      "boolean_policy": {
        "computed": true,
        "description": "A boolean policy is a constraint that is either enforced or not.",
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
        "description": "The name of the Constraint the Policy is configuring, for example, serviceuser.services.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The etag of the organization policy. etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.",
        "description_kind": "plain",
        "type": "string"
      },
      "folder": {
        "description": "The resource name of the folder to set the policy for. Its format is folders/{folder_id}.",
        "description_kind": "plain",
        "required": true,
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
        "description": "A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. ",
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
      "restore_policy": {
        "computed": true,
        "description": "A restore policy is a constraint to restore the default policy.",
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
        "description": "The timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds, representing when the variable was last updated. Example: \"2016-10-09T12:33:37.578138407Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Version of the Policy. Default version is 0.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleFolderOrganizationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFolderOrganizationPolicy), &result)
	return &result
}
