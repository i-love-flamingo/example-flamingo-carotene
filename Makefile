CONTEXT?=dev
.PHONY: test serve

serve:
	go run main.go serve

test:
	go test -v ./...
