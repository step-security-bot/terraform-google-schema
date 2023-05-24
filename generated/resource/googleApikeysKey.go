package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApikeysKey = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Human-readable display name of this API key. Modifiable by user.",
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
      "key_string": {
        "computed": true,
        "description": "Output only. An encrypted and signed value held by this key. This field can be accessed only through the ` + "`" + `GetKeyString` + "`" + ` method.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular expression: ` + "`" + `[a-z]([a-z0-9-]{0,61}[a-z0-9])?` + "`" + `.",
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
      "uid": {
        "computed": true,
        "description": "Output only. Unique id in UUID4 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "restrictions": {
        "block": {
          "block_types": {
            "android_key_restrictions": {
              "block": {
                "block_types": {
                  "allowed_applications": {
                    "block": {
                      "attributes": {
                        "package_name": {
                          "description": "The package name of the application.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "sha1_fingerprint": {
                          "description": "The SHA1 fingerprint of the application. For example, both sha1 formats are acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format is the latter.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of Android applications that are allowed to make API calls with this key.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Android apps that are allowed to use the key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "api_targets": {
              "block": {
                "attributes": {
                  "methods": {
                    "description": "Optional. List of one or more methods that can be called. If empty, all methods for the service are allowed. A wildcard (*) can be used as the last symbol. Valid examples: ` + "`" + `google.cloud.translate.v2.TranslateService.GetSupportedLanguage` + "`" + ` ` + "`" + `TranslateText` + "`" + ` ` + "`" + `Get*` + "`" + ` ` + "`" + `translate.googleapis.com.Get*` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service": {
                    "description": "The service for this restriction. It should be the canonical service name, for example: ` + "`" + `translate.googleapis.com` + "`" + `. You can use ` + "`" + `gcloud services list` + "`" + ` to get a list of services that are enabled in the project.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A restriction for a specific service and optionally one or more specific methods. Requests are allowed if they match any of these restrictions. If no restrictions are specified, all targets are allowed.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "browser_key_restrictions": {
              "block": {
                "attributes": {
                  "allowed_referrers": {
                    "description": "A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The HTTP referrers (websites) that are allowed to use the key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ios_key_restrictions": {
              "block": {
                "attributes": {
                  "allowed_bundle_ids": {
                    "description": "A list of bundle IDs that are allowed when making API calls with this key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The iOS apps that are allowed to use the key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "server_key_restrictions": {
              "block": {
                "attributes": {
                  "allowed_ips": {
                    "description": "A list of the caller IP addresses that are allowed to make API calls with this key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The IP addresses of callers that are allowed to use the key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Key restrictions.",
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

func GoogleApikeysKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApikeysKey), &result)
	return &result
}
