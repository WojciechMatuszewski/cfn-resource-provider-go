# Custom CloudFormation Resource Provider in Go

## Resource Provider vs Custom Resource

External resources

- https://twitter.com/__steele/status/1219102291350786048
- https://youtu.be/MhHxEgCpCws

### Resource Providers

- You operate through a proxy. **The code that you have written will be invoked in an isolated account provided by `CloudFormation`**

- There are tools that generate a lot of boilerplate for you

- The _Namespace_ is different. With a _Resource Provider_ you have the ability to specify a triplet: `MySuper::CustomCustom::HereAsWell`.
  That is not the case with _Custom Resources_

- You have all the benefits of `CloudFormation` at your disposal. Rollbacks, progress events, changesets

- You can easily implement versioning. The resource itself is described by a schema. The schema has versions.
  The versioning is not tied to any kind of _Lambda Function_

### Custom Resources

- You create _Lambda Functions_. You have to worry about their executions roles and such

- The _Lambda Functions_ are created within your account

- Less integrated with `CloudFormation`. I would only use this for simple tasks

## Deployment

1. Run the `make build` command
2. Run the `make deploy` command
3. Use the SAM template within `examples` directory to create the stack

## Gotchas

- The tests events are different than the ones within Java and Python examples. See the `example_inputs/create.json` file

- The credentials that you use to test things have to be real credentials

- The golang plugin does not have the _nameUtils_ exposed by the java or python plugins

- Remember about versioning! The schema (thus you handlers) are versioned.
  You probably want to set the version you are deploying as the default one.
  See the `deploy` step within `Makefile`.
