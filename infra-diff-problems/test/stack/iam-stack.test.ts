import {App} from "aws-cdk-lib";
import {Match, Template} from "aws-cdk-lib/assertions";
import {IamStack} from "../../lib/stack/iam-stack";


const app = new App({
  context: {
    'systemName': 'sin',
    'envType': 'stg'
  }
});
const iamStack = new IamStack(app, 'VpcStack');
const template = Template.fromStack(iamStack);

test('Role', () => {
  template.resourceCountIs('AWS::IAM::Role', 2);
  template.hasResourceProperties('AWS::IAM::Role', {
    AssumeRolePolicyDocument: {
      Statement: [{
        Effect: 'Allow',
        Principal: {
          Service: Match.anyValue()
        },
        Action: 'sts:AssumeRole'
      }]
    },
    ManagedPolicyArns: [
      'arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore'
    ],
    RoleName: 'sin-stg-role-ec2'
  });
  template.hasResourceProperties('AWS::IAM::Role', {
    AssumeRolePolicyDocument: {
      Statement: [{
        Effect: 'Allow',
        Principal: {
          Service: 'monitoring.rds.amazonaws.com'
        },
        Action: 'sts:AssumeRole'
      }]
    },
    ManagedPolicyArns: [
      'arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole'
    ],
    RoleName: 'sin-stg-role-rds'
  });

  template.resourceCountIs('AWS::IAM::InstanceProfile', 1);
  template.hasResourceProperties('AWS::IAM::InstanceProfile', {
    Roles: Match.anyValue(),
    InstanceProfileName: 'sin-stg-role-ec2'
  });
})