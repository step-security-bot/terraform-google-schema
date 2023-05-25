package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDeploymentManagerDeployment = `{
  "block": {
    "attributes": {
      "create_policy": {
        "description": "Set the policy to use for creating new resources. Only used on\ncreate and update. Valid values are 'CREATE_OR_ACQUIRE' (default) or\n'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist,\nthe deployment will fail. Note that updating this field does not\nactually affect the deployment, just how it is updated. Default value: \"CREATE_OR_ACQUIRE\" Possible values: [\"ACQUIRE\", \"CREATE_OR_ACQUIRE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_policy": {
        "description": "Set the policy to use for deleting new resources on update/delete.\nValid values are 'DELETE' (default) or 'ABANDON'. If 'DELETE',\nresource is deleted after removal from Deployment Manager. If\n'ABANDON', the resource is only removed from Deployment Manager\nand is not actually deleted. Note that updating this field does not\nactually change the deployment, just how it is updated. Default value: \"DELETE\" Possible values: [\"ABANDON\", \"DELETE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description": "Unique identifier for deployment. Output only.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional user-provided description of deployment.",
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
      "manifest": {
        "computed": true,
        "description": "Output only. URL of the manifest representing the last manifest that\nwas successfully deployed.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Unique name for the deployment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preview": {
        "description": "If set to true, a deployment is created with \"shell\" resources\nthat are not actually instantiated. This allows you to preview a\ndeployment. It can be updated to false to actually deploy\nwith real resources.\n ~\u003e**NOTE:** Deployment Manager does not allow update\nof a deployment in preview (unless updating to preview=false). Thus,\nTerraform will force-recreate deployments if either preview is updated\nto true or if other fields are updated while preview is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Output only. Server defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description": "Key for label.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description": "Value of label.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Key-value pairs to apply to this labels.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "target": {
        "block": {
          "block_types": {
            "config": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The full YAML contents of your configuration file.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The root configuration file to use for this deployment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "imports": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The full contents of the template that you want to import.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the template to import, as declared in the YAML\nconfiguration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specifies import files for this configuration. This can be\nused to import templates or other files. For example, you might\nimport a text file in order to use the file in a template.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Parameters that define your deployment, including the deployment\nconfiguration and relevant templates.",
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

func GoogleDeploymentManagerDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDeploymentManagerDeployment), &result)
	return &result
}
