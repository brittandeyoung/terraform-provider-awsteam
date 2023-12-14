# Terraform Provider AWSTEAM

This repository is a Work In Progress provider for the [AWS TEAM](https://github.com/aws-samples/iam-identity-center-team) solution. I am currently waiting on the [feature](https://github.com/aws-samples/iam-identity-center-team/issues/33) which will enable programmatic access to the TEAM api. With the current provider, We are only able to read data from the api and not make any changes.  

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See the generated [docs](/docs/index.md)

## Developing the Provider

We welcome any contribution. The easiest way to start would be opening an issue. If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make gen`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources.

```shell
make testacc
```