package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudbuildBitbucketServerConfig = `{
  "block": {
    "attributes": {
      "api_key": {
        "description": "Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.\nChanging this field will result in deleting/ recreating the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_id": {
        "description": "The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_uri": {
        "description": "Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.\nIf you need to change it, please create another BitbucketServerConfig.",
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
      "location": {
        "description": "The location of this bitbucket server config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the config.",
        "description_kind": "plain",
        "type": "string"
      },
      "peered_network": {
        "description": "The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.\nThis should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,\nno network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format\nprojects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_ca": {
        "description": "SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "description": "Username of the account Cloud Build will use on Bitbucket Server.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "webhook_key": {
        "computed": true,
        "description": "Output only. UUID included in webhook requests. The UUID is used to look up the corresponding config.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "connected_repositories": {
        "block": {
          "attributes": {
            "project_key": {
              "description": "Identifier for the project storing the repository.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repo_slug": {
              "description": "Identifier for the repository.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Connected Bitbucket Server repositories for this config.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "secrets": {
        "block": {
          "attributes": {
            "admin_access_token_version_name": {
              "description": "The resource name for the admin access token's secret version.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "read_access_token_version_name": {
              "description": "The resource name for the read access token's secret version.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "webhook_secret_version_name": {
              "description": "Immutable. The resource name for the webhook secret's secret version. Once this field has been set, it cannot be changed.\nChanging this field will result in deleting/ recreating the resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Secret Manager secrets needed by the config.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleCloudbuildBitbucketServerConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudbuildBitbucketServerConfig), &result)
	return &result
}
