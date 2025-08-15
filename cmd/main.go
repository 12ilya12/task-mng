package main

import (
	"encoding/json"
	"net/http"

	"github.com/12ilya12/task-mng/internal/controllers"
	"github.com/12ilya12/task-mng/internal/repos"
	"github.com/12ilya12/task-mng/internal/services"
)

func main() {
	//Первоначальная инициализация
	repo := repos.NewTaskInMemoryRepository()
	taskService := services.NewTaskService(repo)
	taskController := controllers.NewTaskController(taskService)

	//Задаём роуты
	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", taskController.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", taskController.GetTask)
	mux.HandleFunc("GET /tasks", taskController.GetAllTasks)

	//Ручка для проверки работоспособности сервера
	mux.HandleFunc("GET /alive", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Жив, цел, Орёл!")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		//TODO: Логируем ошибку
	}
}
