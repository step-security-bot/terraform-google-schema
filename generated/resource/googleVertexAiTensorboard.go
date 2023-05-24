package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiTensorboard = `{
  "block": {
    "attributes": {
      "blob_storage_path_prefix": {
        "computed": true,
        "description": "Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the Tensorboard was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of this Tensorboard.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User provided name of this Tensorboard.",
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
      "labels": {
        "description": "The labels with user-defined metadata to organize your Tensorboards.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Name of the Tensorboard.",
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
        "description": "The region of the tensorboard. eg us-central1",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "run_count": {
        "computed": true,
        "description": "The number of Runs stored in this Tensorboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the Tensorboard was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.\nHas the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.",
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

func GoogleVertexAiTensorboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiTensorboard), &result)
	return &result
}
