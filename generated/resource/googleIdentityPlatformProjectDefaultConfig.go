package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIdentityPlatformProjectDefaultConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Config resource. Example: \"projects/my-awesome-project/config\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "sign_in": {
        "block": {
          "attributes": {
            "allow_duplicate_emails": {
              "description": "Whether to allow more than one account to have the same email.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hash_config": {
              "computed": true,
              "description": "Output only. Hash config information.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "algorithm": "string",
                    "memory_cost": "number",
                    "rounds": "number",
                    "salt_separator": "string",
                    "signer_key": "string"
                  }
                ]
              ]
            }
          },
          "block_types": {
            "anonymous": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether anonymous user auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration options related to authenticating an anonymous user.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "email": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether email auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "password_required": {
                    "description": "Whether a password is required for email auth or not. If true, both an email and\npassword must be provided to sign in. If false, a user may sign in via either\nemail/password or email link.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration options related to authenticating a user by their email address.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "phone_number": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether phone number auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "test_phone_numbers": {
                    "description": "A map of \u003ctest phone number, fake code\u003e that can be used for phone auth testing.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Configuration options related to authenticated a user by their phone number.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration related to local sign in methods.",
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

func GoogleIdentityPlatformProjectDefaultConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIdentityPlatformProjectDefaultConfig), &result)
	return &result
}
