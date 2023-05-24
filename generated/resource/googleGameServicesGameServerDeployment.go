package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGameServicesGameServerDeployment = `{
  "block": {
    "attributes": {
      "deployment_id": {
        "description": "A unique id for the deployment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Human readable description of the game server deployment.",
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
        "description": "The labels associated with this game server deployment. Each label is a\nkey-value pair.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the Deployment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource id of the game server deployment, eg:\n\n'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'.\nFor example,\n\n'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleGameServicesGameServerDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGameServicesGameServerDeployment), &result)
	return &result
}
