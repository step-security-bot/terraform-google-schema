package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudTasksQueue = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the queue",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The queue name.",
        "description_kind": "plain",
        "optional": true,
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
      "app_engine_routing_override": {
        "block": {
          "attributes": {
            "host": {
              "computed": true,
              "description": "The host that the task is sent to.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance": {
              "description": "App instance.\n\nBy default, the task is sent to an instance which is available when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service": {
              "description": "App service.\n\nBy default, the task is sent to the service which is the default service when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "App version.\n\nBy default, the task is sent to the version which is the default version when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Overrides for task-level appEngineRouting. These settings apply only\nto App Engine tasks in this queue",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rate_limits": {
        "block": {
          "attributes": {
            "max_burst_size": {
              "computed": true,
              "description": "The max burst size.\n\nMax burst size limits how fast tasks in queue are processed when many tasks are\nin the queue and the rate is high. This field allows the queue to have a high\nrate so processing starts shortly after a task is enqueued, but still limits\nresource usage when many tasks are enqueued in a short period of time.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_concurrent_dispatches": {
              "computed": true,
              "description": "The maximum number of concurrent tasks that Cloud Tasks allows to\nbe dispatched for this queue. After this threshold has been\nreached, Cloud Tasks stops dispatching tasks until the number of\nconcurrent requests decreases.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_dispatches_per_second": {
              "computed": true,
              "description": "The maximum rate at which tasks are dispatched from this queue.\n\nIf unspecified when the queue is created, Cloud Tasks will pick the default.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Rate limits for task dispatches.\n\nThe queue's actual dispatch rate is the result of:\n\n* Number of tasks in the queue\n* User-specified throttling: rateLimits, retryConfig, and the queue's state.\n* System throttling due to 429 (Too Many Requests) or 503 (Service\n  Unavailable) responses from the worker, high error rates, or to\n  smooth sudden large traffic spikes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retry_config": {
        "block": {
          "attributes": {
            "max_attempts": {
              "computed": true,
              "description": "Number of attempts per task.\n\nCloud Tasks will attempt the task maxAttempts times (that is, if\nthe first attempt fails, then there will be maxAttempts - 1\nretries). Must be \u003e= -1.\n\nIf unspecified when the queue is created, Cloud Tasks will pick\nthe default.\n\n-1 indicates unlimited attempts.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_backoff": {
              "computed": true,
              "description": "A task will be scheduled for retry between minBackoff and\nmaxBackoff duration after it fails, if the queue's RetryConfig\nspecifies that the task should be retried.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_doublings": {
              "computed": true,
              "description": "The time between retries will double maxDoublings times.\n\nA task's retry interval starts at minBackoff, then doubles maxDoublings times,\nthen increases linearly, and finally retries retries at intervals of maxBackoff\nup to maxAttempts times.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_duration": {
              "computed": true,
              "description": "If positive, maxRetryDuration specifies the time limit for\nretrying a failed task, measured from when the task was first\nattempted. Once maxRetryDuration time has passed and the task has\nbeen attempted maxAttempts times, no further attempts will be\nmade and the task will be deleted.\n\nIf zero, then the task age is unlimited.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_backoff": {
              "computed": true,
              "description": "A task will be scheduled for retry between minBackoff and\nmaxBackoff duration after it fails, if the queue's RetryConfig\nspecifies that the task should be retried.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Settings that determine the retry behavior.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "stackdriver_logging_config": {
        "block": {
          "attributes": {
            "sampling_ratio": {
              "description": "Specifies the fraction of operations to write to Stackdriver Logging.\nThis field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the\ndefault and means that no operations are logged.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Configuration options for writing logs to Stackdriver Logging.",
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

func GoogleCloudTasksQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudTasksQueue), &result)
	return &result
}
