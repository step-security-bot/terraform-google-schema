package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogTagTemplate = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The display name for this template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_delete": {
        "description": "This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the tag template in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Template location region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_template_id": {
        "description": "The id of the tag template to create.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "fields": {
        "block": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description for this field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "computed": true,
              "description": "The display name for this field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "is_required": {
              "computed": true,
              "description": "Whether this is a required field. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "computed": true,
              "description": "The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}",
              "description_kind": "plain",
              "type": "string"
            },
            "order": {
              "computed": true,
              "description": "The order of this field with respect to other fields in this tag template.\nA higher value indicates a more important field. The value can be negative.\nMultiple fields can have the same order, and field orders within a tag do not have to be sequential.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "type": {
              "block": {
                "attributes": {
                  "primitive_type": {
                    "computed": true,
                    "description": "Represents primitive types - string, bool etc.\n Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: [\"DOUBLE\", \"STRING\", \"BOOL\", \"TIMESTAMP\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "enum_type": {
                    "block": {
                      "block_types": {
                        "allowed_values": {
                          "block": {
                            "attributes": {
                              "display_name": {
                                "description": "The display name of the enum value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The set of allowed values for this enum. The display names of the\nvalues must be case-insensitively unique within this set. Currently,\nenum values can only be added to the list of allowed values. Deletion\nand renaming of enum values are not supported.\nCan have up to 500 allowed values.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description": "Represents an enum type.\n Exactly one of 'primitive_type' or 'enum_type' must be set",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The type of value this tag field can contain.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields. The change of field_id will be resulting in re-creating of field. The change of primitive_type will be resulting in re-creating of field, however if the field is a required, you cannot update it.",
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

func GoogleDataCatalogTagTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogTagTemplate), &result)
	return &result
}
