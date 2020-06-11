FROM golang:1.13.4 AS builder

WORKDIR /cloudsqldef

COPY . /cloudsqldef

ENV CGO_ENABLED=0

RUN go install -v github.com/proproto/cloudsqldef/cmd/mysqldef

FROM alpine:3.12.0

COPY --from=builder /go/bin/mysqldef /usr/local/bin/mysqldef

ENTRYPOINT ["mysqldef"]
