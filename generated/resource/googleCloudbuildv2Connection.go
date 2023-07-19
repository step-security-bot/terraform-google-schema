package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudbuildv2Connection = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Allows clients to store small amounts of arbitrary data.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Server assigned timestamp for when the connection was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "installation_state": {
        "computed": true,
        "description": "Output only. Installation state of the Connection.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action_uri": "string",
              "message": "string",
              "stage": "string"
            }
          ]
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Immutable. The resource name of the connection, in the format ` + "`" + `projects/{project}/locations/{location}/connections/{connection_id}` + "`" + `.",
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
      "reconciling": {
        "computed": true,
        "description": "Output only. Set to true when the connection is being set up or updated in the background.",
        "description_kind": "plain",
        "type": "bool"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Server assigned timestamp for when the connection was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "github_config": {
        "block": {
          "attributes": {
            "app_installation_id": {
              "description": "GitHub App installation id.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "oauth_token_secret_version": {
                    "description": "A SecretManager resource containing the OAuth token that authorizes the Cloud Build connection. Format: ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated to this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "OAuth credential of the account that authorized the Cloud Build GitHub App. It is recommended to use a robot account instead of a human user account. The OAuth token must be tied to the Cloud Build GitHub App.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to github.com.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "github_enterprise_config": {
        "block": {
          "attributes": {
            "app_id": {
              "description": "Id of the GitHub App created from the manifest.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "app_installation_id": {
              "description": "ID of the installation of the GitHub App.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "app_slug": {
              "description": "The URL-friendly name of the GitHub App.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_uri": {
              "description": "Required. The URI of the GitHub Enterprise host this connection is for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "private_key_secret_version": {
              "description": "SecretManager resource containing the private key of the GitHub App, formatted as ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_ca": {
              "description": "SSL certificate to use for requests to GitHub Enterprise.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook_secret_secret_version": {
              "description": "SecretManager resource containing the webhook secret of the GitHub App, formatted as ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for using Service Directory to privately connect to a GitHub Enterprise server. This should only be set if the GitHub Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the GitHub Enterprise server will be made over the public internet.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to an instance of GitHub Enterprise.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gitlab_config": {
        "block": {
          "attributes": {
            "host_uri": {
              "computed": true,
              "description": "The URI of the GitLab Enterprise host this connection is for. If not specified, the default value is https://gitlab.com.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_version": {
              "computed": true,
              "description": "Output only. Version of the GitLab Enterprise server running on the ` + "`" + `host_uri` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca": {
              "description": "SSL certificate to use for requests to GitLab Enterprise.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook_secret_secret_version": {
              "description": "Required. Immutable. SecretManager resource containing the webhook secret of a GitLab Enterprise project, formatted as ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated to this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Required. A GitLab personal access token with the ` + "`" + `api` + "`" + ` scope access.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "read_authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: ` + "`" + `projects/*/secrets/*/versions/*` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated to this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Required. A GitLab personal access token with the minimum ` + "`" + `read_api` + "`" + ` scope access.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for using Service Directory to privately connect to a GitLab Enterprise server. This should only be set if the GitLab Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the GitLab Enterprise server will be made over the public internet.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to gitlab.com or an instance of GitLab Enterprise.",
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

func GoogleCloudbuildv2ConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudbuildv2Connection), &result)
	return &result
}
