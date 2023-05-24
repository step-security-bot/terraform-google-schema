package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcareDicomStore = `{
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
        "description": "User-supplied key-value pairs used to organize DICOM stores.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must\nconform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128\nbytes, and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be associated with a given store.\n\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The resource name for the DicomStore.\n\n** Changing this property may recreate the Dicom store (removing all data) **",
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

func GoogleHealthcareDicomStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcareDicomStore), &result)
	return &result
}
