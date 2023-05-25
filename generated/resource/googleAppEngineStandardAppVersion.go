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
      "inbound_services": {
        "description": "A list of the types of messages that this application is able to receive. Possible values: [\"INBOUND_SERVICE_MAIL\", \"INBOUND_SERVICE_MAIL_BOUNCE\", \"INBOUND_SERVICE_XMPP_ERROR\", \"INBOUND_SERVICE_XMPP_MESSAGE\", \"INBOUND_SERVICE_XMPP_SUBSCRIBE\", \"INBOUND_SERVICE_XMPP_PRESENCE\", \"INBOUND_SERVICE_CHANNEL_PRESENCE\", \"INBOUND_SERVICE_WARMUP\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "instance_class": {
        "computed": true,
        "description": "Instance class that is used to run this version. Valid values are\nAutomaticScaling: F1, F2, F4, F4_1G\nBasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8\nDefaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.",
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
        "description": "The version of the API in the given runtime environment.\nPlease see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "AppEngine service resource",
        "description_kind": "plain",
        "required": true,
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
      "automatic_scaling": {
        "block": {
          "attributes": {
            "max_concurrent_requests": {
              "description": "Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.\n\nDefaults to a runtime-specific value.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_idle_instances": {
              "description": "Maximum number of idle instances that should be maintained for this version.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_pending_latency": {
              "description": "Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_idle_instances": {
              "description": "Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_pending_latency": {
              "description": "Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "standard_scheduler_settings": {
              "block": {
                "attributes": {
                  "max_instances": {
                    "description": "Maximum number of instances to run for this version. Set to zero to disable maxInstances configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_instances": {
                    "description": "Minimum number of instances to run for this version. Set to zero to disable minInstances configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_cpu_utilization": {
                    "description": "Target CPU utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_throughput_utilization": {
                    "description": "Target throughput utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Scheduler settings for standard environment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Automatic scaling is based on request rate, response latencies, and other application metrics.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "basic_scaling": {
        "block": {
          "attributes": {
            "idle_timeout": {
              "description": "Duration of time after the last request that an instance must wait before the instance is shut down.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\". Defaults to 900s.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_instances": {
              "description": "Maximum number of instances to create for this version. Must be in the range [1.0, 200.0].",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Manifest of the files stored in Google Cloud Storage that are included as part of this version.\nAll files must be readable using the credentials supplied with this call.",
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
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Zip File",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Code and application artifacts that make up this version.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "entrypoint": {
        "block": {
          "attributes": {
            "shell": {
              "description": "The format should be a shell command that can be fed to bash -c.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The entrypoint for the application.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "handlers": {
        "block": {
          "attributes": {
            "auth_fail_action": {
              "description": "Actions to take when the user is not logged in. Possible values: [\"AUTH_FAIL_ACTION_REDIRECT\", \"AUTH_FAIL_ACTION_UNAUTHORIZED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "login": {
              "description": "Methods to restrict access to a URL based on login status. Possible values: [\"LOGIN_OPTIONAL\", \"LOGIN_ADMIN\", \"LOGIN_REQUIRED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_http_response_code": {
              "description": "30x code to use when performing redirects for the secure field. Possible values: [\"REDIRECT_HTTP_RESPONSE_CODE_301\", \"REDIRECT_HTTP_RESPONSE_CODE_302\", \"REDIRECT_HTTP_RESPONSE_CODE_303\", \"REDIRECT_HTTP_RESPONSE_CODE_307\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_level": {
              "description": "Security (HTTPS) enforcement for this URL. Possible values: [\"SECURE_DEFAULT\", \"SECURE_NEVER\", \"SECURE_OPTIONAL\", \"SECURE_ALWAYS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_regex": {
              "description": "URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.\nAll URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.",
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
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Executes a script to handle the requests that match this URL pattern.\nOnly the auto value is supported for Node.js in the App Engine standard environment, for example \"script:\" \"auto\".",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "static_files": {
              "block": {
                "attributes": {
                  "application_readable": {
                    "description": "Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as\nstatic data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged\nagainst both your code and static data storage resource quotas.",
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
                "description": "Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files. Static file handlers describe which files in the application directory are static files, and which URLs serve them.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "An ordered list of URL-matching patterns that should be applied to incoming requests.\nThe first matching URL handles the request and other request handlers are not attempted.",
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
          "description": "Configuration for third-party Python runtime libraries that are required by the application.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "manual_scaling": {
        "block": {
          "attributes": {
            "instances": {
              "description": "Number of instances to assign to the service at the start.\n\n**Note:** When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2\nModules API set_num_instances() you must use 'lifecycle.ignore_changes = [\"manual_scaling\"[0].instances]' to prevent drift detection.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.",
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
      "vpc_access_connector": {
        "block": {
          "attributes": {
            "name": {
              "description": "Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Enables VPC connectivity for standard apps.",
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

func GoogleAppEngineStandardAppVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineStandardAppVersion), &result)
	return &result
}
