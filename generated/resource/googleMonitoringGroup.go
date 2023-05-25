package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringGroup = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "A user-assigned name for this group, used only for display\npurposes.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter": {
        "description": "The filter used to determine which monitored resources\nbelong to this group.",
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
      "is_cluster": {
        "description": "If true, the members of this group are considered to be a\ncluster. The system can perform additional analysis on\ngroups that are clusters.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for this group. The format is\n\"projects/{project_id_or_number}/groups/{group_id}\".",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_name": {
        "description": "The name of the group's parent, if it has one. The format is\n\"projects/{project_id_or_number}/groups/{group_id}\". For\ngroups with no parent, parentName is the empty string, \"\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func GoogleMonitoringGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringGroup), &result)
	return &result
}
