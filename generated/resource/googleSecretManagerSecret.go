package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerSecret = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the Secret was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels assigned to this Secret.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be assigned to a given resource.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Secret. Format:\n'projects/{{project}}/secrets/{{secret_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_id": {
        "description": "This must be unique within the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ttl": {
        "description": "The TTL for the Secret.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "replication": {
        "block": {
          "attributes": {
            "automatic": {
              "description": "The Secret will automatically be replicated without any restrictions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "user_managed": {
              "block": {
                "block_types": {
                  "replicas": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description": "The canonical IDs of the location to replicate data. For example: \"us-east1\".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "customer_managed_encryption": {
                          "block": {
                            "attributes": {
                              "kms_key_name": {
                                "description": "Describes the Cloud KMS encryption key that will be used to protect destination secret.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Customer Managed Encryption for the secret.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of Replicas for this Secret. Cannot be empty.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Secret will automatically be replicated without any restrictions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The replication policy of the secret data attached to the Secret. It cannot be changed\nafter the Secret has been created.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "rotation": {
        "block": {
          "attributes": {
            "next_rotation_time": {
              "description": "Timestamp in UTC at which the Secret is scheduled to rotate.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rotation_period": {
              "description": "The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years).\nIf rotationPeriod is set, 'next_rotation_time' must be set. 'next_rotation_time' will be advanced by this period when the service automatically sends rotation notifications.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The rotation time and period for a Secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be set to configure rotation.",
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
      },
      "topics": {
        "block": {
          "attributes": {
            "name": {
              "description": "The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.\nFor publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSecretManagerSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerSecret), &result)
	return &result
}
