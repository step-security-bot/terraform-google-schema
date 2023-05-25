package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringAlertPolicy = `{
  "block": {
    "attributes": {
      "combiner": {
        "description": "How to combine the results of multiple conditions to\ndetermine if an incident should be opened.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_record": {
        "computed": true,
        "description": "A read-only record of the creation of the alerting policy.\nIf provided in a call to create or update, this field will\nbe ignored.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mutate_time": "string",
              "mutated_by": "string"
            }
          ]
        ]
      },
      "display_name": {
        "description": "A short name or phrase used to identify the policy in\ndashboards, notifications, and incidents. To avoid confusion, don't use\nthe same display name for multiple policies in the same project. The\nname is limited to 512 Unicode characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description": "Whether or not the policy is enabled. The default is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The unique resource name for this policy.\nIts syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_channels": {
        "description": "Identifies the notification channels to which notifications should be\nsent when incidents are opened or closed or when new violations occur\non an already opened incident. Each element of this array corresponds\nto the name field in each of the NotificationChannel objects that are\nreturned from the notificationChannels.list method. The syntax of the\nentries in this field is\n'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_labels": {
        "description": "This field is intended to be used for organizing and identifying the AlertPolicy\nobjects.The field can contain up to 64 entries. Each key and value is limited\nto 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values\ncan contain only lowercase letters, numerals, underscores, and dashes. Keys\nmust begin with a letter.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "conditions": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "A short name or phrase used to identify the\ncondition in dashboards, notifications, and\nincidents. To avoid confusion, don't use the same\ndisplay name for multiple conditions in the same\npolicy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The unique resource name for this condition.\nIts syntax is:\nprojects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]\n[CONDITION_ID] is assigned by Stackdriver Monitoring when\nthe condition is created as part of a new or updated alerting\npolicy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "condition_absent": {
              "block": {
                "attributes": {
                  "duration": {
                    "description": "The amount of time that a time series must\nfail to report new data to be considered\nfailing. Currently, only values that are a\nmultiple of a minute--e.g. 60s, 120s, or 300s\n--are supported.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter": {
                    "description": "A filter that identifies which time series\nshould be compared with the threshold.The\nfilter is similar to the one that is\nspecified in the\nMetricService.ListTimeSeries request (that\ncall is useful to verify the time series\nthat will be retrieved / processed) and must\nspecify the metric type and optionally may\ncontain restrictions on resource type,\nresource labels, and metric labels. This\nfield may not exceed 2048 Unicode characters\nin length.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "aggregations": {
                    "block": {
                      "attributes": {
                        "alignment_period": {
                          "description": "The alignment period for per-time\nseries alignment. If present,\nalignmentPeriod must be at least\n60 seconds. After per-time series\nalignment, each time series will\ncontain data points only on the\nperiod boundaries. If\nperSeriesAligner is not specified\nor equals ALIGN_NONE, then this\nfield is ignored. If\nperSeriesAligner is specified and\ndoes not equal ALIGN_NONE, then\nthis field must be defined;\notherwise an error is returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_series_reducer": {
                          "description": "The approach to be used to combine\ntime series. Not all reducer\nfunctions may be applied to all\ntime series, depending on the\nmetric type and the value type of\nthe original time series.\nReduction may change the metric\ntype of value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "group_by_fields": {
                          "description": "The set of fields to preserve when\ncrossSeriesReducer is specified.\nThe groupByFields determine how\nthe time series are partitioned\ninto subsets prior to applying the\naggregation function. Each subset\ncontains time series that have the\nsame value for each of the\ngrouping fields. Each individual\ntime series is a member of exactly\none subset. The crossSeriesReducer\nis applied to each subset of time\nseries. It is not possible to\nreduce across different resource\ntypes, so this field implicitly\ncontains resource.type. Fields not\nspecified in groupByFields are\naggregated away. If groupByFields\nis not specified and all the time\nseries have the same resource\ntype, then the time series are\naggregated into a single output\ntime series. If crossSeriesReducer\nis not defined, this field is\nignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "per_series_aligner": {
                          "description": "The approach to be used to align\nindividual time series. Not all\nalignment functions may be applied\nto all time series, depending on\nthe metric type and value type of\nthe original time series.\nAlignment may change the metric\ntype or the value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "trigger": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description": "The absolute number of time series\nthat must fail the predicate for the\ncondition to be triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "percent": {
                          "description": "The percentage of time series that\nmust fail the predicate for the\ncondition to be triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "condition_threshold": {
              "block": {
                "attributes": {
                  "comparison": {
                    "description": "The comparison to apply between the time\nseries (indicated by filter and aggregation)\nand the threshold (indicated by\nthreshold_value). The comparison is applied\non each time series, with the time series on\nthe left-hand side and the threshold on the\nright-hand side. Only COMPARISON_LT and\nCOMPARISON_GT are supported currently.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "denominator_filter": {
                    "description": "A filter that identifies a time series that\nshould be used as the denominator of a ratio\nthat will be compared with the threshold. If\na denominator_filter is specified, the time\nseries specified by the filter field will be\nused as the numerator.The filter is similar\nto the one that is specified in the\nMetricService.ListTimeSeries request (that\ncall is useful to verify the time series\nthat will be retrieved / processed) and must\nspecify the metric type and optionally may\ncontain restrictions on resource type,\nresource labels, and metric labels. This\nfield may not exceed 2048 Unicode characters\nin length.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "duration": {
                    "description": "The amount of time that a time series must\nviolate the threshold to be considered\nfailing. Currently, only values that are a\nmultiple of a minute--e.g., 0, 60, 120, or\n300 seconds--are supported. If an invalid\nvalue is given, an error will be returned.\nWhen choosing a duration, it is useful to\nkeep in mind the frequency of the underlying\ntime series data (which may also be affected\nby any alignments specified in the\naggregations field); a good duration is long\nenough so that a single outlier does not\ngenerate spurious alerts, but short enough\nthat unhealthy states are detected and\nalerted on quickly.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter": {
                    "description": "A filter that identifies which time series\nshould be compared with the threshold.The\nfilter is similar to the one that is\nspecified in the\nMetricService.ListTimeSeries request (that\ncall is useful to verify the time series\nthat will be retrieved / processed) and must\nspecify the metric type and optionally may\ncontain restrictions on resource type,\nresource labels, and metric labels. This\nfield may not exceed 2048 Unicode characters\nin length.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold_value": {
                    "description": "A value against which to compare the time\nseries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "aggregations": {
                    "block": {
                      "attributes": {
                        "alignment_period": {
                          "description": "The alignment period for per-time\nseries alignment. If present,\nalignmentPeriod must be at least\n60 seconds. After per-time series\nalignment, each time series will\ncontain data points only on the\nperiod boundaries. If\nperSeriesAligner is not specified\nor equals ALIGN_NONE, then this\nfield is ignored. If\nperSeriesAligner is specified and\ndoes not equal ALIGN_NONE, then\nthis field must be defined;\notherwise an error is returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_series_reducer": {
                          "description": "The approach to be used to combine\ntime series. Not all reducer\nfunctions may be applied to all\ntime series, depending on the\nmetric type and the value type of\nthe original time series.\nReduction may change the metric\ntype of value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "group_by_fields": {
                          "description": "The set of fields to preserve when\ncrossSeriesReducer is specified.\nThe groupByFields determine how\nthe time series are partitioned\ninto subsets prior to applying the\naggregation function. Each subset\ncontains time series that have the\nsame value for each of the\ngrouping fields. Each individual\ntime series is a member of exactly\none subset. The crossSeriesReducer\nis applied to each subset of time\nseries. It is not possible to\nreduce across different resource\ntypes, so this field implicitly\ncontains resource.type. Fields not\nspecified in groupByFields are\naggregated away. If groupByFields\nis not specified and all the time\nseries have the same resource\ntype, then the time series are\naggregated into a single output\ntime series. If crossSeriesReducer\nis not defined, this field is\nignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "per_series_aligner": {
                          "description": "The approach to be used to align\nindividual time series. Not all\nalignment functions may be applied\nto all time series, depending on\nthe metric type and value type of\nthe original time series.\nAlignment may change the metric\ntype or the value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "denominator_aggregations": {
                    "block": {
                      "attributes": {
                        "alignment_period": {
                          "description": "The alignment period for per-time\nseries alignment. If present,\nalignmentPeriod must be at least\n60 seconds. After per-time series\nalignment, each time series will\ncontain data points only on the\nperiod boundaries. If\nperSeriesAligner is not specified\nor equals ALIGN_NONE, then this\nfield is ignored. If\nperSeriesAligner is specified and\ndoes not equal ALIGN_NONE, then\nthis field must be defined;\notherwise an error is returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_series_reducer": {
                          "description": "The approach to be used to combine\ntime series. Not all reducer\nfunctions may be applied to all\ntime series, depending on the\nmetric type and the value type of\nthe original time series.\nReduction may change the metric\ntype of value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "group_by_fields": {
                          "description": "The set of fields to preserve when\ncrossSeriesReducer is specified.\nThe groupByFields determine how\nthe time series are partitioned\ninto subsets prior to applying the\naggregation function. Each subset\ncontains time series that have the\nsame value for each of the\ngrouping fields. Each individual\ntime series is a member of exactly\none subset. The crossSeriesReducer\nis applied to each subset of time\nseries. It is not possible to\nreduce across different resource\ntypes, so this field implicitly\ncontains resource.type. Fields not\nspecified in groupByFields are\naggregated away. If groupByFields\nis not specified and all the time\nseries have the same resource\ntype, then the time series are\naggregated into a single output\ntime series. If crossSeriesReducer\nis not defined, this field is\nignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "per_series_aligner": {
                          "description": "The approach to be used to align\nindividual time series. Not all\nalignment functions may be applied\nto all time series, depending on\nthe metric type and value type of\nthe original time series.\nAlignment may change the metric\ntype or the value type of the time\nseries.Time series data must be\naligned in order to perform cross-\ntime series reduction. If\ncrossSeriesReducer is specified,\nthen perSeriesAligner must be\nspecified and not equal ALIGN_NONE\nand alignmentPeriod must be\nspecified; otherwise, an error is\nreturned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "trigger": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description": "The absolute number of time series\nthat must fail the predicate for the\ncondition to be triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "percent": {
                          "description": "The percentage of time series that\nmust fail the predicate for the\ncondition to be triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "documentation": {
        "block": {
          "attributes": {
            "content": {
              "description": "The text of the documentation, interpreted according to mimeType.\nThe content may not exceed 8,192 Unicode characters and may not\nexceed more than 10,240 bytes when encoded in UTF-8 format,\nwhichever is smaller.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mime_type": {
              "description": "The format of the content field. Presently, only the value\n\"text/markdown\" is supported.",
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

func GoogleMonitoringAlertPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringAlertPolicy), &result)
	return &result
}
