import {Stack, StackProps} from "aws-cdk-lib";
import {Construct} from "constructs";
import {Role} from "../resource/role";
import {SecurityGroup} from "../resource/security-group";
import {VpcStack} from "./vpc-stack";
import {IamStack} from "./iam-stack";
import {Instance} from "../resource/instance";
import {TargetGroup} from "../resource/target-group";
import {LoadBalancer} from "../resource/load-balancer";

export class Ec2Stack extends Stack {
  public readonly securityGroup: SecurityGroup;

  constructor(
    scope: Construct,
    id: string,
    vpcStack: VpcStack,
    iamStack: IamStack,
    props?: StackProps
  ) {
    super(scope, id, props);

    this.securityGroup = new SecurityGroup(this, vpcStack.vpc);

    const instance = new Instance(this, vpcStack.subnet, iamStack.role, this.securityGroup);

    const targetGroup = new TargetGroup(this, vpcStack.vpc, instance);

    const loadBalancer = new LoadBalancer(this, this.securityGroup, vpcStack.subnet, targetGroup);
  }
}