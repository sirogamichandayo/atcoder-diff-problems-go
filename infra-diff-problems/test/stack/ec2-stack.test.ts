import { App} from "aws-cdk-lib";
import { Match, Template } from "aws-cdk-lib/assertions";
import { VpcStack } from "../../lib/stack/vpc-stack";
import { Ec2Stack } from "../../lib/stack/ec2-stack";
import { IamStack } from "../../lib/stack/iam-stack";

const app = new App({
  context: {
    "systemName": "sin",
    "envType": "stg"
  }
});
const vpcStack = new VpcStack(app, "VpcStack");
const iamStack = new IamStack(app, "IamStack");
const ec2Stack = new Ec2Stack(app, "Ec2Stack", vpcStack, iamStack);
const template = Template.fromStack(ec2Stack);

test("SecurityGroup", () => {
  template.resourceCountIs("AWS::EC2::SecurityGroup", 3);
  template.hasResourceProperties("AWS::EC2::SecurityGroup", {
    GroupDescription: "for ALB",
    GroupName: "sin-stg-sg-alb",
    VpcId: Match.anyValue(),
    Tags: [{ Key: "Name", Value: "sin-stg-sg-alb" }]
  });
  template.hasResourceProperties("AWS::EC2::SecurityGroup", {
    GroupDescription: "for EC2",
    GroupName: "sin-stg-sg-ec2",
    VpcId: Match.anyValue(),
    Tags: [{ Key: "Name", Value: "sin-stg-sg-ec2"}]
  });
  template.hasResourceProperties("AWS::EC2::SecurityGroup", {
    GroupDescription: "for RDS",
    GroupName: "sin-stg-sg-rds",
    VpcId: Match.anyValue(),
    Tags: [{ Key: "Name", Value: "sin-stg-sg-rds"}]
  });

  template.resourceCountIs("AWS::EC2::SecurityGroupIngress", 4);
  template.hasResourceProperties("AWS::EC2::SecurityGroupIngress", {
    IpProtocol: "tcp",
    CidrIp: "0.0.0.0/0",
    FromPort: 80,
    ToPort: 80,
    GroupId: Match.anyValue()
  });
  template.hasResourceProperties("AWS::EC2::SecurityGroupIngress", {
    IpProtocol: "tcp",
    CidrIp: "0.0.0.0/0",
    FromPort: 443,
    ToPort: 443,
    GroupId: Match.anyValue()
  });
  template.hasResourceProperties("AWS::EC2::SecurityGroupIngress", {
    IpProtocol: "tcp",
    FromPort: 80,
    ToPort: 80,
    GroupId: Match.anyValue(),
    SourceSecurityGroupId: Match.anyValue()
  });
  template.hasResourceProperties("AWS::EC2::SecurityGroupIngress", {
    IpProtocol: "tcp",
    FromPort: 3306,
    ToPort: 3306,
    GroupId: Match.anyValue(),
    SourceSecurityGroupId: Match.anyValue()
  });
})

test("Instance", () => {
  template.resourceCountIs("AWS::EC2::Instance", 2);
  template.hasResourceProperties("AWS::EC2::Instance", {
    AvailabilityZone: "ap-northeast-1a",
    IamInstanceProfile: Match.anyValue(),
    ImageId: "ami-0bcc04d20228d0cf6",
    InstanceType: "t2.micro",
    SecurityGroupIds: Match.anyValue(),
    SubnetId: Match.anyValue(),
    Tags: [{ Key: "Name", Value: "sin-stg-ec2-1a"}],
    UserData: Match.anyValue()
  });
  template.hasResourceProperties("AWS::EC2::Instance", {
    AvailabilityZone: "ap-northeast-1c",
    IamInstanceProfile: Match.anyValue(),
    ImageId: "ami-0bcc04d20228d0cf6",
    InstanceType: "t2.micro",
    SecurityGroupIds: Match.anyValue(),
    SubnetId: Match.anyValue(),
    Tags: [{ Key: "Name", Value: "sin-stg-ec2-1c"}],
    UserData: Match.anyValue()
  });
});

test("Target Group", () => {
  template.resourceCountIs("AWS::ElasticLoadBalancingV2::TargetGroup", 1);
  template.hasResourceProperties("AWS::ElasticLoadBalancingV2::TargetGroup", {
    Name: "sin-stg-tg",
    Port: 80,
    Protocol:"HTTP",
    TargetType: "instance",
    Targets: Match.anyValue(),
    VpcId: Match.anyValue()
  });
});

test("Load Balancer", () => {
  template.resourceCountIs("AWS::ElasticLoadBalancingV2::LoadBalancer", 1);
  template.hasResourceProperties("AWS::ElasticLoadBalancingV2::LoadBalancer", {
    IpAddressType: "ipv4",
    Name: "sin-stg-alb",
    Scheme: "internet-facing",
    SecurityGroups: Match.anyValue(),
    Subnets: Match.anyValue(),
    Type: "application"
  });
});

test("Listener", () => {
  template.resourceCountIs("AWS::ElasticLoadBalancingV2::Listener", 1);
  template.hasResourceProperties("AWS::ElasticLoadBalancingV2::Listener", {
    DefaultActions: [{
      Type: "forward",
      ForwardConfig: {
        TargetGroups: [{
          TargetGroupArn: Match.anyValue(),
          Weight: 1
        }]
      }
    }],
    LoadBalancerArn: Match.anyValue(),
    Port: 80,
    Protocol: "HTTP"
  });
})