package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	TOKEN string = ""
)

func main() {

	fmt.Println("main")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", TOKEN),
	})

	s, err := sess.Config.Credentials.Get()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)

}
