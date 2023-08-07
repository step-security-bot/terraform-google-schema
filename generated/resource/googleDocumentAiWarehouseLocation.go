package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDocumentAiWarehouseLocation = `{
  "block": {
    "attributes": {
      "access_control_mode": {
        "description": "The access control mode for accessing the customer data. Possible values: [\"ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI\", \"ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID\", \"ACL_MODE_UNIVERSAL_ACCESS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_type": {
        "description": "The type of database used to store customer data. Possible values: [\"DB_INFRA_SPANNER\", \"DB_CLOUD_SQL_POSTGRES\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "document_creator_default_role": {
        "description": "The default role for the person who create a document. Possible values: [\"DOCUMENT_ADMIN\", \"DOCUMENT_EDITOR\", \"DOCUMENT_VIEWER\"]",
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
      "kms_key": {
        "description": "The KMS key used for CMEK encryption. It is required that\nthe kms key is in the same region as the endpoint. The\nsame key will be used for all provisioned resources, if\nencryption is available. If the kmsKey is left empty, no\nencryption will be enforced.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_number": {
        "description": "The unique identifier of the project.",
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDocumentAiWarehouseLocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDocumentAiWarehouseLocation), &result)
	return &result
}
