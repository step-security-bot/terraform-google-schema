package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingProjectBucketConfig = `{
  "block": {
    "attributes": {
      "bucket_id": {
        "description": "The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description for this bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_analytics": {
        "description": "Enable log analytics for the bucket. Cannot be disabled once enabled.",
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
      "lifecycle_state": {
        "computed": true,
        "description": "The bucket's lifecycle such as active or deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locked": {
        "description": "Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the bucket",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "The parent project that contains the logging bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_days": {
        "description": "Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "cmek_settings": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The resource name for the configured Cloud KMS key.\nKMS key name format:\n\"projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]\"\nTo enable CMEK for the bucket, set this field to a valid kmsKeyName for which the associated service account has the required cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.\nThe Cloud KMS key used by the bucket can be updated by changing the kmsKeyName to a new valid key name. Encryption operations that are in progress will be completed with the key that was in use when they started. Decryption operations will be completed using the key that was used at the time of encryption unless access to that key has been revoked.\nSee [Enabling CMEK for Logging Buckets](https://cloud.google.com/logging/docs/routing/managed-encryption-storage) for more information.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key_version_name": {
              "computed": true,
              "description": "The CryptoKeyVersion resource name for the configured Cloud KMS key.\nKMS key name format:\n\"projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]/cryptoKeyVersions/[VERSION]\"\nFor example:\n\"projects/my-project/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key/cryptoKeyVersions/1\"\nThis is a read-only field used to convey the specific configured CryptoKeyVersion of kms_key that has been configured. It will be populated in cases where the CMEK settings are bound to a single key version.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The resource name of the CMEK settings.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_account_id": {
              "computed": true,
              "description": "The service account associated with a project for which CMEK will apply.\nBefore enabling CMEK for a logging bucket, you must first assign the cloudkms.cryptoKeyEncrypterDecrypter role to the service account associated with the project for which CMEK will apply. Use [v2.getCmekSettings](https://cloud.google.com/logging/docs/reference/v2/rest/v2/TopLevel/getCmekSettings#google.logging.v2.ConfigServiceV2.GetCmekSettings) to obtain the service account ID.\nSee [Enabling CMEK for Logging Buckets](https://cloud.google.com/logging/docs/routing/managed-encryption-storage) for more information.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.",
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

func GoogleLoggingProjectBucketConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingProjectBucketConfig), &result)
	return &result
}
