package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDocumentAiProcessorDefaultVersion = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "processor": {
        "description": "The processor to set the version on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version": {
        "description": "The version to set. Using 'stable' or 'rc' will cause the API to return the latest version in that release channel.\nApply 'lifecycle.ignore_changes' to the 'version' field to suppress this diff.",
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

func GoogleDocumentAiProcessorDefaultVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDocumentAiProcessorDefaultVersion), &result)
	return &result
}
