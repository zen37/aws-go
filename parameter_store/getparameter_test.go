// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

// Define a mock struct to use in unit tests
type mockSSMClient struct {
	ssmiface.SSMAPI
}

func (m *mockSSMClient) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {

	// Check that required inputs exist
	if input.Name == nil || *input.Name == "" {
		return nil, errors.New("GetParameterInput.Name is nil of an empty string")
	}
	value := "hello@collaction.org"
	resp := ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Value: &value,
		},
	}
	fmt.Println("*resp.Parameter.Value:", *resp.Parameter.Value)

	return &resp, nil
}

func TestGetParameter(t *testing.T) {

	// mock resource
	name := "test-param"

	mockSvc := &mockSSMClient{}

	results, err := GetParameter(mockSvc, &name)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("parameter value is " + *results.Parameter.Value)
}
