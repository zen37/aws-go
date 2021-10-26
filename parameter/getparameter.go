// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func init() {
	//https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
	//the configuration settings should be read from the config file
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
}

// GetParameter fetches details of a parameter in SSM
// Inputs:
//     svc is an Amazon SSM service client
//     name is the name of the parameter
// Output:
//     If success, information about the parameter and nil
//     Otherwise, nil and an error from the call to GetParameter
func GetParameter(svc ssmiface.SSMAPI, name *string) (*ssm.GetParameterOutput, error) {
	results, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: name,
	})

	return results, err
}

func main() {
	parameterName := flag.String("n", "", "The name of the parameter")
	flag.Parse()

	if *parameterName == "" {
		fmt.Println("You must supply the name of the parameter")
		fmt.Println("-n NAME")
		return
	}

	sess := session.Must(session.NewSession())

	svc := ssm.New(sess)

	results, err := GetParameter(svc, parameterName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(*results.Parameter.Value)
}
