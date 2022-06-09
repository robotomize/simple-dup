.PHONY: lint
lint:
	golangci-lint run --timeout 5m --fix -v duplicate/*

.PHONY: test
test:
	go test -race -short duplicate/*
