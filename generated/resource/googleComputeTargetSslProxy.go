package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeTargetSslProxy = `{
  "block": {
    "attributes": {
      "backend_service": {
        "description": "A reference to the BackendService resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "proxy_header": {
        "description": "Specifies the type of proxy header to append before sending data to\nthe backend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_certificates": {
        "description": "A list of SslCertificate resources that are used to authenticate\nconnections between users and the load balancer. At least one\nSSL certificate must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ssl_policy": {
        "description": "A reference to the SslPolicy resource that will be associated with\nthe TargetSslProxy resource. If not set, the TargetSslProxy\nresource will not have any SSL policy configured.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleComputeTargetSslProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeTargetSslProxy), &result)
	return &result
}
