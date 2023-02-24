// Well Architected CDK Constructs - Lite
package wacdklite

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Well Architected Bucket Properties.
type WaBucketProps struct {
	// Specifies a canned ACL that grants predefined permissions to the bucket.
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// **Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
	// switching this to `false` in a CDK version *before* `1.126.0` will lead to
	// all objects in the bucket being deleted. Be sure to update your bucket resources
	// by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The block public access configuration of this bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
	//
	BlockPublicAccess awss3.BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// Whether Amazon S3 should use its own intermediary key to generate data keys.
	//
	// Only relevant when using KMS for encryption.
	//
	// - If not enabled, every object GET and PUT will cause an API call to KMS (with the
	//    attendant cost implications of that).
	// - If enabled, S3 will use its own time-limited key instead.
	//
	// Only relevant, when Encryption is set to `BucketEncryption.KMS` or `BucketEncryption.KMS_MANAGED`.
	BucketKeyEnabled *bool `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Physical name of this bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	Cors *[]*awss3.CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	Encryption awss3.BucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The 'encryption' property must be either not specified or set to "Kms".
	// An error will be emitted if encryption is set to "Unencrypted" or
	// "Managed".
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enforces SSL for requests.
	//
	// S3.5 of the AWS Foundational Security Best Practices Regarding S3.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this bucket should send notifications to Amazon EventBridge or not.
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Inteligent Tiering Configurations.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html
	//
	IntelligentTieringConfigurations *[]*awss3.IntelligentTieringConfiguration `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
	//
	Inventories *[]*awss3.Inventory `field:"optional" json:"inventories" yaml:"inventories"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleRules *[]*awss3.LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// The metrics configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
	//
	Metrics *[]*awss3.BucketMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// The role to be used by the notifications handler.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The default retention mode and rules for S3 Object Lock.
	//
	// Default retention can be configured after a bucket is created if the bucket already
	// has object lock enabled. Enabling object lock for existing buckets is not supported.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-bucket-config-enable
	//
	ObjectLockDefaultRetention awss3.ObjectLockRetention `field:"optional" json:"objectLockDefaultRetention" yaml:"objectLockDefaultRetention"`
	// Enable object lock on the bucket.
	//
	// Enabling object lock for existing buckets is not supported. Object lock must be
	// enabled when the bucket is created.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-bucket-config-enable
	//
	ObjectLockEnabled *bool `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// The objectOwnership of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
	//
	ObjectOwnership awss3.ObjectOwnership `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
	// Grants public read access to all objects in the bucket.
	//
	// Similar to calling `bucket.grantPublicAccess()`
	PublicReadAccess *bool `field:"optional" json:"publicReadAccess" yaml:"publicReadAccess"`
	// Policy to apply when the bucket is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Destination bucket for the server access logs.
	ServerAccessLogsBucket awss3.IBucket `field:"optional" json:"serverAccessLogsBucket" yaml:"serverAccessLogsBucket"`
	// Optional log file prefix to use for the bucket's access logs.
	//
	// If defined without "serverAccessLogsBucket", enables access logs to current bucket with this prefix.
	ServerAccessLogsPrefix *string `field:"optional" json:"serverAccessLogsPrefix" yaml:"serverAccessLogsPrefix"`
	// Whether this bucket should have transfer acceleration turned on or not.
	TransferAcceleration *bool `field:"optional" json:"transferAcceleration" yaml:"transferAcceleration"`
	// Whether this bucket should have versioning turned on or not.
	Versioned *bool `field:"optional" json:"versioned" yaml:"versioned"`
	// The name of the error document (e.g. "404.html") for the website. `websiteIndexDocument` must also be set if this is set.
	WebsiteErrorDocument *string `field:"optional" json:"websiteErrorDocument" yaml:"websiteErrorDocument"`
	// The name of the index document (e.g. "index.html") for the website. Enables static website hosting for this bucket.
	WebsiteIndexDocument *string `field:"optional" json:"websiteIndexDocument" yaml:"websiteIndexDocument"`
	// Specifies the redirect behavior of all requests to a website endpoint of a bucket.
	//
	// If you specify this property, you can't specify "websiteIndexDocument", "websiteErrorDocument" nor , "websiteRoutingRules".
	WebsiteRedirect *awss3.RedirectTarget `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
	// Rules that define when a redirect is applied and the redirect behavior.
	WebsiteRoutingRules *[]*awss3.RoutingRule `field:"optional" json:"websiteRoutingRules" yaml:"websiteRoutingRules"`
	// This flag is used to skip the default alarms.
	WaDoNotAddDefaultAlarms *bool `field:"optional" json:"waDoNotAddDefaultAlarms" yaml:"waDoNotAddDefaultAlarms"`
}

