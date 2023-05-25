package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdentityGroup = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the Group was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An extended description to help users determine the purpose of a Group.\nMust not be longer than 4,096 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the Group.",
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
      "initial_group_config": {
        "description": "The initial configuration options for creating a Group.\n\nSee the\n[API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)\nfor possible values. Default value: \"EMPTY\" Possible values: [\"INITIAL_GROUP_CONFIG_UNSPECIFIED\", \"WITH_INITIAL_OWNER\", \"EMPTY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels that apply to the Group.\n\nMust not contain more than one entry. Must contain the entry\n'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or\n'system/groups/external': '' if the Group is an external-identity-mapped group.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Resource name of the Group in the format: groups/{group_id}, where group_id\nis the unique ID assigned to the Group.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the entity under which this Group resides in the\nCloud Identity resource hierarchy.\n\nMust be of the form identitysources/{identity_source_id} for external-identity-mapped\ngroups or customers/{customer_id} for Google Groups.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the Group was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "group_key": {
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
          "description": "EntityKey of the Group.",
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

func GoogleCloudIdentityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdentityGroup), &result)
	return &result
}
