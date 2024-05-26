# `text` - String Utilities Package

The purpose of `text` is to provide string utilities, but with special considerations relating to web-service(s),
cloud providers, and external APIs.

Interfacing with packages such as AWS and Stripe GO SDKs often require working with pointer inputs and outputs. Such
cases can result in undesired, unexpected behavior. When writing API services, tracing issues can become difficult.

That's where the `text` package's [variadic options](./options.go) are perhaps the most notable; configurations
can be set that will log unexpected nil or zeroth input arguments.

## Usage

Refer to the [examples](./example_test.go) for usage and implementation details.

## Contributions

See the [**Contributing Guide**](./CONTRIBUTING.md) for additional details on getting started.
