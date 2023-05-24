package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcChannel = `{
  "block": {
    "attributes": {
      "activation_token": {
        "computed": true,
        "description": "Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key_name": {
        "description": "Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern ` + "`" + `projects/*/locations/*/keyRings/*/cryptoKeys/*` + "`" + `.",
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
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Required. The resource name of the channel. Must be unique within the location on the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pubsub_topic": {
        "computed": true,
        "description": "Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: ` + "`" + `projects/{project}/topics/{topic_id}` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE",
        "description_kind": "plain",
        "type": "string"
      },
      "third_party_provider": {
        "description": "The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: ` + "`" + `projects/{project}/locations/{location}/providers/{provider_id}` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last-modified time.",
        "description_kind": "plain",
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

func GoogleEventarcChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcChannel), &result)
	return &result
}
