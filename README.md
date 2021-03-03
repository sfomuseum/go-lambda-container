# go-lambda-container

Example code for developing Go applications that can be run from the command line, as AWS Lambda functions and containerized Lambda functions.

## Important

This is work in progress. Through documentation to follow.

## What is this?

This is an example package implementing a simple "hello world" application in the Go programming language. The idea is to outline and document the steps and patterns necessary to write a single application that can be run from the command line, as an AWS Lambda function or a containerized Lambda function. 

## Command line

```
$> make cli
go build -mod vendor -o bin/hello-world cmd/hello-world/main.go

$> ./bin/hello-world 
Hello world
```

## Lambda

```
$> make lambda
if test -f main; then rm -f main; fi
if test -f hello-world.zip; then rm -f hello-world.zip; fi
GOOS=linux go build -mod vendor -o main cmd/hello-world/main.go
zip hello-world.zip main
  adding: main (deflated 48%)
rm -f main
```

Create a new Lambda function, in AWS, using `hello-world.zip` as the source code. Ensure that the Lambda handler is configured to be `main`. The function itself does not need any special permissions to the default role, that AWS will create by default, is sufficient.

Ensure the following environment variables are assigned:

| Name | Value |
| --- | --- |
| SFOMUSEUM_MODE | lambda |

Create an empty test (`{}`) and run it. It should succeed with a `null` output, writing the phrase "Hello world" to the function's log file.

## Lambda (using a container image)

```
$> make docker
docker build -f Dockerfile -t hello-world .
Sending build context to Docker daemon  12.82MB

...Docker stuff happens

Successfully built 97c609afd399
Successfully tagged hello-world:latest
```

Tag and push the `hello-world` container image to an AWS ECS repository.

Create a new Lambda function, in AWS, using `hello-world` container image as the source code. The function itself does not need any special permissions to the default role, that AWS will create by default, is sufficient.

Ensure the following container image configuration values:

| Name | Value |
| --- | --- |
| CMD override | /main |

Ensure the following environment variables are assigned:

| Name | Value |
| --- | --- |
| SFOMUSEUM_MODE | lambda |

Create an empty test (`{}`) and run it. It should succeed with a `null` output, writing the phrase "Hello world" to the function's log file.