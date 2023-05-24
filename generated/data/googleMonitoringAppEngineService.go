package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringAppEngineService = `{
  "block": {
    "attributes": {
      "display_name": {
        "computed": true,
        "description": "Name used for UI elements listing this Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "module_id": {
        "description": "The ID of the App Engine module underlying this service. \nCorresponds to the 'moduleId' resource label for a 'gae_app'\nmonitored resource(see https://cloud.google.com/monitoring/api/resources#tag_gae_app)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The full resource name for this service. The syntax is:\nprojects/[PROJECT_ID]/services/[SERVICE_ID].",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description": "An optional service ID to use. If not given, the server will generate a\nservice ID.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Labels which have been used to annotate the service. Label keys must start\nwith a letter. Label keys and values may contain lowercase letters,\nnumbers, underscores, and dashes. Label keys and values have a maximum\nlength of 63 characters, and must be less than 128 bytes in size. Up to 64\nlabel entries may be stored. For labels which do not have a semantic value,\nthe empty string may be supplied for the label value.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMonitoringAppEngineServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringAppEngineService), &result)
	return &result
}
