package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineFlexibleAppVersion = `{
  "block": {
    "attributes": {
      "beta_settings": {
        "description": "Metadata settings that are supplied to this version to enable beta runtime features.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "default_expiration": {
        "description": "Duration that static files should be cached by web proxies and browsers.\nOnly applicable if the corresponding StaticFilesHandler does not specify its own expiration time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_service_on_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "env_variables": {
        "description": "Environment variables available to the application.  As these are not returned in the API request, Terraform will not detect any changes made outside of the Terraform config.",
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
        "description": "Instance class that is used to run this version. Valid values are\nAutomaticScaling: F1, F2, F4, F4_1G\nManualScaling: B1, B2, B4, B8, B4_1G\nDefaults to F1 for AutomaticScaling and B1 for ManualScaling.",
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
      "nobuild_files_regex": {
        "description": "Files that match this pattern will not be built into this version. Only applicable for Go runtimes.",
        "description_kind": "plain",
        "optional": true,
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
        "computed": true,
        "description": "The version of the API in the given runtime environment.\nPlease see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_channel": {
        "description": "The channel of the runtime to use. Only available for some runtimes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_main_executable_path": {
        "description": "The path or name of the app's main executable.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "AppEngine service resource. Can contain numbers, letters, and hyphens.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "serving_status": {
        "description": "Current serving status of this version. Only the versions with a SERVING status create instances and can be billed. Default value: \"SERVING\" Possible values: [\"SERVING\", \"STOPPED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_id": {
        "description": "Relative name of the version within the service. For example, 'v1'. Version names can contain only lowercase letters, numbers, or hyphens.\nReserved names,\"default\", \"latest\", and any name with the prefix \"ah-\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "api_config": {
        "block": {
          "attributes": {
            "auth_fail_action": {
              "description": "Action to take when users access resources that require authentication. Default value: \"AUTH_FAIL_ACTION_REDIRECT\" Possible values: [\"AUTH_FAIL_ACTION_REDIRECT\", \"AUTH_FAIL_ACTION_UNAUTHORIZED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "login": {
              "description": "Level of login required to access this resource. Default value: \"LOGIN_OPTIONAL\" Possible values: [\"LOGIN_OPTIONAL\", \"LOGIN_ADMIN\", \"LOGIN_REQUIRED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script": {
              "description": "Path to the script from the application root directory.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_level": {
              "description": "Security (HTTPS) enforcement for this URL. Possible values: [\"SECURE_DEFAULT\", \"SECURE_NEVER\", \"SECURE_OPTIONAL\", \"SECURE_ALWAYS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description": "URL to serve the endpoint at.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Serving configuration for Google Cloud Endpoints.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "automatic_scaling": {
        "block": {
          "attributes": {
            "cool_down_period": {
              "description": "The time period that the Autoscaler should wait before it starts collecting information from a new instance.\nThis prevents the autoscaler from collecting information when the instance is initializing,\nduring which the collected usage would not be reliable. Default: 120s",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_concurrent_requests": {
              "computed": true,
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
              "description": "Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_total_instances": {
              "description": "Maximum number of instances that should be started to handle requests for this version. Default: 20",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_idle_instances": {
              "description": "Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_pending_latency": {
              "description": "Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_total_instances": {
              "description": "Minimum number of running instances that should be maintained for this version. Default: 2",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "cpu_utilization": {
              "block": {
                "attributes": {
                  "aggregation_window_length": {
                    "description": "Period of time over which CPU utilization is calculated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_utilization": {
                    "description": "Target CPU utilization ratio to maintain when scaling. Must be between 0 and 1.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Target scaling by CPU usage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "disk_utilization": {
              "block": {
                "attributes": {
                  "target_read_bytes_per_second": {
                    "description": "Target bytes read per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_read_ops_per_second": {
                    "description": "Target ops read per seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_write_bytes_per_second": {
                    "description": "Target bytes written per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_write_ops_per_second": {
                    "description": "Target ops written per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Target scaling by disk usage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_utilization": {
              "block": {
                "attributes": {
                  "target_received_bytes_per_second": {
                    "description": "Target bytes received per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_received_packets_per_second": {
                    "description": "Target packets received per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_sent_bytes_per_second": {
                    "description": "Target bytes sent per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_sent_packets_per_second": {
                    "description": "Target packets sent per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Target scaling by network usage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "request_utilization": {
              "block": {
                "attributes": {
                  "target_concurrent_requests": {
                    "description": "Target number of concurrent requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_request_count_per_second": {
                    "description": "Target requests per second.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Target scaling by request utilization.",
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
      "deployment": {
        "block": {
          "block_types": {
            "cloud_build_options": {
              "block": {
                "attributes": {
                  "app_yaml_path": {
                    "description": "Path to the yaml file used in deployment, used to determine runtime configuration details.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cloud_build_timeout": {
                    "description": "The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Options for the build operations performed as a part of the version deployment. Only applicable when creating a version using source code directly.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "container": {
              "block": {
                "attributes": {
                  "image": {
                    "description": "URI to the hosted container image in Google Container Registry. The URI must be fully qualified and include a tag or digest.\nExamples: \"gcr.io/my-project/image:tag\" or \"gcr.io/my-project/image@digest\"",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The Docker image for the container that runs the version.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
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
        "nesting_mode": "list"
      },
      "endpoints_api_service": {
        "block": {
          "attributes": {
            "config_id": {
              "description": "Endpoints service configuration ID as specified by the Service Management API. For example \"2016-09-19r1\".\n\nBy default, the rollout strategy for Endpoints is \"FIXED\". This means that Endpoints starts up with a particular configuration ID.\nWhen a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID\nand is required in this case.\n\nEndpoints also has a rollout strategy called \"MANAGED\". When using this, Endpoints fetches the latest configuration and does not need\nthe configuration ID. In this case, configId must be omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disable_trace_sampling": {
              "description": "Enable or disable trace sampling. By default, this is set to false for enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "Endpoints service name which is the name of the \"service\" resource in the Service Management API.\nFor example \"myapi.endpoints.myproject.cloud.goog\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rollout_strategy": {
              "description": "Endpoints rollout strategy. If FIXED, configId must be specified. If MANAGED, configId must be omitted. Default value: \"FIXED\" Possible values: [\"FIXED\", \"MANAGED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Code and application artifacts that make up this version.",
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
                    "description": "Whether files should also be uploaded as code data. By default, files declared in static file handlers are\nuploaded as static data and are only served to end users; they cannot be read by the application. If enabled,\nuploads are charged against both your code and static data storage resource quotas.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expiration": {
                    "description": "Time a static file served by this handler should be cached by web proxies and browsers.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example \"3.5s\".\nDefault is '0s'",
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
                    "description": "Path to the static files matched by the URL pattern, from the application root directory.\nThe path can refer to text matched in groupings in the URL pattern.",
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
                "description": "Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files.\nStatic file handlers describe which files in the application directory are static files, and which URLs serve them.",
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
      "liveness_check": {
        "block": {
          "attributes": {
            "check_interval": {
              "description": "Interval between health checks.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure_threshold": {
              "description": "Number of consecutive failed checks required before considering the VM unhealthy. Default: 4.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "host": {
              "description": "Host header to send when performing a HTTP Readiness check. Example: \"myapp.appspot.com\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "initial_delay": {
              "description": "The initial delay before starting to execute the checks. Default: \"300s\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description": "The request path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "success_threshold": {
              "description": "Number of consecutive successful checks required before considering the VM healthy. Default: 2.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout": {
              "description": "Time before the check is considered failed. Default: \"4s\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
      "network": {
        "block": {
          "attributes": {
            "forwarded_ports": {
              "description": "List of ports, or port pairs, to forward from the virtual machine to the application container.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "instance_tag": {
              "description": "Tag to apply to the instance during creation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "Google Compute Engine network where the virtual machines are created. Specify the short name, not the resource path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "session_affinity": {
              "description": "Enable session affinity.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "subnetwork": {
              "description": "Google Cloud Platform sub-network where the virtual machines are created. Specify the short name, not the resource path.\n\nIf the network that the instance is being created in is a Legacy network, then the IP address is allocated from the IPv4Range.\nIf the network that the instance is being created in is an auto Subnet Mode Network, then only network name should be specified (not the subnetworkName) and the IP address is created from the IPCidrRange of the subnetwork that exists in that zone for that network.\nIf the network that the instance is being created in is a custom Subnet Mode Network, then the subnetworkName must be specified and the IP address is created from the IPCidrRange of the subnetwork.\nIf specified, the subnetwork must exist in the same region as the App Engine flexible environment application.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Extra network settings",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "readiness_check": {
        "block": {
          "attributes": {
            "app_start_timeout": {
              "description": "A maximum time limit on application initialization, measured from moment the application successfully\nreplies to a healthcheck until it is ready to serve traffic. Default: \"300s\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "check_interval": {
              "description": "Interval between health checks.  Default: \"5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure_threshold": {
              "description": "Number of consecutive failed checks required before removing traffic. Default: 2.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "host": {
              "description": "Host header to send when performing a HTTP Readiness check. Example: \"myapp.appspot.com\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description": "The request path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "success_threshold": {
              "description": "Number of consecutive successful checks required before receiving traffic. Default: 2.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout": {
              "description": "Time before the check is considered failed. Default: \"4s\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "resources": {
        "block": {
          "attributes": {
            "cpu": {
              "description": "Number of CPU cores needed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_gb": {
              "description": "Disk size (GB) needed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory_gb": {
              "description": "Memory (GB) needed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "volumes": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Unique name for the volume.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "size_gb": {
                    "description": "Volume size in gigabytes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "volume_type": {
                    "description": "Underlying volume type, e.g. 'tmpfs'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "List of ports, or port pairs, to forward from the virtual machine to the application container.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Machine resources for a version.",
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

func GoogleAppEngineFlexibleAppVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineFlexibleAppVersion), &result)
	return &result
}
