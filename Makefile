check_install:
	which swagger || (brew tap go-swagger/go-swagger && brew install go-swagger && go get github.com/go-swagger/go-swagger/cmd/swagger)

docs: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

start_db:
	docker compose up


