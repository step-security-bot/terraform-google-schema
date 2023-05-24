package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClouddeployDeliveryPipeline = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "condition": {
        "computed": true,
        "description": "Output only. Information around the state of the Delivery Pipeline.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pipeline_ready_condition": [
                "list",
                [
                  "object",
                  {
                    "status": "bool",
                    "update_time": "string"
                  }
                ]
              ],
              "targets_present_condition": [
                "list",
                [
                  "object",
                  {
                    "missing_targets": [
                      "list",
                      "string"
                    ],
                    "status": "bool",
                    "update_time": "string"
                  }
                ]
              ],
              "targets_type_condition": [
                "list",
                [
                  "object",
                  {
                    "error_details": "string",
                    "status": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Time at which the pipeline was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the ` + "`" + `DeliveryPipeline` + "`" + `. Max length is 255 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be \u003c= 128 bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the ` + "`" + `DeliveryPipeline` + "`" + `. Format is [a-z][a-z0-9\\-]{0,62}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suspended": {
        "description": "When suspended, no new releases or rollouts can be created, but in-progress ones will complete.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "uid": {
        "computed": true,
        "description": "Output only. Unique identifier of the ` + "`" + `DeliveryPipeline` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Most recent time at which the pipeline was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "serial_pipeline": {
        "block": {
          "block_types": {
            "stages": {
              "block": {
                "attributes": {
                  "profiles": {
                    "description": "Skaffold profiles to use when rendering the manifest for this stage's ` + "`" + `Target` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "target_id": {
                    "description": "The target_id to which this stage points. This field refers exclusively to the last segment of a target name. For example, this field would just be ` + "`" + `my-target` + "`" + ` (rather than ` + "`" + `projects/project/locations/location/targets/my-target` + "`" + `). The location of the ` + "`" + `Target` + "`" + ` is inferred to be the same as the location of the ` + "`" + `DeliveryPipeline` + "`" + ` that contains this ` + "`" + `Stage` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "strategy": {
                    "block": {
                      "block_types": {
                        "standard": {
                          "block": {
                            "attributes": {
                              "verify": {
                                "description": "Whether to verify a deployment.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Standard deployment strategy executes a single deploy and allows verifying the deployment.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Optional. The strategy to use for a ` + "`" + `Rollout` + "`" + ` to this stage.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Each stage specifies configuration for a ` + "`" + `Target` + "`" + `. The ordering of this list defines the promotion flow.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "SerialPipeline defines a sequential set of stages for a ` + "`" + `DeliveryPipeline` + "`" + `.",
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

func GoogleClouddeployDeliveryPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClouddeployDeliveryPipeline), &result)
	return &result
}
