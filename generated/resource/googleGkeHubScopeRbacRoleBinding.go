package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubScopeRbacRoleBinding = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the RBAC Role Binding was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Time the RBAC Role Binding was deleted in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "group": {
        "description": "Principal that is be authorized in the cluster (at least of one the oneof\nis required). Updating one will unset the other automatically.\ngroup is the group, as seen by the kubernetes cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels for this ScopeRBACRoleBinding.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The resource name for the RBAC Role Binding",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope_id": {
        "description": "Id of the scope",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope_rbac_role_binding_id": {
        "description": "The client-provided identifier of the RBAC Role Binding.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the RBAC Role Binding resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "Google-generated UUID for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the RBAC Role Binding was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "user": {
        "description": "Principal that is be authorized in the cluster (at least of one the oneof\nis required). Updating one will unset the other automatically.\nuser is the name of the user as seen by the kubernetes cluster, example\n\"alice\" or \"alice@domain.tld\"",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "role": {
        "block": {
          "attributes": {
            "predefined_role": {
              "description": "PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: [\"UNKNOWN\", \"ADMIN\", \"EDIT\", \"VIEW\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Role to bind to the principal.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleGkeHubScopeRbacRoleBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubScopeRbacRoleBinding), &result)
	return &result
}
