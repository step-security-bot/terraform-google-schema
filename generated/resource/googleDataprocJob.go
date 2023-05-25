package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocJob = `{
  "block": {
    "attributes": {
      "driver_controls_files_uri": {
        "computed": true,
        "description": "Output-only. If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.",
        "description_kind": "plain",
        "type": "string"
      },
      "driver_output_resource_uri": {
        "computed": true,
        "description": "Output-only. A URI pointing to the location of the stdout of the job's driver program",
        "description_kind": "plain",
        "type": "string"
      },
      "force_delete": {
        "description": "By default, you can only delete inactive jobs within Dataproc. Setting this to true, and calling destroy, will ensure that the job is first cancelled before issuing the delete.",
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
        "description": "Optional. The labels to associate with this job.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description": "The project in which the cluster can be found and jobs subsequently run against. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The Cloud Dataproc region. This essentially determines which clusters are available for this job to be submitted to. If not specified, defaults to global.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the job.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "details": "string",
              "state": "string",
              "state_start_time": "string",
              "substate": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "hadoop_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The class containing the main method of the driver. Must be in a provided jar or jar that is already on the classpath. Conflicts with main_jar_file_uri",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_jar_file_uri": {
              "description": "The HCFS URI of jar file containing the driver jar. Conflicts with main_class",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "properties": {
              "description": "A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "logging_config": {
              "block": {
                "attributes": {
                  "driver_log_levels": {
                    "description": "Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The runtime logging config of the job",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config of Hadoop job",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hive_config": {
        "block": {
          "attributes": {
            "continue_on_failure": {
              "description": "Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "properties": {
              "description": "A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "query_file_uri": {
              "description": "HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_list": {
              "description": "The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "script_variables": {
              "description": "Mapping of query variable names to values (equivalent to the Hive command: SET name=\"value\";).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "The config of hive job",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pig_config": {
        "block": {
          "attributes": {
            "continue_on_failure": {
              "description": "Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "properties": {
              "description": "A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "query_file_uri": {
              "description": "HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_list": {
              "description": "The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "script_variables": {
              "description": "Mapping of query variable names to values (equivalent to the Pig command: name=[value]).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "logging_config": {
              "block": {
                "attributes": {
                  "driver_log_levels": {
                    "description": "Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The runtime logging config of the job",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config of pag job.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "placement": {
        "block": {
          "attributes": {
            "cluster_name": {
              "description": "The name of the cluster where the job will be submitted",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cluster_uuid": {
              "computed": true,
              "description": "Output-only. A cluster UUID generated by the Cloud Dataproc service when the job is submitted",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The config of job placement.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "pyspark_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "Optional. HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "Optional. The arguments to pass to the driver. Do not include arguments, such as --conf, that can be set as job properties, since a collision may occur that causes an incorrect job submission",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "Optional. HCFS URIs of files to be copied to the working directory of Python drivers and distributed tasks. Useful for naively parallel tasks",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_python_file_uri": {
              "description": "Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "properties": {
              "description": "Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "python_file_uris": {
              "description": "Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "logging_config": {
              "block": {
                "attributes": {
                  "driver_log_levels": {
                    "description": "Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The runtime logging config of the job",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config of pySpark job.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reference": {
        "block": {
          "attributes": {
            "job_id": {
              "computed": true,
              "description": "The job ID, which must be unique within the project. The job ID is generated by the server upon job submission or provided by the user as a means to perform retries without creating duplicate jobs",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The reference of the job",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scheduling": {
        "block": {
          "attributes": {
            "max_failures_per_hour": {
              "description": "Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "max_failures_total": {
              "description": "Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Optional. Job scheduling configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The class containing the main method of the driver. Must be in a provided jar or jar that is already on the classpath. Conflicts with main_jar_file_uri",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_jar_file_uri": {
              "description": "The HCFS URI of jar file containing the driver jar. Conflicts with main_class",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "properties": {
              "description": "A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "logging_config": {
              "block": {
                "attributes": {
                  "driver_log_levels": {
                    "description": "Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The runtime logging config of the job",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config of the Spark job.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sparksql_config": {
        "block": {
          "attributes": {
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to be added to the Spark CLASSPATH.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "properties": {
              "description": "A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "query_file_uri": {
              "description": "The HCFS URI of the script that contains SQL queries. Conflicts with query_list",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_list": {
              "description": "The list of SQL queries or statements to execute as part of the job. Conflicts with query_file_uri",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "script_variables": {
              "description": "Mapping of query variable names to values (equivalent to the Spark SQL command: SET name=\"value\";).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "logging_config": {
              "block": {
                "attributes": {
                  "driver_log_levels": {
                    "description": "Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The runtime logging config of the job",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config of SparkSql job",
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

func GoogleDataprocJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocJob), &result)
	return &result
}
