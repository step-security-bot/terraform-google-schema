package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeTargetHttpsProxy = `{
  "block": {
    "attributes": {
      "certificate_map": {
        "description": "A reference to the CertificateMap resource uri that identifies a certificate map\nassociated with the given target proxy. This field can only be set for global target proxies.\nAccepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "proxy_bind": {
        "computed": true,
        "description": "This field only applies when the forwarding rule that references\nthis target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "proxy_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "quic_override": {
        "description": "Specifies the QUIC override policy for this resource. This determines\nwhether the load balancer will attempt to negotiate QUIC with clients\nor not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is\nspecified, Google manages whether QUIC is used. Default value: \"NONE\" Possible values: [\"NONE\", \"ENABLE\", \"DISABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_certificates": {
        "description": "A list of SslCertificate resources that are used to authenticate\nconnections between users and the load balancer. At least one SSL\ncertificate must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ssl_policy": {
        "description": "A reference to the SslPolicy resource that will be associated with\nthe TargetHttpsProxy resource. If not set, the TargetHttpsProxy\nresource will not have any SSL policy configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "url_map": {
        "description": "A reference to the UrlMap resource that defines the mapping from URL\nto the BackendService.",
        "description_kind": "plain",
        "required": true,
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

func GoogleComputeTargetHttpsProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeTargetHttpsProxy), &result)
	return &result
}
