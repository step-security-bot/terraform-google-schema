package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudSchedulerJob = `{
  "block": {
    "attributes": {
      "attempt_deadline": {
        "description": "The deadline for job attempts. If the request handler does not respond by this deadline then the request is\ncancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in\nexecution logs. Cloud Scheduler will retry the job according to the RetryConfig.\nThe allowed duration for this deadline is:\n* For HTTP targets, between 15 seconds and 30 minutes.\n* For App Engine HTTP targets, between 15 seconds and 24 hours.\n* **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\"",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A human-readable description for the job.\nThis string must not contain more than 500 characters.",
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
      "paused": {
        "computed": true,
        "description": "Sets the job to a paused state. Jobs default to being enabled when this property is not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where the scheduler job resides. If it is not provided, Terraform will use the provider default.",
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
      "state": {
        "computed": true,
        "description": "State of the job.",
        "description_kind": "plain",
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
              "description": "HTTP request body.\nA request body is allowed only if the HTTP method is POST or PUT.\nIt will result in invalid argument error to set a body on a job with an incompatible HttpMethod.\n\nA base64-encoded string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "description": "HTTP request headers.\nThis map contains the header field names and values.\nHeaders can be set when the job is created.",
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
              "description": "The relative URI.\nThe relative URL must begin with \"/\" and must be a valid HTTP relative URL.\nIt can contain a path, query string arguments, and \\# fragments.\nIf the relative URL is empty, then the root path \"/\" will be used.\nNo spaces are allowed, and the maximum length allowed is 2083 characters",
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
                "description": "App Engine Routing setting for the job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "App Engine HTTP target.\nIf the job providers a App Engine HTTP target the cron will\nsend a request to the service instance",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_target": {
        "block": {
          "attributes": {
            "body": {
              "description": "HTTP request body.\nA request body is allowed only if the HTTP method is POST, PUT, or PATCH.\nIt is an error to set body on a job with an incompatible HttpMethod.\n\nA base64-encoded string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "description": "This map contains the header field names and values.\nRepeated headers are not supported, but a header value can contain commas.",
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
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Contains information needed for generating an OAuth token.\nThis type of authorization should be used when sending requests to a GCP endpoint.",
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
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Contains information needed for generating an OpenID Connect token.\nThis type of authorization should be used when sending requests to third party endpoints or Cloud Run.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "HTTP target.\nIf the job providers a http_target the cron will\nsend a request to the targeted url",
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
              "description": "The message payload for PubsubMessage.\nPubsub message must contain either non-empty data, or at least one attribute.\n\n A base64-encoded string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic_name": {
              "description": "The full resource name for the Cloud Pub/Sub topic to which\nmessages will be published when a job is delivered. ~\u003e**NOTE:**\nThe topic name must be in the same format as required by PubSub's\nPublishRequest.name, e.g. 'projects/my-project/topics/my-topic'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Pub/Sub target\nIf the job providers a Pub/Sub target the cron will publish\na message to the provided topic",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retry_config": {
        "block": {
          "attributes": {
            "max_backoff_duration": {
              "computed": true,
              "description": "The maximum amount of time to wait before retrying a job after it fails.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_doublings": {
              "computed": true,
              "description": "The time between retries will double maxDoublings times.\nA job's retry interval starts at minBackoffDuration,\nthen doubles maxDoublings times, then increases linearly,\nand finally retries retries at intervals of maxBackoffDuration up to retryCount times.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_duration": {
              "computed": true,
              "description": "The time limit for retrying a failed job, measured from time when an execution was first attempted.\nIf specified with retryCount, the job will be retried until both limits are reached.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_backoff_duration": {
              "computed": true,
              "description": "The minimum amount of time to wait before retrying a job after it fails.\nA duration in seconds with up to nine fractional digits, terminated by 's'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_count": {
              "computed": true,
              "description": "The number of attempts that the system will make to run a\njob using the exponential backoff procedure described by maxDoublings.\nValues greater than 5 and negative values are not allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "By default, if a job does not complete successfully,\nmeaning that an acknowledgement is not received from the handler,\nthen it will be retried with exponential backoff according to the settings",
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

func GoogleCloudSchedulerJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudSchedulerJob), &result)
	return &result
}
