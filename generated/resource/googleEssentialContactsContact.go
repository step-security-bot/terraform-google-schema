package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEssentialContactsContact = `{
  "block": {
    "attributes": {
      "email": {
        "description": "The email address to send notifications to. This does not need to be a Google account.",
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
      "language_tag": {
        "description": "The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_category_subscriptions": {
        "description": "The categories of notifications that the contact will receive communications for.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "parent": {
        "description": "The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}",
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

func GoogleEssentialContactsContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEssentialContactsContact), &result)
	return &result
}
