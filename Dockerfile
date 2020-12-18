FROM alpine:3.12 as builder
RUN apk add --no-cache ca-certificates tzdata

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY sidecar-proxy /

ENTRYPOINT ["/sidecar-proxy"]
