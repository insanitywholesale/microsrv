all:
	go build -v
install:
	go install -v
forcecheckinstall:
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
checkinstall:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models
