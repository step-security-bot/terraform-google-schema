package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdentityGroupMembership = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the Membership was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "group": {
        "description": "The name of the Group to create this membership in.",
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
      "name": {
        "computed": true,
        "description": "The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the membership.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the Membership was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "preferred_member_key": {
        "block": {
          "attributes": {
            "id": {
              "description": "The ID of the entity.\n\nFor Google-managed entities, the id must be the email address of an existing\ngroup or user.\n\nFor external-identity-mapped entities, the id must be a string conforming\nto the Identity Source's requirements.\n\nMust be unique within a namespace.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The namespace in which the entity exists.\n\nIf not specified, the EntityKey represents a Google-managed entity\nsuch as a Google user or a Google Group.\n\nIf specified, the EntityKey represents an external-identity-mapped group.\nThe namespace must correspond to an identity source created in Admin Console\nand must be in the form of 'identitysources/{identity_source_id}'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "EntityKey of the member.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "roles": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: [\"OWNER\", \"MANAGER\", \"MEMBER\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The MembershipRoles that apply to the Membership.\nMust not contain duplicate MembershipRoles with the same name.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func GoogleCloudIdentityGroupMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdentityGroupMembership), &result)
	return &result
}
