package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringSlo = `{
  "block": {
    "attributes": {
      "calendar_period": {
        "description": "A calendar period, semantically \"since the start of the current\n\u003ccalendarPeriod\u003e\". Possible values: [\"DAY\", \"WEEK\", \"FORTNIGHT\", \"MONTH\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Name used for UI elements listing this SLO.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "goal": {
        "description": "The fraction of service that must be good in order for this objective\nto be met. 0 \u003c goal \u003c= 0.999",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The full resource name for this service. The syntax is:\nprojects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rolling_period_days": {
        "description": "A rolling time period, semantically \"in the past X days\".\nMust be between 1 to 30 days, inclusive.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "service": {
        "description": "ID of the service to which this SLO belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "slo_id": {
        "computed": true,
        "description": "The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "basic_sli": {
        "block": {
          "attributes": {
            "location": {
              "description": "An optional set of locations to which this SLI is relevant.\nTelemetry from other locations will not be used to calculate\nperformance for this SLI. If omitted, this SLI applies to all\nlocations in which the Service has activity. For service types\nthat don't support breaking down by location, setting this\nfield will result in an error.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "method": {
              "description": "An optional set of RPCs to which this SLI is relevant.\nTelemetry from other methods will not be used to calculate\nperformance for this SLI. If omitted, this SLI applies to all\nthe Service's methods. For service types that don't support\nbreaking down by method, setting this field will result in an\nerror.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "version": {
              "description": "The set of API versions to which this SLI is relevant.\nTelemetry from other API versions will not be used to\ncalculate performance for this SLI. If omitted,\nthis SLI applies to all API versions. For service types\nthat don't support breaking down by version, setting this\nfield will result in an error.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "availability": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether an availability SLI is enabled or not. Must be set to true. Defaults to 'true'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Availability based SLI, dervied from count of requests made to this service that return successfully.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "latency": {
              "block": {
                "attributes": {
                  "threshold": {
                    "description": "A duration string, e.g. 10s.\nGood service is defined to be the count of requests made to\nthis service that return in no more than threshold.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Parameters for a latency threshold SLI.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Basic Service-Level Indicator (SLI) on a well-known service type.\nPerformance will be computed on the basis of pre-defined metrics.\n\nSLIs are used to measure and calculate the quality of the Service's\nperformance with respect to a single aspect of service quality.\n\nExactly one of the following must be set:\n'basic_sli', 'request_based_sli', 'windows_based_sli'",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "request_based_sli": {
        "block": {
          "block_types": {
            "distribution_cut": {
              "block": {
                "attributes": {
                  "distribution_filter": {
                    "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\naggregating values to quantify the good service provided.\n\nMust have ValueType = DISTRIBUTION and\nMetricKind = DELTA or MetricKind = CUMULATIVE.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max value for the range (inclusive). If not given,\nwill be set to \"infinity\", defining an open range\n\"\u003e= range.min\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "Min value for the range (inclusive). If not given,\nwill be set to \"-infinity\", defining an open range\n\"\u003c range.max\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Range of numerical values. The computed good_service\nwill be the count of values x in the Distribution such\nthat range.min \u003c= x \u003c= range.max. inclusive of min and\nmax. Open ranges can be defined by setting\njust one of min or max.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Used when good_service is defined by a count of values aggregated in a\nDistribution that fall into a good range. The total_service is the\ntotal count of all values aggregated in the Distribution.\nDefines a distribution TimeSeries filter and thresholds used for\nmeasuring good service and total service.\n\nExactly one of 'distribution_cut' or 'good_total_ratio' can be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "good_total_ratio": {
              "block": {
                "attributes": {
                  "bad_service_filter": {
                    "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying bad service provided, either demanded service that\nwas not provided or demanded service that was of inadequate\nquality.\n\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.\n\nExactly two of 'good_service_filter','bad_service_filter','total_service_filter'\nmust be set (good + bad = total is assumed).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "good_service_filter": {
                    "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying good service provided.\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.\n\nExactly two of 'good_service_filter','bad_service_filter','total_service_filter'\nmust be set (good + bad = total is assumed).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "total_service_filter": {
                    "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying total demanded service.\n\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.\n\nExactly two of 'good_service_filter','bad_service_filter','total_service_filter'\nmust be set (good + bad = total is assumed).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A means to compute a ratio of 'good_service' to 'total_service'.\nDefines computing this ratio with two TimeSeries [monitoring filters](https://cloud.google.com/monitoring/api/v3/filters)\nMust specify exactly two of good, bad, and total service filters.\nThe relationship good_service + bad_service = total_service\nwill be assumed.\n\nExactly one of 'distribution_cut' or 'good_total_ratio' can be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A request-based SLI defines a SLI for which atomic units of\nservice are counted directly.\n\nA SLI describes a good service.\nIt is used to measure and calculate the quality of the Service's\nperformance with respect to a single aspect of service quality.\nExactly one of the following must be set:\n'basic_sli', 'request_based_sli', 'windows_based_sli'",
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
      },
      "windows_based_sli": {
        "block": {
          "attributes": {
            "good_bad_metric_filter": {
              "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nwith ValueType = BOOL. The window is good if any true values\nappear in the window. One of 'good_bad_metric_filter',\n'good_total_ratio_threshold', 'metric_mean_in_range',\n'metric_sum_in_range' must be set for 'windows_based_sli'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "window_period": {
              "description": "Duration over which window quality is evaluated, given as a\nduration string \"{X}s\" representing X seconds. Must be an\ninteger fraction of a day and at least 60s.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "good_total_ratio_threshold": {
              "block": {
                "attributes": {
                  "threshold": {
                    "description": "If window performance \u003e= threshold, the window is counted\nas good.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "basic_sli_performance": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description": "An optional set of locations to which this SLI is relevant.\nTelemetry from other locations will not be used to calculate\nperformance for this SLI. If omitted, this SLI applies to all\nlocations in which the Service has activity. For service types\nthat don't support breaking down by location, setting this\nfield will result in an error.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "method": {
                          "description": "An optional set of RPCs to which this SLI is relevant.\nTelemetry from other methods will not be used to calculate\nperformance for this SLI. If omitted, this SLI applies to all\nthe Service's methods. For service types that don't support\nbreaking down by method, setting this field will result in an\nerror.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "version": {
                          "description": "The set of API versions to which this SLI is relevant.\nTelemetry from other API versions will not be used to\ncalculate performance for this SLI. If omitted,\nthis SLI applies to all API versions. For service types\nthat don't support breaking down by version, setting this\nfield will result in an error.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "availability": {
                          "block": {
                            "attributes": {
                              "enabled": {
                                "description": "Whether an availability SLI is enabled or not. Must be set to 'true. Defaults to 'true'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Availability based SLI, dervied from count of requests made to this service that return successfully.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "latency": {
                          "block": {
                            "attributes": {
                              "threshold": {
                                "description": "A duration string, e.g. 10s.\nGood service is defined to be the count of requests made to\nthis service that return in no more than threshold.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Parameters for a latency threshold SLI.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Basic SLI to evaluate to judge window quality.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "performance": {
                    "block": {
                      "block_types": {
                        "distribution_cut": {
                          "block": {
                            "attributes": {
                              "distribution_filter": {
                                "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\naggregating values to quantify the good service provided.\n\nMust have ValueType = DISTRIBUTION and\nMetricKind = DELTA or MetricKind = CUMULATIVE.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "range": {
                                "block": {
                                  "attributes": {
                                    "max": {
                                      "description": "max value for the range (inclusive). If not given,\nwill be set to \"infinity\", defining an open range\n\"\u003e= range.min\"",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "min": {
                                      "description": "Min value for the range (inclusive). If not given,\nwill be set to \"-infinity\", defining an open range\n\"\u003c range.max\"",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Range of numerical values. The computed good_service\nwill be the count of values x in the Distribution such\nthat range.min \u003c= x \u003c= range.max. inclusive of min and\nmax. Open ranges can be defined by setting\njust one of min or max.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Used when good_service is defined by a count of values aggregated in a\nDistribution that fall into a good range. The total_service is the\ntotal count of all values aggregated in the Distribution.\nDefines a distribution TimeSeries filter and thresholds used for\nmeasuring good service and total service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "good_total_ratio": {
                          "block": {
                            "attributes": {
                              "bad_service_filter": {
                                "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying bad service provided, either demanded service that\nwas not provided or demanded service that was of inadequate\nquality. Exactly two of\ngood, bad, or total service filter must be defined (where\ngood + bad = total is assumed)\n\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "good_service_filter": {
                                "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying good service provided. Exactly two of\ngood, bad, or total service filter must be defined (where\ngood + bad = total is assumed)\n\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "total_service_filter": {
                                "description": "A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nquantifying total demanded service. Exactly two of\ngood, bad, or total service filter must be defined (where\ngood + bad = total is assumed)\n\nMust have ValueType = DOUBLE or ValueType = INT64 and\nmust have MetricKind = DELTA or MetricKind = CUMULATIVE.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A means to compute a ratio of 'good_service' to 'total_service'.\nDefines computing this ratio with two TimeSeries [monitoring filters](https://cloud.google.com/monitoring/api/v3/filters)\nMust specify exactly two of good, bad, and total service filters.\nThe relationship good_service + bad_service = total_service\nwill be assumed.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Request-based SLI to evaluate to judge window quality.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Criterion that describes a window as good if its performance is\nhigh enough. One of 'good_bad_metric_filter',\n'good_total_ratio_threshold', 'metric_mean_in_range',\n'metric_sum_in_range' must be set for 'windows_based_sli'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metric_mean_in_range": {
              "block": {
                "attributes": {
                  "time_series": {
                    "description": "A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nspecifying the TimeSeries to use for evaluating window\nThe provided TimeSeries must have ValueType = INT64 or\nValueType = DOUBLE and MetricKind = GAUGE. Mean value 'X'\nshould satisfy 'range.min \u003c= X \u003c= range.max'\nunder good service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max value for the range (inclusive). If not given,\nwill be set to \"infinity\", defining an open range\n\"\u003e= range.min\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "Min value for the range (inclusive). If not given,\nwill be set to \"-infinity\", defining an open range\n\"\u003c range.max\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Range of numerical values. The computed good_service\nwill be the count of values x in the Distribution such\nthat range.min \u003c= x \u003c= range.max. inclusive of min and\nmax. Open ranges can be defined by setting\njust one of min or max. Mean value 'X' of 'time_series'\nvalues should satisfy 'range.min \u003c= X \u003c= range.max' for a\ngood service.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Criterion that describes a window as good if the metric's value\nis in a good range, *averaged* across returned streams.\nOne of 'good_bad_metric_filter',\n\n'good_total_ratio_threshold', 'metric_mean_in_range',\n'metric_sum_in_range' must be set for 'windows_based_sli'.\nAverage value X of 'time_series' should satisfy\n'range.min \u003c= X \u003c= range.max' for a good window.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metric_sum_in_range": {
              "block": {
                "attributes": {
                  "time_series": {
                    "description": "A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\nspecifying the TimeSeries to use for evaluating window\nquality. The provided TimeSeries must have\nValueType = INT64 or ValueType = DOUBLE and\nMetricKind = GAUGE.\n\nSummed value 'X' should satisfy\n'range.min \u003c= X \u003c= range.max' for a good window.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "range": {
                    "block": {
                      "attributes": {
                        "max": {
                          "description": "max value for the range (inclusive). If not given,\nwill be set to \"infinity\", defining an open range\n\"\u003e= range.min\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min": {
                          "description": "Min value for the range (inclusive). If not given,\nwill be set to \"-infinity\", defining an open range\n\"\u003c range.max\"",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Range of numerical values. The computed good_service\nwill be the count of values x in the Distribution such\nthat range.min \u003c= x \u003c= range.max. inclusive of min and\nmax. Open ranges can be defined by setting\njust one of min or max. Summed value 'X' should satisfy\n'range.min \u003c= X \u003c= range.max' for a good window.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Criterion that describes a window as good if the metric's value\nis in a good range, *summed* across returned streams.\nSummed value 'X' of 'time_series' should satisfy\n'range.min \u003c= X \u003c= range.max' for a good window.\n\nOne of 'good_bad_metric_filter',\n'good_total_ratio_threshold', 'metric_mean_in_range',\n'metric_sum_in_range' must be set for 'windows_based_sli'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A windows-based SLI defines the criteria for time windows.\ngood_service is defined based off the count of these time windows\nfor which the provided service was of good quality.\n\nA SLI describes a good service. It is used to measure and calculate\nthe quality of the Service's performance with respect to a single\naspect of service quality.\n\nExactly one of the following must be set:\n'basic_sli', 'request_based_sli', 'windows_based_sli'",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMonitoringSloSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringSlo), &result)
	return &result
}
