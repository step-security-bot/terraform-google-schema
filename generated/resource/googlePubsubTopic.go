package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubTopic = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "description": "The resource name of the Cloud KMS CryptoKey to be used to protect access\nto messages published on this topic. Your project's PubSub service account\n('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nThe expected format is 'projects/*/locations/*/keyRings/*/cryptoKeys/*'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to this Topic.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "message_retention_duration": {
        "description": "Indicates the minimum duration to retain a message after it is published\nto the topic. If this field is set, messages published to the topic in\nthe last messageRetentionDuration are always available to subscribers.\nFor instance, it allows any attached subscription to seek to a timestamp\nthat is up to messageRetentionDuration in the past. If this field is not\nset, message retention is controlled by settings on individual subscriptions.\nCannot be more than 31 days or less than 10 minutes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the topic.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "message_storage_policy": {
        "block": {
          "attributes": {
            "allowed_persistence_regions": {
              "description": "A list of IDs of GCP regions where messages that are published to\nthe topic may be persisted in storage. Messages published by\npublishers running in non-allowed GCP regions (or running outside\nof GCP altogether) will be routed for storage in one of the\nallowed regions. An empty list means that no regions are allowed,\nand is not a valid configuration.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Policy constraining the set of Google Cloud Platform regions where\nmessages published to the topic may be stored. If not present, then no\nconstraints are in effect.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schema_settings": {
        "block": {
          "attributes": {
            "encoding": {
              "description": "The encoding of messages validated against schema. Default value: \"ENCODING_UNSPECIFIED\" Possible values: [\"ENCODING_UNSPECIFIED\", \"JSON\", \"BINARY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema": {
              "description": "The name of the schema that messages published should be\nvalidated against. Format is projects/{project}/schemas/{schema}.\nThe value of this field will be _deleted-schema_\nif the schema has been deleted.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Settings for validating messages published against a schema.",
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

func GooglePubsubTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubTopic), &result)
	return &result
}
