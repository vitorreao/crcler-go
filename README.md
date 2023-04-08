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
docker container run -d --rm -p 9000:8080 crcler-go:test /crcler-go
```

To get a request in, just something like:

```sh
curl "http://localhost:9000/accounts"
```
