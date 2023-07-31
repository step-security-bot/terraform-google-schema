package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesGateway = `{
  "block": {
    "attributes": {
      "addresses": {
        "description": "Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided,\nan IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'.\nGateways of type 'OPEN_MESH' listen on 0.0.0.0.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "certificate_urls": {
        "description": "A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection.\nThis feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Time the AccessPolicy was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_swg_autogen_router_on_destroy": {
        "description": "When deleting a gateway of type 'SECURE_WEB_GATEWAY', this boolean option will also delete auto generated router by the gateway creation.\nIf there is no other gateway of type 'SECURE_WEB_GATEWAY' remaining for that region and network it will be deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_security_policy": {
        "description": "A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections.\nFor example: 'projects/*/locations/*/gatewaySecurityPolicies/swg-policy'.\nThis policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.",
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
        "description": "Set of label tags associated with the Gateway resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the gateway.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Short name of the Gateway resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The relative resource name identifying the VPC network that is using this configuration.\nFor example: 'projects/*/global/networks/network-1'.\nCurrently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ports": {
        "description": "One or more port numbers (1-65535), on which the Gateway will receive traffic.\nThe proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are\nlimited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "number"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "Immutable. Scope determines how configuration across multiple Gateway instances are merged.\nThe configuration for multiple Gateway instances with the same scope will be merged as presented as\na single coniguration to the proxy/load balancer.\nMax length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_tls_policy": {
        "description": "A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated.\nIf empty, TLS termination is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnetwork": {
        "description": "The relative resource name identifying the subnetwork in which this SWG is allocated.\nFor example: 'projects/*/regions/us-central1/subnetworks/network-1'.\nCurrently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "Immutable. The type of the customer-managed gateway. Possible values are: * OPEN_MESH * SECURE_WEB_GATEWAY. Possible values: [\"TYPE_UNSPECIFIED\", \"OPEN_MESH\", \"SECURE_WEB_GATEWAY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the AccessPolicy was updated in UTC.",
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

func GoogleNetworkServicesGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesGateway), &result)
	return &result
}
