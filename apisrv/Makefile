all:
	go build -v
install:
	go install -v
forcecheckinstall:
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
checkinstall:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
generate_spec:
	swagger generate spec -o ./swagger.yml --scan-models
generate_client:
	cd client && swagger generate client -f ../swagger.yml -A client
