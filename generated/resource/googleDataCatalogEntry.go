package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataCatalogEntry = `{
  "block": {
    "attributes": {
      "bigquery_date_sharded_spec": {
        "computed": true,
        "description": "Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.\nContext: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dataset": "string",
              "shard_count": "number",
              "table_prefix": "string"
            }
          ]
        ]
      },
      "bigquery_table_spec": {
        "computed": true,
        "description": "Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "table_source_type": "string",
              "table_spec": [
                "list",
                [
                  "object",
                  {
                    "grouped_entry": "string"
                  }
                ]
              ],
              "view_spec": [
                "list",
                [
                  "object",
                  {
                    "view_query": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "description": {
        "description": "Entry description, which can consist of several sentences or paragraphs that describe entry contents.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display information such as title and description. A short name to identify the entry,\nfor example, \"Analytics Data - Jan 2011\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entry_group": {
        "description": "The name of the entry group this entry is in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entry_id": {
        "description": "The id of the entry to create.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integrated_system": {
        "computed": true,
        "description": "This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.",
        "description_kind": "plain",
        "type": "string"
      },
      "linked_resource": {
        "computed": true,
        "description": "The resource this metadata entry refers to.\nFor Google Cloud Platform resources, linkedResource is the full name of the resource.\nFor example, the linkedResource for a table resource from BigQuery is:\n//bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId\nOutput only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,\nthis field is optional and defaults to an empty string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The Data Catalog resource name of the entry in URL format.\nExample: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.\nNote that this Entry and its child resources may not actually be stored in the location in this name.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "description": "Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema\nattached to it. See\nhttps://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema\nfor what fields this schema can contain.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of the entry. Only used for Entries with types in the EntryType enum.\nCurrently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType. Possible values: [\"FILESET\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_specified_system": {
        "description": "This field indicates the entry's source system that Data Catalog does not integrate with.\nuserSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,\nand underscores; are case insensitive; must be at least 1 character and at most 64 characters long.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_specified_type": {
        "description": "Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.\nWhen creating an entry, users should check the enum values first, if nothing matches the entry\nto be created, then provide a custom value, for example \"my_special_type\".\nuserSpecifiedType strings must begin with a letter or underscore and can only contain letters,\nnumbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "gcs_fileset_spec": {
        "block": {
          "attributes": {
            "file_patterns": {
              "description": "Patterns to identify a set of files in Google Cloud Storage.\nSee [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)\nfor more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:\n\n* gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.\n* gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.\n* gs://bucket_name/file*: matches files prefixed by file in bucket_name\n* gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name\n* gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name\n* gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name\n* gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b\n* gs://another_bucket/a.txt: matches gs://another_bucket/a.txt",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sample_gcs_file_specs": {
              "computed": true,
              "description": "Sample files contained in this fileset, not all files contained in this fileset are represented here.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "file_path": "string",
                    "size_bytes": "number"
                  }
                ]
              ]
            }
          },
          "description": "Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataCatalogEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataCatalogEntry), &result)
	return &result
}
