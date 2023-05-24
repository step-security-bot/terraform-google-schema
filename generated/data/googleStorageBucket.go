package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucket = `{
  "block": {
    "attributes": {
      "autoclass": {
        "computed": true,
        "description": "The bucket's autoclass configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "cors": {
        "computed": true,
        "description": "The bucket's Cross-Origin Resource Sharing (CORS) configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_age_seconds": "number",
              "method": [
                "list",
                "string"
              ],
              "origin": [
                "list",
                "string"
              ],
              "response_header": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "custom_placement_config": {
        "computed": true,
        "description": "The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_locations": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "default_event_based_hold": {
        "computed": true,
        "description": "Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption": {
        "computed": true,
        "description": "The bucket's encryption configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "default_kms_key_name": "string"
            }
          ]
        ]
      },
      "force_destroy": {
        "computed": true,
        "description": "When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to the bucket.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "lifecycle_rule": {
        "computed": true,
        "description": "The bucket's Lifecycle Rules configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": [
                "set",
                [
                  "object",
                  {
                    "storage_class": "string",
                    "type": "string"
                  }
                ]
              ],
              "condition": [
                "set",
                [
                  "object",
                  {
                    "age": "number",
                    "created_before": "string",
                    "custom_time_before": "string",
                    "days_since_custom_time": "number",
                    "days_since_noncurrent_time": "number",
                    "matches_prefix": [
                      "list",
                      "string"
                    ],
                    "matches_storage_class": [
                      "list",
                      "string"
                    ],
                    "matches_suffix": [
                      "list",
                      "string"
                    ],
                    "noncurrent_time_before": "string",
                    "num_newer_versions": "number",
                    "with_state": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description": "The Google Cloud Storage location",
        "description_kind": "plain",
        "type": "string"
      },
      "logging": {
        "computed": true,
        "description": "The bucket's Access \u0026 Storage Logs configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "log_bucket": "string",
              "log_object_prefix": "string"
            }
          ]
        ]
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
        "type": "string"
      },
      "public_access_prevention": {
        "computed": true,
        "description": "Prevents public access to a bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "requester_pays": {
        "computed": true,
        "description": "Enables Requester Pays on a storage bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "retention_policy": {
        "computed": true,
        "description": "Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_locked": "bool",
              "retention_period": "number"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description": "The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
        "description_kind": "plain",
        "type": "string"
      },
      "uniform_bucket_level_access": {
        "computed": true,
        "description": "Enables uniform bucket-level access on a bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "url": {
        "computed": true,
        "description": "The base URL of the bucket, in the format gs://\u003cbucket-name\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "versioning": {
        "computed": true,
        "description": "The bucket's Versioning configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "website": {
        "computed": true,
        "description": "Configuration if the bucket acts as a website.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "main_page_suffix": "string",
              "not_found_page": "string"
            }
          ]
        ]
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
