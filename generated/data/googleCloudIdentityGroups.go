package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdentityGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description": "List of Cloud Identity groups.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "display_name": "string",
              "group_key": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "namespace": "string"
                  }
                ]
              ],
              "initial_group_config": "string",
              "labels": [
                "map",
                "string"
              ],
              "name": "string",
              "parent": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the entity under which this Group resides in the\nCloud Identity resource hierarchy.\n\nMust be of the form identitysources/{identity_source_id} for external-identity-mapped\ngroups or customers/{customer_id} for Google Groups.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudIdentityGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdentityGroups), &result)
	return &result
}
