FROM golang:1.10.2 as builder

COPY . /go/src/github.com/urbn/ordernumbergenerator/
WORKDIR /go/src/github.com/urbn/ordernumbergenerator/

#Unit test
RUN go test -cover -race ./...

#Build binary
RUN CGO_ENABLED=0 go build -v -o ordernumbergenerator app/cmd/main.go

FROM docker.urbn.com/urbn/alpine:3.7-s6

RUN mkdir /var/log/order-number-generator

EXPOSE 7070

WORKDIR /root/
COPY start ./start
RUN chmod +x start

ARG git_hash
ARG build_number
ARG branch

ENV GIT_HASH=${git_hash} \
    BUILD_NUMBER=${build_number} \
    BRANCH=${branch}

COPY --from=builder /go/src/github.com/urbn/ordernumbergenerator/ordernumbergenerator .

CMD ["./start"]