import {Duration, RemovalPolicy, Stack, StackProps} from "aws-cdk-lib";
import {Construct} from "constructs";
import {Bucket, BucketAccessControl} from "aws-cdk-lib/aws-s3";
import {
  AllowedMethods, CachedMethods, CachePolicy,
  Distribution,
  OriginAccessIdentity, PriceClass, ViewerProtocolPolicy
} from "aws-cdk-lib/aws-cloudfront";
import {CanonicalUserPrincipal, Effect, PolicyStatement} from "aws-cdk-lib/aws-iam";
import {S3Origin} from "aws-cdk-lib/aws-cloudfront-origins";
import {BucketDeployment, Source} from "aws-cdk-lib/aws-s3-deployment";

export class StorageStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const bucket = new Bucket(this, 'WebsiteBucket' ,{
      bucketName: "sin-frontend-bucket",
      accessControl: BucketAccessControl.PRIVATE,
      removalPolicy: RemovalPolicy.DESTROY
    });

    const originAccessIdentity = new OriginAccessIdentity(
      this,
      "OriginAccessIdentify",
      {
        comment: "website-distribution-originAccessIdentity2"
      }
    );

    const websiteBucketPolicyStatement = new PolicyStatement({
      actions: ["s3:GetObject"],
      effect: Effect.ALLOW,
      principals: [
        new CanonicalUserPrincipal(
          originAccessIdentity.cloudFrontOriginAccessIdentityS3CanonicalUserId
        ),
      ],
      resources: [`${bucket.bucketArn}/*`],
    });
    bucket.addToResourcePolicy(websiteBucketPolicyStatement);

    const distribution = new Distribution(this, "distribution", {
      comment: "website-distribution",
      defaultRootObject: "index.html",
      errorResponses: [
        {
          ttl: Duration.seconds(300),
          httpStatus: 403,
          responseHttpStatus: 403,
          responsePagePath: "/error.html"
        },
        {
          ttl: Duration.seconds(300),
          httpStatus: 404,
          responseHttpStatus: 404,
          responsePagePath: "/error.html",
        },
      ],
      defaultBehavior: {
        allowedMethods: AllowedMethods.ALLOW_GET_HEAD,
        cachedMethods: CachedMethods.CACHE_GET_HEAD,
        cachePolicy: CachePolicy.CACHING_OPTIMIZED,
        viewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        origin: new S3Origin(bucket, {
          originAccessIdentity,
        }),
      },
      priceClass: PriceClass.PRICE_CLASS_ALL
    });

    new BucketDeployment(this, "WebsiteDeploy", {
      sources: [Source.asset("../frontend-diff-problems/build")],
      destinationBucket: bucket,
      distribution: distribution,
      distributionPaths: ["/*"],
    });
  }
}