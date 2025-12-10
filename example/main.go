package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kevindiu/gotest2/example/handler"
	"github.com/kevindiu/gotest2/example/model"
	"github.com/kevindiu/gotest2/example/repository"
	"github.com/kevindiu/gotest2/example/service"
)

func main() {
	// 1. Init Repository
	repo := repository.NewMemoryRepository[model.Book, string]()

	// 2. Init Service
	svc := service.NewBookService(repo)

	// 3. Init Handler
	h := handler.NewBookHandler(svc)

	// 4. Register Routes
	http.HandleFunc("/book/create", h.CreateBookHandler)
	http.HandleFunc("/book/get", h.GetBookHandler)
	http.HandleFunc("/book/list", h.ListBooksHandler)

	// 5. Start Server
	port := getPort()
	fmt.Printf("Starting server on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		if port[0] != ':' {
			return ":" + port
		}
		return port
	}
	return ":8080"
}
