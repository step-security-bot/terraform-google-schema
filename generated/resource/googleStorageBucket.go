package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucket = `{
  "block": {
    "attributes": {
      "bucket_policy_only": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_destroy": {
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "predefined_acl": {
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
      "requester_pays": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cors": {
        "block": {
          "attributes": {
            "max_age_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "method": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "origin": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "response_header": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "encryption": {
        "block": {
          "attributes": {
            "default_kms_key_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "lifecycle_rule": {
        "block": {
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "storage_class": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            },
            "condition": {
              "block": {
                "attributes": {
                  "age": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "created_before": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_live": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "matches_storage_class": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "num_newer_versions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "with_state": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "list"
      },
      "logging": {
        "block": {
          "attributes": {
            "log_bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_object_prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "is_locked": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retention_period": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "versioning": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "website": {
        "block": {
          "attributes": {
            "main_page_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "not_found_page": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucket), &result)
	return &result
}
