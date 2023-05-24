package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiEndpoint = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when this Endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployed_models": {
        "computed": true,
        "description": "Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_resources": [
                "list",
                [
                  "object",
                  {
                    "max_replica_count": "number",
                    "min_replica_count": "number"
                  }
                ]
              ],
              "create_time": "string",
              "dedicated_resources": [
                "list",
                [
                  "object",
                  {
                    "autoscaling_metric_specs": [
                      "list",
                      [
                        "object",
                        {
                          "metric_name": "string",
                          "target": "number"
                        }
                      ]
                    ],
                    "machine_spec": [
                      "list",
                      [
                        "object",
                        {
                          "accelerator_count": "number",
                          "accelerator_type": "string",
                          "machine_type": "string"
                        }
                      ]
                    ],
                    "max_replica_count": "number",
                    "min_replica_count": "number"
                  }
                ]
              ],
              "display_name": "string",
              "enable_access_logging": "bool",
              "enable_container_logging": "bool",
              "id": "string",
              "model": "string",
              "model_version_id": "string",
              "private_endpoints": [
                "list",
                [
                  "object",
                  {
                    "explain_http_uri": "string",
                    "health_http_uri": "string",
                    "predict_http_uri": "string",
                    "service_attachment": "string"
                  }
                ]
              ],
              "service_account": "string",
              "shared_resources": "string"
            }
          ]
        ]
      },
      "description": {
        "description": "The description of the Endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_deployment_monitoring_job": {
        "computed": true,
        "description": "Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is network name.",
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
      "region": {
        "description": "The region for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Timestamp when this Endpoint was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: 'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'. The key needs to be in the same region as where the compute resource is created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.",
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

func GoogleVertexAiEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiEndpoint), &result)
	return &result
}
