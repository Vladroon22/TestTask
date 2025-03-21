.PHONY:

tests:
	go test ./internal/utils
	go test ./internal/service

run:
	go build -o ./app cmd/main.go
	./app

compose:
	sudo docker compose up -d

compose-stop:
	sudo docker compose down

compose-delete:
	sudo docker compose down -v
	sudo docker rmi testtask-app