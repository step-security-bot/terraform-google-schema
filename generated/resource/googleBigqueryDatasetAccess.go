package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDatasetAccess = `{
  "block": {
    "attributes": {
      "api_updated_member": {
        "computed": true,
        "description": "If true, represents that that the iam_member in the config was translated to a different member type by the API, and is stored in state as a different member type",
        "description_kind": "plain",
        "type": "bool"
      },
      "dataset_id": {
        "description": "A unique ID for this dataset, without the project name. The ID\nmust contain only letters (a-z, A-Z), numbers (0-9), or\nunderscores (_). The maximum length is 1,024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "iam_member": {
        "description": "Some other type of member that appears in the IAM Policy but isn't a user,\ngroup, domain, or special group. For example: 'allUsers'",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description": "Describes the rights granted to the user specified by the other\nmember of the access object. Basic, predefined, and custom roles are\nsupported. Predefined roles that have equivalent basic roles are\nswapped by the API to their basic counterparts, and will show a diff\npost-create. See\n[official docs](https://cloud.google.com/bigquery/docs/access-control).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "special_group": {
        "description": "A special group to grant access to. Possible values include:\n\n\n* 'projectOwners': Owners of the enclosing project.\n\n\n* 'projectReaders': Readers of the enclosing project.\n\n\n* 'projectWriters': Writers of the enclosing project.\n\n\n* 'allAuthenticatedUsers': All authenticated BigQuery users.",
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
      },
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
          "description": "A view from a different dataset to grant access to. Queries\nexecuted against that view will have read access to tables in\nthis dataset. The role field is not required when this field is\nset. If that view is updated by any user, access to the view\nneeds to be granted again via an update operation.",
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

func GoogleBigqueryDatasetAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDatasetAccess), &result)
	return &result
}
