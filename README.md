# ec2tags

Tagging AWS EC2 instances, particularly via autoscaling groups, is a
great way to convey information about the role and intended behaviour
of the instance.

This simple script enables an instance to discover its own tags and
write them as environment variables for consumption by applications.

## Download

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
