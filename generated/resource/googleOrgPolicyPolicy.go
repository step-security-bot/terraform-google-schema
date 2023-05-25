package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrgPolicyPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * ` + "`" + `projects/{project_number}/policies/{constraint_name}` + "`" + ` * ` + "`" + `folders/{folder_id}/policies/{constraint_name}` + "`" + ` * ` + "`" + `organizations/{organization_id}/policies/{constraint_name}` + "`" + ` For example, \"projects/123/policies/compute.disableSerialPortAccess\". Note: ` + "`" + `projects/{project_id}/policies/{constraint_name}` + "`" + ` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The parent of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "attributes": {
            "etag": {
              "computed": true,
              "description": "An opaque tag indicating the current version of the ` + "`" + `Policy` + "`" + `, used for concurrency control. This field is ignored if used in a ` + "`" + `CreatePolicy` + "`" + ` request. When the ` + "`" + `Policy` + "`" + ` is returned from either a ` + "`" + `GetPolicy` + "`" + ` or a ` + "`" + `ListPolicies` + "`" + ` request, this ` + "`" + `etag` + "`" + ` indicates the version of the current ` + "`" + `Policy` + "`" + ` to use when executing a read-modify-write loop. When the ` + "`" + `Policy` + "`" + ` is returned from a ` + "`" + `GetEffectivePolicy` + "`" + ` request, the ` + "`" + `etag` + "`" + ` will be unset.",
              "description_kind": "plain",
              "type": "string"
            },
            "inherit_from_parent": {
              "description": "Determines the inheritance behavior for this ` + "`" + `Policy` + "`" + `. If ` + "`" + `inherit_from_parent` + "`" + ` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "reset": {
              "description": "Ignores policies set above this resource and restores the ` + "`" + `constraint_default` + "`" + ` enforcement behavior of the specific ` + "`" + `Constraint` + "`" + ` at this resource. This field can be set in policies for either list or boolean constraints. If set, ` + "`" + `rules` + "`" + ` must be empty and ` + "`" + `inherit_from_parent` + "`" + ` must be set to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "update_time": {
              "computed": true,
              "description": "Output only. The time stamp this was previously updated. This represents the last time a call to ` + "`" + `CreatePolicy` + "`" + ` or ` + "`" + `UpdatePolicy` + "`" + ` was made for that ` + "`" + `Policy` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "allow_all": {
                    "description": "Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "deny_all": {
                    "description": "Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enforce": {
                    "description": "If ` + "`" + `true` + "`" + `, then the ` + "`" + `Policy` + "`" + ` is enforced. If ` + "`" + `false` + "`" + `, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "location": {
                          "description": "Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "title": {
                          "description": "Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A condition which determines whether this rule is used in the evaluation of the policy. When set, the ` + "`" + `expression` + "`" + ` field in the ` + "`" + `Expr' must include from 1 to 10 subexpressions, joined by the \"||\" or \"\u0026\u0026\" operators. Each subexpression must be of the form \"resource.matchTag('/tag_key_short_name, 'tag_value_short_name')\". or \"resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')\". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: \"resource.matchTag('123456789/environment, 'prod')\". or \"resource.matchTagId('tagKeys/123', 'tagValues/456')\".",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "values": {
                    "block": {
                      "attributes": {
                        "allowed_values": {
                          "description": "List of values allowed at this resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "denied_values": {
                          "description": "List of values denied at this resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set ` + "`" + `enforced` + "`" + ` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Basic information about the Organization Policy.",
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

func GoogleOrgPolicyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrgPolicyPolicy), &result)
	return &result
}
