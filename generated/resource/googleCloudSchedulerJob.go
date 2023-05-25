package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudSchedulerJob = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description for the job. \nThis string must not contain more than 500 characters.",
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
        "description": "The name of the job.",
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
        "computed": true,
        "description": "Region where the scheduler job resides",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule": {
        "description": "Describes the schedule on which the job will be executed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "description": "Specifies the time zone to be used in interpreting schedule.\nThe value of this field must be a time zone name from the tz database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "app_engine_http_target": {
        "block": {
          "attributes": {
            "body": {
              "description": "HTTP request body. \nA request body is allowed only if the HTTP method is POST or PUT. \nIt will result in invalid argument error to set a body on a job with an incompatible HttpMethod.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "description": "HTTP request headers.\nThis map contains the header field names and values. \nHeaders can be set when the job is created.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "http_method": {
              "description": "Which HTTP method to use for the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "relative_uri": {
              "description": "The relative URI.\nThe relative URL must begin with \"/\" and must be a valid HTTP relative URL. \nIt can contain a path, query string arguments, and \\# fragments. \nIf the relative URL is empty, then the root path \"/\" will be used. \nNo spaces are allowed, and the maximum length allowed is 2083 characters",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "app_engine_routing": {
              "block": {
                "attributes": {
                  "instance": {
                    "description": "App instance.\nBy default, the job is sent to an instance which is available when the job is attempted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service": {
                    "description": "App service.\nBy default, the job is sent to the service which is the default service when the job is attempted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "description": "App version.\nBy default, the job is sent to the version which is the default version when the job is attempted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_target": {
        "block": {
          "attributes": {
            "body": {
              "description": "HTTP request body. \nA request body is allowed only if the HTTP method is POST, PUT, or PATCH. \nIt is an error to set body on a job with an incompatible HttpMethod.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "description": "This map contains the header field names and values. \nRepeated headers are not supported, but a header value can contain commas.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "http_method": {
              "description": "Which HTTP method to use for the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "description": "The full URI path that the request will be sent to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "oauth_token": {
              "block": {
                "attributes": {
                  "scope": {
                    "description": "OAuth scope to be used for generating OAuth access token. If not specified,\n\"https://www.googleapis.com/auth/cloud-platform\" will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "Service account email to be used for generating OAuth token.\nThe service account must be within the same project as the job.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oidc_token": {
              "block": {
                "attributes": {
                  "audience": {
                    "description": "Audience to be used when generating OIDC token. If not specified,\nthe URI specified in target will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "Service account email to be used for generating OAuth token.\nThe service account must be within the same project as the job.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pubsub_target": {
        "block": {
          "attributes": {
            "attributes": {
              "description": "Attributes for PubsubMessage.\nPubsub message must contain either non-empty data, or at least one attribute.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "data": {
              "description": "The message payload for PubsubMessage.\nPubsub message must contain either non-empty data, or at least one attribute.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic_name": {
              "description": "The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. \nThe topic name must be in the same format as required by PubSub's PublishRequest.name, \nfor example projects/PROJECT_ID/topics/TOPIC_ID.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retry_config": {
        "block": {
          "attributes": {
            "max_backoff_duration": {
              "description": "The maximum amount of time to wait before retrying a job after it fails.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_doublings": {
              "description": "The time between retries will double maxDoublings times.\nA job's retry interval starts at minBackoffDuration, \nthen doubles maxDoublings times, then increases linearly, \nand finally retries retries at intervals of maxBackoffDuration up to retryCount times.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_duration": {
              "description": "The time limit for retrying a failed job, measured from time when an execution was first attempted. \nIf specified with retryCount, the job will be retried until both limits are reached.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_backoff_duration": {
              "description": "The minimum amount of time to wait before retrying a job after it fails.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_count": {
              "description": "The number of attempts that the system will make to run a \njob using the exponential backoff procedure described by maxDoublings.\nValues greater than 5 and negative values are not allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
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

func GoogleCloudSchedulerJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudSchedulerJob), &result)
	return &result
}
