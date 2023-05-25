package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringUptimeCheckConfig = `{
  "block": {
    "attributes": {
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
      "is_internal": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "http_check": {
        "block": {
          "attributes": {
            "headers": {
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
              "description": "The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to \"/\").",
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
            "auth_info": {
              "block": {
                "attributes": {
                  "password": {
                    "description": "The password to authenticate.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "The username to authenticate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "internal_checkers": {
        "block": {
          "attributes": {
            "display_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gcp_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "peer_project_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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
              "description": "The resource type of the group members.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
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
