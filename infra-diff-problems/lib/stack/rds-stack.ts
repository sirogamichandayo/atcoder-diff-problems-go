import {VpcStack} from "./vpc-stack";
import {Construct} from "constructs";
import {IamStack} from "./iam-stack";
import {SecretsManagerStack} from "./secrets-manager-stack";
import {Ec2Stack} from "./ec2-stack";
import {Stack, StackProps} from "aws-cdk-lib";
import {RdsSubnetGroup} from "../resource/rds-subnet-group";
import {RdsParameterGroup} from "../resource/rds-parameter-group";
import {RdsDatabase} from "../resource/rds-database";


export class RdsStack extends Stack {
  constructor(
    scope: Construct,
    id: string,
    vpcStack: VpcStack,
    iamStack: IamStack,
    ec2Stack: Ec2Stack,
    secretsManagerStack: SecretsManagerStack,
    props?: StackProps
  ) {
    super(scope, id, props);

    const subnetGroup = new RdsSubnetGroup(this, vpcStack.subnet);

    const parameterGroup = new RdsParameterGroup(this);

    new RdsDatabase(
      this,
      subnetGroup,
      parameterGroup,
      secretsManagerStack.secret,
      ec2Stack.securityGroup,
      iamStack.role
    );
  }
}