package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogPolicyTag = `{
  "block": {
    "attributes": {
      "child_policy_tags": {
        "computed": true,
        "description": "Resource names of child policy tags of this policy tag.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "Description of this policy tag. It must: contain only unicode characters, tabs,\nnewlines, carriage returns and page breaks; and be at most 2000 bytes long when\nencoded in UTF-8. If not set, defaults to an empty description.\nIf not set, defaults to an empty description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User defined name of this policy tag. It must: be unique within the parent\ntaxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;\nnot start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.",
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
        "description": "Resource name of this policy tag, whose format is:\n\"projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}\"",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_policy_tag": {
        "description": "Resource name of this policy tag's parent policy tag.\nIf empty, it means this policy tag is a top level policy tag.\nIf not set, defaults to an empty string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "taxonomy": {
        "description": "Taxonomy the policy tag is associated with",
        "description_kind": "plain",
        "required": true,
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

func GoogleDataCatalogPolicyTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogPolicyTag), &result)
	return &result
}
