package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageTransferJob = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "When the Transfer Job was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_time": {
        "computed": true,
        "description": "When the Transfer Job was deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Unique description to identify the Transfer Job.",
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
      "last_modification_time": {
        "computed": true,
        "description": "When the Transfer Job was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Transfer Job.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description": "Status of the job. Default: ENABLED. NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "notification_config": {
        "block": {
          "attributes": {
            "event_types": {
              "description": "Event types for which a notification is desired. If empty, send notifications for all event types. The valid types are \"TRANSFER_OPERATION_SUCCESS\", \"TRANSFER_OPERATION_FAILED\", \"TRANSFER_OPERATION_ABORTED\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "payload_format": {
              "description": "The desired format of the notification message payloads. One of \"NONE\" or \"JSON\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "pubsub_topic": {
              "description": "The Topic.name of the Pub/Sub topic to which to publish notifications.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Notification configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schedule": {
        "block": {
          "attributes": {
            "repeat_interval": {
              "description": "Interval between the start of each scheduled transfer. If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "schedule_end_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "Month of year. Must be from 1 to 12.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "Year of date. Must be from 1 to 9999.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The last day the recurring transfer will be run. If schedule_end_date is the same as schedule_start_date, the transfer will be executed only once.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule_start_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "Month of year. Must be from 1 to 12.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "Year of date. Must be from 1 to 9999.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The first day the recurring transfer is scheduled to run. If schedule_start_date is in the past, the transfer will run for the first time on the following day.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "start_time_of_day": {
              "block": {
                "attributes": {
                  "hours": {
                    "description": "Hours of day in 24 hour format. Should be from 0 to 23.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "Minutes of hour of day. Must be from 0 to 59.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "nanos": {
                    "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Seconds of minutes of the time. Must normally be from 0 to 59.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "transfer_spec": {
        "block": {
          "attributes": {
            "sink_agent_pool_name": {
              "computed": true,
              "description": "Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_agent_pool_name": {
              "computed": true,
              "description": "Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "aws_s3_data_source": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "S3 Bucket name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
                    "description": "S3 Bucket path in bucket to transfer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description": "The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'. For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "aws_access_key": {
                    "block": {
                      "attributes": {
                        "access_key_id": {
                          "description": "AWS Key ID.",
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "secret_access_key": {
                          "description": "AWS Secret Access Key.",
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description": "AWS credentials block.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "An AWS S3 data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_blob_storage_data_source": {
              "block": {
                "attributes": {
                  "container": {
                    "description": "The container to transfer from the Azure Storage account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_account": {
                    "description": "The name of the Azure Storage account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "azure_credentials": {
                    "block": {
                      "attributes": {
                        "sas_token": {
                          "description": "Azure shared access signature.",
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description": " Credentials used to authenticate API requests to Azure.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "An Azure Blob Storage data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcs_data_sink": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "Google Cloud Storage bucket name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "Google Cloud Storage path in bucket to transfer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A Google Cloud Storage data sink.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcs_data_source": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "Google Cloud Storage bucket name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "Google Cloud Storage path in bucket to transfer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A Google Cloud Storage data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_data_source": {
              "block": {
                "attributes": {
                  "list_url": {
                    "description": "The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A HTTP URL data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "object_conditions": {
              "block": {
                "attributes": {
                  "exclude_prefixes": {
                    "description": "exclude_prefixes must follow the requirements described for include_prefixes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_prefixes": {
                    "description": "If include_refixes is specified, objects that satisfy the object conditions must have names that start with one of the include_prefixes and that do not start with any of the exclude_prefixes. If include_prefixes is not specified, all objects except those that have names starting with one of the exclude_prefixes must satisfy the object conditions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "last_modified_before": {
                    "description": "If specified, only objects with a \"last modification time\" before this timestamp and objects that don't have a \"last modification time\" are transferred. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "last_modified_since": {
                    "description": "If specified, only objects with a \"last modification time\" on or after this timestamp and objects that don't have a \"last modification time\" are transferred. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_time_elapsed_since_last_modification": {
                    "description": "A duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_time_elapsed_since_last_modification": {
                    "description": "A duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' last_modification_time do not exclude objects in a data sink.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "posix_data_sink": {
              "block": {
                "attributes": {
                  "root_directory": {
                    "description": "Root directory path to the filesystem.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A POSIX filesystem data sink.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "posix_data_source": {
              "block": {
                "attributes": {
                  "root_directory": {
                    "description": "Root directory path to the filesystem.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A POSIX filesystem data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "transfer_options": {
              "block": {
                "attributes": {
                  "delete_objects_from_source_after_transfer": {
                    "description": "Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and delete_objects_unique_in_sink are mutually exclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "delete_objects_unique_in_sink": {
                    "description": "Whether objects that exist only in the sink should be deleted. Note that this option and delete_objects_from_source_after_transfer are mutually exclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "overwrite_objects_already_existing_in_sink": {
                    "description": "Whether overwriting objects that already exist in the sink is allowed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "overwrite_when": {
                    "description": "When to overwrite objects that already exist in the sink. If not set, overwrite behavior is determined by overwriteObjectsAlreadyExistingInSink.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Characteristics of how to treat files from datasource and sink during job. If the option delete_objects_unique_in_sink is true, object conditions based on objects' last_modification_time are ignored and do not exclude objects in a data source or a data sink.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Transfer specification.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageTransferJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageTransferJob), &result)
	return &result
}
