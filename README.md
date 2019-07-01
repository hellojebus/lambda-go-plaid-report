# Go Lambda workflow for AWS

This is a basic template repo for deploying an AWS Lambda function in Go. It uses AWS Lambda, S3, and Cloudformation for automating storage and deployment of Go lambda functions.

You can either fork or use as template through Github.   

## Requirements

- [Go Lang](https://golang.org/doc/install)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- Proper `aws configure` settings that will allow you to connect to AWS Lambda, S3, IAM, and CloudFormation  

## Instructions

Most of the commands have been extracted into a make file. To run any of the the following commands type `make COMMAND_GOES_HERE`

List of commands:

- `test` - runs tests using `go test`
- `install` - installs all go dependencies
- `main` - builds the app for environment
- `lambda` - builds app for environment AWS uses in the cloud
- `api` - starts local API using Docker environment provided by `SAM CLI` (runs `build`)
- `package` - uses `sam-cli` to create `packaged.yaml` file for CloudFormation. (Runs `test` and `build` before)
- `deploy` -  uses `sam-cli` to deploy the app on AWS (Runs `test` and `package`)
- `clean` -  cleans up app by removing old `packaged.yaml` file.

## Deploying

In order to run the `make deploy` command successfully, you will need to supply the `S3_BUCKET` and `STACK_NAME` variables. Example usage below:

`S3_BUCKET=YOUR_BUCKET_NAME STACK_NAME=YOUR_STACK_NAME make deploy`
