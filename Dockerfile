FROM golang:alpine
RUN apk add --no-cache curl
ENTRYPOINT ["tail", "-f", "/dev/null"]

