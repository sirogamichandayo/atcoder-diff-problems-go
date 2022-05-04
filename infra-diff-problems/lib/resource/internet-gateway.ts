import {BaseResource} from "./abstract/base-resource";
import {CfnInternetGateway, CfnVPCGatewayAttachment} from "aws-cdk-lib/aws-ec2";
import {Construct} from "constructs";
import {Vpc} from "./vpc";


export class InternetGateway extends BaseResource {
  public readonly igw: CfnInternetGateway;

  constructor(scope: Construct, vpc: Vpc) {
    super();

    this.igw = new CfnInternetGateway(scope, 'InternetGateway', {
      tags: [{
        key: 'Name',
        value: this.createResourceName(scope, 'igw')
      }]
    });

    new CfnVPCGatewayAttachment(scope, 'VpcGatewayAttachment', {
      vpcId: vpc.vpc.ref,
      internetGatewayId: this.igw.ref
    });


  }
}
