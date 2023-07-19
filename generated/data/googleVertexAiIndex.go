package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiIndex = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the Index was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployed_indexes": {
        "computed": true,
        "description": "The pointers to DeployedIndexes created from this Index. An Index can be only deleted if all its DeployedIndexes had been undeployed first.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deployed_index_id": "string",
              "index_endpoint": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "The description of the Index.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.",
        "description_kind": "plain",
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
      "index_stats": {
        "computed": true,
        "description": "Stats of the index resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "shards_count": "number",
              "vectors_count": "string"
            }
          ]
        ]
      },
      "index_update_method": {
        "computed": true,
        "description": "The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.\n* BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.\n* STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "The labels with user-defined metadata to organize your Indexes.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "metadata": {
        "computed": true,
        "description": "An additional information about the Index",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config": [
                "list",
                [
                  "object",
                  {
                    "algorithm_config": [
                      "list",
                      [
                        "object",
                        {
                          "brute_force_config": [
                            "list",
                            [
                              "object",
                              {}
                            ]
                          ],
                          "tree_ah_config": [
                            "list",
                            [
                              "object",
                              {
                                "leaf_node_embedding_count": "number",
                                "leaf_nodes_to_search_percent": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "approximate_neighbors_count": "number",
                    "dimensions": "number",
                    "distance_measure_type": "string",
                    "feature_norm_type": "string",
                    "shard_size": "string"
                  }
                ]
              ],
              "contents_delta_uri": "string",
              "is_complete_overwrite": "bool"
            }
          ]
        ]
      },
      "metadata_schema_uri": {
        "computed": true,
        "description": "Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The resource name of the Index.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region of the index. eg us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the Index was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVertexAiIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiIndex), &result)
	return &result
}
