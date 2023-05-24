package data

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
        "computed": true,
        "description": "The user-provided description of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "docker_config": {
        "computed": true,
        "description": "Docker repository config contains repository level configuration for the repositories of docker type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "immutable_tags": "bool"
            }
          ]
        ]
      },
      "format": {
        "computed": true,
        "description": "The format of packages that are stored in the repository. Supported formats\ncan be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).\nYou can only create alpha formats if you are a member of the\n[alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "The Cloud KMS resource name of the customer managed encryption key that’s\nused to encrypt the contents of the Repository. Has the form:\n'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.\nThis value may not be changed after the Repository has been created.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels with user-defined metadata.\nThis field may contain up to 64 entries. Label keys and values may be no\nlonger than 63 characters. Label keys must begin with a lowercase letter\nand may only contain lowercase letters, numeric characters, underscores,\nand dashes.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location this repository is located in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maven_config": {
        "computed": true,
        "description": "MavenRepositoryConfig is maven related repository details.\nProvides additional configuration details for repositories of the maven\nformat type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_snapshot_overwrites": "bool",
              "version_policy": "string"
            }
          ]
        ]
      },
      "mode": {
        "computed": true,
        "description": "The mode configures the repository to serve artifacts from different sources. Default value: \"STANDARD_REPOSITORY\" Possible values: [\"STANDARD_REPOSITORY\", \"VIRTUAL_REPOSITORY\", \"REMOTE_REPOSITORY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the repository, for example:\n\"repo1\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_repository_config": {
        "computed": true,
        "description": "Configuration specific for a Remote Repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "docker_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": "string"
                  }
                ]
              ],
              "maven_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": "string"
                  }
                ]
              ],
              "npm_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": "string"
                  }
                ]
              ],
              "python_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
      },
      "virtual_repository_config": {
        "computed": true,
        "description": "Configuration specific for a Virtual Repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "upstream_policies": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "priority": "number",
                    "repository": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
