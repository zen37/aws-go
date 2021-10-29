package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func init() {
	//https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
	//the configuration settings should be read from the config file
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
}

func main() {
	/*
		sess, err := session.NewSessionWithOptions(session.Options{
			Config:            aws.Config{Region: aws.String("us-east-1")},
			SharedConfigState: session.SharedConfigEnable,
		})
		if err != nil {
			panic(err)
		}
	*/
	sess := session.Must(session.NewSession())
	//sess, err := session.NewSession()

	ssmsvc := ssm.New(sess)
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/collaction/dev/contact/email"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	value := *param.Parameter.Value
	fmt.Println(value)
}
