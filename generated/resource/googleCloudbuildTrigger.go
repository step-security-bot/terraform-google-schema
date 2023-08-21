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
        "description": "Path, from the source root, to a file whose contents is used for the template.\nEither a filename or build template must be provided. Set this only when using trigger_template or github.\nWhen using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter": {
        "description": "A Common Expression Language string. Used only with Pub/Sub and Webhook.",
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
        "description": "ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match\nextended with support for '**'.\n\nIf ignoredFiles and changed files are both empty, then they are not\nused to determine whether or not to trigger a build.\n\nIf ignoredFiles is not empty, then we ignore any files that match any\nof the ignored_file globs. If the change has no files that are outside\nof the ignoredFiles globs, then we do not trigger a build.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "include_build_logs": {
        "description": "Build logs will be sent back to GitHub as part of the checkrun\nresult.  Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or\nINCLUDE_BUILD_LOGS_WITH_STATUS Possible values: [\"INCLUDE_BUILD_LOGS_UNSPECIFIED\", \"INCLUDE_BUILD_LOGS_WITH_STATUS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "included_files": {
        "description": "ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match\nextended with support for '**'.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is empty, then as far as this filter is concerned, we\nshould trigger the build.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is not empty, then we make sure that at least one of\nthose files matches a includedFiles glob. If not, then we do not trigger\na build.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "description": "The [Cloud Build location](https://cloud.google.com/build/docs/locations) for the trigger.\nIf not specified, \"global\" is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "service_account": {
        "description": "The service account used for all user-controlled operations including\ntriggers.patch, triggers.run, builds.create, and builds.cancel.\n\nIf no service account is set, then the standard Cloud Build service account\n([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.\n\nFormat: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}",
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
      "tags": {
        "description": "Tags for annotation of a BuildTrigger",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
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
      "approval_config": {
        "block": {
          "attributes": {
            "approval_required": {
              "description": "Whether or not approval is needed. If this is set on a build, it will become pending when run,\nand will need to be explicitly approved to start.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configuration for manual approval to start a build invocation of this BuildTrigger.\nBuilds created by this trigger will require approval before they execute.\nAny user with a Cloud Build Approver role for the project can approve a build.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "bitbucket_server_trigger_config": {
        "block": {
          "attributes": {
            "bitbucket_server_config_resource": {
              "description": "The Bitbucket server config resource that this trigger config maps to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_key": {
              "description": "Key of the project that the repo is in. For example: The key for https://mybitbucket.server/projects/TEST/repos/test-repo is \"TEST\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repo_slug": {
              "description": "Slug of the repository. A repository slug is a URL-friendly version of a repository name, automatically generated by Bitbucket for use in the URL.\nFor example, if the repository name is 'test repo', in the URL it would become 'test-repo' as in https://mybitbucket.server/projects/TEST/repos/test-repo.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "pull_request": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.\nThe syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "comment_control": {
                    "description": "Configure builds to run whether a repository owner or collaborator need to comment /gcbrun. Possible values: [\"COMMENTS_DISABLED\", \"COMMENTS_ENABLED\", \"COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "If true, branches that do NOT match the git_ref will trigger a build.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Filter to match changes in pull requests.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "push": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.  Specify only one of branch or tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "When true, only trigger a build if the revision regex does NOT match the gitRef regex.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "Regex of tags to match.  Specify only one of branch or tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Filter to match changes in refs like branches, tags.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
            "logs_bucket": {
              "description": "Google Cloud Storage bucket where logs should be written.\nLogs file names will be of the format ${logsBucket}/log-${build_id}.txt.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "queue_ttl": {
              "description": "TTL in queue for this build. If provided and the build is enqueued longer than this value,\nthe build will expire and the build status will be EXPIRED.\nThe TTL starts ticking from createTime.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
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
            "tags": {
              "description": "Tags for annotation of a Build. These are not docker tags.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "timeout": {
              "description": "Amount of time that this build should be allowed to run, to second granularity.\nIf this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.\nThis timeout must be equal to or greater than the sum of the timeouts for build steps within the build.\nThe expected format is the number of seconds followed by s.\nDefault time is ten minutes (600s).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "artifacts": {
              "block": {
                "attributes": {
                  "images": {
                    "description": "A list of images to be pushed upon the successful completion of all build steps.\n\nThe images will be pushed using the builder service account's credentials.\n\nThe digests of the pushed images will be stored in the Build resource's results field.\n\nIf any of the images fail to be pushed, the build is marked FAILURE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "objects": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description": "Cloud Storage bucket and optional object path, in the form \"gs://bucket/path/to/somewhere/\".\n\nFiles in the workspace matching any path pattern will be uploaded to Cloud Storage with\nthis location as a prefix.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "paths": {
                          "description": "Path globs used to match files in the build's workspace.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "timing": {
                          "computed": true,
                          "description": "Output only. Stores timing information for pushing all artifact objects.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            [
                              "object",
                              {
                                "end_time": "string",
                                "start_time": "string"
                              }
                            ]
                          ]
                        }
                      },
                      "description": "A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.\n\nFiles in the workspace matching specified paths globs will be uploaded to the\nCloud Storage location using the builder service account's credentials.\n\nThe location and generation of the uploaded objects will be stored in the Build resource's results field.\n\nIf any objects fail to be pushed, the build is marked FAILURE.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Artifacts produced by the build that should be uploaded upon successful completion of all build steps.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "available_secrets": {
              "block": {
                "block_types": {
                  "secret_manager": {
                    "block": {
                      "attributes": {
                        "env": {
                          "description": "Environment variable name to associate with the secret. Secret environment\nvariables must be unique across all of a build's secrets, and must be used\nby at least one build step.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version_name": {
                          "description": "Resource name of the SecretVersion. In format: projects/*/secrets/*/versions/*",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Pairs a secret environment variable with a SecretVersion in Secret Manager.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Secrets and secret environment variables.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "options": {
              "block": {
                "attributes": {
                  "disk_size_gb": {
                    "description": "Requested disk size for the VM that runs the build. Note that this is NOT \"disk free\";\nsome of the space will be used by the operating system and build utilities.\nAlso note that this is the minimum disk size that will be allocated for the build --\nthe build may run with a larger disk than requested. At present, the maximum disk size\nis 1000GB; builds that request more than the maximum are rejected with an error.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "dynamic_substitutions": {
                    "description": "Option to specify whether or not to apply bash style string operations to the substitutions.\n\nNOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "env": {
                    "description": "A list of global environment variable definitions that will exist for all build steps\nin this build. If a variable is defined in both globally and in a build step,\nthe variable will use the build step value.\n\nThe elements are of the form \"KEY=VALUE\" for the environment variable \"KEY\" being given the value \"VALUE\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "log_streaming_option": {
                    "description": "Option to define build log streaming behavior to Google Cloud Storage. Possible values: [\"STREAM_DEFAULT\", \"STREAM_ON\", \"STREAM_OFF\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "logging": {
                    "description": "Option to specify the logging mode, which determines if and where build logs are stored. Possible values: [\"LOGGING_UNSPECIFIED\", \"LEGACY\", \"GCS_ONLY\", \"STACKDRIVER_ONLY\", \"CLOUD_LOGGING_ONLY\", \"NONE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "machine_type": {
                    "description": "Compute Engine machine type on which to run the build.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "requested_verify_option": {
                    "description": "Requested verifiability options. Possible values: [\"NOT_VERIFIED\", \"VERIFIED\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_env": {
                    "description": "A list of global environment variables, which are encrypted using a Cloud Key Management\nService crypto key. These values must be specified in the build's Secret. These variables\nwill be available to all build steps in this build.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source_provenance_hash": {
                    "description": "Requested hash for SourceProvenance. Possible values: [\"NONE\", \"SHA256\", \"MD5\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "substitution_option": {
                    "description": "Option to specify behavior when there is an error in the substitution checks.\n\nNOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden\nin the build configuration file. Possible values: [\"MUST_MATCH\", \"ALLOW_LOOSE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "worker_pool": {
                    "description": "Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}\n\nThis field is experimental.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "volumes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the volume to mount.\n\nVolume names must be unique per build step and must be valid names for Docker volumes.\nEach named volume must be used by at least two build steps.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "Path at which to mount the volume.\n\nPaths must be absolute and cannot conflict with other volume paths on the same\nbuild step or with certain reserved volume paths.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Global list of volumes to mount for ALL build steps\n\nEach volume is created as an empty volume prior to starting the build process.\nUpon completion of the build, volumes and their contents are discarded. Global\nvolume names and paths cannot conflict with the volumes defined a build step.\n\nUsing a global volume in a build with only one step is not valid as it is indicative\nof a build request with an incorrect configuration.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Special options for this build.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secret": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "Cloud KMS key name to use to decrypt these envs.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_env": {
                    "description": "Map of environment variable name to its encrypted value.\nSecret environment variables must be unique across all of a build's secrets,\nand must be used by at least one build step. Values can be at most 64 KB in size.\nThere can be at most 100 secret values across all of a build's secrets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Secrets to decrypt using Cloud Key Management Service.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "source": {
              "block": {
                "block_types": {
                  "repo_source": {
                    "block": {
                      "attributes": {
                        "branch_name": {
                          "description": "Regex matching branches to build. Exactly one a of branch name, tag, or commit SHA must be provided.\nThe syntax of the regular expressions accepted is the syntax accepted by RE2 and\ndescribed at https://github.com/google/re2/wiki/Syntax",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "commit_sha": {
                          "description": "Explicit commit SHA to build. Exactly one a of branch name, tag, or commit SHA must be provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "dir": {
                          "description": "Directory, relative to the source root, in which to run the build.\nThis must be a relative path. If a step's dir is specified and is an absolute path,\nthis value is ignored for that step's execution.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "invert_regex": {
                          "description": "Only trigger a build if the revision regex does NOT match the revision regex.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "project_id": {
                          "description": "ID of the project that owns the Cloud Source Repository.\nIf omitted, the project ID requesting the build is assumed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "repo_name": {
                          "description": "Name of the Cloud Source Repository.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "substitutions": {
                          "description": "Substitutions to use in a triggered build. Should only be used with triggers.run",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "tag_name": {
                          "description": "Regex matching tags to build. Exactly one a of branch name, tag, or commit SHA must be provided.\nThe syntax of the regular expressions accepted is the syntax accepted by RE2 and\ndescribed at https://github.com/google/re2/wiki/Syntax",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Location of the source in a Google Cloud Source Repository.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "storage_source": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description": "Google Cloud Storage bucket containing the source.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "generation": {
                          "description": "Google Cloud Storage generation for the object.\nIf the generation is omitted, the latest generation will be used",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "object": {
                          "description": "Google Cloud Storage object containing the source.\nThis object must be a gzipped archive file (.tar.gz) containing source to build.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Location of the source in an archive file in Google Cloud Storage.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The location of the source files to build.\n\nOne of 'storageSource' or 'repoSource' must be provided.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "step": {
              "block": {
                "attributes": {
                  "allow_exit_codes": {
                    "description": "Allow this build step to fail without failing the entire build if and\nonly if the exit code is one of the specified codes.\n\nIf 'allowFailure' is also specified, this field will take precedence.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "allow_failure": {
                    "description": "Allow this build step to fail without failing the entire build.\nIf false, the entire build will fail if this step fails. Otherwise, the\nbuild will succeed, but this step will still have a failure status.\nError information will be reported in the 'failureDetail' field.\n\n'allowExitCodes' takes precedence over this field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
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
                    "description": "The name of the container image that will run this particular build step.\n\nIf the image is available in the host's Docker daemon's cache, it will be\nrun directly. If not, the host will attempt to pull the image first, using\nthe builder service account's credentials if necessary.\n\nThe Docker daemon's cache will already have the latest versions of all of\nthe officially supported build steps (see https://github.com/GoogleCloudPlatform/cloud-builders\nfor images and examples).\nThe Docker daemon will also have cached many of the layers for some popular\nimages, like \"ubuntu\", \"debian\", but they will be refreshed at the time\nyou attempt to use them.\n\nIf you built an image in a previous build step, it will be stored in the\nhost's Docker daemon's cache and is available to use as the name for a\nlater build step.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "script": {
                    "description": "A shell script to be executed in the step.\nWhen script is provided, the user cannot specify the entrypoint or args.",
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
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "Path at which to mount the volume.\n\nPaths must be absolute and cannot conflict with other volume paths on\nthe same build step or with certain reserved volume paths.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "List of volumes to mount into the build step.\n\nEach volume is created as an empty volume prior to execution of the\nbuild step. Upon completion of the build, volumes and their contents\nare discarded.\n\nUsing a named volume in only one step is not valid as it is\nindicative of a build request with an incorrect configuration.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The operations to be performed on the workspace.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Contents of the build template. Either a filename or build template must be provided.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "git_file_source": {
        "block": {
          "attributes": {
            "bitbucket_server_config": {
              "description": "The full resource name of the bitbucket server config.\nFormat: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "github_enterprise_config": {
              "description": "The full resource name of the github enterprise config.\nFormat: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description": "The path of the file, with the repo root as the root of the path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repo_type": {
              "description": "The type of the repo, since it may not be explicit from the repo field (e.g from a URL).\nValues can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: [\"UNKNOWN\", \"CLOUD_SOURCE_REPOSITORIES\", \"GITHUB\", \"BITBUCKET_SERVER\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repository": {
              "description": "The fully qualified resource name of the Repo API repository. The fully qualified resource name of the Repo API repository.\nIf unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "revision": {
              "description": "The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the\nfilename (optional). This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions\nIf unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "description": "The URI of the repo (optional). If unspecified, the repo from which the trigger\ninvocation originated is assumed to be the repo from which to read the specified path.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The file source describing the local or remote Build template.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "github": {
        "block": {
          "attributes": {
            "enterprise_config_resource_name": {
              "description": "The resource name of the github enterprise config that should be applied to this installation.\nFor example: \"projects/{$projectId}/locations/{$locationId}/githubEnterpriseConfigs/{$configId}\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "Name of the repository. For example: The name for\nhttps://github.com/googlecloudplatform/cloud-builders is \"cloud-builders\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "owner": {
              "description": "Owner of the repository. For example: The owner for\nhttps://github.com/googlecloudplatform/cloud-builders is \"googlecloudplatform\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "pull_request": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "comment_control": {
                    "description": "Whether to block builds on a \"/gcbrun\" comment from a repository owner or collaborator. Possible values: [\"COMMENTS_DISABLED\", \"COMMENTS_ENABLED\", \"COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "If true, branches that do NOT match the git_ref will trigger a build.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "filter to match changes in pull requests. Specify only one of 'pull_request' or 'push'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "push": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.  Specify only one of branch or tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "When true, only trigger a build if the revision regex does NOT match the git_ref regex.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "Regex of tags to match.  Specify only one of branch or tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "filter to match changes in refs, like branches or tags. Specify only one of 'pull_request' or 'push'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Describes the configuration of a trigger that creates a build whenever a GitHub event is received.\n\nOne of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pubsub_config": {
        "block": {
          "attributes": {
            "service_account_email": {
              "description": "Service account that will make the push request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "Potential issues with the underlying Pub/Sub subscription configuration.\nOnly populated on get requests.",
              "description_kind": "plain",
              "type": "string"
            },
            "subscription": {
              "computed": true,
              "description": "Output only. Name of the subscription.",
              "description_kind": "plain",
              "type": "string"
            },
            "topic": {
              "description": "The name of the topic from which this subscription is receiving messages.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "PubsubConfig describes the configuration of a trigger that creates\na build whenever a Pub/Sub message is published.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "repository_event_config": {
        "block": {
          "attributes": {
            "repository": {
              "description": "The resource name of the Repo API resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "pull_request": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.\n\nThe syntax of the regular expressions accepted is the syntax accepted by\nRE2 and described at https://github.com/google/re2/wiki/Syntax",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "comment_control": {
                    "description": "Configure builds to run whether a repository owner or collaborator need to comment '/gcbrun'. Possible values: [\"COMMENTS_DISABLED\", \"COMMENTS_ENABLED\", \"COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "If true, branches that do NOT match the git_ref will trigger a build.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Contains filter properties for matching Pull Requests.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "push": {
              "block": {
                "attributes": {
                  "branch": {
                    "description": "Regex of branches to match.\n\nThe syntax of the regular expressions accepted is the syntax accepted by\nRE2 and described at https://github.com/google/re2/wiki/Syntax",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invert_regex": {
                    "description": "If true, only trigger a build if the revision regex does NOT match the git_ref regex.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "Regex of tags to match.\n\nThe syntax of the regular expressions accepted is the syntax accepted by\nRE2 and described at https://github.com/google/re2/wiki/Syntax",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Contains filter properties for matching git pushes.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration of a trigger that creates a build whenever an event from Repo API is received.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_to_build": {
        "block": {
          "attributes": {
            "bitbucket_server_config": {
              "description": "The full resource name of the bitbucket server config.\nFormat: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "github_enterprise_config": {
              "description": "The full resource name of the github enterprise config.\nFormat: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ref": {
              "description": "The branch or tag to use. Must start with \"refs/\" (required).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repo_type": {
              "description": "The type of the repo, since it may not be explicit from the repo field (e.g from a URL).\nValues can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: [\"UNKNOWN\", \"CLOUD_SOURCE_REPOSITORIES\", \"GITHUB\", \"BITBUCKET_SERVER\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repository": {
              "description": "The qualified resource name of the Repo API repository.\nEither uri or repository can be specified and is required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "description": "The URI of the repo.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The repo and ref of the repository from which to build.\nThis field is used only for those triggers that do not respond to SCM events.\nTriggers that respond to such events build source at whatever commit caused the event.\nThis field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
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
            "invert_regex": {
              "description": "Only trigger a build if the revision regex does NOT match the revision regex.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
          "description": "Template describing the types of source changes to trigger a build.\n\nBranch and tag names in trigger templates are interpreted as regular\nexpressions. Any branch or tag change that matches that regular\nexpression will trigger a build.\n\nOne of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "webhook_config": {
        "block": {
          "attributes": {
            "secret": {
              "description": "Resource name for the secret required as a URL parameter.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "Potential issues with the underlying Pub/Sub subscription configuration.\nOnly populated on get requests.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "WebhookConfig describes the configuration of a trigger that creates\na build whenever a webhook is sent to a trigger's webhook URL.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 2
}`

func GoogleCloudbuildTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudbuildTrigger), &result)
	return &result
}
