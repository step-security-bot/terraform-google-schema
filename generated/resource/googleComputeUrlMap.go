package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeUrlMap = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_service": {
        "description": "The backend service or backend bucket to use when none of the given rules match.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. This field is used internally during\nupdates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "map_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "host_rule": {
        "block": {
          "attributes": {
            "description": {
              "description": "An optional description of this HostRule. Provide this property\nwhen you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hosts": {
              "description": "The list of host patterns to match. They must be valid\nhostnames, except * will match any string of ([a-z0-9-.]*). In\nthat case, * must be the first character and must be followed in\nthe pattern by either - or ..",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "path_matcher": {
              "description": "The name of the PathMatcher to use to match the path portion of\nthe URL if the hostRule matches the URL's host portion.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "path_matcher": {
        "block": {
          "attributes": {
            "default_service": {
              "description": "The backend service or backend bucket to use when none of the given paths match.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The name to which this PathMatcher is referred by the HostRule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "path_rule": {
              "block": {
                "attributes": {
                  "paths": {
                    "description": "The list of path patterns to match. Each must start with /\nand the only place a * is allowed is at the end following\na /. The string fed to the path matcher does not include\nany text after the first ? or #, and those chars are not\nallowed here.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "service": {
                    "description": "The backend service or backend bucket to use if any of the given paths match.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "test": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description of this test case.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host": {
              "description": "Host portion of the URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description": "Path portion of the URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service": {
              "description": "The backend service or backend bucket link that should be matched by this test.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func GoogleComputeUrlMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeUrlMap), &result)
	return &result
}
