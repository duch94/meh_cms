infra-up:
	docker-compose up -d mongodb mongo-express

infra-down:
	docker-compose stop mongodb mongo-express

fmt:
	@echo "Autoformatting"
	go fmt cmd/main.go
	go fmt ./src/*
	go vet cmd/main.go
	go vet ./src/*
	go mod tidy
