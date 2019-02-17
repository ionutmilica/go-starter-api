.PHONY: test

test:
	go vet ./... && go fmt ./... && golint -set_exit_status ./... && go test -v -cover ./...
