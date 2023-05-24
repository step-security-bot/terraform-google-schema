package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlBackupRun = `{
  "block": {
    "attributes": {
      "backup_id": {
        "computed": true,
        "description": "The identifier for this backup run. Unique only for a specific Cloud SQL instance. If left empty and multiple backups exist for the instance, most_recent must be set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "Name of the database instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "Location of the backups.",
        "description_kind": "plain",
        "type": "string"
      },
      "most_recent": {
        "description": "Toggles use of the most recent backup run if multiple backups exist for a Cloud SQL instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description": "Project ID of the project that contains the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The time the backup operation actually started in UTC timezone in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this run.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSqlBackupRunSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlBackupRun), &result)
	return &result
}
