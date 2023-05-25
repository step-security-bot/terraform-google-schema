package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionJobTrigger = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of the job trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User set display name of the job trigger.",
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
      "last_run_time": {
        "computed": true,
        "description": "The timestamp of the last time this trigger executed.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the job trigger. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the trigger, either in the format 'projects/{{project}}'\nor 'projects/{{project}}/locations/{{location}}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description": "Whether the trigger is currently active. Default value: \"HEALTHY\" Possible values: [\"PAUSED\", \"HEALTHY\", \"CANCELLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "inspect_job": {
        "block": {
          "attributes": {
            "inspect_template_name": {
              "description": "The name of the template to run when this job is triggered.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "save_findings": {
                    "block": {
                      "block_types": {
                        "output_config": {
                          "block": {
                            "attributes": {
                              "output_schema": {
                                "description": "Schema used for writing the findings for Inspect jobs. This field is only used for\nInspect and must be unspecified for Risk jobs. Columns are derived from the Finding\nobject. If appending to an existing table, any columns from the predefined schema\nthat are missing will be added. No columns in the existing table will be deleted.\n\nIf unspecified, then all available columns will be used for a new table or an (existing)\ntable with no schema, and no changes will be made to an existing table that has a schema.\nOnly for use with external storage. Possible values: [\"BASIC_COLUMNS\", \"GCS_COLUMNS\", \"DATASTORE_COLUMNS\", \"BIG_QUERY_COLUMNS\", \"ALL_COLUMNS\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "table": {
                                "block": {
                                  "attributes": {
                                    "dataset_id": {
                                      "description": "Dataset ID of the table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "project_id": {
                                      "description": "The Google Cloud Platform project ID of the project containing the table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "table_id": {
                                      "description": "Name of the table. If is not set a new one will be generated for you with the following format:\n'dlp_googleapis_yyyy_mm_dd_[dlp_job_id]'. Pacific timezone will be used for generating the date details.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Information on the location of the target BigQuery Table.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Information on where to store output",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Schedule for triggered jobs",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A task to execute on the completion of a job.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "storage_config": {
              "block": {
                "block_types": {
                  "big_query_options": {
                    "block": {
                      "block_types": {
                        "table_reference": {
                          "block": {
                            "attributes": {
                              "dataset_id": {
                                "description": "The dataset ID of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "The Google Cloud Platform project ID of the project containing the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "table_id": {
                                "description": "The name of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Set of files to scan.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining BigQuery table and row identifiers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "cloud_storage_options": {
                    "block": {
                      "attributes": {
                        "bytes_limit_per_file": {
                          "description": "Max number of bytes to scan from a file. If a scanned file's size is bigger than this value\nthen the rest of the bytes are omitted.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "bytes_limit_per_file_percent": {
                          "description": "Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down.\nMust be between 0 and 100, inclusively. Both 0 and 100 means no limit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "file_types": {
                          "description": "List of file type groups to include in the scan. If empty, all files are scanned and available data\nformat processors are applied. In addition, the binary content of the selected files is always scanned as well.\nImages are scanned only as binary if the specified region does not support image inspection and no fileTypes were specified. Possible values: [\"BINARY_FILE\", \"TEXT_FILE\", \"IMAGE\", \"WORD\", \"PDF\", \"AVRO\", \"CSV\", \"TSV\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "files_limit_percent": {
                          "description": "Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down.\nMust be between 0 and 100, inclusively. Both 0 and 100 means no limit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "sample_method": {
                          "description": "How to sample bytes if not all bytes are scanned. Meaningful only when used in conjunction with bytesLimitPerFile.\nIf not specified, scanning would start from the top. Possible values: [\"TOP\", \"RANDOM_START\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "file_set": {
                          "block": {
                            "attributes": {
                              "url": {
                                "description": "The Cloud Storage url of the file(s) to scan, in the format 'gs://\u003cbucket\u003e/\u003cpath\u003e'. Trailing wildcard\nin the path is allowed.\n\nIf the url ends in a trailing slash, the bucket or directory represented by the url will be scanned\nnon-recursively (content in sub-directories will not be scanned). This means that 'gs://mybucket/' is\nequivalent to 'gs://mybucket/*', and 'gs://mybucket/directory/' is equivalent to 'gs://mybucket/directory/*'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "regex_file_set": {
                                "block": {
                                  "attributes": {
                                    "bucket_name": {
                                      "description": "The name of a Cloud Storage bucket.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "exclude_regex": {
                                      "description": "A list of regular expressions matching file paths to exclude. All files in the bucket that match at\nleast one of these regular expressions will be excluded from the scan.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "include_regex": {
                                      "description": "A list of regular expressions matching file paths to include. All files in the bucket\nthat match at least one of these regular expressions will be included in the set of files,\nexcept for those that also match an item in excludeRegex. Leaving this field empty will\nmatch all files by default (this is equivalent to including .* in the list)",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "The regex-filtered set of files to scan.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Set of files to scan.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining a file or a set of files within a Google Cloud Storage bucket.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "datastore_options": {
                    "block": {
                      "block_types": {
                        "kind": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The name of the Datastore kind.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A representation of a Datastore kind.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "partition_id": {
                          "block": {
                            "attributes": {
                              "namespace_id": {
                                "description": "If not empty, the ID of the namespace to which the entities belong.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "The ID of the project to which the entities belong.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Datastore partition ID. A partition ID identifies a grouping of entities. The grouping\nis always by project and namespace, however the namespace ID may be empty.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining a data set within Google Cloud Datastore.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "timespan_config": {
                    "block": {
                      "attributes": {
                        "enable_auto_population_of_timespan_config": {
                          "description": "When the job is started by a JobTrigger we will automatically figure out a valid startTime to avoid\nscanning files that have not been modified since the last time the JobTrigger executed. This will\nbe based on the time of the execution of the last run of the JobTrigger.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "end_time": {
                          "description": "Exclude files or rows newer than this value. If set to zero, no upper time limit is applied.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start_time": {
                          "description": "Exclude files or rows older than this value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "timestamp_field": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery.\n\nFor BigQuery: Required to filter out rows based on the given start and end times. If not specified and the table was\nmodified between the given start and end times, the entire table will be scanned. The valid data types of the timestamp\nfield are: INTEGER, DATE, TIMESTAMP, or DATETIME BigQuery column.\n\nFor Datastore. Valid data types of the timestamp field are: TIMESTAMP. Datastore entity will be scanned if the\ntimestamp property does not exist or its value is empty or invalid.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Information on where to inspect",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Information on where to inspect",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Information on where to inspect",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Controls what and how to inspect for findings.",
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
      "triggers": {
        "block": {
          "block_types": {
            "schedule": {
              "block": {
                "attributes": {
                  "recurrence_period_duration": {
                    "description": "With this option a job is started a regular periodic basis. For example: every day (86400 seconds).\n\nA scheduled start time will be skipped if the previous execution has not ended when its scheduled time occurs.\n\nThis value must be set to a time duration greater than or equal to 1 day and can be no longer than 60 days.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Schedule for triggered jobs",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "What event needs to occur for a new job to be started.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataLossPreventionJobTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionJobTrigger), &result)
	return &result
}
