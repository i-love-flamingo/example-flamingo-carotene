CONTEXT?=dev
.PHONY: test serve frontend-build

serve:
	go run main.go serve

test:
	go test -v ./...


frontend-build:
	bash -c 'cd frontend && npm ci && npx flamingo-carotene build'