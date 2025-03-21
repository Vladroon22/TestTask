package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gthub.com/Vladroon22/TestTask/internal/database"
	"gthub.com/Vladroon22/TestTask/internal/handlers"
	"gthub.com/Vladroon22/TestTask/internal/repository"
	"gthub.com/Vladroon22/TestTask/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewDB().Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.NewRepo(db)
	srv := service.NewService(repo)
	h := handlers.NewHandler(srv)

	router := mux.NewRouter()

	router.HandleFunc("/users", h.CreateAccount).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", h.UpdateAccount).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", h.GetAccount).Methods("GET")

	go func() {
		if err := http.ListenAndServe(os.Getenv("addr"), router); err != nil {
			log.Fatalln(err)
		}
	}()

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, syscall.SIGINT, syscall.SIGTERM)

	<-killSig

	go func() {
		var wg sync.WaitGroup
		defer wg.Done()

		wg.Add(1)
		if err := db.Close(); err != nil {
			log.Println(err)
			return
		}
		wg.Wait()
	}()

	log.Println("Graceful shutdown")
}
