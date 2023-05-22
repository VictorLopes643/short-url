package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	// "github.com/aws/aws-lambda-go/lambda"
)

type AppServerlessCdkGoStackProps struct {
	awscdk.StackProps
}

func NewAppServerlessCdkGoStack(scope constructs.Construct, id string, props *AppServerlessCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create Lambda functionss
	urlHash := awslambda.NewFunction(stack, jsii.String("urlHash"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("main"),
		Code:    awslambda.Code_FromAsset(jsii.String("../go/cmd/urls/"), nil),
	})

	redirectionUrl := awslambda.NewFunction(stack, jsii.String("redirectionUrl"),&awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("main"),
		Code: awslambda.Code_FromAsset(jsii.String("../go/cmd/redirectionUrl/"), nil),
	})
	
	operations := awslambda.NewFunction(stack, jsii.String("operationsUrl"),&awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("main"),
		Code: awslambda.Code_FromAsset(jsii.String("../go/cmd/operations/"), nil),
	})
	// create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("ml-api"), &awsapigateway.RestApiProps{
		RestApiName:    jsii.String("ml-api"),
		CloudWatchRole: jsii.Bool(false),
	})

	
	// create API Gateway resource
	restApi.Root().AddResource(jsii.String("urlHash"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(urlHash, &awsapigateway.LambdaIntegrationOptions{}),
		restApi.Root().DefaultMethodOptions(),
	)
	restApi.Root().AddResource(jsii.String("redirectionUrl"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(redirectionUrl, &awsapigateway.LambdaIntegrationOptions{}),
		restApi.Root().DefaultMethodOptions(),
	)

	operationsUrlResource := restApi.Root().AddResource(jsii.String("operationsUrl"), nil)

		operationsUrlResource.AddMethod(
			jsii.String("GET"),
			awsapigateway.NewLambdaIntegration(operations, nil),
			restApi.Root().DefaultMethodOptions(),
		)

		operationsUrlResource.AddMethod(
			jsii.String("DELETE"),
			awsapigateway.NewLambdaIntegration(operations, nil),
			restApi.Root().DefaultMethodOptions(),
		)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewAppServerlessCdkGoStack(app, "StackML-", &AppServerlessCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}