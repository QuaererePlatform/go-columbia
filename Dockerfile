ARG ALPINE_VERSION=3.10
ARG GOLANG_VERSION=1.13

FROM alpine:${ALPINE_VERSION} as certs
RUN apk --update add ca-certificates

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as builder
ENV CGO_ENABLED=0
WORKDIR /webapp
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o columbia ./cmd/server

FROM alpine:${ALPINE_VERSION}
WORKDIR /webapp
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /webapp/columbia .
COPY ./build/docker/prod/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
