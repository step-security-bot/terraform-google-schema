package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiFeaturestoreEntitytypeFeature = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the entity type was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the feature.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entitytype": {
        "description": "The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Used to perform consistent read-modify-write updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to the feature.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the entity type was most recently updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "value_type": {
        "description": "Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType",
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

func GoogleVertexAiFeaturestoreEntitytypeFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiFeaturestoreEntitytypeFeature), &result)
	return &result
}
