# Go Startup Lambda
[![Go Report Card](https://goreportcard.com/badge/github.com/judegiordano/startup)](https://goreportcard.com/report/github.com/judegiordano/startup)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/judegiordano/startup/blob/main/LICENSE)
![example workflow](https://github.com/judegiordano/startup/actions/workflows/deploy_main.yml/badge.svg)
<!-- ![Lines of code](https://img.shields.io/tokei/lines/github.com/judegiordano/startup) -->

## Go Resources
---
[project layout](https://github.com/golang-standards/project-layout)

[packages](https://pkg.go.dev)

[go lambda api proxy](https://github.com/awslabs/aws-lambda-go-api-proxy)

[fiber](https://docs.gofiber.io)

[uber style guide](https://github.com/uber-go/guide/blob/master/style.md)

[go docs](https://go.dev/doc/)

[go by example](https://gobyexample.com/)

[practical go](https://www.practical-go-lessons.com/)

[effective go](https://go.dev/doc/effective_go)

[awesome go](https://github.com/avelino/awesome-go)

[go report card](https://goreportcard.com)

## Commands
---
```sh
# ginkgo
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@latest
ginkgo -r
# lint
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
golangci-lint run
```
## Example Dockerfile

```Dockerfile
FROM golang:1.19 as build

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV GOPATH /

RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/startup

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

CMD ["/app"]
	
```

```yml
# dependabot.yml
version: 2
updates:
  - package-ecosystem: gomod
    directory: /
    schedule:
      interval: daily
```
