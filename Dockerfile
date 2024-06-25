FROM alpine:latest as certs
RUN apk --update add ca-certificates
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY go-contact-service ./
COPY config.env ./
EXPOSE 1327/tcp
ENTRYPOINT ["/go-contact-service"]