.PHONY:

run:
	go build -o ./app cmd/main.go
	./app

compose:
	sudo docker compose up -d

compose-stop:
	sudo docker compose down

compose-delete:
	sudo docker compose down -v
	sudo docker rmi cvmaker-cvmake