package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerServicePerimeter = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the AccessPolicy was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the ServicePerimeter and its use. Does not affect\nbehavior.",
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
        "description": "Resource name for the ServicePerimeter. The short_name component must\nbegin with a letter and only include alphanumeric and '_'.\nFormat: accessPolicies/{policy_id}/servicePerimeters/{short_name}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The AccessPolicy this ServicePerimeter lives in.\nFormat: accessPolicies/{policy_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "perimeter_type": {
        "description": "Specifies the type of the Perimeter. There are two types: regular and\nbridge. Regular Service Perimeter contains resources, access levels,\nand restricted services. Every resource can be in at most\nONE regular Service Perimeter.\n\nIn addition to being in a regular service perimeter, a resource can also\nbe in zero or more perimeter bridges. A perimeter bridge only contains\nresources. Cross project operations are permitted if all effected\nresources share some perimeter (whether bridge or regular). Perimeter\nBridge does not contain access levels or services: those are governed\nentirely by the regular perimeter that resource is in.\n\nPerimeter Bridges are typically useful when building more complex\ntopologies with many independent perimeters that need to share some data\nwith a common perimeter, but should not be able to share data among\nthemselves. Default value: \"PERIMETER_TYPE_REGULAR\" Possible values: [\"PERIMETER_TYPE_REGULAR\", \"PERIMETER_TYPE_BRIDGE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "description": "Human readable title. Must be unique within the Policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the AccessPolicy was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_explicit_dry_run_spec": {
        "description": "Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists\nfor all Service Perimeters, and that spec is identical to the status for those\nService Perimeters. When this flag is set, it inhibits the generation of the\nimplicit spec, thereby allowing the user to explicitly provide a\nconfiguration (\"spec\") to use in a dry-run version of the Service Perimeter.\nThis allows the user to test changes to the enforced config (\"status\") without\nactually enforcing them. This testing is done through analyzing the differences\nbetween currently enforced and suggested restrictions. useExplicitDryRunSpec must\nbet set to True if any of the fields in the spec are set to non-default values.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "attributes": {
            "access_levels": {
              "description": "A list of AccessLevel resource names that allow resources within\nthe ServicePerimeter to be accessed from the internet.\nAccessLevels listed must be in the same policy as this\nServicePerimeter. Referencing a nonexistent AccessLevel is a\nsyntax error. If no AccessLevel names are listed, resources within\nthe perimeter can only be accessed via GCP calls with request\norigins within the perimeter. For Service Perimeter Bridge, must\nbe empty.\n\nFormat: accessPolicies/{policy_id}/accessLevels/{access_level_name}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "description": "A list of GCP resources that are inside of the service perimeter.\nCurrently only projects are allowed.\nFormat: projects/{project_number}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "restricted_services": {
              "description": "GCP services that are subject to the Service Perimeter\nrestrictions. Must contain a list of services. For example, if\n'storage.googleapis.com' is specified, access to the storage\nbuckets inside the perimeter must meet the perimeter's access\nrestrictions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "egress_policies": {
              "block": {
                "block_types": {
                  "egress_from": {
                    "block": {
                      "attributes": {
                        "identities": {
                          "description": "A list of identities that are allowed access through this 'EgressPolicy'. \nShould be in the format of email address. The email address should \nrepresent individual user or service account only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "identity_type": {
                          "description": "Specifies the type of identities that are allowed access to outside the \nperimeter. If left unspecified, then members of 'identities' field will \nbe allowed access. Possible values: [\"IDENTITY_TYPE_UNSPECIFIED\", \"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Defines conditions on the source of a request causing this 'EgressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "egress_to": {
                    "block": {
                      "attributes": {
                        "resources": {
                          "description": "A list of resources, currently only projects in the form \n'projects/\u003cprojectnumber\u003e', that match this to stanza. A request matches \nif it contains a resource in this list. If * is specified for resources, \nthen this 'EgressTo' rule will authorize access to all resources outside \nthe perimeter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "operations": {
                          "block": {
                            "attributes": {
                              "service_name": {
                                "description": "The name of the API whose methods or permissions the 'IngressPolicy' or \n'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName \nfield set to '*' will allow all methods AND permissions for all services.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "method_selectors": {
                                "block": {
                                  "attributes": {
                                    "method": {
                                      "description": "Value for 'method' should be a valid method name for the corresponding \n'serviceName' in 'ApiOperation'. If '*' used as value for method, \nthen ALL methods and permissions are allowed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "permission": {
                                      "description": "Value for permission should be a valid Cloud IAM permission for the \ncorresponding 'serviceName' in 'ApiOperation'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "API methods or permissions to allow. Method or permission must belong \nto the service specified by 'serviceName' field. A single MethodSelector \nentry with '*' specified for the 'method' field will allow all methods \nAND permissions for the service specified in 'serviceName'.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of 'ApiOperations' that this egress rule applies to. A request matches \nif it contains an operation/service in this list.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the 'ApiOperation' and destination resources that \ncause this 'EgressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of EgressPolicies to apply to the perimeter. A perimeter may \nhave multiple EgressPolicies, each of which is evaluated separately.\nAccess is granted if any EgressPolicy grants it. Must be empty for \na perimeter bridge.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ingress_policies": {
              "block": {
                "block_types": {
                  "ingress_from": {
                    "block": {
                      "attributes": {
                        "identities": {
                          "description": "A list of identities that are allowed access through this ingress policy.\nShould be in the format of email address. The email address should represent \nindividual user or service account only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "identity_type": {
                          "description": "Specifies the type of identities that are allowed access from outside the \nperimeter. If left unspecified, then members of 'identities' field will be \nallowed access. Possible values: [\"IDENTITY_TYPE_UNSPECIFIED\", \"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "sources": {
                          "block": {
                            "attributes": {
                              "access_level": {
                                "description": "An 'AccessLevel' resource name that allow resources within the \n'ServicePerimeters' to be accessed from the internet. 'AccessLevels' listed \nmust be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent\n'AccessLevel' will cause an error. If no 'AccessLevel' names are listed, \nresources within the perimeter can only be accessed via Google Cloud calls \nwith request origins within the perimeter. \nExample 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.' \nIf * is specified, then all IngressSources will be allowed.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "resource": {
                                "description": "A Google Cloud resource that is allowed to ingress the perimeter. \nRequests from these resources will be allowed to access perimeter data. \nCurrently only projects are allowed. Format 'projects/{project_number}' \nThe project may be in any Google Cloud organization, not just the \norganization that the perimeter is defined in. '*' is not allowed, the case \nof allowing all Google Cloud resources only is not supported.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Sources that this 'IngressPolicy' authorizes access from.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the source of a request causing this 'IngressPolicy'\nto apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ingress_to": {
                    "block": {
                      "attributes": {
                        "resources": {
                          "description": "A list of resources, currently only projects in the form \n'projects/\u003cprojectnumber\u003e', protected by this 'ServicePerimeter'\nthat are allowed to be accessed by sources defined in the\ncorresponding 'IngressFrom'. A request matches if it contains\na resource in this list. If '*' is specified for resources,\nthen this 'IngressTo' rule will authorize access to all \nresources inside the perimeter, provided that the request\nalso matches the 'operations' field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "operations": {
                          "block": {
                            "attributes": {
                              "service_name": {
                                "description": "The name of the API whose methods or permissions the 'IngressPolicy' or \n'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName' \nfield set to '*' will allow all methods AND permissions for all services.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "method_selectors": {
                                "block": {
                                  "attributes": {
                                    "method": {
                                      "description": "Value for method should be a valid method name for the corresponding \nserviceName in 'ApiOperation'. If '*' used as value for 'method', then \nALL methods and permissions are allowed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "permission": {
                                      "description": "Value for permission should be a valid Cloud IAM permission for the \ncorresponding 'serviceName' in 'ApiOperation'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "API methods or permissions to allow. Method or permission must belong to \nthe service specified by serviceName field. A single 'MethodSelector' entry \nwith '*' specified for the method field will allow all methods AND \npermissions for the service specified in 'serviceName'.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom' \nare allowed to perform in this 'ServicePerimeter'.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the 'ApiOperation' and request destination that cause\nthis 'IngressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of 'IngressPolicies' to apply to the perimeter. A perimeter may\nhave multiple 'IngressPolicies', each of which is evaluated\nseparately. Access is granted if any 'Ingress Policy' grants it.\nMust be empty for a perimeter bridge.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_accessible_services": {
              "block": {
                "attributes": {
                  "allowed_services": {
                    "description": "The list of APIs usable within the Service Perimeter.\nMust be empty unless 'enableRestriction' is True.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable_restriction": {
                    "description": "Whether to restrict API calls within the Service Perimeter to the\nlist of APIs specified in 'allowedServices'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Specifies how APIs are allowed to communicate within the Service\nPerimeter.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Proposed (or dry run) ServicePerimeter configuration.\nThis configuration allows to specify and test ServicePerimeter configuration\nwithout enforcing actual access restrictions. Only allowed to be set when\nthe 'useExplicitDryRunSpec' flag is set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "status": {
        "block": {
          "attributes": {
            "access_levels": {
              "description": "A list of AccessLevel resource names that allow resources within\nthe ServicePerimeter to be accessed from the internet.\nAccessLevels listed must be in the same policy as this\nServicePerimeter. Referencing a nonexistent AccessLevel is a\nsyntax error. If no AccessLevel names are listed, resources within\nthe perimeter can only be accessed via GCP calls with request\norigins within the perimeter. For Service Perimeter Bridge, must\nbe empty.\n\nFormat: accessPolicies/{policy_id}/accessLevels/{access_level_name}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "description": "A list of GCP resources that are inside of the service perimeter.\nCurrently only projects are allowed.\nFormat: projects/{project_number}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "restricted_services": {
              "description": "GCP services that are subject to the Service Perimeter\nrestrictions. Must contain a list of services. For example, if\n'storage.googleapis.com' is specified, access to the storage\nbuckets inside the perimeter must meet the perimeter's access\nrestrictions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "egress_policies": {
              "block": {
                "block_types": {
                  "egress_from": {
                    "block": {
                      "attributes": {
                        "identities": {
                          "description": "A list of identities that are allowed access through this 'EgressPolicy'. \nShould be in the format of email address. The email address should \nrepresent individual user or service account only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "identity_type": {
                          "description": "Specifies the type of identities that are allowed access to outside the \nperimeter. If left unspecified, then members of 'identities' field will \nbe allowed access. Possible values: [\"IDENTITY_TYPE_UNSPECIFIED\", \"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Defines conditions on the source of a request causing this 'EgressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "egress_to": {
                    "block": {
                      "attributes": {
                        "resources": {
                          "description": "A list of resources, currently only projects in the form \n'projects/\u003cprojectnumber\u003e', that match this to stanza. A request matches \nif it contains a resource in this list. If * is specified for resources, \nthen this 'EgressTo' rule will authorize access to all resources outside \nthe perimeter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "operations": {
                          "block": {
                            "attributes": {
                              "service_name": {
                                "description": "The name of the API whose methods or permissions the 'IngressPolicy' or \n'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName \nfield set to '*' will allow all methods AND permissions for all services.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "method_selectors": {
                                "block": {
                                  "attributes": {
                                    "method": {
                                      "description": "Value for 'method' should be a valid method name for the corresponding \n'serviceName' in 'ApiOperation'. If '*' used as value for method, \nthen ALL methods and permissions are allowed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "permission": {
                                      "description": "Value for permission should be a valid Cloud IAM permission for the \ncorresponding 'serviceName' in 'ApiOperation'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "API methods or permissions to allow. Method or permission must belong \nto the service specified by 'serviceName' field. A single MethodSelector \nentry with '*' specified for the 'method' field will allow all methods \nAND permissions for the service specified in 'serviceName'.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of 'ApiOperations' that this egress rule applies to. A request matches \nif it contains an operation/service in this list.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the 'ApiOperation' and destination resources that \ncause this 'EgressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of EgressPolicies to apply to the perimeter. A perimeter may \nhave multiple EgressPolicies, each of which is evaluated separately.\nAccess is granted if any EgressPolicy grants it. Must be empty for \na perimeter bridge.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ingress_policies": {
              "block": {
                "block_types": {
                  "ingress_from": {
                    "block": {
                      "attributes": {
                        "identities": {
                          "description": "A list of identities that are allowed access through this ingress policy.\nShould be in the format of email address. The email address should represent \nindividual user or service account only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "identity_type": {
                          "description": "Specifies the type of identities that are allowed access from outside the \nperimeter. If left unspecified, then members of 'identities' field will be \nallowed access. Possible values: [\"IDENTITY_TYPE_UNSPECIFIED\", \"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "sources": {
                          "block": {
                            "attributes": {
                              "access_level": {
                                "description": "An 'AccessLevel' resource name that allow resources within the \n'ServicePerimeters' to be accessed from the internet. 'AccessLevels' listed \nmust be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent\n'AccessLevel' will cause an error. If no 'AccessLevel' names are listed, \nresources within the perimeter can only be accessed via Google Cloud calls \nwith request origins within the perimeter. \nExample 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.' \nIf * is specified, then all IngressSources will be allowed.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "resource": {
                                "description": "A Google Cloud resource that is allowed to ingress the perimeter. \nRequests from these resources will be allowed to access perimeter data. \nCurrently only projects are allowed. Format 'projects/{project_number}' \nThe project may be in any Google Cloud organization, not just the \norganization that the perimeter is defined in. '*' is not allowed, the case \nof allowing all Google Cloud resources only is not supported.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Sources that this 'IngressPolicy' authorizes access from.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the source of a request causing this 'IngressPolicy'\nto apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ingress_to": {
                    "block": {
                      "attributes": {
                        "resources": {
                          "description": "A list of resources, currently only projects in the form \n'projects/\u003cprojectnumber\u003e', protected by this 'ServicePerimeter'\nthat are allowed to be accessed by sources defined in the\ncorresponding 'IngressFrom'. A request matches if it contains\na resource in this list. If '*' is specified for resources,\nthen this 'IngressTo' rule will authorize access to all \nresources inside the perimeter, provided that the request\nalso matches the 'operations' field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "operations": {
                          "block": {
                            "attributes": {
                              "service_name": {
                                "description": "The name of the API whose methods or permissions the 'IngressPolicy' or \n'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName' \nfield set to '*' will allow all methods AND permissions for all services.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "method_selectors": {
                                "block": {
                                  "attributes": {
                                    "method": {
                                      "description": "Value for method should be a valid method name for the corresponding \nserviceName in 'ApiOperation'. If '*' used as value for 'method', then \nALL methods and permissions are allowed.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "permission": {
                                      "description": "Value for permission should be a valid Cloud IAM permission for the \ncorresponding 'serviceName' in 'ApiOperation'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "API methods or permissions to allow. Method or permission must belong to \nthe service specified by serviceName field. A single 'MethodSelector' entry \nwith '*' specified for the method field will allow all methods AND \npermissions for the service specified in 'serviceName'.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom' \nare allowed to perform in this 'ServicePerimeter'.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the conditions on the 'ApiOperation' and request destination that cause\nthis 'IngressPolicy' to apply.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of 'IngressPolicies' to apply to the perimeter. A perimeter may\nhave multiple 'IngressPolicies', each of which is evaluated\nseparately. Access is granted if any 'Ingress Policy' grants it.\nMust be empty for a perimeter bridge.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_accessible_services": {
              "block": {
                "attributes": {
                  "allowed_services": {
                    "description": "The list of APIs usable within the Service Perimeter.\nMust be empty unless 'enableRestriction' is True.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "enable_restriction": {
                    "description": "Whether to restrict API calls within the Service Perimeter to the\nlist of APIs specified in 'allowedServices'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Specifies how APIs are allowed to communicate within the Service\nPerimeter.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "ServicePerimeter configuration. Specifies sets of resources,\nrestricted services and access levels that determine\nperimeter content and boundaries.",
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

func GoogleAccessContextManagerServicePerimeterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerServicePerimeter), &result)
	return &result
}
