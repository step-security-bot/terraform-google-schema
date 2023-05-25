package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingProjectSink = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of this sink. The maximum length of the description is 8000 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "description": "The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: \"storage.googleapis.com/[GCS_BUCKET]\" \"bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]\" \"pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]\" The writer associated with the sink must have access to write to the above resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disabled": {
        "description": "If set to True, then this sink is disabled and it does not export any log entries.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter": {
        "description": "The filter to apply when exporting logs. Only log entries that match the filter are exported.",
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
      "name": {
        "description": "The name of the logging sink.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project to create the sink in. If omitted, the project associated with the provider is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unique_writer_identity": {
        "description": "Whether or not to create a unique identity associated with this sink. If false (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true, then a unique service account is created and used for this sink. If you wish to publish logs across projects, you must set unique_writer_identity to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "writer_identity": {
        "computed": true,
        "description": "The identity associated with this sink. This identity must be granted write access to the configured destination.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "bigquery_options": {
        "block": {
          "attributes": {
            "use_partitioned_tables": {
              "description": "Whether to use BigQuery's partition tables. By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Options that affect sinks exporting data to BigQuery.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "exclusions": {
        "block": {
          "attributes": {
            "description": {
              "description": "A description of this exclusion.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disabled": {
              "description": "If set to True, then this exclusion is disabled and it does not exclude any log entries",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "filter": {
              "description": "An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "A client-assigned identifier, such as \"load-balancer-exclusion\". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingProjectSinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingProjectSink), &result)
	return &result
}
