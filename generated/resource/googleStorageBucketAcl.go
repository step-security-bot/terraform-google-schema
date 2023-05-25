package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketAcl = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the bucket it applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_acl": {
        "description": "Configure this ACL to be the default ACL.",
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
      "predefined_acl": {
        "description": "The canned GCS ACL to apply. Must be set if role_entity is not.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_entity": {
        "computed": true,
        "description": "List of role/entity pairs in the form ROLE:entity. See GCS Bucket ACL documentation  for more details. Must be set if predefined_acl is not.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageBucketAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucketAcl), &result)
	return &result
}
