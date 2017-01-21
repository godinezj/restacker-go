package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	cloudformation.New(sess, &aws.Config{Region: aws.String("us-west-2")})
}
