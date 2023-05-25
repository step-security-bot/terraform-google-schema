package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSourcerepoRepository = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Resource name of the repository, of the form '{{repo}}'.\nThe repo name may contain slashes. eg, 'name/with/slash'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "The disk usage of the repo, in bytes.",
        "description_kind": "plain",
        "type": "number"
      },
      "url": {
        "computed": true,
        "description": "URL to clone the repository from Google Cloud Source Repositories.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "pubsub_configs": {
        "block": {
          "attributes": {
            "message_format": {
              "description": "The format of the Cloud Pub/Sub messages. \n- PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.\n- JSON: The message payload is a JSON string of SourceRepoEvent. Possible values: [\"PROTOBUF\", \"JSON\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_account_email": {
              "computed": true,
              "description": "Email address of the service account used for publishing Cloud Pub/Sub messages. \nThis service account needs to be in the same project as the PubsubConfig. When added, \nthe caller needs to have iam.serviceAccounts.actAs permission on this service account. \nIf unspecified, it defaults to the compute engine default service account.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "How this repository publishes a change in the repository through Cloud Pub/Sub. \nKeyed by the topic names.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleSourcerepoRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSourcerepoRepository), &result)
	return &result
}
