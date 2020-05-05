# build
FROM golang as build
WORKDIR /go/src/microsrv
COPY . .
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
COPY --from=build /go/bin/microsrv /go/bin/microsrv
EXPOSE 9090
CMD ["/go/bin/microsrv"]
