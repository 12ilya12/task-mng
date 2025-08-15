package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/12ilya12/task-mng/internal/services"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskCreateDto struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&taskCreateDto); err != nil {
		http.Error(w, "Некорректное тело запроса для создания задачи", http.StatusBadRequest)
		return
	}

	task, err := tc.service.CreateTask(r.Context(), taskCreateDto.Title, taskCreateDto.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Требуется идентификатор", http.StatusBadRequest)
		return
	}

	task, err := tc.service.GetTaskByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	tasks, err := tc.service.GetAllTasks(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
