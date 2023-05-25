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
        "optional": true,
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
      "notification_pubsub_topic": {
        "description": "Pub/Sub topic where notifications will be sent after transfer runs\nassociated with this transfer config finish.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "service_account_name": {
        "description": "Optional service account name. If this field is set, transfer config will\nbe created with this service account credentials. It requires that\nrequesting user calling this API has permissions to act as this service account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "email_preferences": {
        "block": {
          "attributes": {
            "enable_failure_email": {
              "description": "If true, email notifications will be sent on transfer run failures.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Email notifications will be sent according to these preferences to the\nemail address of the user who owns this transfer config.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schedule_options": {
        "block": {
          "attributes": {
            "disable_auto_scheduling": {
              "description": "If true, automatic scheduling of data transfer runs for this\nconfiguration will be disabled. The runs can be started on ad-hoc\nbasis using transferConfigs.startManualRuns API. When automatic\nscheduling is disabled, the TransferConfig.schedule field will\nbe ignored.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "end_time": {
              "description": "Defines time to stop scheduling transfer runs. A transfer run cannot be\nscheduled at or after the end time. The end time can be changed at any\nmoment. The time when a data transfer can be triggered manually is not\nlimited by this option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "description": "Specifies time to start scheduling transfer runs. The first run will be\nscheduled at or after the start time according to a recurrence pattern\ndefined in the schedule string. The start time can be changed at any\nmoment. The time when a data transfer can be triggered manually is not\nlimited by this option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Options customizing the data transfer schedule.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sensitive_params": {
        "block": {
          "attributes": {
            "secret_access_key": {
              "description": "The Secret Access Key of the AWS account transferring data from.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Different parameters are configured primarily using the the 'params' field on this\nresource. This block contains the parameters which contain secrets or passwords so that they can be marked\nsensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key\nin the 'params' map in the api request.\n\nCredentials may not be specified in both locations and will cause an error. Changing from one location\nto a different credential configuration in the config will require an apply to update state.",
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

func GoogleBigqueryDataTransferConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDataTransferConfig), &result)
	return &result
}
