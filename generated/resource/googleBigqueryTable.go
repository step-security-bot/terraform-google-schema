package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryTable = `{
  "block": {
    "attributes": {
      "clustering": {
        "description": "Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "creation_time": {
        "computed": true,
        "description": "The time when this table was created, in milliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description": "The dataset ID to create the table in. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "The field description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "A hash of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description": "The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "friendly_name": {
        "description": "A descriptive name for the table.",
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
        "description": "A mapping of labels to assign to the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time when this table was last modified, in milliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "computed": true,
        "description": "The geographic location where the table resides. This value is inherited from the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "num_bytes": {
        "computed": true,
        "description": "The geographic location where the table resides. This value is inherited from the dataset.",
        "description_kind": "plain",
        "type": "number"
      },
      "num_long_term_bytes": {
        "computed": true,
        "description": "The number of bytes in the table that are considered \"long-term storage\".",
        "description_kind": "plain",
        "type": "number"
      },
      "num_rows": {
        "computed": true,
        "description": "The number of rows of data in this table, excluding any data in the streaming buffer.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description": "A JSON schema for the table.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_id": {
        "description": "A unique ID for the resource. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Describes the table type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The self link or full name of a key which should be used to encrypt this table. Note that the default bigquery service account will need to have encrypt/decrypt permissions on this key - you may want to see the google_bigquery_default_service_account datasource and the google_kms_crypto_key_iam_binding resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key_version": {
              "computed": true,
              "description": "The self link or full name of the kms key version used to encrypt this table.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "external_data_configuration": {
        "block": {
          "attributes": {
            "autodetect": {
              "description": "Let BigQuery try to autodetect the schema and format of the table.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "compression": {
              "description": "The compression type of the data source. Valid values are \"NONE\" or \"GZIP\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_id": {
              "description": "The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3. The connectionId can have the form \"{{project}}.{{location}}.{{connection_id}}\" or \"projects/{{project}}/locations/{{location}}/connections/{{connection_id}}\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_unknown_values": {
              "description": "Indicates if BigQuery should allow extra values that are not represented in the table schema. If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_bad_records": {
              "description": "The maximum number of bad records that BigQuery can ignore when reading data.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metadata_cache_mode": {
              "description": "Metadata Cache Mode for the table. Set this to enable caching of metadata from external data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "object_metadata": {
              "description": "Object Metadata is used to create Object Tables. Object Tables contain a listing of objects (with their metadata) found at the sourceUris. If ObjectMetadata is set, sourceFormat should be omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reference_file_schema_uri": {
              "description": "When creating an external table, the user can provide a reference file with the table schema. This is enabled for the following formats: AVRO, PARQUET, ORC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema": {
              "computed": true,
              "description": "A JSON schema for the external table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_format": {
              "description": " Please see sourceFormat under ExternalDataConfiguration in Bigquery's public API documentation (https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration) for supported formats. To use \"GOOGLE_SHEETS\" the scopes must include \"googleapis.com/auth/drive.readonly\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_uris": {
              "description": "A list of the fully-qualified URIs that point to your data in Google Cloud.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "avro_options": {
              "block": {
                "attributes": {
                  "use_avro_logical_types": {
                    "description": "If sourceFormat is set to \"AVRO\", indicates whether to interpret logical types as the corresponding BigQuery data type (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Additional options if source_format is set to \"AVRO\"",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "csv_options": {
              "block": {
                "attributes": {
                  "allow_jagged_rows": {
                    "description": "Indicates if BigQuery should accept rows that are missing trailing optional columns.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "allow_quoted_newlines": {
                    "description": "Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encoding": {
                    "description": "The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "field_delimiter": {
                    "description": "The separator for fields in a CSV file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "quote": {
                    "description": "The value that is used to quote data sections in a CSV file. If your data does not contain quoted sections, set the property value to an empty string. If your data contains quoted newline characters, you must also set the allow_quoted_newlines property to true. The API-side default is \", specified in Terraform escaped as \\\". Due to limitations with Terraform default values, this value is required to be explicitly set.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description": "The number of rows at the top of a CSV file that BigQuery will skip when reading the data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Additional properties to set if source_format is set to \"CSV\".",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_sheets_options": {
              "block": {
                "attributes": {
                  "range": {
                    "description": "Range of a sheet to query from. Only used when non-empty. At least one of range or skip_leading_rows must be set. Typical format: \"sheet_name!top_left_cell_id:bottom_right_cell_id\" For example: \"sheet1!A1:B20\"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description": "The number of rows at the top of the sheet that BigQuery will skip when reading the data. At least one of range or skip_leading_rows must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Additional options if source_format is set to \"GOOGLE_SHEETS\".",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hive_partitioning_options": {
              "block": {
                "attributes": {
                  "mode": {
                    "description": "When set, what mode of hive partitioning to use when reading data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_partition_filter": {
                    "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_uri_prefix": {
                    "description": "When hive partition detection is requested, a common for all source uris must be required. The prefix must end immediately before the partition key encoding begins.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "When set, configures hive partitioning support. Not all storage formats support hive partitioning -- requesting hive partitioning on an unsupported format will lead to an error, as will providing an invalid specification.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "materialized_view": {
        "block": {
          "attributes": {
            "enable_refresh": {
              "description": "Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "query": {
              "description": "A query whose result is persisted.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "refresh_interval_ms": {
              "description": "Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "If specified, configures this table as a materialized view.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "range_partitioning": {
        "block": {
          "attributes": {
            "field": {
              "description": "The field used to determine how to create a range-based partition.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "range": {
              "block": {
                "attributes": {
                  "end": {
                    "description": "End of the range partitioning, exclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "interval": {
                    "description": "The width of each range within the partition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "start": {
                    "description": "Start of the range partitioning, inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Information required to partition based on ranges. Structure is documented below.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "If specified, configures range-based partitioning for this table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "time_partitioning": {
        "block": {
          "attributes": {
            "expiration_ms": {
              "computed": true,
              "description": "Number of milliseconds for which to keep the storage for a partition.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "field": {
              "description": "The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_partition_filter": {
              "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description": "The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "If specified, configures time-based partitioning for this table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "view": {
        "block": {
          "attributes": {
            "query": {
              "description": "A query that BigQuery executes when the view is referenced.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_legacy_sql": {
              "description": "Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "If specified, configures this table as a view.",
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

func GoogleBigqueryTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryTable), &result)
	return &result
}
