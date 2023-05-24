package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingSink = `{
  "block": {
    "attributes": {
      "bigquery_options": {
        "computed": true,
        "description": "Options that affect sinks exporting data to BigQuery.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "use_partitioned_tables": "bool"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "A description of this sink. The maximum length of the description is 8000 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "computed": true,
        "description": "The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: \"storage.googleapis.com/[GCS_BUCKET]\" \"bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]\" \"pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]\" The writer associated with the sink must have access to write to the above resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description": "If set to True, then this sink is disabled and it does not export any log entries.",
        "description_kind": "plain",
        "type": "bool"
      },
      "exclusions": {
        "computed": true,
        "description": "Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion's filters, it will not be exported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "disabled": "bool",
              "filter": "string",
              "name": "string"
            }
          ]
        ]
      },
      "filter": {
        "computed": true,
        "description": "The filter to apply when exporting logs. Only log entries that match the filter are exported.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Required. An identifier for the resource in format: \"projects/[PROJECT_ID]/sinks/[SINK_NAME]\", \"organizations/[ORGANIZATION_ID]/sinks/[SINK_NAME]\", \"billingAccounts/[BILLING_ACCOUNT_ID]/sinks/[SINK_NAME]\", \"folders/[FOLDER_ID]/sinks/[SINK_NAME]\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the logging sink.",
        "description_kind": "plain",
        "type": "string"
      },
      "writer_identity": {
        "computed": true,
        "description": "The identity associated with this sink. This identity must be granted write access to the configured destination.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingSinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingSink), &result)
	return &result
}
