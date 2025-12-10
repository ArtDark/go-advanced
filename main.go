package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
)

// main является точкой входа в программу
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	addr := ":8084"

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Создаем канал готовности
	ready := make(chan struct{})

	// Выполняем net.Listen для нужного адреса
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	// Сигнализируем готовность закрытием канала
	close(ready)

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("Starting server on %s\n", addr)
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
		log.Println("Server stopped")
	}()

	// Дожидаемся готовности сервера
	<-ready

	// Выполняем клиентский запрос
	testUrl := url.URL{
		Scheme: "http",
		Host:   addr,
		Path:   "test",
	}

	resp, err := http.Get(testUrl.String())
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Status: %d\n", resp.StatusCode)
	}

	// Запуск graceful shutdown в отдельной горутине
	go gracefulShutdown(srv)

}

// gracefulShutdown корректно завершает работу HTTP сервера при получении сигнала прерывания
func gracefulShutdown(srv *http.Server) {
	go func() {

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()
}
