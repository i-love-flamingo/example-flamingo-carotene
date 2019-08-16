CONTEXT?=dev
.PHONY: test serve frontend-build frontend-ci frontend

serve:
	go run main.go serve

test:
	go test -v ./...


frontend-build:
	bash -c 'cd frontend && npx flamingo-carotene build'

frontend-ci:
	bash -c 'npm ci'


frontend: frontend-build
	bash -c 'cd frontend && npx flamingo-carotene dev'