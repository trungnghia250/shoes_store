package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Sess *session.Session

func ConnectAws() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			"",
			"",
			""),
	})
	if err != nil {
		return err
	}

	Sess = sess

	return nil
}
