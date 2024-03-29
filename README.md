# ec2tags

## DEPRECATED

This project is no longer necessary, as tags are now available from instance metadata:

https://aws.amazon.com/about-aws/whats-new/2022/01/instance-tags-amazon-ec2-instance-metadata-service/

## Introduction

Tagging AWS EC2 instances, particularly via autoscaling groups, is a
great way to convey information about the role and intended behaviour
of the instance.

This simple script enables an instance to discover its own tags and
write them as environment variables for consumption by applications.

## Download

```
curl -L https://github.com/rlister/ec2tags/releases/download/0.1/ec2tags-linux-amd64-0.1.gz | gunzip > ec2tags
chmod 755 ec2tags
```

## Authentication

It is preferable to authenticate an instance via IAM roles. The
following should suffice:

```json
"Action": [
  "ec2:Describe*"
],
"Resource": [
  "*"
]
```

If IAM roles are not possible for some reason, or for testing from a
non-AWS system, you may authenticate via environment variables in the
AWS-standard way by setting:

```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION
```

## Usage

```
ec2tags [instance_id]
```

If an `instance_id` is not provided, it will be discovered from
instance metadata.

Tags are uppercased and formatted as environment variables, in the
form: `FOO=bar`.

## Docker

Comes with a Dockerfile for building a minimalist from-scratch
image. To build with a static binary:

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' ec2tags.go
docker build -t rlister/ec2tags .
```

If you are not on linux architecture you will need a golang built with
cross-compiler support.

To run:

```
docker run \
  -e AWS_ACCESS_KEY_ID \
  -e AWS_SECRET_ACCESS_KEY \
  -e AWS_DEFAULT_REGION \
  rlister/ec2tags [instance-id]
```

The AWS environment vars are unncessary if using IAM roles.
