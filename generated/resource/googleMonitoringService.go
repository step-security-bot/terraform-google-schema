package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringService = `{
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
        "description": "An optional service ID to use. If not given, the server will generate a\nservice ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "telemetry": {
        "computed": true,
        "description": "Configuration for how to query telemetry on a Service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_name": "string"
            }
          ]
        ]
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
      "basic_service": {
        "block": {
          "attributes": {
            "service_labels": {
              "description": "Labels that specify the resource that emits the monitoring data\nwhich is used for SLO reporting of this 'Service'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "service_type": {
              "description": "The type of service that this basic service defines, e.g.\nAPP_ENGINE service type",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A well-known service type, defined by its service type and service labels.\nValid values of service types and services labels are described at\nhttps://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli",
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

func GoogleMonitoringServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringService), &result)
	return &result
}
