FROM golang:1.13.4 AS builder

WORKDIR /cloudsqldef

COPY . /cloudsqldef

ENV CGO_ENABLED=0

RUN go install -v github.com/proproto/cloudsqldef/cmd/mysqldef

FROM gcr.io/distroless/static

COPY --from=builder /go/bin/mysqldef /usr/local/bin/mysqldef

ENTRYPOINT ["mysqldef"]
