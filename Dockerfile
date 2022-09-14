FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/doorman
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/doorman /go/src/doorman


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/doorman /usr/local/bin/doorman
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["doorman"]
