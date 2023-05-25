package data

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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pubsub_configs": {
        "computed": true,
        "description": "How this repository publishes a change in the repository through Cloud Pub/Sub. \nKeyed by the topic names.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "message_format": "string",
              "service_account_email": "string",
              "topic": "string"
            }
          ]
        ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSourcerepoRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSourcerepoRepository), &result)
	return &result
}
