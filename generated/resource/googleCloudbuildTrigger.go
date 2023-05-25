package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudbuildTrigger = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time when the trigger was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Human-readable description of the trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the trigger is disabled or not. If true, the trigger will never result in a build.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filename": {
        "description": "Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.",
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
      "ignored_files": {
        "description": "ignoredFiles and includedFiles are file glob matches using http://godoc/pkg/path/filepath#Match\nextended with support for '**'.\n\nIf ignoredFiles and changed files are both empty, then they are not\nused to determine whether or not to trigger a build.\n\nIf ignoredFiles is not empty, then we ignore any files that match any\nof the ignored_file globs. If the change has no files that are outside\nof the ignoredFiles globs, then we do not trigger a build.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "included_files": {
        "description": "ignoredFiles and includedFiles are file glob matches using http://godoc/pkg/path/filepath#Match\nextended with support for '**'.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is empty, then as far as this filter is concerned, we\nshould trigger the build.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is not empty, then we make sure that at least one of\nthose files matches a includedFiles glob. If not, then we do not trigger\na build.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Name of the trigger. Must be unique within the project.",
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
      "substitutions": {
        "description": "Substitutions data for Build resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "trigger_id": {
        "computed": true,
        "description": "The unique identifier for the trigger.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "build": {
        "block": {
          "attributes": {
            "images": {
              "description": "A list of images to be pushed upon the successful completion of all build steps.\nThe images are pushed using the builder service account's credentials.\nThe digests of the pushed images will be stored in the Build resource's results field.\nIf any of the images fail to be pushed, the build status is marked FAILURE.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tags": {
              "description": "Tags for annotation of a Build. These are not docker tags.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "step": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "A list of arguments that will be presented to the step when it is started.\n\nIf the image used to run the step's container has an entrypoint, the args\nare used as arguments to that entrypoint. If the image does not define an\nentrypoint, the first element in args is used as the entrypoint, and the\nremainder will be used as arguments.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dir": {
                    "description": "Working directory to use when running this step's container.\n\nIf this value is a relative path, it is relative to the build's working\ndirectory. If this value is absolute, it may be outside the build's working\ndirectory, in which case the contents of the path may not be persisted\nacross build step executions, unless a 'volume' for that path is specified.\n\nIf the build specifies a 'RepoSource' with 'dir' and a step with a\n'dir',\nwhich specifies an absolute path, the 'RepoSource' 'dir' is ignored\nfor the step's execution.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "entrypoint": {
                    "description": "Entrypoint to be used instead of the build step image's\ndefault entrypoint.\nIf unset, the image's default entrypoint is used",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "env": {
                    "description": "A list of environment variable definitions to be used when\nrunning a step.\n\nThe elements are of the form \"KEY=VALUE\" for the environment variable\n\"KEY\" being given the value \"VALUE\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "id": {
                    "description": "Unique identifier for this build step, used in 'wait_for' to\nreference this build step as a dependency.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the container image that will run this particular build step.\n\nIf the image is available in the host's Docker daemon's cache, it will be\nrun directly. If not, the host will attempt to pull the image first, using\nthe builder service account's credentials if necessary.\n\nThe Docker daemon's cache will already have the latest versions of all of\nthe officially supported build steps (https://github.com/GoogleCloudPlatform/cloud-builders).\nThe Docker daemon will also have cached many of the layers for some popular\nimages, like \"ubuntu\", \"debian\", but they will be refreshed at the time\nyou attempt to use them.\n\nIf you built an image in a previous build step, it will be stored in the\nhost's Docker daemon's cache and is available to use as the name for a\nlater build step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_env": {
                    "description": "A list of environment variables which are encrypted using\na Cloud Key\nManagement Service crypto key. These values must be specified in\nthe build's 'Secret'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "timeout": {
                    "description": "Time limit for executing this build step. If not defined,\nthe step has no\ntime limit and will be allowed to continue to run until either it\ncompletes or the build itself times out.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timing": {
                    "description": "Output only. Stores timing information for executing this\nbuild step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "wait_for": {
                    "description": "The ID(s) of the step(s) that this build step depends on.\n\nThis build step will not start until all the build steps in 'wait_for'\nhave completed successfully. If 'wait_for' is empty, this build step\nwill start when all previous build steps in the 'Build.Steps' list\nhave completed successfully.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "volumes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the volume to mount.\n\nVolume names must be unique per build step and must be valid names for\nDocker volumes. Each named volume must be used by at least two build steps.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "Path at which to mount the volume.\n\nPaths must be absolute and cannot conflict with other volume paths on\nthe same build step or with certain reserved volume paths.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
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
      "trigger_template": {
        "block": {
          "attributes": {
            "branch_name": {
              "description": "Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.\nThis field is a regular expression.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "commit_sha": {
              "description": "Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dir": {
              "description": "Directory, relative to the source root, in which to run the build.\n\nThis must be a relative path. If a step's dir is specified and\nis an absolute path, this value is ignored for that step's\nexecution.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_id": {
              "computed": true,
              "description": "ID of the project that owns the Cloud Source Repository. If\nomitted, the project ID requesting the build is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "repo_name": {
              "description": "Name of the Cloud Source Repository. If omitted, the name \"default\" is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_name": {
              "description": "Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.\nThis field is a regular expression.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleCloudbuildTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudbuildTrigger), &result)
	return &result
}
