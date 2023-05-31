package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionCommitment = `{
  "block": {
    "attributes": {
      "auto_renew": {
        "computed": true,
        "description": "Specifies whether to enable automatic renewal for the commitment.\nThe default value is false if not specified.\nIf the field is set to true, the commitment will be automatically renewed for either\none or three years according to the terms of the existing commitment.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "category": {
        "computed": true,
        "description": "The category of the commitment. Category MACHINE specifies commitments composed of\nmachine resources such as VCPU or MEMORY, listed in resources. Category LICENSE\nspecifies commitments composed of software licenses, listed in licenseResources.\nNote that only MACHINE commitments should have a Type specified. Possible values: [\"LICENSE\", \"MACHINE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "commitment_id": {
        "computed": true,
        "description": "Unique identifier for the resource.",
        "description_kind": "plain",
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
      "end_timestamp": {
        "computed": true,
        "description": "Commitment end time in RFC3339 text format.",
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
        "description": "Name of the resource. The name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan": {
        "description": "The plan for this commitment, which determines duration and discount rate.\nThe currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years). Possible values: [\"TWELVE_MONTH\", \"THIRTY_SIX_MONTH\"]",
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
      "region": {
        "computed": true,
        "description": "URL of the region where this commitment may be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_timestamp": {
        "computed": true,
        "description": "Commitment start time in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the commitment with regards to eventual expiration\n(each commitment has an end date defined).",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "A human-readable explanation of the status.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of commitment, which affects the discount rate and the eligible resources.\nThe type could be one of the following value: 'MEMORY_OPTIMIZED', 'ACCELERATOR_OPTIMIZED',\n'GENERAL_PURPOSE_N1', 'GENERAL_PURPOSE_N2', 'GENERAL_PURPOSE_N2D', 'GENERAL_PURPOSE_E2',\n'GENERAL_PURPOSE_T2D', 'GENERAL_PURPOSE_C3', 'COMPUTE_OPTIMIZED_C2', 'COMPUTE_OPTIMIZED_C2D' and\n'GRAPHICS_OPTIMIZED_G2'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "license_resource": {
        "block": {
          "attributes": {
            "amount": {
              "description": "The number of licenses purchased.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cores_per_license": {
              "description": "Specifies the core range of the instance for which this license applies.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "license": {
              "description": "Any applicable license URI.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The license specification required as part of a license commitment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resources": {
        "block": {
          "attributes": {
            "accelerator_type": {
              "description": "Name of the accelerator type resource. Applicable only when the type is ACCELERATOR.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "amount": {
              "description": "The amount of the resource purchased (in a type-dependent unit,\nsuch as bytes). For vCPUs, this can just be an integer. For memory,\nthis must be provided in MB. Memory must be a multiple of 256 MB,\nwith up to 6.5GB of memory per every vCPU.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Type of resource for which this commitment applies.\nPossible values are VCPU, MEMORY, LOCAL_SSD, and ACCELERATOR.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A list of commitment amounts for particular resources.\nNote that VCPU and MEMORY resource commitments must occur together.",
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

func GoogleComputeRegionCommitmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionCommitment), &result)
	return &result
}
