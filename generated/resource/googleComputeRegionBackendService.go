package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionBackendService = `{
  "block": {
    "attributes": {
      "connection_draining_timeout_sec": {
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this\nobject. This field is used in optimistic locking.",
        "description_kind": "plain",
        "type": "string"
      },
      "health_checks": {
        "description": "The set of URLs to HealthCheck resources for health checking\nthis RegionBackendService. Currently at most one health\ncheck can be specified, and a health check is required.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_scheme": {
        "description": "Indicates what kind of load balancing this regional backend service\nwill be used for. A backend service created for one type of load\nbalancing cannot be used with the other(s). Must be 'INTERNAL' or\n'INTERNAL_MANAGED'. Defaults to 'INTERNAL'.",
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
      "protocol": {
        "computed": true,
        "description": "The protocol this RegionBackendService uses to communicate with backends.\nPossible values are HTTP, HTTPS, HTTP2, SSL, TCP, and UDP. The default is\nHTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer\ntypes and may result in errors if used with the GA API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created backend service should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_sec": {
        "computed": true,
        "description": "How many seconds to wait for the backend before considering it a\nfailed request. Default is 30 seconds. Valid range is [1, 86400].",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "backend": {
        "block": {
          "attributes": {
            "balancing_mode": {
              "description": "Specifies the balancing mode for this backend. Defaults to CONNECTION.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "capacity_scaler": {
              "computed": true,
              "description": "A multiplier applied to the group's maximum servicing capacity\n(based on UTILIZATION, RATE or CONNECTION).\n\nA setting of 0 means the group is completely drained, offering\n0% of its available Capacity. Valid range is [0.0,1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "description": {
              "description": "An optional description of this resource.\nProvide this property when you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group": {
              "description": "The fully-qualified URL of an Instance Group or Network Endpoint\nGroup resource. In case of instance group this defines the list\nof instances that serve traffic. Member virtual machine\ninstances from each instance group must live in the same zone as\nthe instance group itself. No two backends in a backend service\nare allowed to use same Instance Group resource.\n\nFor Network Endpoint Groups this defines list of endpoints. All\nendpoints of Network Endpoint Group must be hosted on instances\nlocated in the same zone as the Network Endpoint Group.\n\nBackend services cannot mix Instance Group and\nNetwork Endpoint Group backends.\n\nWhen the 'load_balancing_scheme' is INTERNAL, only instance groups\nare supported.\n\nNote that you must specify an Instance Group or Network Endpoint\nGroup resource using the fully-qualified URL, rather than a\npartial URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_connections": {
              "description": "The max number of simultaneous connections for the group. Can\nbe used with either CONNECTION or UTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or one\nof maxConnectionsPerInstance or maxConnectionsPerEndpoint,\nas appropriate for group type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_endpoint": {
              "description": "The max number of simultaneous connections that a single backend\nnetwork endpoint can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either\nmaxConnections or maxConnectionsPerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_instance": {
              "description": "The max number of simultaneous connections that a single\nbackend instance can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or\nmaxConnectionsPerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate": {
              "description": "The max requests per second (RPS) of the group.\n\nCan be used with either RATE or UTILIZATION balancing modes,\nbut required if RATE mode. Either maxRate or one\nof maxRatePerInstance or maxRatePerEndpoint, as appropriate for\ngroup type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_endpoint": {
              "description": "The max requests per second (RPS) that a single backend network\nendpoint can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_instance": {
              "description": "The max requests per second (RPS) that a single backend\ninstance can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_utilization": {
              "description": "Used when balancingMode is UTILIZATION. This ratio defines the\nCPU utilization target for the group. Valid range is [0.0, 1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
  "version": 1
}`

func GoogleComputeRegionBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionBackendService), &result)
	return &result
}
