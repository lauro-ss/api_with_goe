package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lauro-ss/api_with_goe/internal/data"
	"github.com/lauro-ss/api_with_goe/internal/handlers"
	"github.com/lauro-ss/api_with_goe/internal/service"
)

func main() {
	dns := "user=postgres password=postgres host=localhost port=5432 database=postgres"
	db, err := data.OpenAndMigrate(dns)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := service.NewUserRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /user", handlers.UserList(userRepository))
	mux.HandleFunc("POST /user", handlers.UserCreate(userRepository))
	mux.HandleFunc("PUT /user/:id", handlers.UserUpdate(userRepository))
	mux.HandleFunc("DELETE /user/:id", handlers.UserDelete(userRepository))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("Server is running on 8080")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	log.Println("Closing the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	log.Println("Server gracefully shutdown")

}
