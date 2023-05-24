package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkEndpointGroup = `{
  "block": {
    "attributes": {
      "app_engine": {
        "computed": true,
        "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "service": "string",
              "url_mask": "string",
              "version": "string"
            }
          ]
        ]
      },
      "cloud_function": {
        "computed": true,
        "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "function": "string",
              "url_mask": "string"
            }
          ]
        ]
      },
      "cloud_run": {
        "computed": true,
        "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "service": "string",
              "tag": "string",
              "url_mask": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "This field is only used for PSC.\nThe URL of the network to which all network endpoints in the NEG belong. Uses\n\"default\" project network if unspecified.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_endpoint_type": {
        "computed": true,
        "description": "Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: \"SERVERLESS\" Possible values: [\"SERVERLESS\", \"PRIVATE_SERVICE_CONNECT\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_target_service": {
        "computed": true,
        "description": "The target service url used to set up private service connection to\na Google API or a PSC Producer Service Attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "description": "A reference to the region where the Serverless NEGs Reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "This field is only used for PSC.\nOptional URL of the subnetwork to which all network endpoints in the NEG belong.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRegionNetworkEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkEndpointGroup), &result)
	return &result
}
