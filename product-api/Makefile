check_install:
	which swagger || GOO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	Go111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models