package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIdentityPlatformOauthIdpConfig = `{
  "block": {
    "attributes": {
      "client_id": {
        "description": "The client id of an OAuth client.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_secret": {
        "description": "The client secret of the OAuth client, to enable OIDC code flow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Human friendly display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "description": "If this config allows users to sign in with the provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "issuer": {
        "description": "For OIDC Idps, the issuer identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the OauthIdpConfig. Must start with 'oidc.'.",
        "description_kind": "plain",
        "required": true,
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

func GoogleIdentityPlatformOauthIdpConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIdentityPlatformOauthIdpConfig), &result)
	return &result
}
