FROM alpine:latest as certs
RUN apk --no-cache add openssl
WORKDIR /certs
RUN /usr/bin/openssl req -x509 -newkey rsa:2048 -nodes -out cert.pem -keyout key.pem -days 365 -subj "/C=SK/ST=The Reach/L=Rifton/O=Theives Guild"

FROM cosmtrek/air:latest
COPY --from=certs /certs /tls
ENTRYPOINT ["/go/bin/air"]
