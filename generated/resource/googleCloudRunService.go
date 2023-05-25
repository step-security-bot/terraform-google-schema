package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunService = `{
  "block": {
    "attributes": {
      "autogenerate_revision_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the cloud run instance. eg us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name must be unique within a namespace, within a Cloud Run region.\nIs required when creating resources. Name is primarily intended\nfor creation idempotence and configuration definition. Cannot be updated.\nMore info: http://kubernetes.io/docs/user-guide/identifiers#names",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the Service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "message": "string",
                    "reason": "string",
                    "status": "string",
                    "type": "string"
                  }
                ]
              ],
              "latest_created_revision_name": "string",
              "latest_ready_revision_name": "string",
              "observed_generation": "number",
              "url": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "computed": true,
              "description": "Annotations is a key value map stored with a resource that\nmay be set by external tools to store and retrieve arbitrary metadata. More\ninfo: http://kubernetes.io/docs/user-guide/annotations\n\n**Note**: The Cloud Run API may add additional annotations that were not provided in your config.\nIf terraform plan shows a diff where a server-side annotation is added, you can add it to your config\nor apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.\n\nCloud Run (fully managed) uses the following annotation keys to configure features on a Service:\n\n- 'run.googleapis.com/ingress' sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)\n  for the Service. For example, '\"run.googleapis.com/ingress\" = \"all\"'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "generation": {
              "computed": true,
              "description": "A sequence number representing a specific generation of the desired state.",
              "description_kind": "plain",
              "type": "number"
            },
            "labels": {
              "computed": true,
              "description": "Map of string keys and values that can be used to organize and categorize\n(scope and select) objects. May match selectors of replication controllers\nand routes.\nMore info: http://kubernetes.io/docs/user-guide/labels",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "namespace": {
              "computed": true,
              "description": "In Cloud Run the namespace must be equal to either the\nproject ID or project number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this object that\ncan be used by clients to determine when objects have changed. May be used\nfor optimistic concurrency, change detection, and the watch operation on a\nresource or set of resources. They may only be valid for a\nparticular resource or set of resources.\n\nMore info:\nhttps://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "self_link": {
              "computed": true,
              "description": "SelfLink is a URL representing this object.",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "UID is a unique id generated by the server on successful creation of a resource and is not\nallowed to change on PUT operations.\n\nMore info: http://kubernetes.io/docs/user-guide/identifiers#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Metadata associated with this Service, including name, namespace, labels,\nand annotations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "template": {
        "block": {
          "block_types": {
            "metadata": {
              "block": {
                "attributes": {
                  "annotations": {
                    "computed": true,
                    "description": "Annotations is a key value map stored with a resource that\nmay be set by external tools to store and retrieve arbitrary metadata. More\ninfo: http://kubernetes.io/docs/user-guide/annotations\n\n**Note**: The Cloud Run API may add additional annotations that were not provided in your config.\nIf terraform plan shows a diff where a server-side annotation is added, you can add it to your config\nor apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "generation": {
                    "computed": true,
                    "description": "A sequence number representing a specific generation of the desired state.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "labels": {
                    "description": "Map of string keys and values that can be used to organize and categorize\n(scope and select) objects. May match selectors of replication controllers\nand routes.\nMore info: http://kubernetes.io/docs/user-guide/labels",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "name": {
                    "computed": true,
                    "description": "Name must be unique within a namespace, within a Cloud Run region.\nIs required when creating resources. Name is primarily intended\nfor creation idempotence and configuration definition. Cannot be updated.\nMore info: http://kubernetes.io/docs/user-guide/identifiers#names",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespace": {
                    "computed": true,
                    "description": "In Cloud Run the namespace must be equal to either the\nproject ID or project number. It will default to the resource's project.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_version": {
                    "computed": true,
                    "description": "An opaque value that represents the internal version of this object that\ncan be used by clients to determine when objects have changed. May be used\nfor optimistic concurrency, change detection, and the watch operation on a\nresource or set of resources. They may only be valid for a\nparticular resource or set of resources.\n\nMore info:\nhttps://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "self_link": {
                    "computed": true,
                    "description": "SelfLink is a URL representing this object.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "UID is a unique id generated by the server on successful creation of a resource and is not\nallowed to change on PUT operations.\n\nMore info: http://kubernetes.io/docs/user-guide/identifiers#uids",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Optional metadata for this Revision, including labels and annotations.\nName will be generated by the Configuration. To set minimum instances\nfor this revision, use the \"autoscaling.knative.dev/minScale\" annotation\nkey. To set maximum instances for this revision, use the\n\"autoscaling.knative.dev/maxScale\" annotation key. To set Cloud SQL\nconnections for the revision, use the \"run.googleapis.com/cloudsql-instances\"\nannotation key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "spec": {
              "block": {
                "attributes": {
                  "container_concurrency": {
                    "computed": true,
                    "description": "ContainerConcurrency specifies the maximum allowed in-flight (concurrent)\nrequests per container of the Revision. Values are:\n- '0' thread-safe, the system should manage the max concurrency. This is\n    the default value.\n- '1' not-thread-safe. Single concurrency\n- '2-N' thread-safe, max concurrency of N",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_account_name": {
                    "description": "Email address of the IAM service account associated with the revision of the\nservice. The service account represents the identity of the running revision,\nand determines what permissions the revision has. If not provided, the revision\nwill use the project's default service account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "serving_state": {
                    "computed": true,
                    "deprecated": true,
                    "description": "ServingState holds a value describing the state the resources\nare in for this Revision.\nIt is expected\nthat the system will manipulate this based on routability and load.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description": "TimeoutSeconds holds the max duration the instance is allowed for responding to a request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "containers": {
                    "block": {
                      "attributes": {
                        "args": {
                          "description": "Arguments to the entrypoint.\nThe docker image's CMD is used if this is not provided.\nVariable references $(VAR_NAME) are expanded using the container's\nenvironment. If a variable cannot be resolved, the reference in the input\nstring will be unchanged. The $(VAR_NAME) syntax can be escaped with a\ndouble $$, ie: $$(VAR_NAME). Escaped references will never be expanded,\nregardless of whether the variable exists or not.\nMore info:\nhttps://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "description": "Entrypoint array. Not executed within a shell.\nThe docker image's ENTRYPOINT is used if this is not provided.\nVariable references $(VAR_NAME) are expanded using the container's\nenvironment. If a variable cannot be resolved, the reference in the input\nstring will be unchanged. The $(VAR_NAME) syntax can be escaped with a\ndouble $$, ie: $$(VAR_NAME). Escaped references will never be expanded,\nregardless of whether the variable exists or not.\nMore info:\nhttps://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "image": {
                          "description": "Docker image name. This is most often a reference to a container located\nin the container registry, such as gcr.io/cloudrun/hello\nMore info: https://kubernetes.io/docs/concepts/containers/images",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "working_dir": {
                          "deprecated": true,
                          "description": "Container's working directory.\nIf not specified, the container runtime's default will be used, which\nmight be configured in the container image.",
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
                                "description": "Name of the environment variable.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Variable references $(VAR_NAME) are expanded\nusing the previous defined environment variables in the container and\nany route environment variables. If a variable cannot be resolved,\nthe reference in the input string will be unchanged. The $(VAR_NAME)\nsyntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped\nreferences will never be expanded, regardless of whether the variable\nexists or not.\nDefaults to \"\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "List of environment variables to set in the container.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
                        },
                        "env_from": {
                          "block": {
                            "attributes": {
                              "prefix": {
                                "description": "An optional identifier to prepend to each key in the ConfigMap.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "config_map_ref": {
                                "block": {
                                  "attributes": {
                                    "optional": {
                                      "description": "Specify whether the ConfigMap must be defined",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "local_object_reference": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the referent.\nMore info:\nhttps://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The ConfigMap to select from.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The ConfigMap to select from.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "secret_ref": {
                                "block": {
                                  "attributes": {
                                    "optional": {
                                      "description": "Specify whether the Secret must be defined",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "local_object_reference": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the referent.\nMore info:\nhttps://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The Secret to select from.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The Secret to select from.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "deprecated": true,
                            "description": "List of sources to populate environment variables in the container.\nAll invalid keys will be reported as an event when the container is starting.\nWhen a key exists in multiple sources, the value associated with the last source will\ntake precedence. Values defined by an Env with a duplicate key will take\nprecedence.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "ports": {
                          "block": {
                            "attributes": {
                              "container_port": {
                                "description": "Port number.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "name": {
                                "computed": true,
                                "description": "Name of the port.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "protocol": {
                                "description": "Protocol used on port. Defaults to TCP.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "List of open ports in the container.\nMore Info: \nhttps://cloud.google.com/run/docs/reference/rest/v1/RevisionSpec#ContainerPort",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "resources": {
                          "block": {
                            "attributes": {
                              "limits": {
                                "computed": true,
                                "description": "Limits describes the maximum amount of compute resources allowed.\nThe values of the map is string form of the 'quantity' k8s type:\nhttps://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "requests": {
                                "description": "Requests describes the minimum amount of compute resources required.\nIf Requests is omitted for a container, it defaults to Limits if that is\nexplicitly specified, otherwise to an implementation-defined value.\nThe values of the map is string form of the 'quantity' k8s type:\nhttps://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "Compute Resources required by this container. Used to set values such as max memory\nMore info:\nhttps://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Container defines the unit of execution for this Revision.\nIn the context of a Revision, we disallow a number of the fields of\nthis Container, including: name, ports, and volumeMounts.\nThe runtime contract is documented here:\nhttps://github.com/knative/serving/blob/master/docs/runtime-contract.md",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "RevisionSpec holds the desired state of the Revision (from the client).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "template holds the latest specification for the Revision to\nbe stamped out. The template references the container image, and may also\ninclude labels and annotations that should be attached to the Revision.\nTo correlate a Revision, and/or to force a Revision to be created when the\nspec doesn't otherwise change, a nonce label may be provided in the\ntemplate metadata. For more details, see:\nhttps://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions\n\nCloud Run does not currently support referencing a build that is\nresponsible for materializing the container image from source.",
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
      "traffic": {
        "block": {
          "attributes": {
            "latest_revision": {
              "description": "LatestRevision may be optionally provided to indicate that the latest ready\nRevision of the Configuration should be used for this traffic target. When\nprovided LatestRevision must be true if RevisionName is empty; it must be\nfalse when RevisionName is non-empty.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "percent": {
              "description": "Percent specifies percent of the traffic to this Revision or Configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "revision_name": {
              "description": "RevisionName of a specific revision to which to send this portion of traffic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Traffic specifies how to distribute traffic over a collection of Knative Revisions\nand Configurations",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleCloudRunServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunService), &result)
	return &result
}
