package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBillingBudget = `{
  "block": {
    "attributes": {
      "billing_account": {
        "description": "ID of the billing account to set a budget on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "User data for display name in UI. Must be \u003c= 60 chars.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Resource name of the budget. The resource name\nimplies the scope of a budget. Values are of the form\nbillingAccounts/{billingAccountId}/budgets/{budgetId}.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "all_updates_rule": {
        "block": {
          "attributes": {
            "disable_default_iam_recipients": {
              "description": "Boolean. When set to true, disables default notifications sent\nwhen a threshold is exceeded. Default recipients are\nthose with Billing Account Administrators and Billing\nAccount Users IAM roles for the target account.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "monitoring_notification_channels": {
              "description": "The full resource name of a monitoring notification\nchannel in the form\nprojects/{project_id}/notificationChannels/{channel_id}.\nA maximum of 5 channels are allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "pubsub_topic": {
              "description": "The name of the Cloud Pub/Sub topic where budget related\nmessages will be published, in the form\nprojects/{project_id}/topics/{topic_id}. Updates are sent\nat regular intervals to the topic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema_version": {
              "description": "The schema version of the notification. Only \"1.0\" is\naccepted. It represents the JSON schema as defined in\nhttps://cloud.google.com/billing/docs/how-to/budgets#notification_format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Defines notifications that are sent on every update to the\nbilling account's spend, regardless of the thresholds defined\nusing threshold rules.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "amount": {
        "block": {
          "attributes": {
            "last_period_amount": {
              "description": "Configures a budget amount that is automatically set to 100% of\nlast period's spend.\nBoolean. Set value to true to use. Do not set to false, instead\nuse the 'specified_amount' block.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "specified_amount": {
              "block": {
                "attributes": {
                  "currency_code": {
                    "computed": true,
                    "description": "The 3-letter currency code defined in ISO 4217.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "nanos": {
                    "description": "Number of nano (10^-9) units of the amount.\nThe value must be between -999,999,999 and +999,999,999\ninclusive. If units is positive, nanos must be positive or\nzero. If units is zero, nanos can be positive, zero, or\nnegative. If units is negative, nanos must be negative or\nzero. For example $-1.75 is represented as units=-1 and\nnanos=-750,000,000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "units": {
                    "description": "The whole units of the amount. For example if currencyCode\nis \"USD\", then 1 unit is one US dollar.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A specified amount to use as the budget. currencyCode is\noptional. If specified, it must match the currency of the\nbilling account. The currencyCode is provided on output.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The budgeted amount for each usage period.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "budget_filter": {
        "block": {
          "attributes": {
            "calendar_period": {
              "description": "A CalendarPeriod represents the abstract concept of a recurring time period that has a\ncanonical start. Grammatically, \"the start of the current CalendarPeriod\".\nAll calendar times begin at 12 AM US and Canadian Pacific Time (UTC-8).\n\nExactly one of 'calendar_period', 'custom_period' must be provided. Possible values: [\"MONTH\", \"QUARTER\", \"YEAR\", \"CALENDAR_PERIOD_UNSPECIFIED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "credit_types": {
              "computed": true,
              "description": "Optional. If creditTypesTreatment is INCLUDE_SPECIFIED_CREDITS,\nthis is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See a list of acceptable credit type values.\nIf creditTypesTreatment is not INCLUDE_SPECIFIED_CREDITS, this field must be empty.\n\n**Note:** If the field has a value in the config and needs to be removed, the field has to be an emtpy array in the config.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "credit_types_treatment": {
              "description": "Specifies how credits should be treated when determining spend\nfor threshold calculations. Default value: \"INCLUDE_ALL_CREDITS\" Possible values: [\"INCLUDE_ALL_CREDITS\", \"EXCLUDE_ALL_CREDITS\", \"INCLUDE_SPECIFIED_CREDITS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description": "A single label and value pair specifying that usage from only\nthis set of labeled resources should be included in the budget.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "projects": {
              "description": "A set of projects of the form projects/{project_number},\nspecifying that usage from only this set of projects should be\nincluded in the budget. If omitted, the report will include\nall usage for the billing account, regardless of which project\nthe usage occurred on.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "services": {
              "computed": true,
              "description": "A set of services of the form services/{service_id},\nspecifying that usage from only this set of services should be\nincluded in the budget. If omitted, the report will include\nusage for all the services. The service names are available\nthrough the Catalog API:\nhttps://cloud.google.com/billing/v1/how-tos/catalog-api.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subaccounts": {
              "computed": true,
              "description": "A set of subaccounts of the form billingAccounts/{account_id},\nspecifying that usage from only this set of subaccounts should\nbe included in the budget. If a subaccount is set to the name of\nthe parent account, usage from the parent account will be included.\nIf the field is omitted, the report will include usage from the parent\naccount and all subaccounts, if they exist.\n\n**Note:** If the field has a value in the config and needs to be removed, the field has to be an emtpy array in the config.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "custom_period": {
              "block": {
                "block_types": {
                  "end_date": {
                    "block": {
                      "attributes": {
                        "day": {
                          "description": "Day of a month. Must be from 1 to 31 and valid for the year and month.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "month": {
                          "description": "Month of a year. Must be from 1 to 12.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "year": {
                          "description": "Year of the date. Must be from 1 to 9999.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Optional. The end date of the time period. Budgets with elapsed end date won't be processed.\nIf unset, specifies to track all usage incurred since the startDate.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "start_date": {
                    "block": {
                      "attributes": {
                        "day": {
                          "description": "Day of a month. Must be from 1 to 31 and valid for the year and month.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "month": {
                          "description": "Month of a year. Must be from 1 to 12.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "year": {
                          "description": "Year of the date. Must be from 1 to 9999.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "A start date is required. The start date must be after January 1, 2017.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies to track usage from any start date (required) to any end date (optional).\nThis time period is static, it does not recur.\n\nExactly one of 'calendar_period', 'custom_period' must be provided.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Filters that define which resources are used to compute the actual\nspend against the budget.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "threshold_rules": {
        "block": {
          "attributes": {
            "spend_basis": {
              "description": "The type of basis used to determine if spend has passed\nthe threshold. Default value: \"CURRENT_SPEND\" Possible values: [\"CURRENT_SPEND\", \"FORECASTED_SPEND\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold_percent": {
              "description": "Send an alert when this threshold is exceeded. This is a\n1.0-based percentage, so 0.5 = 50%. Must be \u003e= 0.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Rules that trigger alerts (notifications of thresholds being\ncrossed) when spend exceeds the specified percentages of the\nbudget.",
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
  "version": 1
}`

func GoogleBillingBudgetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBillingBudget), &result)
	return &result
}
