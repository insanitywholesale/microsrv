# build
FROM golang as build
WORKDIR /go/src/microsrv
COPY . .
WORKDIR /go/src/microsrv/imagesrv
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN go get -d -v ./...
RUN go install -v ./...

# run
FROM busybox:glibc
RUN adduser -D -u 5000 app
WORKDIR /imagestore
RUN chown app:app /imagestore
USER app:app
COPY --from=build /go/bin/imagesrv /imagesrv
COPY --from=build /go/src/microsrv/imagesrv/imagestore /imagestore
WORKDIR /
EXPOSE 9091
CMD ["/imagesrv"]
