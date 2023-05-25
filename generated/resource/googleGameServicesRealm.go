package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGameServicesRealm = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Human readable description of the realm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "ETag of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels associated with this realm. Each label is a key-value pair.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the Realm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource id of the realm, of the form:\n'projects/{project_id}/locations/{location}/realms/{realm_id}'. For\nexample, 'projects/my-project/locations/{location}/realms/my-realm'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "realm_id": {
        "description": "GCP region of the Realm.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_zone": {
        "description": "Required. Time zone where all realm-specific policies are evaluated. The value of\nthis field must be from the IANA time zone database:\nhttps://www.iana.org/time-zones.",
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

func GoogleGameServicesRealmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGameServicesRealm), &result)
	return &result
}
