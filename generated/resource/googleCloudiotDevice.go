package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudiotDevice = `{
  "block": {
    "attributes": {
      "blocked": {
        "description": "If a device is blocked, connections or requests from this device will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "config": {
        "computed": true,
        "description": "The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "binary_data": "string",
              "cloud_update_time": "string",
              "device_ack_time": "string",
              "version": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_config_ack_time": {
        "computed": true,
        "description": "The last time a cloud-to-device config version acknowledgment was received from the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_config_send_time": {
        "computed": true,
        "description": "The last time a cloud-to-device config version was sent to the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_error_status": {
        "computed": true,
        "description": "The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "details": [
                "list",
                [
                  "map",
                  "string"
                ]
              ],
              "message": "string",
              "number": "number"
            }
          ]
        ]
      },
      "last_error_time": {
        "computed": true,
        "description": "The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_event_time": {
        "computed": true,
        "description": "The last time a telemetry event was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_heartbeat_time": {
        "computed": true,
        "description": "The last time an MQTT PINGREQ was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_state_time": {
        "computed": true,
        "description": "The last time a state event was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_level": {
        "description": "The logging verbosity for device activity. Possible values: [\"NONE\", \"ERROR\", \"INFO\", \"DEBUG\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata": {
        "description": "The metadata key-value pairs assigned to the device.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "A unique name for the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "num_id": {
        "computed": true,
        "description": "A server-defined unique numeric ID for the device.\nThis is a more compact way to identify devices, and it is globally unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry": {
        "description": "The name of the device registry where this device should be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state most recently received from the device.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "binary_data": "string",
              "update_time": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "credentials": {
        "block": {
          "attributes": {
            "expiration_time": {
              "computed": true,
              "description": "The time at which this credential becomes invalid.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "public_key": {
              "block": {
                "attributes": {
                  "format": {
                    "description": "The format of the key. Possible values: [\"RSA_PEM\", \"RSA_X509_PEM\", \"ES256_PEM\", \"ES256_X509_PEM\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "The key data.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A public key used to verify the signature of JSON Web Tokens (JWTs).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The credentials used to authenticate this device.",
          "description_kind": "plain"
        },
        "max_items": 3,
        "nesting_mode": "list"
      },
      "gateway_config": {
        "block": {
          "attributes": {
            "gateway_auth_method": {
              "description": "Indicates whether the device is a gateway. Possible values: [\"ASSOCIATION_ONLY\", \"DEVICE_AUTH_TOKEN_ONLY\", \"ASSOCIATION_AND_DEVICE_AUTH_TOKEN\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway_type": {
              "description": "Indicates whether the device is a gateway. Default value: \"NON_GATEWAY\" Possible values: [\"GATEWAY\", \"NON_GATEWAY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_accessed_gateway_id": {
              "computed": true,
              "description": "The ID of the gateway the device accessed most recently.",
              "description_kind": "plain",
              "type": "string"
            },
            "last_accessed_gateway_time": {
              "computed": true,
              "description": "The most recent time at which the device accessed the gateway specified in last_accessed_gateway.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Gateway-related configuration and state.",
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

func GoogleCloudiotDeviceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudiotDevice), &result)
	return &result
}
