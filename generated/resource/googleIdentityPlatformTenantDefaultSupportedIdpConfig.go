package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIdentityPlatformTenantDefaultSupportedIdpConfig = `{
  "block": {
    "attributes": {
      "client_id": {
        "description": "OAuth client ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_secret": {
        "description": "OAuth client secret",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description": "If this IDP allows the user to sign in",
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
      "idp_id": {
        "description": "ID of the IDP. Possible values include:\n\n* 'apple.com'\n\n* 'facebook.com'\n\n* 'gc.apple.com'\n\n* 'github.com'\n\n* 'google.com'\n\n* 'linkedin.com'\n\n* 'microsoft.com'\n\n* 'playgames.google.com'\n\n* 'twitter.com'\n\n* 'yahoo.com'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the default supported IDP config resource",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tenant": {
        "description": "The name of the tenant where this DefaultSupportedIdpConfig resource exists",
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

func GoogleIdentityPlatformTenantDefaultSupportedIdpConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIdentityPlatformTenantDefaultSupportedIdpConfig), &result)
	return &result
}
