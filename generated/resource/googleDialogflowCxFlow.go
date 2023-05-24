package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxFlow = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the flow.",
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
      "language_code": {
        "description": "The language of the following fields in flow:\nFlow.event_handlers.trigger_fulfillment.messages\nFlow.event_handlers.trigger_fulfillment.conditional_cases\nFlow.transition_routes.trigger_fulfillment.messages\nFlow.transition_routes.trigger_fulfillment.conditional_cases\nIf not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the flow.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a flow for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transition_route_groups": {
        "description": "A flow's transition route group serve two purposes:\nThey are responsible for matching the user's first utterances in the flow.\nThey are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.\nFormat:projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/transitionRouteGroups/\u003cTransitionRouteGroup ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "event_handlers": {
        "block": {
          "attributes": {
            "event": {
              "description": "The name of the event to handle.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The unique identifier of this event handler.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_flow": {
              "description": "The target flow to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_page": {
              "description": "The target page to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger_fulfillment": {
              "block": {
                "attributes": {
                  "return_partial_responses": {
                    "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "webhook": {
                    "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "messages": {
                    "block": {
                      "block_types": {
                        "text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "text": {
                                "description": "A collection of text responses.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The text response message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of rich message responses to present to the user.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A flow's event handlers serve two purposes:\nThey are responsible for handling events (e.g. no match, webhook errors) in the flow.\nThey are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.\nUnlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "nlu_settings": {
        "block": {
          "attributes": {
            "classification_threshold": {
              "description": "To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.\nIf the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "model_training_mode": {
              "description": "Indicates NLU model training mode.\n* MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.\n* MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: [\"MODEL_TRAINING_MODE_AUTOMATIC\", \"MODEL_TRAINING_MODE_MANUAL\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_type": {
              "description": "Indicates the type of NLU model.\n* MODEL_TYPE_STANDARD: Use standard NLU model.\n* MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: [\"MODEL_TYPE_STANDARD\", \"MODEL_TYPE_ADVANCED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "NLU related settings of the flow.",
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
      },
      "transition_routes": {
        "block": {
          "attributes": {
            "condition": {
              "description": "The condition to evaluate against form parameters or session parameters.\nAt least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "intent": {
              "description": "The unique identifier of an Intent.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/intents/\u003cIntent ID\u003e. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The unique identifier of this transition route.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_flow": {
              "description": "The target flow to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_page": {
              "description": "The target page to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger_fulfillment": {
              "block": {
                "attributes": {
                  "return_partial_responses": {
                    "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "webhook": {
                    "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "messages": {
                    "block": {
                      "block_types": {
                        "text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "text": {
                                "description": "A collection of text responses.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The text response message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of rich message responses to present to the user.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A flow's transition routes serve two purposes:\nThey are responsible for matching the user's first utterances in the flow.\nThey are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying \"help\" or \"can I talk to a human?\", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.\n\nTransitionRoutes are evalauted in the following order:\n  TransitionRoutes with intent specified.\n  TransitionRoutes with only condition specified.\n  TransitionRoutes with intent specified are inherited by pages in the flow.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowCxFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxFlow), &result)
	return &result
}
