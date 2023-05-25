package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBinaryAuthorizationPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A descriptive comment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_policy_evaluation_mode": {
        "computed": true,
        "description": "Controls the evaluation of a Google-maintained global admission policy\nfor common system-level images. Images not covered by the global\npolicy will be subject to the project admission policy.",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "admission_whitelist_patterns": {
        "block": {
          "attributes": {
            "name_pattern": {
              "description": "An image name pattern to whitelist, in the form\n'registry/path/to/image'. This supports a trailing * as a\nwildcard, but this is allowed only in text after the registry/\npart.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "cluster_admission_rules": {
        "block": {
          "attributes": {
            "cluster": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enforcement_mode": {
              "description": "The action when a pod creation is denied by the admission rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "evaluation_mode": {
              "description": "How this admission rule will be evaluated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_attestations_by": {
              "description": "The resource names of the attestors that must attest to a\ncontainer image. If the attestor is in a different project from the\npolicy, it should be specified in the format 'projects/*/attestors/*'.\nEach attestor must exist before a policy can reference it. To add an\nattestor to a policy the principal issuing the policy change\nrequest must be able to read the attestor resource.\n\nNote: this field must be non-empty when the evaluation_mode field\nspecifies REQUIRE_ATTESTATION, otherwise it must be empty.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "default_admission_rule": {
        "block": {
          "attributes": {
            "enforcement_mode": {
              "description": "The action when a pod creation is denied by the admission rule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "evaluation_mode": {
              "description": "How this admission rule will be evaluated.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "require_attestations_by": {
              "description": "The resource names of the attestors that must attest to a\ncontainer image. If the attestor is in a different project from the\npolicy, it should be specified in the format 'projects/*/attestors/*'.\nEach attestor must exist before a policy can reference it. To add an\nattestor to a policy the principal issuing the policy change\nrequest must be able to read the attestor resource.\n\nNote: this field must be non-empty when the evaluation_mode field\nspecifies REQUIRE_ATTESTATION, otherwise it must be empty.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleBinaryAuthorizationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBinaryAuthorizationPolicy), &result)
	return &result
}
