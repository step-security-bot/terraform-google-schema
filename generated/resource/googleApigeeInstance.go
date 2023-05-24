package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeInstance = `{
  "block": {
    "attributes": {
      "consumer_accept_list": {
        "computed": true,
        "description": "Optional. Customer accept list represents the list of projects (id/number) on customer\nside that can privately connect to the service attachment. It is an optional field\nwhich the customers can provide during the instance creation. By default, the customer\nproject associated with the Apigee organization will be included to the list.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "Description of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_encryption_key_name": {
        "description": "Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.\nUse the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "computed": true,
        "description": "Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_range": {
        "description": "IP range represents the customer-provided CIDR block of length 22 that will be used for\nthe Apigee instance creation. This optional range, if provided, should be freely\navailable as part of larger named range the customer has allocated to the Service\nNetworking peering. If this is not provided, Apigee will automatically request for any\navailable /22 CIDR block from Service Networking. The customer should use this CIDR block\nfor configuring their firewall needs to allow traffic from Apigee.\nInput format: \"a.b.c.d/22\"",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Required. Compute Engine location where the instance resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Resource ID of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering_cidr_range": {
        "computed": true,
        "description": "The size of the CIDR block range that will be reserved by the instance. For valid values,\nsee [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "Output only. Port number of the exposed Apigee endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_attachment": {
        "computed": true,
        "description": "Output only. Resource name of the service attachment created for the instance in\nthe format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately\nforward traffic to this service attachment using the PSC endpoints.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleApigeeInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeInstance), &result)
	return &result
}
