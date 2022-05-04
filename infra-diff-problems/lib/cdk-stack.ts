import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import {VpcStack} from "./stack/vpc-stack";
import {IamStack} from "./stack/iam-stack";
import {Ec2Stack} from "./stack/ec2-stack";
import {SecretsManagerStack} from "./stack/secrets-manager-stack";
import {RdsStack} from "./stack/rds-stack";
import {StorageStack} from "./stack/storage-stack";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    new StorageStack(scope, "StorageStack", {
      stackName: this.createStackName(scope, "storage")
    });

    const vpcStack = new VpcStack(scope, 'VpcStack', {
      stackName: this.createStackName(scope, 'vpc')
    });

    const iamStack = new IamStack(scope, 'IamStack', {
      stackName: this.createStackName(scope, 'iam')
    });

    const ec2Stack = new Ec2Stack(scope, 'Ec2Stack', vpcStack, iamStack, {
      stackName: this.createStackName(scope, 'ec2')
    });

    const secretsManagerStack = new SecretsManagerStack(scope, 'SecretsManagerStack', {
      stackName: this.createStackName(scope, 'secrets-manager')
    });

    new RdsStack(
      scope, "RdsStack", vpcStack, iamStack, ec2Stack, secretsManagerStack,
      { stackName: this.createStackName(scope, "rds")}
    )
  }

  private createStackName(scope: Construct, originalName: string): string {
    const systemName = scope.node.tryGetContext('systemName');
    const envType = scope.node.tryGetContext('envType');
    const stackNamePrefix = `${systemName}-${envType}-stack-`;

    return `${stackNamePrefix}${originalName}`;
  }
}
