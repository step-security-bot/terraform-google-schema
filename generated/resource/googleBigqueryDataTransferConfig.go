package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDataTransferConfig = `{
  "block": {
    "attributes": {
      "data_refresh_window_days": {
        "description": "The number of days to look back to automatically refresh the data.\nFor example, if dataRefreshWindowDays = 10, then every day BigQuery\nreingests data for [today-10, today-1], rather than ingesting data for\njust [today-1]. Only valid if the data source supports the feature.\nSet the value to 0 to use the default value.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_source_id": {
        "description": "The data source id. Cannot be changed once the transfer config is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_dataset_id": {
        "description": "The BigQuery target dataset id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disabled": {
        "description": "When set to true, no runs are scheduled for a given transfer.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The user specified display name for the transfer config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The geographic location where the transfer config should reside.\nExamples: US, EU, asia-northeast1. The default value is US.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the transfer config. Transfer config names have the\nform projects/{projectId}/locations/{location}/transferConfigs/{configId}.\nWhere configId is usually a uuid, but this is not required.\nThe name is ignored when creating a transfer config.",
        "description_kind": "plain",
        "type": "string"
      },
      "params": {
        "description": "These parameters are specific to each data source.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule": {
        "description": "Data transfer schedule. If the data source does not support a custom\nschedule, this should be empty. If it is empty, the default value for\nthe data source will be used. The specified times are in UTC. Examples\nof valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,\njun 13:15, and first sunday of quarter 00:00. See more explanation\nabout the format here:\nhttps://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format\nNOTE: the granularity should be at least 8 hours, or less frequent.",
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

func GoogleBigqueryDataTransferConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDataTransferConfig), &result)
	return &result
}
