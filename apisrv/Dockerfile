# build
FROM golang:1.14 as build
WORKDIR /go/src/microsrv
COPY . .
WORKDIR /go/src/microsrv/apisrv
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN go get -d -v ./...
RUN go install -v ./...

# run
FROM busybox:glibc
RUN adduser -D -u 5000 app
USER app:app
WORKDIR /go/bin/
COPY --from=build /go/bin/apisrv /microsrv
EXPOSE 9090
CMD ["/microsrv"]
