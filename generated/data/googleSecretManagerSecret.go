package data

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
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "The labels assigned to this Secret.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be assigned to a given resource.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication": {
        "computed": true,
        "description": "The replication policy of the secret data attached to the Secret. It cannot be changed\nafter the Secret has been created.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic": "bool",
              "user_managed": [
                "list",
                [
                  "object",
                  {
                    "replicas": [
                      "list",
                      [
                        "object",
                        {
                          "customer_managed_encryption": [
                            "list",
                            [
                              "object",
                              {
                                "kms_key_name": "string"
                              }
                            ]
                          ],
                          "location": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "rotation": {
        "computed": true,
        "description": "The rotation time and period for a Secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be set to configure rotation.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "next_rotation_time": "string",
              "rotation_period": "string"
            }
          ]
        ]
      },
      "secret_id": {
        "description": "This must be unique within the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "topics": {
        "computed": true,
        "description": "A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "ttl": {
        "computed": true,
        "description": "The TTL for the Secret.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "type": "string"
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
