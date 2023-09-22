check_install:
	which swagger || go get github.com/go-swagger/go-swagger/cmd/swagger

docs: check_install
	swagger generate spec -o ./swagger.yaml --scan-models