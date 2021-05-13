# Welcome to your CDK Go project

This is a blank project for Go development with CDK.

**NOTICE**: Go support is still in Developer Preview. This implies that APIs may
change while we address early feedback from the community. We would love to hear
about your experience through GitHub issues.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

## Bootstrap

One-time:

```bash
cdk bootstrap aws://238220415119/us-west-2 --profile personal
```

## Steps to test and deploy

```bash
go test
cdk deploy
```
