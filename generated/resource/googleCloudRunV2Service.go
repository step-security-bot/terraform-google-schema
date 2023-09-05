package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunV2Service = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.\n\nCloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected in new resources.\nAll system annotations in v1 now have a corresponding field in v2 Service.\n\nThis field follows Kubernetes annotations' namespacing, limits, and rules.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "client": {
        "description": "Arbitrary identifier for the API client.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_version": {
        "description": "Arbitrary version identifier for the API client.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "conditions": {
        "computed": true,
        "description": "The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "execution_reason": "string",
              "last_transition_time": "string",
              "message": "string",
              "reason": "string",
              "revision_reason": "string",
              "severity": "string",
              "state": "string",
              "type": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "Email address of the authenticated creator.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The deletion time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description of the Service. This field currently has a 512-character limit.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "For a deleted resource, the time after which it will be permamently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "generation": {
        "computed": true,
        "description": "A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a string instead of an integer.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingress": {
        "computed": true,
        "description": "Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active. Possible values: [\"INGRESS_TRAFFIC_ALL\", \"INGRESS_TRAFFIC_INTERNAL_ONLY\", \"INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,\nenvironment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.\n\nCloud Run API v2 does not support labels with  'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system labels in v1 now have a corresponding field in v2 Service.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_modifier": {
        "computed": true,
        "description": "Email address of the last authenticated modifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_created_revision": {
        "computed": true,
        "description": "Name of the last created revision. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_ready_revision": {
        "computed": true,
        "description": "Name of the latest revision that is serving traffic. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_stage": {
        "computed": true,
        "description": "The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.\nIf no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.\n\nFor example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: [\"UNIMPLEMENTED\", \"PRELAUNCH\", \"EARLY_ACCESS\", \"ALPHA\", \"BETA\", \"GA\", \"DEPRECATED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the cloud run service",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "observed_generation": {
        "computed": true,
        "description": "The generation of this Service currently serving traffic. See comments in reconciling for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a string instead of an integer.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Returns true if the Service is currently being acted upon by the system to bring it into the desired state.\n\nWhen a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, observedGeneration, latest_ready_revison, trafficStatuses, and uri will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.\n\nIf reconciliation succeeded, the following fields will match: traffic and trafficStatuses, observedGeneration and generation, latestReadyRevision and latestCreatedRevision.\n\nIf reconciliation failed, trafficStatuses, observedGeneration, and latestReadyRevision will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in terminalCondition and conditions.",
        "description_kind": "plain",
        "type": "bool"
      },
      "terminal_condition": {
        "computed": true,
        "description": "The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "execution_reason": "string",
              "last_transition_time": "string",
              "message": "string",
              "reason": "string",
              "revision_reason": "string",
              "severity": "string",
              "state": "string",
              "type": "string"
            }
          ]
        ]
      },
      "traffic_statuses": {
        "computed": true,
        "description": "Detailed status information for corresponding traffic targets. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "percent": "number",
              "revision": "string",
              "tag": "string",
              "type": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last-modified time.",
        "description_kind": "plain",
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "The main URI in which this Service is serving traffic.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "binary_authorization": {
        "block": {
          "attributes": {
            "breakglass_justification": {
              "description": "If present, indicates to use Breakglass using this justification. If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_default": {
              "description": "If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings for the Binary Authorization feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "template": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.\n\nCloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system annotations in v1 now have a corresponding field in v2 RevisionTemplate.\n\nThis field follows Kubernetes annotations' namespacing, limits, and rules.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "encryption_key": {
              "description": "A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_environment": {
              "description": "The sandbox environment to host this Revision. Possible values: [\"EXECUTION_ENVIRONMENT_GEN1\", \"EXECUTION_ENVIRONMENT_GEN2\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description": "Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc.\nFor more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.\n\nCloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system labels in v1 now have a corresponding field in v2 RevisionTemplate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "max_instance_request_concurrency": {
              "computed": true,
              "description": "Sets the maximum number of requests that each serving instance can receive.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "revision": {
              "description": "The unique name for the revision. If this field is omitted, it will be automatically generated based on the Service name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account": {
              "computed": true,
              "description": "Email address of the IAM service account associated with the revision of the service. The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "session_affinity": {
              "description": "Enables session affinity. For more information, go to https://cloud.google.com/run/docs/configuring/session-affinity",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "timeout": {
              "computed": true,
              "description": "Max allowed time for an instance to respond to a request.\n\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "containers": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "image": {
                    "description": "URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the container specified as a DNS_LABEL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "working_dir": {
                    "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "env": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\", and the maximum length is 32768 bytes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "value_source": {
                          "block": {
                            "block_types": {
                              "secret_key_ref": {
                                "block": {
                                  "attributes": {
                                    "secret": {
                                      "description": "The name of the secret in Cloud Secret Manager. Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "version": {
                                      "description": "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Selects a secret and a specific version from Cloud Secret Manager.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Source for the environment variable's value.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of environment variables to set in the container.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "liveness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before the probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. Must be greater or equal than timeoutSeconds",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. Maximum value is 3600. Must be smaller than periodSeconds. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "The name of the service to place in the gRPC HealthCheckRequest\n(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).\nIf this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "path": {
                                "description": "Path to access on the HTTP server. Defaults to '/'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "http_headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Custom headers to set in the request. HTTP allows repeated headers.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "HTTPGet specifies the http request to perform.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Port number to access on the container. Must be in the range 1 to 65535. If not specified, defaults to 8080.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "deprecated": true,
                            "description": "TCPSocket specifies an action involving a TCP port. This field is not supported in liveness probe currently.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Periodic probe of container liveness. Container will be restarted if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ports": {
                    "block": {
                      "attributes": {
                        "container_port": {
                          "description": "Port number the container listens on. This must be a valid TCP port number, 0 \u003c containerPort \u003c 65536.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "name": {
                          "computed": true,
                          "description": "If specified, used to specify which protocol to use. Allowed values are \"http1\" and \"h2c\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.\n\nIf omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "resources": {
                    "block": {
                      "attributes": {
                        "cpu_idle": {
                          "description": "Determines whether CPU should be throttled or not outside of requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "limits": {
                          "computed": true,
                          "description": "Only memory and CPU are supported. Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "startup_cpu_boost": {
                          "description": "Determines whether CPU should be boosted on startup of a new container instance above the requested CPU threshold, this can help reduce cold-start latency.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "startup_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds after the container has started before the probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. Must be greater or equal than timeoutSeconds",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. Maximum value is 3600. Must be smaller than periodSeconds. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "The name of the service to place in the gRPC HealthCheckRequest\n(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).\nIf this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GRPC specifies an action involving a GRPC port.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "path": {
                                "description": "Path to access on the HTTP server. Defaults to '/'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "Port number to access on the container. Must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "http_headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Custom headers to set in the request. HTTP allows repeated headers.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "HTTPGet specifies the http request to perform. Exactly one of HTTPGet or TCPSocket must be specified.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "Port number to access on the container. Must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "TCPSocket specifies an action involving a TCP port. Exactly one of HTTPGet or TCPSocket must be specified.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Startup probe of application within the container. All other probes are disabled if a startup probe is provided, until it succeeds. Container will not be added to service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "volume_mounts": {
                    "block": {
                      "attributes": {
                        "mount_path": {
                          "description": "Path within the container at which the volume should be mounted. Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must otherwise be /cloudsql. All instances defined in the Volume will be available as /cloudsql/[instance]. For more information on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "This must match the Name of a Volume.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Volume to mount into the container's filesystem.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Holds the containers that define the unit of execution for this Service.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scaling": {
              "block": {
                "attributes": {
                  "max_instance_count": {
                    "description": "Maximum number of serving instances that this resource should have.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_instance_count": {
                    "description": "Minimum number of serving instances that this resource should have.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Scaling settings for this Revision.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "volumes": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Volume's name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "cloud_sql_instance": {
                    "block": {
                      "attributes": {
                        "instances": {
                          "description": "The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret": {
                    "block": {
                      "attributes": {
                        "default_mode": {
                          "description": "Integer representation of mode bits to use on created files by default. Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "secret": {
                          "description": "The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "items": {
                          "block": {
                            "attributes": {
                              "mode": {
                                "description": "Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal). If 0 or not set, the Volume's default mode will be used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "path": {
                                "description": "The relative path of the secret in the container.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description": "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "If unspecified, the volume will expose a file whose name is the secret, relative to VolumeMount.mount_path. If specified, the key will be used as the version to fetch from Cloud Secret Manager and the path will be the name of the file exposed in the volume. When items are defined, they must specify a path and a version.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of Volumes to make available to containers.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_access": {
              "block": {
                "attributes": {
                  "connector": {
                    "description": "VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "egress": {
                    "description": "Traffic VPC egress settings. Possible values: [\"ALL_TRAFFIC\", \"PRIVATE_RANGES_ONLY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The template used to create revisions for this Service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
      "traffic": {
        "block": {
          "attributes": {
            "percent": {
              "computed": true,
              "description": "Specifies percent of the traffic to this Revision. This defaults to zero if unspecified.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "revision": {
              "description": "Revision to which to send this portion of traffic, if traffic allocation is by revision.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag": {
              "description": "Indicates a string to be part of the URI to exclusively reference this target.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "The allocation type for this traffic target. Possible values: [\"TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST\", \"TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest Ready Revision.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudRunV2ServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunV2Service), &result)
	return &result
}
