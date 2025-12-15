package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ArtDark/go-advanced/internal/auth"
	"github.com/ArtDark/go-advanced/internal/config"
)

func main() {

	conf := config.New().Load().Init()

	mux := http.NewServeMux()

	auth.NewAuthHandler(mux, auth.AuthHandlerDeps{
		Auth: &conf.Auth,
	})
	addr := ":8084"

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Выполняем net.Listen для нужного адреса
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	// Запускаем сервер в отдельной горутине
	log.Printf("Starting server on %s\n", addr)
	if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
	log.Println("Server stopped")

}
