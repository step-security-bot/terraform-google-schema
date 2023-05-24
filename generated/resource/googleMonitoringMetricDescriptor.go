package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringMetricDescriptor = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A detailed description of the metric, which can be used in documentation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example \"Request count\".",
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
      "launch_stage": {
        "description": "The launch stage of the metric definition. Possible values: [\"LAUNCH_STAGE_UNSPECIFIED\", \"UNIMPLEMENTED\", \"PRELAUNCH\", \"EARLY_ACCESS\", \"ALPHA\", \"BETA\", \"GA\", \"DEPRECATED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_kind": {
        "description": "Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported. Possible values: [\"METRIC_KIND_UNSPECIFIED\", \"GAUGE\", \"DELTA\", \"CUMULATIVE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitored_resource_types": {
        "computed": true,
        "description": "If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here. This field allows time series to be associated with the intersection of this metric type and the monitored resource types in this list.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The resource name of the metric descriptor.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "unit": {
        "description": "The units in which the metric value is reported. It is only applicable if the\nvalueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of\nthe stored metric values.\n\nDifferent systems may scale the values to be more easily displayed (so a value of\n0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as\n3.5MBy). However, if the unit is KBy, then the value of the metric is always in\nthousands of bytes, no matter how it may be displayed.\n\nIf you want a custom metric to record the exact number of CPU-seconds used by a job,\nyou can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently\n1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as\n12005.\n\nAlternatively, if you want a custom metric to record data in a more granular way, you\ncan create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value\n12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).\nThe supported units are a subset of The Unified Code for Units of Measure standard.\nMore info can be found in the API documentation\n(https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "value_type": {
        "description": "Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported. Possible values: [\"BOOL\", \"INT64\", \"DOUBLE\", \"STRING\", \"DISTRIBUTION\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "labels": {
        "block": {
          "attributes": {
            "description": {
              "description": "A human-readable description for the label.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "description": "The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value_type": {
              "description": "The type of data that can be assigned to the label. Default value: \"STRING\" Possible values: [\"STRING\", \"BOOL\", \"INT64\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "metadata": {
        "block": {
          "attributes": {
            "ingest_delay": {
              "description": "The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?\u0026_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_period": {
              "description": "The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?\u0026_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Metadata which can be used to guide usage of the metric.",
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

func GoogleMonitoringMetricDescriptorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringMetricDescriptor), &result)
	return &result
}
