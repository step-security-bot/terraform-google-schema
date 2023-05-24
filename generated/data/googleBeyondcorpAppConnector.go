package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpAppConnector = `{
  "block": {
    "attributes": {
      "display_name": {
        "computed": true,
        "description": "An arbitrary user-provided name for the AppConnector.",
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
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "ID of the AppConnector.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_info": {
        "computed": true,
        "description": "Principal information about the Identity of the AppConnector.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "service_account": [
                "list",
                [
                  "object",
                  {
                    "email": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region of the AppConnector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Represents the different states of a AppConnector.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBeyondcorpAppConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBeyondcorpAppConnector), &result)
	return &result
}
