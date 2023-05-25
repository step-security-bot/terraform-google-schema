package data

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
        "computed": true,
        "description": "The resource name of the Cloud KMS CryptoKey to be used to protect access\nto messages published on this topic. Your project's PubSub service account\n('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nThe expected format is 'projects/*/locations/*/keyRings/*/cryptoKeys/*'",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to this Topic.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "message_storage_policy": {
        "computed": true,
        "description": "Policy constraining the set of Google Cloud Platform regions where\nmessages published to the topic may be stored. If not present, then no\nconstraints are in effect.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_persistence_regions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "Name of the topic.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema_settings": {
        "computed": true,
        "description": "Settings for validating messages published against a schema.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "encoding": "string",
              "schema": "string"
            }
          ]
        ]
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
