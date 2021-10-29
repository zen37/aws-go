//https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/ses/VerifyAddress/VerifyAddress.go

package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// SendVerification sends an email to verify the recipient's address
// Inputs:
//     svc is the Amazon SES service client
//     recipient is the email address for the To value
// Output:
//     If success, nil
//     Otherwise, an error from the call to VerifyEmailAddress
func SendVerification(svc sesiface.SESAPI, recipient *string) error {
	_, err := svc.VerifyEmailAddress(&ses.VerifyEmailAddressInput{
		EmailAddress: recipient,
	})

	return err
}

func main() {

	recipient := flag.String("r", "", "The email address of the recipient")
	flag.Parse()

	if *recipient == "" {
		fmt.Println("You must supply an email address for the recipient")
		fmt.Println("-r RECIPIENT")
		return
	}
	/*
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	*/

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc := ses.New(sess)

	err = SendVerification(svc, recipient)
	if err != nil {
		fmt.Println("Got an error sending a validation email to " + *recipient)
		fmt.Println(err)
		return
	}

	fmt.Println("Email address: " + *recipient + " verified")
}
