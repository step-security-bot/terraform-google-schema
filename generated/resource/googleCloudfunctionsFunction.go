package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctionsFunction = `{
  "block": {
    "attributes": {
      "available_memory_mb": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entry_point": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "https_trigger_url": {
        "computed": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "max_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
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
      "region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_email": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_archive_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_archive_object": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "trigger_bucket": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trigger_http": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "trigger_topic": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_connector": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "event_trigger": {
        "block": {
          "attributes": {
            "event_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "failure_policy": {
              "block": {
                "attributes": {
                  "retry": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_repository": {
        "block": {
          "attributes": {
            "deployed_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
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
            "read": {
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

func GoogleCloudfunctionsFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctionsFunction), &result)
	return &result
}
