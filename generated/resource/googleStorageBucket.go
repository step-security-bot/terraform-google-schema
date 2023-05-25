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
        "deprecated": true,
        "description": "Enables Bucket Policy Only access to a bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "default_event_based_hold": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_destroy": {
        "description": "When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run.",
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
        "description": "A set of key/value label pairs to assign to the bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The Google Cloud Storage location",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requester_pays": {
        "description": "Enables Requester Pays on a storage bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "description": "The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uniform_bucket_level_access": {
        "computed": true,
        "description": "Enables uniform bucket-level access on a bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "url": {
        "computed": true,
        "description": "The base URL of the bucket, in the format gs://\u003cbucket-name\u003e.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cors": {
        "block": {
          "attributes": {
            "max_age_seconds": {
              "description": "The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "method": {
              "description": "The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: \"*\" is permitted in the list of methods, and means \"any method\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "origin": {
              "description": "The list of Origins eligible to receive CORS response headers. Note: \"*\" is permitted in the list of origins, and means \"any Origin\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "response_header": {
              "description": "The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "The bucket's Cross-Origin Resource Sharing (CORS) configuration.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "encryption": {
        "block": {
          "attributes": {
            "default_kms_key_name": {
              "description": "A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified. You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The bucket's encryption configuration.",
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
                    "description": "The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The type of the action of this Lifecycle Rule. Supported values include: Delete and SetStorageClass.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The Lifecycle Rule's action configuration. A single block of this type is supported.",
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
                    "description": "Minimum age of an object in days to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "created_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_time_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days_since_custom_time": {
                    "description": "Number of days elapsed since the user-specified timestamp set on an object.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "days_since_noncurrent_time": {
                    "description": "Number of days elapsed since the noncurrent timestamp of an object. This\n\t\t\t\t\t\t\t\t\t\tcondition is relevant only for versioned objects.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "matches_storage_class": {
                    "description": "Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "noncurrent_time_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_newer_versions": {
                    "description": "Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "with_state": {
                    "computed": true,
                    "description": "Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: \"LIVE\", \"ARCHIVED\", \"ANY\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Lifecycle Rule's condition configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "The bucket's Lifecycle Rules configuration.",
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "list"
      },
      "logging": {
        "block": {
          "attributes": {
            "log_bucket": {
              "description": "The bucket that will receive log objects.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_object_prefix": {
              "computed": true,
              "description": "The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The bucket's Access \u0026 Storage Logs configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "is_locked": {
              "description": "If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retention_period": {
              "description": "The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "versioning": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "While set to true, versioning is fully enabled for this bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The bucket's Versioning configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "website": {
        "block": {
          "attributes": {
            "main_page_suffix": {
              "description": "Behaves as the bucket's directory index where missing objects are treated as potential directories.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "not_found_page": {
              "description": "The custom object to return when a requested resource is not found.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration if the bucket acts as a website.",
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
