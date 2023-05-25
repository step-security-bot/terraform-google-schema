package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineStandardAppVersion = `{
  "block": {
    "attributes": {
      "delete_service_on_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "env_variables": {
        "description": "Environment variables available to the application.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_class": {
        "description": "Instance class that is used to run this version. Valid values are\nAutomaticScaling F1, F2, F4, F4_1G\n(Only AutomaticScaling is supported at the moment)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Full path to the Version resource in the API. Example, \"v1\".",
        "description_kind": "plain",
        "type": "string"
      },
      "noop_on_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "description": "Desired runtime. Example python27.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_api_version": {
        "description": "The version of the API in the given runtime environment. \nPlease see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "AppEngine service resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threadsafe": {
        "description": "Whether multiple requests can be dispatched to this version at once.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "version_id": {
        "description": "Relative name of the version within the service. For example, 'v1'. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,\"default\", \"latest\", and any name with the prefix \"ah-\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "deployment": {
        "block": {
          "block_types": {
            "files": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sha1_sum": {
                    "description": "SHA1 checksum of the file",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_url": {
                    "description": "Source URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "zip": {
              "block": {
                "attributes": {
                  "files_count": {
                    "description": "files count",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "source_url": {
                    "description": "Source URL",
                    "description_kind": "plain",
                    "optional": true,
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
      "entrypoint": {
        "block": {
          "attributes": {
            "shell": {
              "description": "The format should be a shell command that can be fed to bash -c.",
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
      "handlers": {
        "block": {
          "attributes": {
            "auth_fail_action": {
              "description": "Actions to take when the user is not logged in.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "login": {
              "description": "Methods to restrict access to a URL based on login status.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_http_response_code": {
              "description": "Redirect codes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_level": {
              "description": "Security (HTTPS) enforcement for this URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_regex": {
              "description": "URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings. \nAll URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "script": {
              "block": {
                "attributes": {
                  "script_path": {
                    "description": "Path to the script from the application root directory.",
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
            "static_files": {
              "block": {
                "attributes": {
                  "application_readable": {
                    "description": "Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged against both your code and static data storage resource quotas.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expiration": {
                    "description": "Time a static file served by this handler should be cached by web proxies and browsers.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_headers": {
                    "description": "HTTP headers to use for all responses from these URLs.\nAn object containing a list of \"key:value\" value pairs.\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "mime_type": {
                    "description": "MIME type used to serve all files served by this handler.\nDefaults to file-specific MIME types, which are derived from each file's filename extension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description": "Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_matching_file": {
                    "description": "Whether this handler should match the request if the file referenced by the handler does not exist.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "upload_path_regex": {
                    "description": "Regular expression that matches the file paths for all files that should be referenced by this handler.",
                    "description_kind": "plain",
                    "optional": true,
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
        "nesting_mode": "list"
      },
      "libraries": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the library. Example \"django\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "Version of the library to select, or \"latest\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
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
  "version": 0
}`

func GoogleAppEngineStandardAppVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineStandardAppVersion), &result)
	return &result
}
