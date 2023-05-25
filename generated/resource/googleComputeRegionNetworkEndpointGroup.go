package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkEndpointGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_endpoint_type": {
        "description": "Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: \"SERVERLESS\" Possible values: [\"SERVERLESS\"]",
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
        "description": "A reference to the region where the Serverless NEGs Reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "app_engine": {
        "block": {
          "attributes": {
            "service": {
              "description": "Optional serving service.\nThe service name must be 1-63 characters long, and comply with RFC1035.\nExample value: \"default\", \"my-service\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_mask": {
              "description": "A template to parse service and version fields from a request URL.\nURL mask allows for routing to multiple App Engine services without\nhaving to create multiple Network Endpoint Groups and backend services.\n\nFor example, the request URLs \"foo1-dot-appname.appspot.com/v1\" and\n\"foo1-dot-appname.appspot.com/v2\" can be backed by the same Serverless NEG with\nURL mask \"-dot-appname.appspot.com/\". The URL mask will parse\nthem to { service = \"foo1\", version = \"v1\" } and { service = \"foo1\", version = \"v2\" } respectively.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "Optional serving version.\nThe version must be 1-63 characters long, and comply with RFC1035.\nExample value: \"v1\", \"v2\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine or cloud_function may be set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloud_function": {
        "block": {
          "attributes": {
            "function": {
              "description": "A user-defined name of the Cloud Function.\nThe function name is case-sensitive and must be 1-63 characters long.\nExample value: \"func1\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_mask": {
              "description": "A template to parse function field from a request URL. URL mask allows\nfor routing to multiple Cloud Functions without having to create\nmultiple Network Endpoint Groups and backend services.\n\nFor example, request URLs \"mydomain.com/function1\" and \"mydomain.com/function2\"\ncan be backed by the same Serverless NEG with URL mask \"/\". The URL mask\nwill parse them to { function = \"function1\" } and { function = \"function2\" } respectively.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine or cloud_function may be set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloud_run": {
        "block": {
          "attributes": {
            "service": {
              "description": "Cloud Run service is the main resource of Cloud Run.\nThe service must be 1-63 characters long, and comply with RFC1035.\nExample value: \"run-service\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag": {
              "description": "Cloud Run tag represents the \"named-revision\" to provide\nadditional fine-grained traffic routing information.\nThe tag must be 1-63 characters long, and comply with RFC1035.\nExample value: \"revision-0010\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_mask": {
              "description": "A template to parse service and tag fields from a request URL.\nURL mask allows for routing to multiple Run services without having\nto create multiple network endpoint groups and backend services.\n\nFor example, request URLs \"foo1.domain.com/bar1\" and \"foo1.domain.com/bar2\"\nan be backed by the same Serverless Network Endpoint Group (NEG) with\nURL mask \".domain.com/\". The URL mask will parse them to { service=\"bar1\", tag=\"foo1\" }\nand { service=\"bar2\", tag=\"foo2\" } respectively.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Only valid when networkEndpointType is \"SERVERLESS\".\nOnly one of cloud_run, app_engine or cloud_function may be set.",
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

func GoogleComputeRegionNetworkEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkEndpointGroup), &result)
	return &result
}
