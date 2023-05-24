package data

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
        "type": "number"
      },
      "bigquery_config": {
        "computed": true,
        "description": "If delivery to BigQuery is used with this subscription, this field is used to configure it.\nEither pushConfig or bigQueryConfig can be set, but not both.\nIf both are empty, then the subscriber will pull and ack messages using API methods.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "drop_unknown_fields": "bool",
              "table": "string",
              "use_topic_schema": "bool",
              "write_metadata": "bool"
            }
          ]
        ]
      },
      "dead_letter_policy": {
        "computed": true,
        "description": "A policy that specifies the conditions for dead lettering messages in\nthis subscription. If dead_letter_policy is not set, dead lettering\nis disabled.\n\nThe Cloud Pub/Sub service account associated with this subscription's\nparent project (i.e.,\nservice-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have\npermission to Acknowledge() messages on this subscription.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dead_letter_topic": "string",
              "max_delivery_attempts": "number"
            }
          ]
        ]
      },
      "enable_exactly_once_delivery": {
        "computed": true,
        "description": "If 'true', Pub/Sub provides the following guarantees for the delivery\nof a message with a given value of messageId on this Subscriptions':\n\n- The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.\n\n- An acknowledged message will not be resent to a subscriber.\n\nNote that subscribers may still receive multiple copies of a message when 'enable_exactly_once_delivery'\nis true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_message_ordering": {
        "computed": true,
        "description": "If 'true', messages published with the same orderingKey in PubsubMessage will be delivered to\nthe subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they\nmay be delivered in any order.",
        "description_kind": "plain",
        "type": "bool"
      },
      "expiration_policy": {
        "computed": true,
        "description": "A policy that specifies the conditions for this subscription's expiration.\nA subscription is considered active as long as any connected subscriber\nis successfully consuming messages from the subscription or is issuing\noperations on the subscription. If expirationPolicy is not set, a default\npolicy with ttl of 31 days will be used.  If it is set but ttl is \"\", the\nresource never expires.  The minimum allowed value for expirationPolicy.ttl\nis 1 day.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ttl": "string"
            }
          ]
        ]
      },
      "filter": {
        "computed": true,
        "description": "The subscription only delivers the messages that match the filter.\nPub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages\nby their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,\nyou can't modify the filter.",
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
        "computed": true,
        "description": "A set of key/value label pairs to assign to this Subscription.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "message_retention_duration": {
        "computed": true,
        "description": "How long to retain unacknowledged messages in the subscription's\nbacklog, from the moment a message is published. If\nretain_acked_messages is true, then this also configures the retention\nof acknowledged messages, and thus configures how far back in time a\nsubscriptions.seek can be done. Defaults to 7 days. Cannot be more\nthan 7 days ('\"604800s\"') or less than 10 minutes ('\"600s\"').\n\nA duration in seconds with up to nine fractional digits, terminated\nby 's'. Example: '\"600.5s\"'.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the subscription.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "push_config": {
        "computed": true,
        "description": "If push delivery is used with this subscription, this field is used to\nconfigure it. An empty pushConfig signifies that the subscriber will\npull and ack messages using API methods.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attributes": [
                "map",
                "string"
              ],
              "oidc_token": [
                "list",
                [
                  "object",
                  {
                    "audience": "string",
                    "service_account_email": "string"
                  }
                ]
              ],
              "push_endpoint": "string"
            }
          ]
        ]
      },
      "retain_acked_messages": {
        "computed": true,
        "description": "Indicates whether to retain acknowledged messages. If 'true', then\nmessages are not expunged from the subscription's backlog, even if\nthey are acknowledged, until they fall out of the\nmessageRetentionDuration window.",
        "description_kind": "plain",
        "type": "bool"
      },
      "retry_policy": {
        "computed": true,
        "description": "A policy that specifies how Pub/Sub retries message delivery for this subscription.\n\nIf not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.\nRetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "maximum_backoff": "string",
              "minimum_backoff": "string"
            }
          ]
        ]
      },
      "topic": {
        "computed": true,
        "description": "A reference to a Topic resource.",
        "description_kind": "plain",
        "type": "string"
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
