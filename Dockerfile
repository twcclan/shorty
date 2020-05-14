FROM golang:alpine as builder

RUN apk --no-cache add git build-base

# need to be outside of GOPATH for module support
WORKDIR /code

# have this separate for caching purposes
COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY . .
RUN mkdir /shorty

RUN go build -o /shorty/binary .

FROM alpine as runner

WORKDIR /root

COPY --from=builder /shorty/binary /root/

ENTRYPOINT ["/root/binary"]
