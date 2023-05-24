package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexZone = `{
  "block": {
    "attributes": {
      "asset_status": {
        "computed": true,
        "description": "Output only. Aggregated status of the underlying assets of the zone.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_assets": "number",
              "security_policy_applying_assets": "number",
              "update_time": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the zone was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. Description of the zone.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "description": "Optional. User defined labels for the zone.",
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
        "description": "The name of the zone.",
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
      "state": {
        "computed": true,
        "description": "Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the zone was last updated.",
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
              "computed": true,
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
          "description": "Required. Specification of the discovery feature applied to data in this zone.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "resource_spec": {
        "block": {
          "attributes": {
            "location_type": {
              "description": "Required. Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Required. Immutable. Specification of the resources that are referenced by the assets within this zone.",
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

func GoogleDataplexZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexZone), &result)
	return &result
}
