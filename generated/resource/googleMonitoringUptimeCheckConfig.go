package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringUptimeCheckConfig = `{
  "block": {
    "attributes": {
      "checker_type": {
        "computed": true,
        "description": "The checker type to use for the check. If the monitored resource type is servicedirectory_service, checkerType must be set to VPC_CHECKERS. Possible values: [\"STATIC_IP_CHECKERS\", \"VPC_CHECKERS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A unique resource name for this UptimeCheckConfig. The format is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].",
        "description_kind": "plain",
        "type": "string"
      },
      "period": {
        "description": "How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.",
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
      "selected_regions": {
        "description": "The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "timeout": {
        "description": "The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uptime_check_id": {
        "computed": true,
        "description": "The id of the uptime check",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "content_matchers": {
        "block": {
          "attributes": {
            "content": {
              "description": "String or regex content to match (max 1024 bytes)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "matcher": {
              "description": "The type of content matcher that will be applied to the server output, compared to the content string when the check is run. Default value: \"CONTAINS_STRING\" Possible values: [\"CONTAINS_STRING\", \"NOT_CONTAINS_STRING\", \"MATCHES_REGEX\", \"NOT_MATCHES_REGEX\", \"MATCHES_JSON_PATH\", \"NOT_MATCHES_JSON_PATH\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "json_path_matcher": {
              "block": {
                "attributes": {
                  "json_matcher": {
                    "description": "Options to perform JSONPath content matching. Default value: \"EXACT_MATCH\" Possible values: [\"EXACT_MATCH\", \"REGEX_MATCH\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "json_path": {
                    "description": "JSONPath within the response output pointing to the expected 'ContentMatcher::content' to match against.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Information needed to perform a JSONPath content match. Used for 'ContentMatcherOption::MATCHES_JSON_PATH' and 'ContentMatcherOption::NOT_MATCHES_JSON_PATH'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "http_check": {
        "block": {
          "attributes": {
            "body": {
              "description": "The request body associated with the HTTP POST request. If contentType is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the requestMethod is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. \"foo=bar\" in URL-encoded form is \"foo%3Dbar\" and in base64 encoding is \"Zm9vJTI1M0RiYXI=\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type": {
              "description": "The content type to use for the check. Possible values: [\"TYPE_UNSPECIFIED\", \"URL_ENCODED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "computed": true,
              "description": "The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "mask_headers": {
              "description": "Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "path": {
              "description": "The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with \"/\", a \"/\" will be prepended automatically. Optional (defaults to \"/\").",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "request_method": {
              "description": "The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then requestMethod defaults to GET. Default value: \"GET\" Possible values: [\"METHOD_UNSPECIFIED\", \"GET\", \"POST\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_ssl": {
              "description": "If true, use HTTPS instead of HTTP to run the check.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "validate_ssl": {
              "description": "Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "accepted_response_status_codes": {
              "block": {
                "attributes": {
                  "status_class": {
                    "description": "A class of status codes to accept. Possible values: [\"STATUS_CLASS_1XX\", \"STATUS_CLASS_2XX\", \"STATUS_CLASS_3XX\", \"STATUS_CLASS_4XX\", \"STATUS_CLASS_5XX\", \"STATUS_CLASS_ANY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_value": {
                    "description": "A status code to accept.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "If present, the check will only pass if the HTTP response status code is in this set of status codes. If empty, the HTTP status code will only pass if the HTTP status code is 200-299.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "auth_info": {
              "block": {
                "attributes": {
                  "password": {
                    "description": "The password to authenticate.",
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "The username to authenticate.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The authentication information. Optional when creating an HTTP check; defaults to empty.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Contains information needed to make an HTTP or HTTPS check.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitored_resource": {
        "block": {
          "attributes": {
            "labels": {
              "description": "Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels \"project_id\", \"instance_id\", and \"zone\".",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "description": "The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance aws_elb_load_balancer  k8s_service  servicedirectory_service",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resource_group": {
        "block": {
          "attributes": {
            "group_id": {
              "description": "The group of resources being monitored. Should be the 'name' of a group",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_type": {
              "description": "The resource type of the group members. Possible values: [\"RESOURCE_TYPE_UNSPECIFIED\", \"INSTANCE\", \"AWS_ELB_LOAD_BALANCER\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The group resource associated with the configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "synthetic_monitor": {
        "block": {
          "block_types": {
            "cloud_function_v2": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The fully qualified name of the cloud function resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Target a Synthetic Monitor GCFv2 Instance",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A Synthetic Monitor deployed to a Cloud Functions V2 instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tcp_check": {
        "block": {
          "attributes": {
            "port": {
              "description": "The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Contains information needed to make a TCP check.",
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

func GoogleMonitoringUptimeCheckConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringUptimeCheckConfig), &result)
	return &result
}
