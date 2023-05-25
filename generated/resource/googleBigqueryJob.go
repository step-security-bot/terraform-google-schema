package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryJob = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_id": {
        "description": "The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_timeout_ms": {
        "description": "Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_type": {
        "computed": true,
        "description": "The type of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "The labels associated with this job. You can use these to organize and group your jobs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The geographic location of the job. The default value is US.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "error_result": [
                "list",
                [
                  "object",
                  {
                    "location": "string",
                    "message": "string",
                    "reason": "string"
                  }
                ]
              ],
              "errors": [
                "list",
                [
                  "object",
                  {
                    "location": "string",
                    "message": "string",
                    "reason": "string"
                  }
                ]
              ],
              "state": "string"
            }
          ]
        ]
      },
      "user_email": {
        "computed": true,
        "description": "Email address of the user who ran the job.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "copy": {
        "block": {
          "attributes": {
            "create_disposition": {
              "description": "Specifies whether the job is allowed to create new tables. The following values are supported:\nCREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.\nCREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.\nCreation, truncation and append actions occur as one atomic update upon job completion Default value: \"CREATE_IF_NEEDED\" Possible values: [\"CREATE_IF_NEEDED\", \"CREATE_NEVER\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "write_disposition": {
              "description": "Specifies the action that occurs if the destination table already exists. The following values are supported:\nWRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.\nWRITE_APPEND: If the table already exists, BigQuery appends the data to the table.\nWRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.\nEach action is atomic and only occurs if BigQuery is able to complete the job successfully.\nCreation, truncation and append actions occur as one atomic update upon job completion. Default value: \"WRITE_EMPTY\" Possible values: [\"WRITE_TRUNCATE\", \"WRITE_APPEND\", \"WRITE_EMPTY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination_encryption_configuration": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.\nThe BigQuery Service Account associated with your project requires access to this encryption key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_version": {
                    "computed": true,
                    "description": "Describes the Cloud KMS encryption key version used to protect destination BigQuery table.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Custom encryption configuration (e.g., Cloud KMS keys)",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "destination_table": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "computed": true,
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The destination table.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "source_tables": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "computed": true,
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Source tables to copy.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Copies a table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "extract": {
        "block": {
          "attributes": {
            "compression": {
              "description": "The compression type to use for exported files. Possible values include GZIP, DEFLATE, SNAPPY, and NONE.\nThe default value is NONE. DEFLATE and SNAPPY are only supported for Avro.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_format": {
              "computed": true,
              "description": "The exported file format. Possible values include CSV, NEWLINE_DELIMITED_JSON and AVRO for tables and SAVED_MODEL for models.\nThe default value for tables is CSV. Tables with nested or repeated fields cannot be exported as CSV.\nThe default value for models is SAVED_MODEL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_uris": {
              "description": "A list of fully-qualified Google Cloud Storage URIs where the extracted table should be written.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "field_delimiter": {
              "computed": true,
              "description": "When extracting data in CSV format, this defines the delimiter to use between fields in the exported data.\nDefault is ','",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "print_header": {
              "description": "Whether to print out a header row in the results. Default is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_avro_logical_types": {
              "description": "Whether to use logical types when extracting to AVRO format.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "source_model": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The ID of the dataset containing this model.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "model_id": {
                    "description": "The ID of the model.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the project containing this model.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A reference to the model being exported.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "source_table": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "computed": true,
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A reference to the table being exported.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configures an extract job.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "load": {
        "block": {
          "attributes": {
            "allow_jagged_rows": {
              "description": "Accept rows that are missing trailing optional columns. The missing values are treated as nulls.\nIf false, records with missing trailing columns are treated as bad records, and if there are too many bad records,\nan invalid error is returned in the job result. The default value is false. Only applicable to CSV, ignored for other formats.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_quoted_newlines": {
              "description": "Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file.\nThe default value is false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "autodetect": {
              "description": "Indicates if we should automatically infer the options and schema for CSV and JSON sources.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_disposition": {
              "description": "Specifies whether the job is allowed to create new tables. The following values are supported:\nCREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.\nCREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.\nCreation, truncation and append actions occur as one atomic update upon job completion Default value: \"CREATE_IF_NEEDED\" Possible values: [\"CREATE_IF_NEEDED\", \"CREATE_NEVER\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encoding": {
              "description": "The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.\nThe default value is UTF-8. BigQuery decodes the data after the raw, binary data\nhas been split using the values of the quote and fieldDelimiter properties.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_delimiter": {
              "computed": true,
              "description": "The separator for fields in a CSV file. The separator can be any ISO-8859-1 single-byte character.\nTo use a character in the range 128-255, you must encode the character as UTF8. BigQuery converts\nthe string to ISO-8859-1 encoding, and then uses the first byte of the encoded string to split the\ndata in its raw, binary state. BigQuery also supports the escape sequence \"\\t\" to specify a tab separator.\nThe default value is a comma (',').",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_unknown_values": {
              "description": "Indicates if BigQuery should allow extra values that are not represented in the table schema.\nIf true, the extra values are ignored. If false, records with extra columns are treated as bad records,\nand if there are too many bad records, an invalid error is returned in the job result.\nThe default value is false. The sourceFormat property determines what BigQuery treats as an extra value:\nCSV: Trailing columns\nJSON: Named values that don't match any column names",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_bad_records": {
              "description": "The maximum number of bad records that BigQuery can ignore when running the job. If the number of bad records exceeds this value,\nan invalid error is returned in the job result. The default value is 0, which requires that all records are valid.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "null_marker": {
              "description": "Specifies a string that represents a null value in a CSV file. For example, if you specify \"\\N\", BigQuery interprets \"\\N\" as a null value\nwhen loading a CSV file. The default value is the empty string. If you set this property to a custom value, BigQuery throws an error if an\nempty string is present for all data types except for STRING and BYTE. For STRING and BYTE columns, BigQuery interprets the empty string as\nan empty value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "projection_fields": {
              "description": "If sourceFormat is set to \"DATASTORE_BACKUP\", indicates which entity properties to load into BigQuery from a Cloud Datastore backup.\nProperty names are case sensitive and must be top-level properties. If no properties are specified, BigQuery loads all properties.\nIf any named property isn't found in the Cloud Datastore backup, an invalid error is returned in the job result.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "quote": {
              "computed": true,
              "description": "The value that is used to quote data sections in a CSV file. BigQuery converts the string to ISO-8859-1 encoding,\nand then uses the first byte of the encoded string to split the data in its raw, binary state.\nThe default value is a double-quote ('\"'). If your data does not contain quoted sections, set the property value to an empty string.\nIf your data contains quoted newline characters, you must also set the allowQuotedNewlines property to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema_update_options": {
              "description": "Allows the schema of the destination table to be updated as a side effect of the load job if a schema is autodetected or\nsupplied in the job configuration. Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;\nwhen writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table, specified by partition decorators.\nFor normal tables, WRITE_TRUNCATE will always overwrite the schema. One or more of the following values are specified:\nALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.\nALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "skip_leading_rows": {
              "description": "The number of rows at the top of a CSV file that BigQuery will skip when loading the data.\nThe default value is 0. This property is useful if you have header rows in the file that should be skipped.\nWhen autodetect is on, the behavior is the following:\nskipLeadingRows unspecified - Autodetect tries to detect headers in the first row. If they are not detected,\nthe row is read as data. Otherwise data is read starting from the second row.\nskipLeadingRows is 0 - Instructs autodetect that there are no headers and data should be read starting from the first row.\nskipLeadingRows = N \u003e 0 - Autodetect skips N-1 rows and tries to detect headers in row N. If headers are not detected,\nrow N is just skipped. Otherwise row N is used to extract column names for the detected schema.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "source_format": {
              "description": "The format of the data files. For CSV files, specify \"CSV\". For datastore backups, specify \"DATASTORE_BACKUP\".\nFor newline-delimited JSON, specify \"NEWLINE_DELIMITED_JSON\". For Avro, specify \"AVRO\". For parquet, specify \"PARQUET\".\nFor orc, specify \"ORC\". [Beta] For Bigtable, specify \"BIGTABLE\".\nThe default value is CSV.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_uris": {
              "description": "The fully-qualified URIs that point to your data in Google Cloud.\nFor Google Cloud Storage URIs: Each URI can contain one '*' wildcard character\nand it must come after the 'bucket' name. Size limits related to load jobs apply\nto external data sources. For Google Cloud Bigtable URIs: Exactly one URI can be\nspecified and it has be a fully specified and valid HTTPS URL for a Google Cloud Bigtable table.\nFor Google Cloud Datastore backups: Exactly one URI can be specified. Also, the '*' wildcard character is not allowed.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "write_disposition": {
              "description": "Specifies the action that occurs if the destination table already exists. The following values are supported:\nWRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.\nWRITE_APPEND: If the table already exists, BigQuery appends the data to the table.\nWRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.\nEach action is atomic and only occurs if BigQuery is able to complete the job successfully.\nCreation, truncation and append actions occur as one atomic update upon job completion. Default value: \"WRITE_EMPTY\" Possible values: [\"WRITE_TRUNCATE\", \"WRITE_APPEND\", \"WRITE_EMPTY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination_encryption_configuration": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.\nThe BigQuery Service Account associated with your project requires access to this encryption key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_version": {
                    "computed": true,
                    "description": "Describes the Cloud KMS encryption key version used to protect destination BigQuery table.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Custom encryption configuration (e.g., Cloud KMS keys)",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "destination_table": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "computed": true,
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The destination table to load the data into.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "time_partitioning": {
              "block": {
                "attributes": {
                  "expiration_ms": {
                    "description": "Number of milliseconds for which to keep the storage for a partition. A wrapper is used here because 0 is an invalid value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "field": {
                    "description": "If not set, the table is partitioned by pseudo column '_PARTITIONTIME'; if set, the table is partitioned by this field.\nThe field must be a top-level TIMESTAMP or DATE field. Its mode must be NULLABLE or REQUIRED.\nA wrapper is used here because an empty string is an invalid value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The only type supported is DAY, which will generate one partition per day. Providing an empty string used to cause an error,\nbut in OnePlatform the field will be treated as unset.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Time-based partitioning specification for the destination table.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configures a load job.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "query": {
        "block": {
          "attributes": {
            "allow_large_results": {
              "description": "If true and query uses legacy SQL dialect, allows the query to produce arbitrarily large result tables at a slight cost in performance.\nRequires destinationTable to be set. For standard SQL queries, this flag is ignored and large results are always allowed.\nHowever, you must still set destinationTable when result size exceeds the allowed maximum response size.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_disposition": {
              "description": "Specifies whether the job is allowed to create new tables. The following values are supported:\nCREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.\nCREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.\nCreation, truncation and append actions occur as one atomic update upon job completion Default value: \"CREATE_IF_NEEDED\" Possible values: [\"CREATE_IF_NEEDED\", \"CREATE_NEVER\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flatten_results": {
              "description": "If true and query uses legacy SQL dialect, flattens all nested and repeated fields in the query results.\nallowLargeResults must be true if this is set to false. For standard SQL queries, this flag is ignored and results are never flattened.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "maximum_billing_tier": {
              "description": "Limits the billing tier for this job. Queries that have resource usage beyond this tier will fail (without incurring a charge).\nIf unspecified, this will be set to your project default.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maximum_bytes_billed": {
              "description": "Limits the bytes billed for this job. Queries that will have bytes billed beyond this limit will fail (without incurring a charge).\nIf unspecified, this will be set to your project default.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_mode": {
              "description": "Standard SQL only. Set to POSITIONAL to use positional (?) query parameters or to NAMED to use named (@myparam) query parameters in this query.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description": "Specifies a priority for the query. Default value: \"INTERACTIVE\" Possible values: [\"INTERACTIVE\", \"BATCH\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query": {
              "description": "SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.\n*NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)\n('DELETE', 'UPDATE', 'MERGE', 'INSERT') must specify 'create_disposition = \"\"' and 'write_disposition = \"\"'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schema_update_options": {
              "description": "Allows the schema of the destination table to be updated as a side effect of the query job.\nSchema update options are supported in two cases: when writeDisposition is WRITE_APPEND;\nwhen writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table,\nspecified by partition decorators. For normal tables, WRITE_TRUNCATE will always overwrite the schema.\nOne or more of the following values are specified:\nALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.\nALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "use_legacy_sql": {
              "description": "Specifies whether to use BigQuery's legacy SQL dialect for this query. The default value is true.\nIf set to false, the query will use BigQuery's standard SQL.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_query_cache": {
              "description": "Whether to look for the result in the query cache. The query cache is a best-effort cache that will be flushed whenever\ntables in the query are modified. Moreover, the query cache is only available when a query does not have a destination table specified.\nThe default value is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "write_disposition": {
              "description": "Specifies the action that occurs if the destination table already exists. The following values are supported:\nWRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.\nWRITE_APPEND: If the table already exists, BigQuery appends the data to the table.\nWRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.\nEach action is atomic and only occurs if BigQuery is able to complete the job successfully.\nCreation, truncation and append actions occur as one atomic update upon job completion. Default value: \"WRITE_EMPTY\" Possible values: [\"WRITE_TRUNCATE\", \"WRITE_APPEND\", \"WRITE_EMPTY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_dataset": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The dataset. Can be specified '{{dataset_id}}' if 'project_id' is also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the default dataset to use for unqualified table names in the query. Note that this does not alter behavior of unqualified dataset names.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "destination_encryption_configuration": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.\nThe BigQuery Service Account associated with your project requires access to this encryption key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_version": {
                    "computed": true,
                    "description": "Describes the Cloud KMS encryption key version used to protect destination BigQuery table.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Custom encryption configuration (e.g., Cloud KMS keys)",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "destination_table": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "computed": true,
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set,\nor of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Describes the table where the query results should be stored.\nThis property must be set for large results that exceed the maximum response size.\nFor queries that produce anonymous (cached) results, this field will be populated by BigQuery.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "script_options": {
              "block": {
                "attributes": {
                  "key_result_statement": {
                    "description": "Determines which statement in the script represents the \"key result\",\nused to populate the schema and query results of the script job. Possible values: [\"LAST\", \"FIRST_SELECT\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statement_byte_budget": {
                    "description": "Limit on the number of bytes billed per statement. Exceeding this budget results in an error.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statement_timeout_ms": {
                    "description": "Timeout period for each statement in a script.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Options controlling the execution of scripts.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "user_defined_function_resources": {
              "block": {
                "attributes": {
                  "inline_code": {
                    "description": "An inline resource that contains code for a user-defined function (UDF).\nProviding a inline code resource is equivalent to providing a URI for a file containing the same code.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_uri": {
                    "description": "A code resource to load from a Google Cloud Storage URI (gs://bucket/path).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Describes user-defined function resources used in the query.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configures a query job.",
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

func GoogleBigqueryJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryJob), &result)
	return &result
}
