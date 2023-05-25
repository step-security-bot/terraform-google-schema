package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocCluster = `{
  "block": {
    "attributes": {
      "graceful_decommission_timeout": {
        "description": "The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a terraform apply",
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
      "labels": {
        "computed": true,
        "description": "The list of labels (key/value pairs) to be applied to instances in the cluster. GCP generates some itself including goog-dataproc-cluster-name which is the name of the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the cluster, unique within the project and zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region in which the cluster and associated nodes will be created in. Defaults to global.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cluster_config": {
        "block": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": " The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If staging_bucket is specified, it will contain this value, otherwise it will be the auto generated name.",
              "description_kind": "plain",
              "type": "string"
            },
            "staging_bucket": {
              "description": "The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster. Note: If you don't explicitly specify a staging_bucket then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "temp_bucket": {
              "computed": true,
              "description": "The Cloud Storage temp bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. Note: If you don't explicitly specify a temp_bucket then GCP will auto create / assign one for you.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "autoscaling_config": {
              "block": {
                "attributes": {
                  "policy_uri": {
                    "description": "The autoscaling policy used by the cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The autoscaling policy config associated with the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The Customer managed encryption keys settings for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gce_cluster_config": {
              "block": {
                "attributes": {
                  "internal_ip_only": {
                    "description": "By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. If set to true, all instances in the cluster will only have internal IP addresses. Note: Private Google Access (also known as privateIpGoogleAccess) must be enabled on the subnetwork that the cluster will be launched in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "metadata": {
                    "description": "A map of the Compute Engine metadata entries to add to all instances",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "network": {
                    "computed": true,
                    "description": "The name or self_link of the Google Compute Engine network to the cluster will be part of. Conflicts with subnetwork. If neither is specified, this defaults to the \"default\" network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account": {
                    "description": "The service account to be used by the Node VMs. If not specified, the \"default\" service account is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_scopes": {
                    "computed": true,
                    "description": "The set of Google API scopes to be made available on all of the node VMs under the service_account specified. These can be either FQDNs, or scope aliases.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnetwork": {
                    "description": "The name or self_link of the Google Compute Engine subnetwork the cluster will be part of. Conflicts with network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "description": "The list of instance tags applied to instances in the cluster. Tags are used to identify valid sources or targets for network firewalls.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "The GCP zone where your data is stored and used (i.e. where the master and the worker nodes will be created in). If region is set to 'global' (default) then zone is mandatory, otherwise GCP is able to make use of Auto Zone Placement to determine this automatically for you. Note: This setting additionally determines and restricts which computing resources are available for use with other configs such as cluster_config.master_config.machine_type and cluster_config.worker_config.machine_type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "shielded_instance_config": {
                    "block": {
                      "attributes": {
                        "enable_integrity_monitoring": {
                          "description": "Defines whether instances have integrity monitoring enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_secure_boot": {
                          "description": "Defines whether instances have Secure Boot enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_vtpm": {
                          "description": "Defines whether instances have the vTPM enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Shielded Instance Config for clusters using Compute Engine Shielded VMs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Common config settings for resources of Google Compute Engine cluster instances, applicable to all instances in the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "initialization_action": {
              "block": {
                "attributes": {
                  "script": {
                    "description": "The script to be executed during initialization of the cluster. The script must be a GCS file with a gs:// prefix.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout_sec": {
                    "description": "The maximum duration (in seconds) which script is allowed to take to execute its action. GCP will default to a predetermined computed value if not set (currently 300).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Commands to execute on each node after config is completed. You can specify multiple versions of these.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "master_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description": "The URI for the image to use for this master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description": "List of master/worker instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The name of a Google Compute Engine machine type to create for the master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "The name of a minimum generation of CPU family for the master/worker. If not specified, GCP will default to a predetermined computed value for each zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of master/worker nodes to create. If not specified, GCP will default to a predetermined computed value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each node. One of \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Google Compute Engine config settings for the master/worker instances in a cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "preemptible_worker_config": {
              "block": {
                "attributes": {
                  "instance_names": {
                    "computed": true,
                    "description": "List of preemptible instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of preemptible nodes to create. Defaults to 0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each preemptible worker node, specified in GB. The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each preemptible worker node. One of \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each preemptible worker node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Google Compute Engine config settings for the additional (aka preemptible) instances in a cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "security_config": {
              "block": {
                "block_types": {
                  "kerberos_config": {
                    "block": {
                      "attributes": {
                        "cross_realm_trust_admin_server": {
                          "description": "The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_kdc": {
                          "description": "The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_realm": {
                          "description": "The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_shared_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster\nKerberos realm and the remote trusted realm, in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enable_kerberos": {
                          "description": "Flag to indicate whether to Kerberize the cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "kdc_db_key_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keystore_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing\nthe password to the user provided keystore. For the self-signed certificate, this password is generated\nby Dataproc",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keystore_uri": {
                          "description": "The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key_uri": {
                          "description": "The uri of the KMS key used to encrypt various sensitive files.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "realm": {
                          "description": "The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "root_principal_password_uri": {
                          "description": "The cloud Storage URI of a KMS encrypted file containing the root principal password.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "tgt_lifetime_hours": {
                          "description": "The lifetime of the ticket granting ticket, in hours.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "truststore_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "truststore_uri": {
                          "description": "The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Kerberos related configuration",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Security related configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "software_config": {
              "block": {
                "attributes": {
                  "image_version": {
                    "computed": true,
                    "description": "The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters. If not specified, defaults to the latest version.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "optional_components": {
                    "description": "The set of optional components to activate on the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "override_properties": {
                    "description": "A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "properties": {
                    "computed": true,
                    "description": "A list of the properties used to set the daemon config files. This will include any values supplied by the user via cluster_config.software_config.override_properties",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The config settings for software inside the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "worker_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description": "The URI for the image to use for this master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description": "List of master/worker instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The name of a Google Compute Engine machine type to create for the master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "The name of a minimum generation of CPU family for the master/worker. If not specified, GCP will default to a predetermined computed value for each zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of master/worker nodes to create. If not specified, GCP will default to a predetermined computed value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each node. One of \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Google Compute Engine config settings for the master/worker instances in a cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Allows you to configure various aspects of the cluster.",
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

func GoogleDataprocClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocCluster), &result)
	return &result
}
