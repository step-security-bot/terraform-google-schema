package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctionsFunction = `{
  "block": {
    "attributes": {
      "available_memory_mb": {
        "computed": true,
        "description": "Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.",
        "description_kind": "plain",
        "type": "number"
      },
      "build_environment_variables": {
        "computed": true,
        "description": " A set of key/value environment variable pairs available during build time.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "build_worker_pool": {
        "computed": true,
        "description": "Name of the Cloud Build Custom Worker Pool that should be used to build the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "docker_registry": {
        "computed": true,
        "description": "Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.",
        "description_kind": "plain",
        "type": "string"
      },
      "docker_repository": {
        "computed": true,
        "description": "User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry for storing images built with Cloud Build.",
        "description_kind": "plain",
        "type": "string"
      },
      "entry_point": {
        "computed": true,
        "description": "Name of the function that will be executed when the Google Cloud Function is triggered.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "A set of key/value environment variable pairs to assign to the function.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "event_trigger": {
        "computed": true,
        "description": "A source that fires events in response to a condition in another service. Cannot be used with trigger_http.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "event_type": "string",
              "failure_policy": [
                "list",
                [
                  "object",
                  {
                    "retry": "bool"
                  }
                ]
              ],
              "resource": "string"
            }
          ]
        ]
      },
      "https_trigger_security_level": {
        "computed": true,
        "description": "The security level for the function. Defaults to SECURE_OPTIONAL. Valid only if trigger_http is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "https_trigger_url": {
        "computed": true,
        "description": "URL which triggers function execution. Returned only if trigger_http is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingress_settings": {
        "computed": true,
        "description": "String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "max_instances": {
        "computed": true,
        "description": "The limit on the maximum number of function instances that may coexist at a given time.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_instances": {
        "computed": true,
        "description": "The limit on the minimum number of function instances that may coexist at a given time.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "A user-defined name of the function. Function names must be unique globally.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "Project of the function. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "Region of function. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description": "The runtime in which the function is going to run. Eg. \"nodejs8\", \"nodejs10\", \"python37\", \"go111\".",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_environment_variables": {
        "computed": true,
        "description": "Secret environment variables configuration",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key": "string",
              "project_id": "string",
              "secret": "string",
              "version": "string"
            }
          ]
        ]
      },
      "secret_volumes": {
        "computed": true,
        "description": "Secret volumes configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mount_path": "string",
              "project_id": "string",
              "secret": "string",
              "versions": [
                "list",
                [
                  "object",
                  {
                    "path": "string",
                    "version": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "service_account_email": {
        "computed": true,
        "description": " If provided, the self-provided service account to run the function with.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_archive_bucket": {
        "computed": true,
        "description": "The GCS bucket containing the zip archive which contains the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_archive_object": {
        "computed": true,
        "description": "The source archive object (file) in archive bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_repository": {
        "computed": true,
        "description": "Represents parameters related to source repository where a function is hosted. Cannot be set alongside source_archive_bucket or source_archive_object.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deployed_url": "string",
              "url": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description": "Describes the current stage of a deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout": {
        "computed": true,
        "description": "Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "trigger_http": {
        "computed": true,
        "description": "Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with trigger_bucket and trigger_topic.",
        "description_kind": "plain",
        "type": "bool"
      },
      "vpc_connector": {
        "computed": true,
        "description": "The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is projects/*/locations/*/connectors/*.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_connector_egress_settings": {
        "computed": true,
        "description": "The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudfunctionsFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctionsFunction), &result)
	return &result
}
