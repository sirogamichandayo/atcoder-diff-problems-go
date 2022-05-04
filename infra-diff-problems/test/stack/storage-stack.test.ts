import {App, RemovalPolicy} from "aws-cdk-lib";
import {VpcStack} from "../../lib/stack/vpc-stack";
import {StorageStack} from "../../lib/stack/storage-stack";
import {Match, Template} from "aws-cdk-lib/assertions";
import {BucketAccessControl} from "aws-cdk-lib/aws-s3";
import {Effect} from "aws-cdk-lib/aws-iam";
import {PriceClass, ViewerProtocolPolicy} from "aws-cdk-lib/aws-cloudfront";

const app = new App({
  context: {
    "systemName": "sin",
    "envType": "stg"
  }
});
const storageStack = new StorageStack(app, "StorageStack");
const template = Template.fromStack(storageStack);

test("Bucket", () => {
  template.resourceCountIs("AWS::S3::Bucket", 1);
  template.hasResourceProperties("AWS::S3::Bucket", {
    BucketName: "sin-frontend-bucket",
    AccessControl: BucketAccessControl.PRIVATE
  });
});

test("OriginAccessIdentity", () => {
  template.resourceCountIs("AWS::CloudFront::CloudFrontOriginAccessIdentity", 1);
  template.hasResourceProperties("AWS::CloudFront::CloudFrontOriginAccessIdentity", {
    Comment: "website-distribution-originAccessIdentity"
  });
});

test("BucketPolicy", () => {
  template.resourceCountIs("AWS::S3::BucketPolicy", 1);
  template.hasResourceProperties("AWS::S3::BucketPolicy", {
    Bucket: {
      Ref: Match.anyValue()
    }
  });
});

test("Distribution", () => {
  template.resourceCountIs("AWS::CloudFront::Distribution", 1);
  template.hasResourceProperties("AWS::CloudFront::Distribution", {
    DistributionConfig: {
      Comment: "website-distribution",
      CustomErrorResponses: [
        {
          ErrorCachingMinTTL: 300,
          ErrorCode: 403,
          ResponseCode: 403,
          ResponsePagePath: "/error.html"
        },
        {
          ErrorCachingMinTTL: 300,
          ErrorCode: 404,
          ResponseCode: 404,
          ResponsePagePath: "/error.html"
        }
      ],
      DefaultCacheBehavior: {
        AllowedMethods: [
          "GET",
          "HEAD"
        ],
        CachePolicyId: "658327ea-f89d-4fab-a63d-7e88639e58f6",
        CachedMethods: [
          "GET",
          "HEAD"
        ],
        Compress: true,
        TargetOriginId: "StorageStackdistributionOrigin194558D29",
        ViewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS
      },
      DefaultRootObject: "index.html",
      Enabled: true,
      HttpVersion: "http2",
      IPV6Enabled: true,
      Origins: [
        {
          DomainName: {
            "Fn::GetAtt": [
              "WebsiteBucket75C24D94",
              "RegionalDomainName"
            ]
          },
          Id: "StorageStackdistributionOrigin194558D29",
          S3OriginConfig: {
            OriginAccessIdentity: {
              "Fn::Join": [
                "",
                [
                  "origin-access-identity/cloudfront/",
                  {
                    Ref: "OriginAccessIdentify442FA33F"
                  }
                ]
              ]
            }
          }
        }
      ],
      PriceClass: PriceClass.PRICE_CLASS_ALL
    }
  });
});

