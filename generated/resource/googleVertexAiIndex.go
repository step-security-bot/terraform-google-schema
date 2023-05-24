package resource

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
        "description": "The description of the Index.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.",
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
        "description": "The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.\n* BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.\n* STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels with user-defined metadata to organize your Indexes.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "metadata_schema_uri": {
        "computed": true,
        "description": "Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Index.",
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
        "description": "The region of the index. eg us-central1",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the Index was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "contents_delta_uri": {
              "description": "Allows inserting, updating  or deleting the contents of the Matching Engine Index.\nThe string must be a valid Cloud Storage directory path. If this\nfield is set when calling IndexService.UpdateIndex, then no other\nIndex field can be also updated as part of the same call.\nThe expected structure and format of the files this URI points to is\ndescribed at https://cloud.google.com/vertex-ai/docs/matching-engine/using-matching-engine#input-data-format",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_complete_overwrite": {
              "description": "If this field is set together with contentsDeltaUri when calling IndexService.UpdateIndex,\nthen existing content of the Index will be replaced by the data from the contentsDeltaUri.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "config": {
              "block": {
                "attributes": {
                  "approximate_neighbors_count": {
                    "description": "The default number of neighbors to find via approximate search before exact reordering is\nperformed. Exact reordering is a procedure where results returned by an\napproximate search algorithm are reordered via a more expensive distance computation.\nRequired if tree-AH algorithm is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "dimensions": {
                    "description": "The number of dimensions of the input vectors.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "distance_measure_type": {
                    "description": "The distance measure used in nearest neighbor search. The value must be one of the followings:\n* SQUARED_L2_DISTANCE: Euclidean (L_2) Distance\n* L1_DISTANCE: Manhattan (L_1) Distance\n* COSINE_DISTANCE: Cosine Distance. Defined as 1 - cosine similarity.\n* DOT_PRODUCT_DISTANCE: Dot Product Distance. Defined as a negative of the dot product",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "feature_norm_type": {
                    "description": "Type of normalization to be carried out on each vector. The value must be one of the followings:\n* UNIT_L2_NORM: Unit L2 normalization type\n* NONE: No normalization type is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "algorithm_config": {
                    "block": {
                      "block_types": {
                        "brute_force_config": {
                          "block": {
                            "description": "Configuration options for using brute force search, which simply implements the\nstandard linear search in the database for each query.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tree_ah_config": {
                          "block": {
                            "attributes": {
                              "leaf_node_embedding_count": {
                                "description": "Number of embeddings on each leaf node. The default value is 1000 if not set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "leaf_nodes_to_search_percent": {
                                "description": "The default percentage of leaf nodes that any query may be searched. Must be in\nrange 1-100, inclusive. The default value is 10 (means 10%) if not set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Configuration options for using the tree-AH algorithm (Shallow tree + Asymmetric Hashing).\nPlease refer to this paper for more details: https://arxiv.org/abs/1908.10396",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The configuration with regard to the algorithms used for efficient search.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration of the Matching Engine Index.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "An additional information about the Index",
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

func GoogleVertexAiIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiIndex), &result)
	return &result
}
