import { Template } from 'aws-cdk-lib/assertions';
import { DocBoxStack } from '../src/doc-box-stack';
import { TestApp } from './cdk-test-helper';


test('DocBoxStackSnapshot', () => {
  const projectName = 'torenken-doc-box';

  const app = new TestApp();
  const stack = new DocBoxStack(app, 'DocBoxStack', {
    documentBucketName: `${projectName}-storage`,
  });

  const template = Template.fromStack(stack);

  template.hasResourceProperties('AWS::Lambda::Function', {
    Environment: {
      Variables: {
        APPLICATION_CONTEXT: 'DocBox#createDocument',
        DOCUMENT_TABLE_NAME: {
          Ref: 'DocumentTable9FE6D880',
        },
      },
    },
  });

  template.hasResourceProperties('AWS::Lambda::Function', {
    Environment: {
      Variables: {
        APPLICATION_CONTEXT: 'DocBox#attachDocument',
        DOCUMENT_STORAGE_NAME: {
          Ref: 'DocumentBucketAE41E5A9',
        },
      },
    },
  });


  expect(template.toJSON()).toMatchSnapshot();
});