package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGameServicesGameServerDeploymentRollout = `{
  "block": {
    "attributes": {
      "default_game_server_config": {
        "computed": true,
        "description": "This field points to the game server config that is\napplied by default to all realms and clusters. For example,\n\n'projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config'.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_id": {
        "description": "The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "game_server_config_overrides": {
        "computed": true,
        "description": "The game_server_config_overrides contains the per game server config\noverrides. The overrides are processed in the order they are listed. As\nsoon as a match is found for a cluster, the rest of the list is not\nprocessed.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_version": "string",
              "realms_selector": [
                "list",
                [
                  "object",
                  {
                    "realms": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
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
      "name": {
        "computed": true,
        "description": "The resource id of the game server deployment\n\neg: 'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGameServicesGameServerDeploymentRolloutSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGameServicesGameServerDeploymentRollout), &result)
	return &result
}
