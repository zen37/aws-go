// https://raw.githubusercontent.com/awsdocs/aws-doc-sdk-examples/main/go/example_code/lambda/aws-go-sdk-lambda-example-show-functions.go

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"fmt"
	"os"
)

func init() {
	//https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
	//the configuration settings should be read from the config file
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
}

// Lists all of your Lambda functions in us-west-2
func main() {
	// Initialize a session
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(sess)

	// Create Lambda service client
	svc := lambda.New(sess)

	result, err := svc.ListFunctions(nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot list functions")
		os.Exit(0)
	}

	fmt.Println("Functions:")

	for _, f := range result.Functions {
		fmt.Println("Name:        " + aws.StringValue(f.FunctionName))
		fmt.Println("Description: " + aws.StringValue(f.Description))
		fmt.Println("")
	}
}
