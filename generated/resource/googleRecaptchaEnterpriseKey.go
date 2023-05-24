package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleRecaptchaEnterpriseKey = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp corresponding to the creation of this Key.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Human-readable display name of this key. Modifiable by user.",
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
        "description": "See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The resource name for the Key in the format \"projects/{project}/keys/{key}\".",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "android_settings": {
        "block": {
          "attributes": {
            "allow_all_package_names": {
              "description": "If set to true, it means allowed_package_names will not be enforced.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allowed_package_names": {
              "description": "Android package names of apps allowed to use the key. Example: 'com.companyname.appname'",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Settings for keys that can be used by Android apps.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ios_settings": {
        "block": {
          "attributes": {
            "allow_all_bundle_ids": {
              "description": "If set to true, it means allowed_bundle_ids will not be enforced.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allowed_bundle_ids": {
              "description": "iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Settings for keys that can be used by iOS apps.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "testing_options": {
        "block": {
          "attributes": {
            "testing_challenge": {
              "computed": true,
              "description": "For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "testing_score": {
              "description": "All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Options for user acceptance testing.",
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
      "web_settings": {
        "block": {
          "attributes": {
            "allow_all_domains": {
              "description": "If set to true, it means allowed_domains will not be enforced.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_amp_traffic": {
              "description": "If set to true, the key can be used on AMP (Accelerated Mobile Pages) websites. This is supported only for the SCORE integration type.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allowed_domains": {
              "description": "Domains or subdomains of websites allowed to use the key. All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com'",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "challenge_security_preference": {
              "computed": true,
              "description": "Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED, USABILITY, BALANCE, SECURITY",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "integration_type": {
              "description": "Required. Describes how this key is integrated with the website. Possible values: SCORE, CHECKBOX, INVISIBLE",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Settings for keys that can be used by websites.",
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

func GoogleRecaptchaEnterpriseKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleRecaptchaEnterpriseKey), &result)
	return &result
}
