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
