package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketObjectContent = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the containing bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cache_control": {
        "computed": true,
        "description": "Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "description": "Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_disposition": {
        "computed": true,
        "description": "Content-Disposition of the object data.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_encoding": {
        "computed": true,
        "description": "Content-Encoding of the object data.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_language": {
        "computed": true,
        "description": "Content-Language of the object data.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "Content-Type of the object data. Defaults to \"application/octet-stream\" or \"text/plain; charset=utf-8\".",
        "description_kind": "plain",
        "type": "string"
      },
      "crc32c": {
        "computed": true,
        "description": "Base 64 CRC32 hash of the uploaded data.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_encryption": {
        "computed": true,
        "description": "Encryption key; encoded using base64.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "encryption_algorithm": "string",
              "encryption_key": "string"
            }
          ]
        ]
      },
      "detect_md5hash": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_based_hold": {
        "computed": true,
        "description": "Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName value, if any.",
        "description_kind": "plain",
        "type": "string"
      },
      "md5hash": {
        "computed": true,
        "description": "Base 64 MD5 hash of the uploaded data.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_link": {
        "computed": true,
        "description": "A url reference to download this object.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "User-provided metadata, in key/value pairs.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the object. If you're interpolating the name of this object, see output_name instead.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_name": {
        "computed": true,
        "description": "The name of the object. Use this field in interpolations with google_storage_object_acl to recreate google_storage_object_acl resources when your google_storage_bucket_object is recreated.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "A url reference to this object.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "A path to the data you want to upload. Must be defined if content is not.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description": "The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.",
        "description_kind": "plain",
        "type": "string"
      },
      "temporary_hold": {
        "computed": true,
        "description": "Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageBucketObjectContentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucketObjectContent), &result)
	return &result
}
