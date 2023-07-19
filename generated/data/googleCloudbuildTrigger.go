package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudbuildTrigger = `{
  "block": {
    "attributes": {
      "approval_config": {
        "computed": true,
        "description": "Configuration for manual approval to start a build invocation of this BuildTrigger.\nBuilds created by this trigger will require approval before they execute.\nAny user with a Cloud Build Approver role for the project can approve a build.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "approval_required": "bool"
            }
          ]
        ]
      },
      "bitbucket_server_trigger_config": {
        "computed": true,
        "description": "BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bitbucket_server_config_resource": "string",
              "project_key": "string",
              "pull_request": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "comment_control": "string",
                    "invert_regex": "bool"
                  }
                ]
              ],
              "push": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "invert_regex": "bool",
                    "tag": "string"
                  }
                ]
              ],
              "repo_slug": "string"
            }
          ]
        ]
      },
      "build": {
        "computed": true,
        "description": "Contents of the build template. Either a filename or build template must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "artifacts": [
                "list",
                [
                  "object",
                  {
                    "images": [
                      "list",
                      "string"
                    ],
                    "objects": [
                      "list",
                      [
                        "object",
                        {
                          "location": "string",
                          "paths": [
                            "list",
                            "string"
                          ],
                          "timing": [
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
                      ]
                    ]
                  }
                ]
              ],
              "available_secrets": [
                "list",
                [
                  "object",
                  {
                    "secret_manager": [
                      "list",
                      [
                        "object",
                        {
                          "env": "string",
                          "version_name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "images": [
                "list",
                "string"
              ],
              "logs_bucket": "string",
              "options": [
                "list",
                [
                  "object",
                  {
                    "disk_size_gb": "number",
                    "dynamic_substitutions": "bool",
                    "env": [
                      "list",
                      "string"
                    ],
                    "log_streaming_option": "string",
                    "logging": "string",
                    "machine_type": "string",
                    "requested_verify_option": "string",
                    "secret_env": [
                      "list",
                      "string"
                    ],
                    "source_provenance_hash": [
                      "list",
                      "string"
                    ],
                    "substitution_option": "string",
                    "volumes": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "path": "string"
                        }
                      ]
                    ],
                    "worker_pool": "string"
                  }
                ]
              ],
              "queue_ttl": "string",
              "secret": [
                "list",
                [
                  "object",
                  {
                    "kms_key_name": "string",
                    "secret_env": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "source": [
                "list",
                [
                  "object",
                  {
                    "repo_source": [
                      "list",
                      [
                        "object",
                        {
                          "branch_name": "string",
                          "commit_sha": "string",
                          "dir": "string",
                          "invert_regex": "bool",
                          "project_id": "string",
                          "repo_name": "string",
                          "substitutions": [
                            "map",
                            "string"
                          ],
                          "tag_name": "string"
                        }
                      ]
                    ],
                    "storage_source": [
                      "list",
                      [
                        "object",
                        {
                          "bucket": "string",
                          "generation": "string",
                          "object": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "step": [
                "list",
                [
                  "object",
                  {
                    "allow_exit_codes": [
                      "list",
                      "number"
                    ],
                    "allow_failure": "bool",
                    "args": [
                      "list",
                      "string"
                    ],
                    "dir": "string",
                    "entrypoint": "string",
                    "env": [
                      "list",
                      "string"
                    ],
                    "id": "string",
                    "name": "string",
                    "script": "string",
                    "secret_env": [
                      "list",
                      "string"
                    ],
                    "timeout": "string",
                    "timing": "string",
                    "volumes": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "path": "string"
                        }
                      ]
                    ],
                    "wait_for": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "substitutions": [
                "map",
                "string"
              ],
              "tags": [
                "list",
                "string"
              ],
              "timeout": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Time when the trigger was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Human-readable description of the trigger.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description": "Whether the trigger is disabled or not. If true, the trigger will never result in a build.",
        "description_kind": "plain",
        "type": "bool"
      },
      "filename": {
        "computed": true,
        "description": "Path, from the source root, to a file whose contents is used for the template.\nEither a filename or build template must be provided. Set this only when using trigger_template or github.\nWhen using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter": {
        "computed": true,
        "description": "A Common Expression Language string. Used only with Pub/Sub and Webhook.",
        "description_kind": "plain",
        "type": "string"
      },
      "git_file_source": {
        "computed": true,
        "description": "The file source describing the local or remote Build template.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "github_enterprise_config": "string",
              "path": "string",
              "repo_type": "string",
              "repository": "string",
              "revision": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "github": {
        "computed": true,
        "description": "Describes the configuration of a trigger that creates a build whenever a GitHub event is received.\n\nOne of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enterprise_config_resource_name": "string",
              "name": "string",
              "owner": "string",
              "pull_request": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "comment_control": "string",
                    "invert_regex": "bool"
                  }
                ]
              ],
              "push": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "invert_regex": "bool",
                    "tag": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignored_files": {
        "computed": true,
        "description": "ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match\nextended with support for '**'.\n\nIf ignoredFiles and changed files are both empty, then they are not\nused to determine whether or not to trigger a build.\n\nIf ignoredFiles is not empty, then we ignore any files that match any\nof the ignored_file globs. If the change has no files that are outside\nof the ignoredFiles globs, then we do not trigger a build.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "include_build_logs": {
        "computed": true,
        "description": "Build logs will be sent back to GitHub as part of the checkrun\nresult.  Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or\nINCLUDE_BUILD_LOGS_WITH_STATUS Possible values: [\"INCLUDE_BUILD_LOGS_UNSPECIFIED\", \"INCLUDE_BUILD_LOGS_WITH_STATUS\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "included_files": {
        "computed": true,
        "description": "ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match\nextended with support for '**'.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is empty, then as far as this filter is concerned, we\nshould trigger the build.\n\nIf any of the files altered in the commit pass the ignoredFiles filter\nand includedFiles is not empty, then we make sure that at least one of\nthose files matches a includedFiles glob. If not, then we do not trigger\na build.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "description": "The [Cloud Build location](https://cloud.google.com/build/docs/locations) for the trigger.\nIf not specified, \"global\" is used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the trigger. Must be unique within the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pubsub_config": {
        "computed": true,
        "description": "PubsubConfig describes the configuration of a trigger that creates\na build whenever a Pub/Sub message is published.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "service_account_email": "string",
              "state": "string",
              "subscription": "string",
              "topic": "string"
            }
          ]
        ]
      },
      "repository_event_config": {
        "computed": true,
        "description": "The configuration of a trigger that creates a build whenever an event from Repo API is received.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pull_request": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "comment_control": "string",
                    "invert_regex": "bool"
                  }
                ]
              ],
              "push": [
                "list",
                [
                  "object",
                  {
                    "branch": "string",
                    "invert_regex": "bool",
                    "tag": "string"
                  }
                ]
              ],
              "repository": "string"
            }
          ]
        ]
      },
      "service_account": {
        "computed": true,
        "description": "The service account used for all user-controlled operations including\ntriggers.patch, triggers.run, builds.create, and builds.cancel.\n\nIf no service account is set, then the standard Cloud Build service account\n([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.\n\nFormat: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}",
        "description_kind": "plain",
        "type": "string"
      },
      "source_to_build": {
        "computed": true,
        "description": "The repo and ref of the repository from which to build.\nThis field is used only for those triggers that do not respond to SCM events.\nTriggers that respond to such events build source at whatever commit caused the event.\nThis field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "github_enterprise_config": "string",
              "ref": "string",
              "repo_type": "string",
              "repository": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "substitutions": {
        "computed": true,
        "description": "Substitutions data for Build resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags for annotation of a BuildTrigger",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "trigger_id": {
        "description": "The unique identifier for the trigger.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trigger_template": {
        "computed": true,
        "description": "Template describing the types of source changes to trigger a build.\n\nBranch and tag names in trigger templates are interpreted as regular\nexpressions. Any branch or tag change that matches that regular\nexpression will trigger a build.\n\nOne of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "branch_name": "string",
              "commit_sha": "string",
              "dir": "string",
              "invert_regex": "bool",
              "project_id": "string",
              "repo_name": "string",
              "tag_name": "string"
            }
          ]
        ]
      },
      "webhook_config": {
        "computed": true,
        "description": "WebhookConfig describes the configuration of a trigger that creates\na build whenever a webhook is sent to a trigger's webhook URL.\n\nOne of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "secret": "string",
              "state": "string"
            }
          ]
        ]
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
