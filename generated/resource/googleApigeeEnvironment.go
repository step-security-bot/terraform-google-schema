package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvironment = `{
  "block": {
    "attributes": {
      "api_proxy_type": {
        "computed": true,
        "description": "Optional. API Proxy type supported by the environment. The type can be set when creating\nthe Environment and cannot be changed. Possible values: [\"API_PROXY_TYPE_UNSPECIFIED\", \"PROGRAMMABLE\", \"CONFIGURABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description": "Optional. Deployment type supported by the environment. The deployment type can be\nset when creating the environment and cannot be changed. When you enable archive\ndeployment, you will be prevented from performing a subset of actions within the\nenvironment, including:\nManaging the deployment of API proxy or shared flow revisions;\nCreating, updating, or deleting resource files;\nCreating, updating, or deleting target servers. Possible values: [\"DEPLOYMENT_TYPE_UNSPECIFIED\", \"PROXY\", \"ARCHIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of the environment.",
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
      "name": {
        "description": "The resource ID of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee environment,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "node_config": {
        "block": {
          "attributes": {
            "current_aggregate_node_count": {
              "computed": true,
              "description": "The current total number of gateway nodes that each environment currently has across\nall instances.",
              "description_kind": "plain",
              "type": "string"
            },
            "max_node_count": {
              "description": "The maximum total number of gateway nodes that the is reserved for all instances that\nhas the specified environment. If not specified, the default is determined by the\nrecommended maximum number of nodes for that gateway.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_node_count": {
              "description": "The minimum total number of gateway nodes that the is reserved for all instances that\nhas the specified environment. If not specified, the default is determined by the\nrecommended minimum number of nodes for that gateway.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "NodeConfig for setting the min/max number of nodes associated with the environment.",
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

func GoogleApigeeEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvironment), &result)
	return &result
}
