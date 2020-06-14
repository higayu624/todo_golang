FROM golang:1.14-alpine AS build-env

COPY . /clean-architecture-sample-src
WORKDIR /clean-architecture-sample-src/src/cmd/clean-architecture-sample

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -o clean-architecture-sample \
    && mv ./cleanarchitecture-sample /


FROM scratch

COPY --from=build-env /clean-architecture-sample /usr/local/bin/

EXPOSE 8080

