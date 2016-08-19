FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-rollbar /bin/
ENTRYPOINT ["/bin/drone-rollbar"]
