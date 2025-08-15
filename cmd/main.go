package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/12ilya12/task-mng/internal/controllers"
	"github.com/12ilya12/task-mng/internal/logger"
	"github.com/12ilya12/task-mng/internal/repos"
	"github.com/12ilya12/task-mng/internal/services"
)

func main() {
	//Запуск логировалки
	loggerInst := logger.NewLogger()
	loggerInst.Start()
	defer loggerInst.Stop()

	//Инициализация компонентов
	repo := repos.NewTaskInMemoryRepository()
	taskService := services.NewTaskService(repo, loggerInst)
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

	//Инициализация сервера
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	//Канал, отслеживающий остановку сервера для поддержки graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		loggerInst.Log("SERVER", "Запуск сервера :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggerInst.Log("SERVER", "Сбой в работе сервера: "+err.Error())
			os.Exit(1)
		}
	}()

	//Ожидаем остановки сервера
	<-shutdownChan
	loggerInst.Log("SERVER", "Завершение работы сервера...")

	//Ожидаем три секунды для завершения текущих задач (в нашем случае задач логирования)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		loggerInst.Log("SERVER", "Ошибка при завершении работы сервера: "+err.Error())
	} else {
		loggerInst.Log("SERVER", "Завершение работы сервера выполнено успешно")
	}
}
