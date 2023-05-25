package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctionsFunction = `{
  "block": {
    "attributes": {
      "available_memory_mb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entry_point": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "event_trigger": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "event_type": "string",
              "failure_policy": [
                "list",
                [
                  "object",
                  {
                    "retry": "bool"
                  }
                ]
              ],
              "resource": "string"
            }
          ]
        ]
      },
      "https_trigger_url": {
        "computed": true,
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
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "max_instances": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_account_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_archive_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_archive_object": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_repository": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deployed_url": "string",
              "url": "string"
            }
          ]
        ]
      },
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "trigger_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trigger_http": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "trigger_topic": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_connector": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudfunctionsFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctionsFunction), &result)
	return &result
}
