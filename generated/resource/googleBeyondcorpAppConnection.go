package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpAppConnection = `{
  "block": {
    "attributes": {
      "connectors": {
        "description": "List of AppConnectors that are authorised to be associated with this AppConnection",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "An arbitrary user-provided name for the AppConnection.",
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
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "ID of the AppConnection.",
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
        "description": "The region of the AppConnection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of network connectivity used by the AppConnection. Refer to\nhttps://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type\nfor a list of possible values.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "application_endpoint": {
        "block": {
          "attributes": {
            "host": {
              "description": "Hostname or IP address of the remote application endpoint.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "Port of the remote application endpoint.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Address of the remote application endpoint for the BeyondCorp AppConnection.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "gateway": {
        "block": {
          "attributes": {
            "app_gateway": {
              "description": "AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ingress_port": {
              "computed": true,
              "description": "Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443.",
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "description": "The type of hosting used by the gateway. Refer to\nhttps://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1\nfor a list of possible values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "Server-defined URI for this resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Gateway used by the AppConnection.",
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

func GoogleBeyondcorpAppConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBeyondcorpAppConnection), &result)
	return &result
}
