package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDataset = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time when this dataset was created, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description": "A unique ID for this dataset, without the project name. The ID\nmust contain only letters (a-z, A-Z), numbers (0-9), or\nunderscores (_). The maximum length is 1,024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_partition_expiration_ms": {
        "description": "The default partition expiration for all partitioned tables in\nthe dataset, in milliseconds.\n\n\nOnce this property is set, all newly-created partitioned tables in\nthe dataset will have an 'expirationMs' property in the 'timePartitioning'\nsettings set to this value, and changing the value will only\naffect new tables, not existing ones. The storage in a partition will\nhave an expiration time of its partition time plus this value.\nSetting this property overrides the use of 'defaultTableExpirationMs'\nfor partitioned tables: only one of 'defaultTableExpirationMs' and\n'defaultPartitionExpirationMs' will be used for any new partitioned\ntable. If you provide an explicit 'timePartitioning.expirationMs' when\ncreating or updating a partitioned table, that value takes precedence\nover the default partition expiration time indicated by this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_table_expiration_ms": {
        "description": "The default lifetime of all tables in the dataset, in milliseconds.\nThe minimum value is 3600000 milliseconds (one hour).\n\n\nOnce this property is set, all newly-created tables in the dataset\nwill have an 'expirationTime' property set to the creation time plus\nthe value in this property, and changing the value will only affect\nnew tables, not existing ones. When the 'expirationTime' for a given\ntable is reached, that table will be deleted automatically.\nIf a table's 'expirationTime' is modified or removed before the\ntable expires, or if you provide an explicit 'expirationTime' when\ncreating a table, that value takes precedence over the default\nexpiration time indicated by this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "delete_contents_on_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A user-friendly description of the dataset",
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
      "friendly_name": {
        "description": "A descriptive name for the dataset",
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
        "description": "The labels associated with this dataset. You can use these to\norganize and group your datasets",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date when this dataset or any of its tables was last modified, in\nmilliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "description": "The geographic location where the dataset should reside.\nSee [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).\n\n\nThere are two types of locations, regional or multi-regional. A regional\nlocation is a specific geographic place, such as Tokyo, and a multi-regional\nlocation is a large geographic area, such as the United States, that\ncontains at least two geographic places.\n\n\nPossible regional values include: 'asia-east1', 'asia-northeast1',\n'asia-southeast1', 'australia-southeast1', 'europe-north1',\n'europe-west2' and 'us-east4'.\n\n\nPossible multi-regional values: 'EU' and 'US'.\n\n\nThe default value is multi-regional location 'US'.\nChanging this forces a new resource to be created.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "access": {
        "block": {
          "attributes": {
            "domain": {
              "description": "A domain to grant access to. Any users signed in with the\ndomain specified will be granted the specified access",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_by_email": {
              "description": "An email address of a Google Group to grant access to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role": {
              "description": "Describes the rights granted to the user specified by the other\nmember of the access object. Primitive, Predefined and custom\nroles are supported. Predefined roles that have equivalent\nprimitive roles are swapped by the API to their Primitive\ncounterparts, and will show a diff post-create. See\n[official docs](https://cloud.google.com/bigquery/docs/access-control).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "special_group": {
              "description": "A special group to grant access to.\n\n\nPossible values include:\n\n\n* 'projectOwners': Owners of the enclosing project.\n\n\n* 'projectReaders': Readers of the enclosing project.\n\n\n* 'projectWriters': Writers of the enclosing project.\n\n\n* 'allAuthenticatedUsers': All authenticated BigQuery users.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_by_email": {
              "description": "An email address of a user to grant access to. For example:\nfred@example.com",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "view": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The ID of the table. The ID must contain only letters (a-z,\nA-Z), numbers (0-9), or underscores (_). The maximum length\nis 1,024 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "default_encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Describes the Cloud KMS encryption key that will be used to protect destination\nBigQuery table. The BigQuery Service Account associated with your project requires\naccess to this encryption key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
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

func GoogleBigqueryDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDataset), &result)
	return &result
}
