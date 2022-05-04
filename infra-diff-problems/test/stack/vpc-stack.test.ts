import {App} from "aws-cdk-lib";
import {VpcStack} from "../../lib/stack/vpc-stack";
import {Match, Template} from "aws-cdk-lib/assertions";


const app = new App({
  context: {
    'systemName': 'sin',
    'envType': 'stg'
  }
});
const vpcStack = new VpcStack(app, 'VpcStack');
const template = Template.fromStack(vpcStack);

test('VPC', () => {
  template.resourceCountIs('AWS::EC2::VPC', 1);
  template.hasResourceProperties('AWS::EC2::VPC', {
    CidrBlock: '10.0.0.0/16',
    Tags: [{Key: 'Name', Value: 'sin-stg-vpc'}]
  });
})

test('Subnet', () => {
  template.resourceCountIs('AWS::EC2::Subnet', 6);

  const az_a = 'ap-northeast-1a';
  const az_c = 'ap-northeast-1c';
  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.11.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_a,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-public-1a'}]
  })
  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.12.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_c,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-public-1c'}]
  })

  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.21.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_a,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-app-1a'}]
  })
  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.22.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_c,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-app-1c'}]
  })

  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.31.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_a,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-db-1a'}]
  })
  template.hasResourceProperties('AWS::EC2::Subnet', {
    CidrBlock: '10.0.32.0/24',
    VpcId: Match.anyValue(),
    AvailabilityZone: az_c,
    Tags: [{ Key: 'Name', Value: 'sin-stg-subnet-db-1c'}]
  })
})

test('Internet Gateway', () => {
  template.resourceCountIs('AWS::EC2::InternetGateway', 1);
  template.hasResourceProperties('AWS::EC2::VPCGatewayAttachment', {
    VpcId: Match.anyValue(),
    InternetGatewayId: Match.anyValue()
  })
})

test('Elastic Ip', () => {
  template.resourceCountIs('AWS::EC2::EIP', 2);
  template.hasResourceProperties('AWS::EC2::EIP', {
    Domain: 'vpc',
    Tags: [{Key: 'Name', Value: 'sin-stg-eip-ngw-1a'}]
  })
  template.hasResourceProperties('AWS::EC2::EIP', {
    Domain: 'vpc',
    Tags: [{Key: 'Name', Value: 'sin-stg-eip-ngw-1c'}]
  })
})

test('NAT Gateway', () => {
  template.resourceCountIs('AWS::EC2::NatGateway', 2);
  template.hasResourceProperties('AWS::EC2::NatGateway', {
    AllocationId: Match.anyValue(),
    SubnetId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-ngw-1a'}]
  });
  template.hasResourceProperties('AWS::EC2::NatGateway', {
    AllocationId: Match.anyValue(),
    SubnetId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-ngw-1c'}]
  });
})

test('RouteTable', () => {
  template.resourceCountIs('AWS::EC2::RouteTable', 4);
  template.hasResourceProperties('AWS::EC2::RouteTable', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-rtb-public'}]
  });
  template.hasResourceProperties('AWS::EC2::RouteTable', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-rtb-app-1a'}]
  });
  template.hasResourceProperties('AWS::EC2::RouteTable', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-rtb-app-1c'}]
  });
  template.hasResourceProperties('AWS::EC2::RouteTable', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-rtb-db'}]
  });

  template.resourceCountIs('AWS::EC2::Route', 3);
  template.hasResourceProperties('AWS::EC2::Route', {
    RouteTableId: Match.anyValue(),
    DestinationCidrBlock: '0.0.0.0/0',
    GatewayId: Match.anyValue()
  });
  template.hasResourceProperties('AWS::EC2::Route', {
    RouteTableId: Match.anyValue(),
    DestinationCidrBlock: '0.0.0.0/0',
    NatGatewayId: Match.anyValue()
  })

  template.resourceCountIs('AWS::EC2::SubnetRouteTableAssociation', 6);
  template.hasResourceProperties('AWS::EC2::SubnetRouteTableAssociation', {
    RouteTableId: Match.anyValue(),
    SubnetId: Match.anyValue()
  });
})

test('Network ACL', () => {
  template.resourceCountIs('AWS::EC2::NetworkAcl', 3);
  template.hasResourceProperties('AWS::EC2::NetworkAcl', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-nacl-public'}]
  });
  template.hasResourceProperties('AWS::EC2::NetworkAcl', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-nacl-app'}]
  });
  template.hasResourceProperties('AWS::EC2::NetworkAcl', {
    VpcId: Match.anyValue(),
    Tags: [{ Key: 'Name', Value: 'sin-stg-nacl-db'}]
  });

  template.resourceCountIs('AWS::EC2::NetworkAclEntry', 6);
  template.hasResourceProperties('AWS::EC2::NetworkAclEntry', {
    NetworkAclId: Match.anyValue(),
    Protocol: -1,
    RuleAction: 'allow',
    RuleNumber: 100,
    CidrBlock: '0.0.0.0/0',
    Egress: false
  });
  template.hasResourceProperties('AWS::EC2::NetworkAclEntry', {
    NetworkAclId: Match.anyValue(),
    Protocol: -1,
    RuleAction: 'allow',
    RuleNumber: 100,
    CidrBlock: '0.0.0.0/0',
    Egress: true
  });
  template.resourceCountIs('AWS::EC2::SubnetNetworkAclAssociation', 6);
  template.hasResourceProperties('AWS::EC2::SubnetNetworkAclAssociation', {
    NetworkAclId: Match.anyValue(),
    SubnetId: Match.anyValue()
  });
})