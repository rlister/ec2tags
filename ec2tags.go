package main

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
	"os"
	"strings"
)

func main() {

	// accept instance ID as arg, or get from aws metadata
	var id string
	if len(os.Args) > 1 {
		id = os.Args[1]
	} else {
		instance_id, _ := aws.GetMetaData("instance-id")
		id = string(instance_id)
	}

	// get creds from env or IAM role
	auth, err := aws.GetAuth("", "")
	if err != nil {
		panic(err.Error())
	}

	// get region from env or instance metadata
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		az, _ := aws.GetMetaData("placement/availability-zone")
		region = string(az[:len(az)-1]) // trim last char
	}

	// get instance by id
	e := ec2.New(auth, aws.Region(aws.Regions[region]))
	resp, err := e.Instances([]string{id}, nil)
	if err != nil {
		panic(err.Error())
	}

	// emit tags in uppercase
	for _, t := range resp.Reservations[0].Instances[0].Tags {
		fmt.Printf("%s=%s\n", strings.ToUpper(strings.Replace(t.Key, ":", "_", -1)), t.Value)
	}
}
