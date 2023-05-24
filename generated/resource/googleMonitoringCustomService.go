package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringCustomService = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Name used for UI elements listing this Service.",
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
      "name": {
        "computed": true,
        "description": "The full resource name for this service. The syntax is:\nprojects/[PROJECT_ID]/services/[SERVICE_ID].",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description": "An optional service ID to use. If not given, the server will generate a\nservice ID.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_labels": {
        "description": "Labels which have been used to annotate the service. Label keys must start\nwith a letter. Label keys and values may contain lowercase letters,\nnumbers, underscores, and dashes. Label keys and values have a maximum\nlength of 63 characters, and must be less than 128 bytes in size. Up to 64\nlabel entries may be stored. For labels which do not have a semantic value,\nthe empty string may be supplied for the label value.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "telemetry": {
        "block": {
          "attributes": {
            "resource_name": {
              "description": "The full name of the resource that defines this service.\nFormatted as described in\nhttps://cloud.google.com/apis/design/resource_names.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration for how to query telemetry on a Service.",
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

func GoogleMonitoringCustomServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringCustomService), &result)
	return &result
}
