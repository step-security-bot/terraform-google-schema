package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamAccessBoundaryPolicy = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The display name of the rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The hash of the resource. Used internally during updates.",
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
        "description": "The name of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The attachment point is identified by its URL-encoded full resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rules": {
        "block": {
          "attributes": {
            "description": {
              "description": "The description of the rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "access_boundary_rule": {
              "block": {
                "attributes": {
                  "available_permissions": {
                    "description": "A list of permissions that may be allowed for use on the specified resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "available_resource": {
                    "description": "The full resource name of a Google Cloud resource entity.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "availability_condition": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "Description of the expression. This is a longer text which describes the expression,\ne.g. when hovered over it in a UI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "location": {
                          "description": "String indicating the location of the expression for error reporting,\ne.g. a file name and a position in the file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "title": {
                          "description": "Title for the expression, i.e. a short string describing its purpose.\nThis can be used e.g. in UIs which allow to enter the expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The availability condition further constrains the access allowed by the access boundary rule.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "An access boundary rule in an IAM policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Rules to be applied.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleIamAccessBoundaryPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamAccessBoundaryPolicy), &result)
	return &result
}
