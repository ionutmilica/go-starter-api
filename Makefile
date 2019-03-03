.PHONY: test run-all run-deps

test:
	go vet ./... && go fmt ./... && golint -set_exit_status ./... && go test -v -cover ./...

run-all:
	./dev.sh start_all

run-deps:
	./dev.sh start_deps
