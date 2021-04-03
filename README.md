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

## Learnings

- Event shape are different than anywhere on the internet, you should dig into a file called `events.go`

- WTF is protocol 2.0.0 ?

- the schemas (and the resource itself) have versions? If you want to update you should probably use `cfn submit --set-default`

- `TAGS=logging make` to make sure the logs are properly forwarded to cloudwatch

- wtf is happening with logs?

- the golang plugin does not have the _nameUtils_ exposed by the java or python plugins

## Deployment

1. Run the `make build` command
2. Run the `make deploy` command
3. Use the SAM template within `examples` directory to create the stack

## Gotchas

- The tests events are different than the ones within Java and Python examples. See the `example_inputs/create.json` file

- The credentials that you use to test things have to be real credentials
