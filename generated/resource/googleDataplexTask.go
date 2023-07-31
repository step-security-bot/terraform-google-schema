package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexTask = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the task was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description of the task.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User friendly display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_status": {
        "computed": true,
        "description": "Configuration for the cluster",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "latest_job": [
                "list",
                [
                  "object",
                  {
                    "end_time": "string",
                    "message": "string",
                    "name": "string",
                    "retry_count": "number",
                    "service": "string",
                    "service_job": "string",
                    "start_time": "string",
                    "state": "string",
                    "uid": "string"
                  }
                ]
              ],
              "update_time": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the task.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "lake": {
        "description": "The lake in which the task will be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location in which the task will be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The relative resource name of the task, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/ tasks/{name}.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current state of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_id": {
        "description": "The task Id of the task.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with the same name.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the task was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "execution_spec": {
        "block": {
          "attributes": {
            "args": {
              "description": "The arguments to pass to the task. The args can use placeholders of the format ${placeholder} as part of key/value string. These will be interpolated before passing the args to the driver. Currently supported placeholders: - ${taskId} - ${job_time} To pass positional args, set the key as TASK_ARGS. The value should be a comma-separated string of all the positional arguments. To use a delimiter other than comma, refer to https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of other keys being present in the args, then TASK_ARGS will be passed as the last argument. An object containing a list of 'key': value pairs. Example: { 'name': 'wrench', 'mass': '1.3kg', 'count': '3' }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "kms_key": {
              "description": "The Cloud KMS key to use for encryption, of the form: projects/{project_number}/locations/{locationId}/keyRings/{key-ring-name}/cryptoKeys/{key-name}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_job_execution_lifetime": {
              "description": "The maximum duration after which the job execution is expired. A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description": "The project in which jobs are run. By default, the project containing the Lake is used. If a project is provided, the ExecutionSpec.service_account must belong to this project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account": {
              "description": "Service account to use to execute a task. If not provided, the default Compute service account for the project is used.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for the cluster",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "notebook": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "Cloud Storage URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "Cloud Storage URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "notebook": {
              "description": "Path to input notebook. This can be the Cloud Storage URI of the notebook file or the path to a Notebook Content. The execution args are accessible as environment variables (TASK_key=value).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "infrastructure_spec": {
              "block": {
                "block_types": {
                  "batch": {
                    "block": {
                      "attributes": {
                        "executors_count": {
                          "description": "Total number of job executors. Executor Count should be between 2 and 100. [Default=2]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_executors_count": {
                          "description": "Max configurable executors. If maxExecutorsCount \u003e executorsCount, then auto-scaling is enabled. Max Executor Count should be between 2 and 1000. [Default=1000]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Compute resources needed for a Task when using Dataproc Serverless.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "container_image": {
                    "block": {
                      "attributes": {
                        "image": {
                          "description": "Container image to use.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "java_jars": {
                          "description": "A list of Java JARS to add to the classpath. Valid input includes Cloud Storage URIs to Jar binaries. For example, gs://bucket-name/my/path/to/file.jar",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "properties": {
                          "description": "Override to common configuration of open source components installed on the Dataproc cluster. The properties to set on daemon config files. Property keys are specified in prefix:property format, for example core:hadoop.tmp.dir. For more information, see Cluster properties.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "python_packages": {
                          "description": "A list of python packages to be installed. Valid formats include Cloud Storage URI to a PIP installable library. For example, gs://bucket-name/my/path/to/lib.tar.gz",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Container Image Runtime Configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "vpc_network": {
                    "block": {
                      "attributes": {
                        "network": {
                          "description": "The Cloud VPC network in which the job is run. By default, the Cloud VPC network named Default within the project is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "network_tags": {
                          "description": "List of network tags to apply to the job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "sub_network": {
                          "description": "The Cloud VPC sub-network in which the job is run.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Vpc network.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Infrastructure specification for the execution.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "Cloud Storage URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "Cloud Storage URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in jar_file_uris. The execution args are passed in as a sequence of named process arguments (--key=value).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_jar_file_uri": {
              "description": "The Cloud Storage URI of the jar file that contains the main class. The execution args are passed in as a sequence of named process arguments (--key=value).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_script_file": {
              "description": "The Gcloud Storage URI of the main Python file to use as the driver. Must be a .py file. The execution args are passed in as a sequence of named process arguments (--key=value).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_script": {
              "description": "The query text. The execution args are used to declare a set of script variables (set key='value';).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_script_file": {
              "description": "A reference to a query file. This can be the Cloud Storage URI of the query file or it can the path to a SqlScript Content. The execution args are used to declare a set of script variables (set key='value';).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "infrastructure_spec": {
              "block": {
                "block_types": {
                  "batch": {
                    "block": {
                      "attributes": {
                        "executors_count": {
                          "description": "Total number of job executors. Executor Count should be between 2 and 100. [Default=2]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_executors_count": {
                          "description": "Max configurable executors. If maxExecutorsCount \u003e executorsCount, then auto-scaling is enabled. Max Executor Count should be between 2 and 1000. [Default=1000]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Compute resources needed for a Task when using Dataproc Serverless.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "container_image": {
                    "block": {
                      "attributes": {
                        "image": {
                          "description": "Container image to use.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "java_jars": {
                          "description": "A list of Java JARS to add to the classpath. Valid input includes Cloud Storage URIs to Jar binaries. For example, gs://bucket-name/my/path/to/file.jar",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "properties": {
                          "description": "Override to common configuration of open source components installed on the Dataproc cluster. The properties to set on daemon config files. Property keys are specified in prefix:property format, for example core:hadoop.tmp.dir. For more information, see Cluster properties.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "python_packages": {
                          "description": "A list of python packages to be installed. Valid formats include Cloud Storage URI to a PIP installable library. For example, gs://bucket-name/my/path/to/lib.tar.gz",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Container Image Runtime Configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "vpc_network": {
                    "block": {
                      "attributes": {
                        "network": {
                          "description": "The Cloud VPC network in which the job is run. By default, the Cloud VPC network named Default within the project is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "network_tags": {
                          "description": "List of network tags to apply to the job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "sub_network": {
                          "description": "The Cloud VPC sub-network in which the job is run.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Vpc network.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Infrastructure specification for the execution.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.",
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
      "trigger_spec": {
        "block": {
          "attributes": {
            "disabled": {
              "description": "Prevent the task from executing. This does not cancel already running tasks. It is intended to temporarily disable RECURRING tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_retries": {
              "description": "Number of retry attempts before aborting. Set to zero to never attempt to retry a failed task.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schedule": {
              "description": "Cron schedule (https://en.wikipedia.org/wiki/Cron) for running tasks periodically. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: 'CRON_TZ=${IANA_TIME_ZONE}' or 'TZ=${IANA_TIME_ZONE}'. The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, CRON_TZ=America/New_York 1 * * * *, or TZ=America/New_York 1 * * * *. This field is required for RECURRING tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "description": "The first run of the task will be after this time. If not specified, the task will run shortly after being submitted if ON_DEMAND and based on the schedule if RECURRING.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Trigger type of the user-specified Task Possible values: [\"ON_DEMAND\", \"RECURRING\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for the cluster",
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

func GoogleDataplexTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexTask), &result)
	return &result
}
