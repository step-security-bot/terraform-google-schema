package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTpuNode = `{
  "block": {
    "attributes": {
      "accelerator_type": {
        "description": "The type of hardware accelerators associated with this node.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description": "The CIDR block that the TPU node will use when selecting an IP\naddress. This CIDR block must be a /29 block; the Compute Engine\nnetworks API forbids a smaller block, and using a larger block would\nbe wasteful (a node can only consume one IP address).\n\nErrors will occur if the CIDR block has already been used for a\ncurrently existing TPU node, the CIDR block conflicts with any\nsubnetworks in the user's provided network, or the provided network\nis peered with another network that is using that CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "The user-supplied description of the TPU. Maximum of 512 characters.",
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
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The immutable name of the TPU.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The name of a network to peer the TPU node to. It must be a\npreexisting Compute Engine network inside of the project on which\nthis API has been activated. If none is provided, \"default\" will be\nused.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_endpoints": {
        "computed": true,
        "description": "The network endpoints where TPU workers can be accessed and sent work.\nIt is recommended that Tensorflow clients of the node first reach out\nto the first (index 0) entry.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "port": "number"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "The service account used to run the tensor flow services within the\nnode. To share resources, including Google Cloud Storage data, with\nthe Tensorflow job running in the Node, this account must have\npermissions to that data.",
        "description_kind": "plain",
        "type": "string"
      },
      "tensorflow_version": {
        "description": "The version of Tensorflow running in the Node.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "use_service_networking": {
        "description": "Whether the VPC peering for the node is set up through Service Networking API.\nThe VPC Peering should be set up before provisioning the node. If this field is set,\ncidr_block field should not be specified. If the network that you want to peer the\nTPU Node to is a Shared VPC network, the node must be created with this this field enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone": {
        "computed": true,
        "description": "The GCP location for the TPU. If it is not provided, the provider zone is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "scheduling_config": {
        "block": {
          "attributes": {
            "preemptible": {
              "description": "Defines whether the TPU instance is preemptible.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Sets the scheduling options for this TPU instance.",
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

func GoogleTpuNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTpuNode), &result)
	return &result
}
