package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogTag = `{
  "block": {
    "attributes": {
      "column": {
        "description": "Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an\nindividual column based on that schema.\n\nFor attaching a tag to a nested column, use '.' to separate the column names. Example:\n'outer_column.inner_column'",
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
      "name": {
        "computed": true,
        "description": "The resource name of the tag in URL format. Example:\nprojects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or\nprojects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id}\nwhere tag_id is a system-generated identifier. Note that this Tag may not actually be stored in the location in this name.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to\nall entries in that group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template": {
        "description": "The resource name of the tag template that this tag uses. Example:\nprojects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}\nThis field cannot be modified after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_displayname": {
        "computed": true,
        "description": "The display name of the tag template.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "fields": {
        "block": {
          "attributes": {
            "bool_value": {
              "description": "Holds the value for a tag field with boolean type.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "display_name": {
              "computed": true,
              "description": "The display name of this field",
              "description_kind": "plain",
              "type": "string"
            },
            "double_value": {
              "description": "Holds the value for a tag field with double type.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enum_value": {
              "description": "The display name of the enum value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "order": {
              "computed": true,
              "description": "The order of this field with respect to other fields in this tag. For example, a higher value can indicate\na more important field. The value can be negative. Multiple fields can have the same order, and field orders\nwithin a tag do not have to be sequential.",
              "description_kind": "plain",
              "type": "number"
            },
            "string_value": {
              "description": "Holds the value for a tag field with string type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timestamp_value": {
              "description": "Holds the value for a tag field with timestamp type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "This maps the ID of a tag field to the value of and additional information about that field.\nValid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.",
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

func GoogleDataCatalogTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogTag), &result)
	return &result
}
