package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexAsset = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the asset was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataplex_zone": {
        "description": "The zone for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Optional. Description of the asset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discovery_status": {
        "computed": true,
        "description": "Output only. Status of the discovery feature applied to data referenced by this asset.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "last_run_duration": "string",
              "last_run_time": "string",
              "message": "string",
              "state": "string",
              "stats": [
                "list",
                [
                  "object",
                  {
                    "data_items": "number",
                    "data_size": "number",
                    "filesets": "number",
                    "tables": "number"
                  }
                ]
              ],
              "update_time": "string"
            }
          ]
        ]
      },
      "display_name": {
        "description": "Optional. User friendly display name.",
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
        "description": "Optional. User defined labels for the asset.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "lake": {
        "description": "The lake for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the asset.",
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
      "resource_status": {
        "computed": true,
        "description": "Output only. Status of the resource referenced by this asset.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "message": "string",
              "state": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "security_status": {
        "computed": true,
        "description": "Output only. Status of the security policy applied to resource referenced by this asset.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "message": "string",
              "state": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the asset was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "discovery_spec": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Required. Whether discovery is enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "exclude_patterns": {
              "description": "Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "include_patterns": {
              "description": "Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "schedule": {
              "description": "Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: \"CRON_TZ=${IANA_TIME_ZONE}\" or TZ=${IANA_TIME_ZONE}\". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, \"CRON_TZ=America/New_York 1 * * * *\", or \"TZ=America/New_York 1 * * * *\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "csv_options": {
              "block": {
                "attributes": {
                  "delimiter": {
                    "description": "Optional. The delimiter being used to separate values. This defaults to ','.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disable_type_inference": {
                    "description": "Optional. Whether to disable the inference of data type for CSV data. If true, all columns will be registered as strings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encoding": {
                    "description": "Optional. The character encoding of the data. The default is UTF-8.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "header_rows": {
                    "description": "Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Optional. Configuration for CSV data.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "json_options": {
              "block": {
                "attributes": {
                  "disable_type_inference": {
                    "description": "Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encoding": {
                    "description": "Optional. The character encoding of the data. The default is UTF-8.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Configuration for Json data.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "resource_spec": {
        "block": {
          "attributes": {
            "name": {
              "description": "Immutable. Relative name of the cloud resource that contains the data that is being managed within a lake. For example: ` + "`" + `projects/{project_number}/buckets/{bucket_id}` + "`" + ` ` + "`" + `projects/{project_number}/datasets/{dataset_id}` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read_access_mode": {
              "computed": true,
              "description": "Optional. Determines how read permissions are handled for each asset and their associated tables. Only available to storage buckets assets. Possible values: DIRECT, MANAGED",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Required. Immutable. Specification of the resource that is referenced by this asset.",
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

func GoogleDataplexAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexAsset), &result)
	return &result
}
