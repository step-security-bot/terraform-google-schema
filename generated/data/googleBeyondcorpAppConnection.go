package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpAppConnection = `{
  "block": {
    "attributes": {
      "application_endpoint": {
        "computed": true,
        "description": "Address of the remote application endpoint for the BeyondCorp AppConnection.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "port": "number"
            }
          ]
        ]
      },
      "connectors": {
        "computed": true,
        "description": "List of AppConnectors that are authorised to be associated with this AppConnection",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "computed": true,
        "description": "An arbitrary user-provided name for the AppConnection.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway": {
        "computed": true,
        "description": "Gateway used by the AppConnection.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_gateway": "string",
              "ingress_port": "number",
              "type": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The type of network connectivity used by the AppConnection. Refer to\nhttps://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type\nfor a list of possible values.",
        "description_kind": "plain",
        "type": "string"
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
