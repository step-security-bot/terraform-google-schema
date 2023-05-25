package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubSubscription = `{
  "block": {
    "attributes": {
      "ack_deadline_seconds": {
        "computed": true,
        "description": "This value is the maximum time after a subscriber receives a message\nbefore the subscriber should acknowledge the message. After message\ndelivery but before the ack deadline expires and before the message is\nacknowledged, it is an outstanding message and will not be delivered\nagain during that time (on a best-effort basis).\n\nFor pull subscriptions, this value is used as the initial value for\nthe ack deadline. To override this value for a given message, call\nsubscriptions.modifyAckDeadline with the corresponding ackId if using\npull. The minimum custom deadline you can specify is 10 seconds. The\nmaximum custom deadline you can specify is 600 seconds (10 minutes).\nIf this parameter is 0, a default value of 10 seconds is used.\n\nFor push delivery, this value is also used to set the request timeout\nfor the call to the push endpoint.\n\nIf the subscriber never acknowledges the message, the Pub/Sub system\nwill eventually redeliver the message.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to this Subscription.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "message_retention_duration": {
        "description": "How long to retain unacknowledged messages in the subscription's\nbacklog, from the moment a message is published. If\nretainAckedMessages is true, then this also configures the retention\nof acknowledged messages, and thus configures how far back in time a\nsubscriptions.seek can be done. Defaults to 7 days. Cannot be more\nthan 7 days ('\"604800s\"') or less than 10 minutes ('\"600s\"').\n\nA duration in seconds with up to nine fractional digits, terminated\nby 's'. Example: '\"600.5s\"'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the subscription.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retain_acked_messages": {
        "description": "Indicates whether to retain acknowledged messages. If 'true', then\nmessages are not expunged from the subscription's backlog, even if\nthey are acknowledged, until they fall out of the\nmessageRetentionDuration window.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "topic": {
        "description": "A reference to a Topic resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "expiration_policy": {
        "block": {
          "attributes": {
            "ttl": {
              "description": "Specifies the \"time-to-live\" duration for an associated resource. The\nresource expires if it is not active for a period of ttl.\nIf ttl is not set, the associated resource never expires.\nA duration in seconds with up to nine fractional digits, terminated by 's'.\nExample - \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "push_config": {
        "block": {
          "attributes": {
            "attributes": {
              "description": "Endpoint configuration attributes.\n\nEvery endpoint has a set of API supported attributes that can\nbe used to control different aspects of the message delivery.\n\nThe currently supported attribute is x-goog-version, which you\ncan use to change the format of the pushed message. This\nattribute indicates the version of the data expected by\nthe endpoint. This controls the shape of the pushed message\n(i.e., its fields and metadata). The endpoint version is\nbased on the version of the Pub/Sub API.\n\nIf not present during the subscriptions.create call,\nit will default to the version of the API used to make\nsuch call. If not present during a subscriptions.modifyPushConfig\ncall, its value will not be changed. subscriptions.get\ncalls will always return a valid version, even if the\nsubscription was created without this attribute.\n\nThe possible values for this attribute are:\n\n- v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.\n- v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "push_endpoint": {
              "description": "A URL locating the endpoint to which messages should be pushed.\nFor example, a Webhook endpoint might use\n\"https://example.com/push\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "oidc_token": {
              "block": {
                "attributes": {
                  "audience": {
                    "description": "Audience to be used when generating OIDC token. The audience claim\nidentifies the recipients that the JWT is intended for. The audience\nvalue is a single case-sensitive string. Having multiple values (array)\nfor the audience field is not supported. More info about the OIDC JWT\ntoken audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3\nNote: if not specified, the Push endpoint URL will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "Service account email to be used for generating the OIDC token.\nThe caller (for subscriptions.create, subscriptions.patch, and\nsubscriptions.modifyPushConfig RPCs) must have the\niam.serviceAccounts.actAs permission for the service account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
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

func GooglePubsubSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubSubscription), &result)
	return &result
}
