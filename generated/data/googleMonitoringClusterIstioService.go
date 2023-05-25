package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringClusterIstioService = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description": "The name of the Kubernetes cluster in which this Istio service is defined. \n                        Corresponds to the clusterName resource label in k8s_cluster resources.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Name used for UI elements listing this Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the Kubernetes cluster in which this Istio service is defined. \n                        Corresponds to the location resource label in k8s_cluster resources.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The full resource name for this service. The syntax is:\nprojects/[PROJECT_ID]/services/[SERVICE_ID].",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description": "An optional service ID to use. If not given, the server will generate a\nservice ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "description": "The name of the Istio service underlying this service. \n                        Corresponds to the destination_service_name metric label in Istio metrics.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_namespace": {
        "description": "The namespace of the Istio service underlying this service. \n                        Corresponds to the destination_service_namespace metric label in Istio metrics.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "telemetry": {
        "computed": true,
        "description": "Configuration for how to query telemetry on a Service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMonitoringClusterIstioServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringClusterIstioService), &result)
	return &result
}
