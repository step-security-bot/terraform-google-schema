package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNotebooksEnvironment = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Instance creation time",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A brief description of this environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of this environment for the UI.",
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
        "description": "A reference to the zone where the machine resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name specified for the Environment instance.\nFormat: projects/{project_id}/locations/{location}/environments/{environmentId}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "post_startup_script": {
        "description": "Path to a Bash script that automatically runs after a notebook instance fully boots up.\nThe path must be a URL or Cloud Storage path. Example: \"gs://path-to-file/file-name\"",
        "description_kind": "plain",
        "optional": true,
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
      "container_image": {
        "block": {
          "attributes": {
            "repository": {
              "description": "The path to the container image repository.\nFor example: gcr.io/{project_id}/{imageName}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tag": {
              "description": "The tag of the container image. If not specified, this defaults to the latest tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Use a container image to start the notebook instance.",
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
      },
      "vm_image": {
        "block": {
          "attributes": {
            "image_family": {
              "description": "Use this VM image family to find the image; the newest image in this family will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_name": {
              "description": "Use VM image name to find the image.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description": "The name of the Google Cloud project that this VM image belongs to.\nFormat: projects/{project_id}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Use a Compute Engine VM image to start the notebook instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleNotebooksEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNotebooksEnvironment), &result)
	return &result
}
