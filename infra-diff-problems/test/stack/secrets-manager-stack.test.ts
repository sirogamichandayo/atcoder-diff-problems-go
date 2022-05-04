import { App } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { SecretsManagerStack } from '../../lib/stack/secrets-manager-stack';

const app = new App({
  context: {
    'systemName': 'sin',
    'envType': 'stg'
  }
});
const secretsManagerStack = new SecretsManagerStack(app, 'SecretsManagerStack');
const template = Template.fromStack(secretsManagerStack);

test("Secrets Manager", () => {
  template.resourceCountIs("AWS::SecretsManager::Secret", 1);
  template.hasResourceProperties("AWS::SecretsManager::Secret", {
    Description: "for RDS cluster",
    GenerateSecretString: {
      ExcludeCharacters: `\"@/\\\'`,
      GenerateStringKey: "MasterUserPassword",
      PasswordLength: 16,
      SecretStringTemplate: `{"MasterUsername": "admin"}`
    },
    Name: "sin-stg-secret-rds-cluster"
  });
});

