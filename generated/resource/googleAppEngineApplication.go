package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineApplication = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description": "Identifier of the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_domain": {
        "computed": true,
        "description": "The domain to authenticate users with when using App Engine's User API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_bucket": {
        "computed": true,
        "description": "The GCS bucket code is being stored in for this app.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_bucket": {
        "computed": true,
        "description": "The GCS bucket content is being stored in for this app.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_hostname": {
        "computed": true,
        "description": "The default hostname for this app.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcr_domain": {
        "computed": true,
        "description": "The GCR domain used for storing managed Docker images for this app.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location_id": {
        "description": "The location to serve the app from.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Unique name of the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project ID to create the application under.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serving_status": {
        "computed": true,
        "description": "The serving status of the app.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "url_dispatch_rule": {
        "computed": true,
        "description": "A list of dispatch rule blocks. Each block has a domain, path, and service field.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain": "string",
              "path": "string",
              "service": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "feature_settings": {
        "block": {
          "attributes": {
            "split_health_checks": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "A block of optional settings to configure specific App Engine features:",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "iap": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Adapted for use with the app",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "oauth2_client_id": {
              "description": "OAuth2 client ID to use for the authentication flow.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "oauth2_client_secret": {
              "description": "OAuth2 client secret to use for the authentication flow. The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "oauth2_client_secret_sha256": {
              "computed": true,
              "description": "Hex-encoded SHA-256 hash of the client secret.",
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Settings for enabling Cloud Identity Aware Proxy",
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

func GoogleAppEngineApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineApplication), &result)
	return &result
}
