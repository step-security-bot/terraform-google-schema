package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcareHl7V2Store = `{
  "block": {
    "attributes": {
      "dataset": {
        "description": "Identifies the dataset addressed by this request. Must be in the format\n'projects/{project}/locations/{location}/datasets/{dataset}'",
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
      "labels": {
        "description": "User-supplied key-value pairs used to organize HL7v2 stores.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must\nconform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128\nbytes, and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be associated with a given store.\n\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The resource name for the Hl7V2Store.\n\n** Changing this property may recreate the Hl7v2 store (removing all data) **",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The fully qualified name of this dataset",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "notification_config": {
        "block": {
          "attributes": {
            "pubsub_topic": {
              "description": "The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.\nPubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.\nIt is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message\nwas published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a\nproject. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given\nCloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A nested object resource",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_configs": {
        "block": {
          "attributes": {
            "filter": {
              "description": "Restricts notifications sent for messages matching a filter. If this is empty, all messages\nare matched. Syntax: https://cloud.google.com/appengine/docs/standard/python/search/query_strings\n\nFields/functions available for filtering are:\n\n* messageType, from the MSH-9.1 field. For example, NOT messageType = \"ADT\".\n* send_date or sendDate, the YYYY-MM-DD date the message was sent in the dataset's timeZone, from the MSH-7 segment. For example, send_date \u003c \"2017-01-02\".\n* sendTime, the timestamp when the message was sent, using the RFC3339 time format for comparisons, from the MSH-7 segment. For example, sendTime \u003c \"2017-01-02T00:00:00-05:00\".\n* sendFacility, the care center that the message came from, from the MSH-4 segment. For example, sendFacility = \"ABC\".\n* PatientId(value, type), which matches if the message lists a patient having an ID of the given value and type in the PID-2, PID-3, or PID-4 segments. For example, PatientId(\"123456\", \"MRN\").\n* labels.x, a string value of the label with key x as set using the Message.labels map. For example, labels.\"priority\"=\"high\". The operator :* can be used to assert the existence of a label. For example, labels.\"priority\":*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pubsub_topic": {
              "description": "The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.\nPubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.\nIt is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message\nwas published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a\nproject. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given\nCloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.\n\nIf a notification cannot be published to Cloud Pub/Sub, errors will be logged to Stackdriver",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of notification configs. Each configuration uses a filter to determine whether to publish a\nmessage (both Ingest \u0026 Create) on the corresponding notification destination. Only the message name\nis sent as part of the notification. Supplied by the client.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "parser_config": {
        "block": {
          "attributes": {
            "allow_null_header": {
              "description": "Determines whether messages with no header are allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "schema": {
              "description": "JSON encoded string for schemas used to parse messages in this\nstore if schematized parsing is desired.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "segment_terminator": {
              "description": "Byte(s) to be used as the segment terminator. If this is unset, '\\r' will be used as segment terminator.\n\nA base64-encoded string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "The version of the unschematized parser to be used when a custom 'schema' is not set. Default value: \"V1\" Possible values: [\"V1\", \"V2\", \"V3\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource",
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

func GoogleHealthcareHl7V2StoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcareHl7V2Store), &result)
	return &result
}
