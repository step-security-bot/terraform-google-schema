package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkManagementConnectivityTest = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The user-supplied description of the Connectivity Test.\nMaximum of 512 characters.",
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
      "labels": {
        "description": "Resource labels to represent user-provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Unique name for the connectivity test.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description": "IP Protocol of the test. When not provided, \"TCP\" is assumed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "related_projects": {
        "description": "Other projects that may be relevant for reachability analysis.\nThis is applicable to scenarios where a test can cross project\nboundaries.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "destination": {
        "block": {
          "attributes": {
            "instance": {
              "description": "A Compute Engine instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "description": "The IP address of the endpoint, which can be an external or\ninternal IP. An IPv6 address is only allowed when the test's\ndestination is a global load balancer VIP.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "A Compute Engine network URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The IP protocol port of the endpoint. Only applicable when\nprotocol is TCP or UDP.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "project_id": {
              "description": "Project ID where the endpoint is located. The Project ID can be\nderived from the URI if you provide a VM instance or network URI.\nThe following are two cases where you must provide the project ID:\n1. Only the IP address is specified, and the IP address is within\na GCP project. 2. When you are using Shared VPC and the IP address\nthat you provide is from the service project. In this case, the\nnetwork that the IP address resides in is defined in the host\nproject.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Required. Destination specification of the Connectivity Test.\n\nYou can use a combination of destination IP address, Compute\nEngine VM instance, or VPC network to uniquely identify the\ndestination location.\n\nEven if the destination IP address is not unique, the source IP\nlocation is unique. Usually, the analysis can infer the destination\nendpoint from route information.\n\nIf the destination you specify is a VM instance and the instance has\nmultiple network interfaces, then you must also specify either a\ndestination IP address or VPC network to identify the destination\ninterface.\n\nA reachability analysis proceeds even if the destination location\nis ambiguous. However, the result can include endpoints that you\ndon't intend to test.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source": {
        "block": {
          "attributes": {
            "instance": {
              "description": "A Compute Engine instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "description": "The IP address of the endpoint, which can be an external or\ninternal IP. An IPv6 address is only allowed when the test's\ndestination is a global load balancer VIP.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "A Compute Engine network URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_type": {
              "description": "Type of the network where the endpoint is located. Possible values: [\"GCP_NETWORK\", \"NON_GCP_NETWORK\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The IP protocol port of the endpoint. Only applicable when\nprotocol is TCP or UDP.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "project_id": {
              "description": "Project ID where the endpoint is located. The Project ID can be\nderived from the URI if you provide a VM instance or network URI.\nThe following are two cases where you must provide the project ID:\n\n1. Only the IP address is specified, and the IP address is\n   within a GCP project.\n2. When you are using Shared VPC and the IP address\n   that you provide is from the service project. In this case,\n   the network that the IP address resides in is defined in the\n   host project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Required. Source specification of the Connectivity Test.\n\nYou can use a combination of source IP address, virtual machine\n(VM) instance, or Compute Engine network to uniquely identify the\nsource location.\n\nExamples: If the source IP address is an internal IP address within\na Google Cloud Virtual Private Cloud (VPC) network, then you must\nalso specify the VPC network. Otherwise, specify the VM instance,\nwhich already contains its internal IP address and VPC network\ninformation.\n\nIf the source of the test is within an on-premises network, then\nyou must provide the destination VPC network.\n\nIf the source endpoint is a Compute Engine VM instance with multiple\nnetwork interfaces, the instance itself is not sufficient to\nidentify the endpoint. So, you must also specify the source IP\naddress or VPC network.\n\nA reachability analysis proceeds even if the source location is\nambiguous. However, the test result may include endpoints that\nyou don't intend to test.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleNetworkManagementConnectivityTestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkManagementConnectivityTest), &result)
	return &result
}
