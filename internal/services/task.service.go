package services

import (
	"context"
	"strconv"
	"sync"

	"github.com/12ilya12/task-mng/internal/logger"
	"github.com/12ilya12/task-mng/internal/models"
	"github.com/12ilya12/task-mng/internal/repos"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*models.Task, error)
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	GetAllTasks(ctx context.Context, status string) ([]*models.Task, error)
}

type taskService struct {
	repo   repos.TaskRepo
	logger *logger.Logger
}

func NewTaskService(repo repos.TaskRepo, logger *logger.Logger) TaskService {
	return &taskService{repo: repo, logger: logger}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string) (*models.Task, error) {
	id := generateID()

	task := &models.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "New",
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	s.logger.Log("TASK", "Создана задача с идентификатором "+id)
	return task, nil
}

func (s *taskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Log("TASK", "Запрос на получение задачи с идентификатором "+id)
	return task, nil
}

func (s *taskService) GetAllTasks(ctx context.Context, status string) ([]*models.Task, error) {
	tasks, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	if status != "" {
		s.logger.Log("TASK", "Запрос на получение задач со статусом "+status)
	} else {
		s.logger.Log("TASK", "Запрос на получение всех задач")
	}
	return tasks, nil
}

var (
	idCounter int
	idMutex   sync.Mutex
)

func generateID() string {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return strconv.Itoa(idCounter)
}
