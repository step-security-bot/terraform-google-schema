package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubLiteTopic = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the topic.",
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
        "description": "The region of the pubsub lite topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "description": "The zone of the pubsub lite topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "partition_config": {
        "block": {
          "attributes": {
            "count": {
              "description": "The number of partitions in the topic. Must be at least 1.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "capacity": {
              "block": {
                "attributes": {
                  "publish_mib_per_sec": {
                    "description": "Subscribe throughput capacity per partition in MiB/s. Must be \u003e= 4 and \u003c= 16.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "subscribe_mib_per_sec": {
                    "description": "Publish throughput capacity per partition in MiB/s. Must be \u003e= 4 and \u003c= 16.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The capacity configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The settings for this topic's partitions.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reservation_config": {
        "block": {
          "attributes": {
            "throughput_reservation": {
              "description": "The Reservation to use for this topic's throughput capacity.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The settings for this topic's Reservation usage.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_config": {
        "block": {
          "attributes": {
            "per_partition_bytes": {
              "description": "The provisioned storage, in bytes, per partition. If the number of bytes stored\nin any of the topic's partitions grows beyond this value, older messages will be\ndropped to make room for newer ones, regardless of the value of period.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "period": {
              "description": "How long a published message is retained. If unset, messages will be retained as\nlong as the bytes retained for each partition is below perPartitionBytes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The settings for a topic's message retention.",
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

func GooglePubsubLiteTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubLiteTopic), &result)
	return &result
}
