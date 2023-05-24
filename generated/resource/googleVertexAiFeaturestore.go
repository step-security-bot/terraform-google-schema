package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiFeaturestore = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the featurestore was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Used to perform consistent read-modify-write updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "force_destroy": {
        "description": "If set to true, any EntityTypes and Features for this Featurestore will also be deleted",
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
      "labels": {
        "description": "A set of key/value label pairs to assign to this Featurestore.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The region of the dataset. eg us-central1",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the featurestore was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "If set, both of the online and offline data storage will be secured by this key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "online_serving_config": {
        "block": {
          "attributes": {
            "fixed_node_count": {
              "description": "The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "scaling": {
              "block": {
                "attributes": {
                  "max_node_count": {
                    "description": "The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "min_node_count": {
                    "description": "The minimum number of nodes to scale down to. Must be greater than or equal to 1.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Config for online serving resources.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleVertexAiFeaturestoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiFeaturestore), &result)
	return &result
}
