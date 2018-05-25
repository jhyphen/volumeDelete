package main

// EBS Volume delete script *** written by Level 99 DevOps Engineer
// Will delete all EBS volumes with state available

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var region = "us-east-1"
var id string

func main() {

	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("status"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}

	res, err := svc.DescribeVolumes(input)
	if err != nil {
		fmt.Println("you fucked up", err.Error())
	}
	for _, vol := range res.Volumes {
		id := aws.StringValue(vol.VolumeId)
		if aws.StringValue(vol.State) == ec2.VolumeStateAvailable {
			fmt.Printf("volume %s no longer attached\n", id)
		}

		svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
		params := &ec2.DeleteVolumeInput{
			VolumeId: aws.String(id),
		}

		result, err := svc.DeleteVolume(params)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)

	}

}
