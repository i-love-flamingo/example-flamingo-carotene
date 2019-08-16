CONTEXT?=dev
.PHONY: test serve frontend-build frontend-ci

serve:
	go run main.go serve

test:
	go test -v ./...


frontend-build:
	bash -c 'cd frontend && npx flamingo-carotene build'

frontend-ci:
	bash -c 'npm ci'