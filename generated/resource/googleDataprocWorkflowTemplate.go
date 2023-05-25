package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocWorkflowTemplate = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time template was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dag_timeout": {
        "description": "Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes (\"600s\") to 24 hours (\"86400s\"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a [managed cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.",
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
      "labels": {
        "description": "Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. * For ` + "`" + `projects.regions.workflowTemplates` + "`" + `, the resource name of the template has the following format: ` + "`" + `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` + "`" + ` * For ` + "`" + `projects.locations.workflowTemplates` + "`" + `, the resource name of the template has the following format: ` + "`" + `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}` + "`" + `",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time template was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "deprecated": true,
        "description": "Output only. The current version of this workflow template.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "jobs": {
        "block": {
          "attributes": {
            "labels": {
              "description": "Optional. The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given job.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "prerequisite_step_ids": {
              "description": "Optional. The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "step_id": {
              "description": "Required. The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job ` + "`" + `goog-dataproc-workflow-step-id` + "`" + ` label, and in prerequisiteStepIds field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "hadoop_job": {
              "block": {
                "attributes": {
                  "archive_uris": {
                    "description": "Optional. HCFS URIs of archives to be extracted in the working directory of Hadoop drivers and tasks. Supported file types: .jar, .tar, .tar.gz, .tgz, or .zip.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "args": {
                    "description": "Optional. The arguments to pass to the driver. Do not include arguments, such as ` + "`" + `-libjars` + "`" + ` or ` + "`" + `-Dfoo=bar` + "`" + `, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "file_uris": {
                    "description": "Optional. HCFS (Hadoop Compatible Filesystem) URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "jar_file_uris": {
                    "description": "Optional. Jar file URIs to add to the CLASSPATHs of the Hadoop driver and tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "main_class": {
                    "description": "The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in ` + "`" + `jar_file_uris` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "main_jar_file_uri": {
                    "description": "The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site and classes in user code.",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a Hadoop job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hive_job": {
              "block": {
                "attributes": {
                  "continue_on_failure": {
                    "description": "Optional. Whether to continue executing queries if a query fails. The default value is ` + "`" + `false` + "`" + `. Setting to ` + "`" + `true` + "`" + ` can be useful when executing independent parallel queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "jar_file_uris": {
                    "description": "Optional. HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "query_file_uri": {
                    "description": "The HCFS URI of the script that contains Hive queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_variables": {
                    "description": "Optional. Mapping of query variable names to values (equivalent to the Hive command: ` + "`" + `SET name=\"value\";` + "`" + `).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "query_list": {
                    "block": {
                      "attributes": {
                        "queries": {
                          "description": "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "A list of queries.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a Hive job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pig_job": {
              "block": {
                "attributes": {
                  "continue_on_failure": {
                    "description": "Optional. Whether to continue executing queries if a query fails. The default value is ` + "`" + `false` + "`" + `. Setting to ` + "`" + `true` + "`" + ` can be useful when executing independent parallel queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "jar_file_uris": {
                    "description": "Optional. HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "query_file_uri": {
                    "description": "The HCFS URI of the script that contains the Pig queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_variables": {
                    "description": "Optional. Mapping of query variable names to values (equivalent to the Pig command: ` + "`" + `name=[value]` + "`" + `).",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "query_list": {
                    "block": {
                      "attributes": {
                        "queries": {
                          "description": "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "A list of queries.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a Pig job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "presto_job": {
              "block": {
                "attributes": {
                  "client_tags": {
                    "description": "Optional. Presto client tags to attach to this query",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "continue_on_failure": {
                    "description": "Optional. Whether to continue executing queries if a query fails. The default value is ` + "`" + `false` + "`" + `. Setting to ` + "`" + `true` + "`" + ` can be useful when executing independent parallel queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "output_format": {
                    "description": "Optional. The format in which query output will be displayed. See the Presto documentation for supported output formats",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values. Used to set Presto [session properties](https://prestodb.io/docs/current/sql/set-session.html) Equivalent to using the --session flag in the Presto CLI",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "query_file_uri": {
                    "description": "The HCFS URI of the script that contains SQL queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "logging_config": {
                    "block": {
                      "attributes": {
                        "driver_log_levels": {
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "query_list": {
                    "block": {
                      "attributes": {
                        "queries": {
                          "description": "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "A list of queries.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a Presto job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pyspark_job": {
              "block": {
                "attributes": {
                  "archive_uris": {
                    "description": "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "args": {
                    "description": "Optional. The arguments to pass to the driver. Do not include arguments, such as ` + "`" + `--conf` + "`" + `, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "file_uris": {
                    "description": "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "jar_file_uris": {
                    "description": "Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "main_python_file_uri": {
                    "description": "Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "python_file_uris": {
                    "description": "Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a PySpark job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "scheduling": {
              "block": {
                "attributes": {
                  "max_failures_per_hour": {
                    "description": "Optional. Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. A job may be reported as thrashing if driver exits with non-zero code 4 times within 10 minute window. Maximum value is 10.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_failures_total": {
                    "description": "Optional. Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. Maximum value is 240.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Optional. Job scheduling configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "spark_job": {
              "block": {
                "attributes": {
                  "archive_uris": {
                    "description": "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "args": {
                    "description": "Optional. The arguments to pass to the driver. Do not include arguments, such as ` + "`" + `--conf` + "`" + `, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "file_uris": {
                    "description": "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "jar_file_uris": {
                    "description": "Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "main_class": {
                    "description": "The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in ` + "`" + `jar_file_uris` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "main_jar_file_uri": {
                    "description": "The HCFS URI of the jar file that contains the main class.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a Spark job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "spark_r_job": {
              "block": {
                "attributes": {
                  "archive_uris": {
                    "description": "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "args": {
                    "description": "Optional. The arguments to pass to the driver. Do not include arguments, such as ` + "`" + `--conf` + "`" + `, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "file_uris": {
                    "description": "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "main_r_file_uri": {
                    "description": "Required. The HCFS URI of the main R file to use as the driver. Must be a .R file.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure SparkR. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a SparkR job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "spark_sql_job": {
              "block": {
                "attributes": {
                  "jar_file_uris": {
                    "description": "Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "properties": {
                    "description": "Optional. A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Dataproc API may be overwritten.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "query_file_uri": {
                    "description": "The HCFS URI of the script that contains SQL queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_variables": {
                    "description": "Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET ` + "`" + `name=\"value\";` + "`" + `).",
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
                          "description": "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The runtime log config for job execution.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "query_list": {
                    "block": {
                      "attributes": {
                        "queries": {
                          "description": "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "A list of queries.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Job is a SparkSql job.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. The Directed Acyclic Graph of Jobs to submit.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "parameters": {
        "block": {
          "attributes": {
            "description": {
              "description": "Optional. Brief description of the parameter. Must not exceed 1024 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fields": {
              "description": "Required. Paths to all fields that the parameter replaces. A field is allowed to appear in at most one parameter's list of field paths. A field path is similar in syntax to a google.protobuf.FieldMask. For example, a field path that references the zone field of a workflow template's cluster selector would be specified as ` + "`" + `placement.clusterSelector.zone` + "`" + `. Also, field paths can reference fields using the following syntax: * Values in maps can be referenced by key: * labels['key'] * placement.clusterSelector.clusterLabels['key'] * placement.managedCluster.labels['key'] * placement.clusterSelector.clusterLabels['key'] * jobs['step-id'].labels['key'] * Jobs in the jobs list can be referenced by step-id: * jobs['step-id'].hadoopJob.mainJarFileUri * jobs['step-id'].hiveJob.queryFileUri * jobs['step-id'].pySparkJob.mainPythonFileUri * jobs['step-id'].hadoopJob.jarFileUris[0] * jobs['step-id'].hadoopJob.archiveUris[0] * jobs['step-id'].hadoopJob.fileUris[0] * jobs['step-id'].pySparkJob.pythonFileUris[0] * Items in repeated fields can be referenced by a zero-based index: * jobs['step-id'].sparkJob.args[0] * Other examples: * jobs['step-id'].hadoopJob.properties['key'] * jobs['step-id'].hadoopJob.args[0] * jobs['step-id'].hiveJob.scriptVariables['key'] * jobs['step-id'].hadoopJob.mainJarFileUri * placement.clusterSelector.zone It may not be possible to parameterize maps and repeated fields in their entirety since only individual map values and individual items in repeated fields can be referenced. For example, the following field paths are invalid: - placement.clusterSelector.clusterLabels - jobs['step-id'].sparkJob.args",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "description": "Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "validation": {
              "block": {
                "block_types": {
                  "regex": {
                    "block": {
                      "attributes": {
                        "regexes": {
                          "description": "Required. RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Validation based on regular expressions.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "values": {
                    "block": {
                      "attributes": {
                        "values": {
                          "description": "Required. List of allowed values for the parameter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Validation based on a list of allowed values.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Validation rules to be applied to this parameter's value.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "placement": {
        "block": {
          "block_types": {
            "cluster_selector": {
              "block": {
                "attributes": {
                  "cluster_labels": {
                    "description": "Required. The cluster labels. Cluster must have all labels to match.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "Optional. The zone where workflow process executes. This parameter does not affect the selection of the cluster. If unspecified, the zone of the first cluster matching the selector is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Optional. A selector that chooses target cluster for jobs based on metadata. The selector is evaluated at the time each job is submitted.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "managed_cluster": {
              "block": {
                "attributes": {
                  "cluster_name": {
                    "description": "Required. The cluster name prefix. A unique cluster name will be formed by appending a random suffix. The name must contain only lower-case letters (a-z), numbers (0-9), and hyphens (-). Must begin with a letter. Cannot begin or end with hyphen. Must consist of between 2 and 35 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "labels": {
                    "description": "Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "config": {
                    "block": {
                      "attributes": {
                        "staging_bucket": {
                          "description": "Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "temp_bucket": {
                          "description": "Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "autoscaling_config": {
                          "block": {
                            "attributes": {
                              "policy": {
                                "description": "Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` + "`" + ` * ` + "`" + `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` + "`" + ` Note that the policy must be in the same project and Dataproc region.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "encryption_config": {
                          "block": {
                            "attributes": {
                              "gce_pd_kms_key_name": {
                                "description": "Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Optional. Encryption settings for the cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "endpoint_config": {
                          "block": {
                            "attributes": {
                              "enable_http_port_access": {
                                "description": "Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "http_ports": {
                                "computed": true,
                                "description": "Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.",
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "Optional. Port/endpoint configuration for this cluster",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "gce_cluster_config": {
                          "block": {
                            "attributes": {
                              "internal_ip_only": {
                                "computed": true,
                                "description": "Optional. If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This ` + "`" + `internal_ip_only` + "`" + ` restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "metadata": {
                                "description": "The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "network": {
                                "description": "Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither ` + "`" + `network_uri` + "`" + ` nor ` + "`" + `subnetwork_uri` + "`" + ` is specified, the \"default\" network of the project is used, if it exists. Cannot be a \"Custom Subnet Network\" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` + "`" + ` * ` + "`" + `projects/[project_id]/regions/global/default` + "`" + ` * ` + "`" + `default` + "`" + `",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "private_ipv6_google_access": {
                                "description": "Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "service_account": {
                                "description": "Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "service_account_scopes": {
                                "description": "Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnetwork": {
                                "description": "Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` + "`" + ` * ` + "`" + `projects/[project_id]/regions/us-east1/subnetworks/sub0` + "`" + ` * ` + "`" + `sub0` + "`" + `",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "tags": {
                                "description": "The Compute Engine tags to add to all instances (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "zone": {
                                "computed": true,
                                "description": "Optional. The zone where the Compute Engine cluster will be located. On a create request, it is required in the \"global\" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` + "`" + ` * ` + "`" + `projects/[project_id]/zones/[zone]` + "`" + ` * ` + "`" + `us-central1-f` + "`" + `",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "node_group_affinity": {
                                "block": {
                                  "attributes": {
                                    "node_group": {
                                      "description": "Required. The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` + "`" + ` * ` + "`" + `node-group-1` + "`" + `",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Optional. Node Group Affinity for sole-tenant clusters.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "reservation_affinity": {
                                "block": {
                                  "attributes": {
                                    "consume_reservation_type": {
                                      "description": "Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "description": "Optional. Corresponds to the label key of reservation resource.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "values": {
                                      "description": "Optional. Corresponds to the label values of reservation resource.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Optional. Reservation Affinity for consuming Zonal reservation.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Optional. The shared Compute Engine config settings for all instances in a cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "initialization_actions": {
                          "block": {
                            "attributes": {
                              "executable_file": {
                                "description": "Required. Cloud Storage URI of executable file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "execution_timeout": {
                                "description": "Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's ` + "`" + `role` + "`" + ` metadata to run an executable on a master or worker node, as shown below using ` + "`" + `curl` + "`" + ` (you can also use ` + "`" + `wget` + "`" + `): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if [[ \"${ROLE}\" == 'Master' ]]; then ... master specific actions ... else ... worker specific actions ... fi",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "lifecycle_config": {
                          "block": {
                            "attributes": {
                              "auto_delete_time": {
                                "description": "Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "auto_delete_ttl": {
                                "description": "Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "idle_delete_ttl": {
                                "description": "Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "idle_start_time": {
                                "computed": true,
                                "description": "Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "description": "Optional. Lifecycle setting for the cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "master_config": {
                          "block": {
                            "attributes": {
                              "image": {
                                "description": "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `image-id` + "`" + ` Image family examples. Dataproc will use the most recent image from the family: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` If the URI is unspecified, it will be inferred from ` + "`" + `SoftwareConfig.image_version` + "`" + ` or the system default.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "instance_names": {
                                "computed": true,
                                "description": "Output only. The list of instance names. Dataproc derives the names from ` + "`" + `cluster_name` + "`" + `, ` + "`" + `num_instances` + "`" + `, and the instance group.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "is_preemptible": {
                                "computed": true,
                                "description": "Output only. Specifies that this instance group contains preemptible instances.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "machine_type": {
                                "description": "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `n1-standard-2` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, ` + "`" + `n1-standard-2` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "managed_group_config": {
                                "computed": true,
                                "description": "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "instance_group_manager_name": "string",
                                      "instance_template_name": "string"
                                    }
                                  ]
                                ]
                              },
                              "min_cpu_platform": {
                                "computed": true,
                                "description": "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -\u003e Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_instances": {
                                "description": "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "preemptibility": {
                                "description": "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is ` + "`" + `NON_PREEMPTIBLE` + "`" + `. This default cannot be changed. The default value for secondary instances is ` + "`" + `PREEMPTIBLE` + "`" + `. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "accelerators": {
                                "block": {
                                  "attributes": {
                                    "accelerator_count": {
                                      "description": "The number of the accelerator cards of this type exposed to this instance.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "accelerator_type": {
                                      "description": "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `nvidia-tesla-k80` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, ` + "`" + `nvidia-tesla-k80` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Optional. The Compute Engine accelerator configuration for these instances.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "disk_config": {
                                "block": {
                                  "attributes": {
                                    "boot_disk_size_gb": {
                                      "description": "Optional. Size in GB of the boot disk (default is 500GB).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "boot_disk_type": {
                                      "description": "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "num_local_ssds": {
                                      "computed": true,
                                      "description": "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Optional. Disk option config settings.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Optional. The Compute Engine config settings for worker instances in a cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "secondary_worker_config": {
                          "block": {
                            "attributes": {
                              "image": {
                                "description": "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `image-id` + "`" + ` Image family examples. Dataproc will use the most recent image from the family: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` If the URI is unspecified, it will be inferred from ` + "`" + `SoftwareConfig.image_version` + "`" + ` or the system default.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "instance_names": {
                                "computed": true,
                                "description": "Output only. The list of instance names. Dataproc derives the names from ` + "`" + `cluster_name` + "`" + `, ` + "`" + `num_instances` + "`" + `, and the instance group.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "is_preemptible": {
                                "computed": true,
                                "description": "Output only. Specifies that this instance group contains preemptible instances.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "machine_type": {
                                "description": "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `n1-standard-2` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, ` + "`" + `n1-standard-2` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "managed_group_config": {
                                "computed": true,
                                "description": "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "instance_group_manager_name": "string",
                                      "instance_template_name": "string"
                                    }
                                  ]
                                ]
                              },
                              "min_cpu_platform": {
                                "computed": true,
                                "description": "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -\u003e Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_instances": {
                                "description": "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "preemptibility": {
                                "description": "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is ` + "`" + `NON_PREEMPTIBLE` + "`" + `. This default cannot be changed. The default value for secondary instances is ` + "`" + `PREEMPTIBLE` + "`" + `. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "accelerators": {
                                "block": {
                                  "attributes": {
                                    "accelerator_count": {
                                      "description": "The number of the accelerator cards of this type exposed to this instance.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "accelerator_type": {
                                      "description": "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `nvidia-tesla-k80` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, ` + "`" + `nvidia-tesla-k80` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Optional. The Compute Engine accelerator configuration for these instances.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "disk_config": {
                                "block": {
                                  "attributes": {
                                    "boot_disk_size_gb": {
                                      "description": "Optional. Size in GB of the boot disk (default is 500GB).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "boot_disk_type": {
                                      "description": "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "num_local_ssds": {
                                      "computed": true,
                                      "description": "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Optional. Disk option config settings.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Optional. The Compute Engine config settings for worker instances in a cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "security_config": {
                          "block": {
                            "block_types": {
                              "kerberos_config": {
                                "block": {
                                  "attributes": {
                                    "cross_realm_trust_admin_server": {
                                      "description": "Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "cross_realm_trust_kdc": {
                                      "description": "Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "cross_realm_trust_realm": {
                                      "description": "Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "cross_realm_trust_shared_password": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "enable_kerberos": {
                                      "description": "Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "kdc_db_key": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key_password": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "keystore": {
                                      "description": "Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "keystore_password": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "kms_key": {
                                      "description": "Optional. The uri of the KMS key used to encrypt various sensitive files.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "realm": {
                                      "description": "Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "root_principal_password": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "tgt_lifetime_hours": {
                                      "description": "Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "truststore": {
                                      "description": "Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "truststore_password": {
                                      "description": "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Optional. Kerberos related configuration.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Optional. Security settings for the cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "software_config": {
                          "block": {
                            "attributes": {
                              "image_version": {
                                "description": "Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions), such as \"1.2\" (including a subminor version, such as \"1.2.29\"), or the [\"preview\" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "optional_components": {
                                "description": "Optional. The set of components to activate on the cluster.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "properties": {
                                "description": "Optional. The properties to set on daemon config files. Property keys are specified in ` + "`" + `prefix:property` + "`" + ` format, for example ` + "`" + `core:hadoop.tmp.dir` + "`" + `. The following are supported prefixes and their mappings: * capacity-scheduler: ` + "`" + `capacity-scheduler.xml` + "`" + ` * core: ` + "`" + `core-site.xml` + "`" + ` * distcp: ` + "`" + `distcp-default.xml` + "`" + ` * hdfs: ` + "`" + `hdfs-site.xml` + "`" + ` * hive: ` + "`" + `hive-site.xml` + "`" + ` * mapred: ` + "`" + `mapred-site.xml` + "`" + ` * pig: ` + "`" + `pig.properties` + "`" + ` * spark: ` + "`" + `spark-defaults.conf` + "`" + ` * yarn: ` + "`" + `yarn-site.xml` + "`" + ` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "Optional. The config settings for software inside the cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "worker_config": {
                          "block": {
                            "attributes": {
                              "image": {
                                "description": "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/[image-id]` + "`" + ` * ` + "`" + `image-id` + "`" + ` Image family examples. Dataproc will use the most recent image from the family: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` * ` + "`" + `projects/[project_id]/global/images/family/[custom-image-family-name]` + "`" + ` If the URI is unspecified, it will be inferred from ` + "`" + `SoftwareConfig.image_version` + "`" + ` or the system default.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "instance_names": {
                                "computed": true,
                                "description": "Output only. The list of instance names. Dataproc derives the names from ` + "`" + `cluster_name` + "`" + `, ` + "`" + `num_instances` + "`" + `, and the instance group.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "is_preemptible": {
                                "computed": true,
                                "description": "Output only. Specifies that this instance group contains preemptible instances.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "machine_type": {
                                "description": "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * ` + "`" + `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` + "`" + ` * ` + "`" + `n1-standard-2` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, ` + "`" + `n1-standard-2` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "managed_group_config": {
                                "computed": true,
                                "description": "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "instance_group_manager_name": "string",
                                      "instance_template_name": "string"
                                    }
                                  ]
                                ]
                              },
                              "min_cpu_platform": {
                                "computed": true,
                                "description": "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -\u003e Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_instances": {
                                "description": "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "preemptibility": {
                                "description": "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is ` + "`" + `NON_PREEMPTIBLE` + "`" + `. This default cannot be changed. The default value for secondary instances is ` + "`" + `PREEMPTIBLE` + "`" + `. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "accelerators": {
                                "block": {
                                  "attributes": {
                                    "accelerator_count": {
                                      "description": "The number of the accelerator cards of this type exposed to this instance.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "accelerator_type": {
                                      "description": "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * ` + "`" + `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` + "`" + ` * ` + "`" + `nvidia-tesla-k80` + "`" + ` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, ` + "`" + `nvidia-tesla-k80` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Optional. The Compute Engine accelerator configuration for these instances.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "disk_config": {
                                "block": {
                                  "attributes": {
                                    "boot_disk_size_gb": {
                                      "description": "Optional. Size in GB of the boot disk (default is 500GB).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "boot_disk_type": {
                                      "description": "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "num_local_ssds": {
                                      "computed": true,
                                      "description": "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Optional. Disk option config settings.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Optional. The Compute Engine config settings for worker instances in a cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Required. The cluster configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A cluster that is managed by the workflow.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. WorkflowTemplate scheduling information.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleDataprocWorkflowTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocWorkflowTemplate), &result)
	return &result
}
