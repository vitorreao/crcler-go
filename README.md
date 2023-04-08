# crcler-go

This repository contains the source code for the crcler backend API, writen in
[Go](https://go.dev/).

## Running locally

To run this project locally, the easiest way is to use
[Docker](https://docker.com/). You can build the docker image by running the
following commands:

```sh
docker image build -t crcler-go:test .
```

You can also use any other tag name you want, instead of `test`.

Then, with the image built, you can run a container by executing the following:

```sh
docker container run -d --rm -v ~/.aws-lambda-rie:/aws-lambda \
--entrypoint /aws-lambda/aws-lambda-rie -p 9000:8080 crcler-go:test /crcler-go
```

Before you do so, don't forget to download the AWS Lambda Runtime Interface
Emulator (RIE). It must be on the docker volume we attach to the container.
More information on that can be found
[here](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html#go-image-other).

To get a request in, just run:

```sh
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" \
-d '{"Name": "Doc"}'
```
