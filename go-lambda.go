package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type LambdaStackProps struct {
	awscdk.StackProps
}

func NewLambdaStack(scope constructs.Construct, id string, props *LambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awslambda.NewFunction(stack, jsii.String("Singleton"), &awslambda.FunctionProps{
		Code:         awslambda.NewAssetCode(jsii.String("lambda"), nil),
		Handler:      jsii.String("handler.main"),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(900)),
		MemorySize:   jsii.Number(128.0),
		Description:  jsii.String("Simple Lambda"),
		FunctionName: jsii.String("MyLambda"),
		Runtime:      awslambda.Runtime_PYTHON_3_8(),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewLambdaStack(app, "GoLambdaStack", &LambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("238220415119"),
		Region:  jsii.String("us-west-2"),
	}
}
