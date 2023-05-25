package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGameServicesGameServerConfig = `{
  "block": {
    "attributes": {
      "config_id": {
        "description": "A unique id for the deployment config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deployment_id": {
        "description": "A unique id for the deployment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the game server config.",
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
        "description": "The labels associated with this game server config. Each label is a\nkey-value pair.",
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
        "description": "The resource name of the game server config, in the form:\n\n'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}/configs/{config_id}'.",
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
      "fleet_configs": {
        "block": {
          "attributes": {
            "fleet_spec": {
              "description": "The fleet spec, which is sent to Agones to configure fleet.\nThe spec can be passed as inline json but it is recommended to use a file reference\ninstead. File references can contain the json or yaml format of the fleet spec. Eg:\n\n* fleet_spec = jsonencode(yamldecode(file(\"fleet_configs.yaml\")))\n* fleet_spec = file(\"fleet_configs.json\")\n\nThe format of the spec can be found :\n'https://agones.dev/site/docs/reference/fleet/'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the FleetConfig.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The fleet config contains list of fleet specs. In the Single Cloud, there\nwill be only one.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "scaling_configs": {
        "block": {
          "attributes": {
            "fleet_autoscaler_spec": {
              "description": "Fleet autoscaler spec, which is sent to Agones.\nExample spec can be found :\nhttps://agones.dev/site/docs/reference/fleetautoscaler/",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "The name of the ScalingConfig",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "schedules": {
              "block": {
                "attributes": {
                  "cron_job_duration": {
                    "description": "The duration for the cron job event. The duration of the event is effective\nafter the cron job's start time.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cron_spec": {
                    "description": "The cron definition of the scheduled event. See\nhttps://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as\ndefined by the realm.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "end_time": {
                    "description": "The end time of the event.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description": "The start time of the event.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The schedules to which this scaling config applies.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "selectors": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Set of labels to group by.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Labels used to identify the clusters to which this scaling config\napplies. A cluster is subject to this scaling config if its labels match\nany of the selector entries.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Optional. This contains the autoscaling settings.",
          "description_kind": "plain"
        },
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

func GoogleGameServicesGameServerConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGameServicesGameServerConfig), &result)
	return &result
}
