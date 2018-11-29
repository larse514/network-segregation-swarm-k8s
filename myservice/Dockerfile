# build stage pattern from https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a
FROM golang:alpine AS build-env
ADD /src /src
RUN go get net/http
RUN go get encoding/json
WORKDIR /src
RUN go build -o hello .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/hello /app/
ENTRYPOINT ./hello