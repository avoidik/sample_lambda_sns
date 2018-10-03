# AWS Lambda in Go

## What's inside

* Sample AWS Lambda function written in Golang with DynamoDB as a storage backend
* AWS SNS Topic used as a trigger
* Dependencies management with Dep
* Serverless framework for infrastructure deployment
* Operations with Makefile

## Dependencies

1. [serverless framework](https://serverless.com/)
1. [AWS SAM](https://docs.aws.amazon.com/lambda/latest/dg/serverless_app.html)
1. docker & docker-machine for AWS SAM

## Workflow

1. `make`
1. `make run`
1. `make deploy`
1. `test.sh`
1. `make undeploy`

## Windows notes

Additional dependency

```
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

Issue with AWS SAM and path detection in python

```
https://github.com/awslabs/aws-sam-cli/issues/461#issuecomment-403231130
```
