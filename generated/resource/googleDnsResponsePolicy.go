package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsResponsePolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description of the response policy, such as 'My new response policy'.",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "response_policy_name": {
        "description": "The user assigned name for this Response Policy, such as 'myresponsepolicy'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "gke_clusters": {
        "block": {
          "attributes": {
            "gke_cluster_name": {
              "description": "The resource name of the cluster to bind this ManagedZone to.\nThis should be specified in the format like\n'projects/*/locations/*/clusters/*'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of Google Kubernetes Engine clusters that can see this zone.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "networks": {
        "block": {
          "attributes": {
            "network_url": {
              "description": "The fully qualified URL of the VPC network to bind to.\nThis should be formatted like\n'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of network names specifying networks to which this policy is applied.",
          "description_kind": "plain"
        },
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

func GoogleDnsResponsePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsResponsePolicy), &result)
	return &result
}
