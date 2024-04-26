package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/molestov/go_final_project/database"
	"github.com/molestov/go_final_project/handlers"
)

func getPort() string {
	const default_port = "7540"
	if val, exists := os.LookupEnv("TODO_PORT"); exists {
		return ":" + val
	}
	return ":" + default_port
}

func main() {
	const webDir = "./web"

	fmt.Println("DB connect")
	database.ConnectDB()
	fmt.Println("DB connected")
	myHandler := chi.NewRouter()

	fmt.Println("Register handlers")
	myHandler.Mount("/", http.FileServer(http.Dir(webDir)))
	myHandler.Get("/api/nextdate", handlers.NextDateGET)
	myHandler.Post("/api/task", handlers.TaskPost)
	myHandler.Get("/api/tasks", handlers.TasksRead)
	myHandler.Get("/api/task", handlers.TaskReadByID)
	myHandler.Put("/api/task", handlers.TaskUpdate)
	myHandler.Post("/api/task/done", handlers.TaskDone)
	myHandler.Delete("/api/task", handlers.TaskDelete)

	fmt.Printf("Starting server on port %s\n", getPort())

	s := &http.Server{
		Addr:           getPort(),
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
