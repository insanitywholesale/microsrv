# build
FROM golang as build
WORKDIR /go/src/microsrv
COPY . .
#RUN apk update
#RUN apk add git
ENV CGO_ENABLED 0
RUN go get -d -v ./...
RUN go build -v -ldflags '-extldflags "-static"' ./...
RUN go install -v ./...

# run
FROM busybox
RUN adduser -D -u 5000 app
USER app:app
WORKDIR /go/bin/
ENV GOPATH /go
COPY --from=build /go/bin/microsrv /go/bin/microsrv
EXPOSE 9090
CMD ["/go/bin/microsrv"]
