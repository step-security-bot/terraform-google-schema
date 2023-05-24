package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFolderOrganizationPolicy = `{
  "block": {
    "attributes": {
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
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "boolean_policy": {
        "block": {
          "attributes": {
            "enforced": {
              "description": "If true, then the Policy is enforced. If false, then any configuration is acceptable.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "A boolean policy is a constraint that is either enforced or not.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "list_policy": {
        "block": {
          "attributes": {
            "inherit_from_parent": {
              "description": "If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "suggested_value": {
              "computed": true,
              "description": "The Google Cloud Console will try to default to a configuration that matches the value specified in this field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "allow": {
              "block": {
                "attributes": {
                  "all": {
                    "description": "The policy allows or denies all values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "values": {
                    "description": "The policy can define specific values that are allowed or denied.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "One or the other must be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "deny": {
              "block": {
                "attributes": {
                  "all": {
                    "description": "The policy allows or denies all values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "values": {
                    "description": "The policy can define specific values that are allowed or denied.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "One or the other must be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. ",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_policy": {
        "block": {
          "attributes": {
            "default": {
              "description": "May only be set to true. If set, then the default Policy is restored.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "A restore policy is a constraint to restore the default policy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func GoogleFolderOrganizationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFolderOrganizationPolicy), &result)
	return &result
}
