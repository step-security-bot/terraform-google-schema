package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcTrigger = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.",
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
        "description": "Optional. User labels attached to the triggers that can be used to group resources.",
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
        "description": "Required. The resource name of the trigger. Must be unique within the location on the project and must be in ` + "`" + `projects/{project}/locations/{location}/triggers/{trigger}` + "`" + ` format.",
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
      "service_account": {
        "description": "Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have ` + "`" + `iam.serviceAccounts.actAs` + "`" + ` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have ` + "`" + `roles/eventarc.eventReceiver` + "`" + ` IAM role.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last-modified time.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "destination": {
        "block": {
          "attributes": {
            "cloud_function": {
              "description": "The Cloud Function resource name. Only Cloud Functions V2 is supported. Format: projects/{project}/locations/{location}/functions/{function}",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloud_run_service": {
              "block": {
                "attributes": {
                  "path": {
                    "description": "Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: \"/route\", \"route\", \"route/subroute\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "Required. The region the Cloud Run service is deployed in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service": {
                    "description": "Required. The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project of the trigger object can be addressed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. Destination specifies where the events should be sent to.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "matching_criteria": {
        "block": {
          "attributes": {
            "attribute": {
              "description": "Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Required. The value for the attribute.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.",
          "description_kind": "plain"
        },
        "min_items": 1,
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
      },
      "transport": {
        "block": {
          "block_types": {
            "pubsub": {
              "block": {
                "attributes": {
                  "subscription": {
                    "computed": true,
                    "description": "Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: ` + "`" + `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}` + "`" + `.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "topic": {
                    "description": "Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: ` + "`" + `projects/{PROJECT_ID}/topics/{TOPIC_NAME You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished` + "`" + ` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleEventarcTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcTrigger), &result)
	return &result
}
