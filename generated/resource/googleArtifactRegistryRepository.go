package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleArtifactRegistryRepository = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the repository was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The user-provided description of the repository.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "format": {
        "description": "The format of packages that are stored in the repository. Supported formats\ncan be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).\nYou can only create alpha formats if you are a member of the\n[alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).",
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
      "kms_key_name": {
        "description": "The Cloud KMS resource name of the customer managed encryption key that’s\nused to encrypt the contents of the Repository. Has the form:\n'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.\nThis value may not be changed after the Repository has been created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels with user-defined metadata.\nThis field may contain up to 64 entries. Label keys and values may be no\nlonger than 63 characters. Label keys must begin with a lowercase letter\nand may only contain lowercase letters, numeric characters, underscores,\nand dashes.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description": "The name of the location this repository is located in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mode": {
        "description": "The mode configures the repository to serve artifacts from different sources. Default value: \"STANDARD_REPOSITORY\" Possible values: [\"STANDARD_REPOSITORY\", \"VIRTUAL_REPOSITORY\", \"REMOTE_REPOSITORY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the repository, for example:\n\"repo1\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_id": {
        "description": "The last part of the repository name, for example:\n\"repo1\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the repository was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "docker_config": {
        "block": {
          "attributes": {
            "immutable_tags": {
              "description": "The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Docker repository config contains repository level configuration for the repositories of docker type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maven_config": {
        "block": {
          "attributes": {
            "allow_snapshot_overwrites": {
              "description": "The repository with this flag will allow publishing the same\nsnapshot versions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "version_policy": {
              "description": "Version policy defines the versions that the registry will accept. Default value: \"VERSION_POLICY_UNSPECIFIED\" Possible values: [\"VERSION_POLICY_UNSPECIFIED\", \"RELEASE\", \"SNAPSHOT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "MavenRepositoryConfig is maven related repository details.\nProvides additional configuration details for repositories of the maven\nformat type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "remote_repository_config": {
        "block": {
          "attributes": {
            "description": {
              "description": "The description of the remote source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "docker_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"DOCKER_HUB\" Possible values: [\"DOCKER_HUB\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specific settings for a Docker remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "maven_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"MAVEN_CENTRAL\" Possible values: [\"MAVEN_CENTRAL\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specific settings for a Maven remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "npm_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"NPMJS\" Possible values: [\"NPMJS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specific settings for an Npm remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "python_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"PYPI\" Possible values: [\"PYPI\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specific settings for a Python remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration specific for a Remote Repository.",
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
      "virtual_repository_config": {
        "block": {
          "block_types": {
            "upstream_policies": {
              "block": {
                "attributes": {
                  "id": {
                    "description": "The user-provided ID of the upstream policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "Entries with a greater priority value take precedence in the pull order.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "repository": {
                    "description": "A reference to the repository resource, for example:\n\"projects/p1/locations/us-central1/repository/repo1\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Policies that configure the upstream artifacts distributed by the Virtual\nRepository. Upstream policies cannot be set on a standard repository.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configuration specific for a Virtual Repository.",
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

func GoogleArtifactRegistryRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleArtifactRegistryRepository), &result)
	return &result
}
