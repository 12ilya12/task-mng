package repos

import (
	"errors"
	"sync"

	"github.com/12ilya12/task-mng/internal/models"
)

type taskInMemoryRepo struct {
	mu    sync.RWMutex
	tasks map[string]*models.Task
}

func NewTaskInMemoryRepository() TaskRepo {
	return &taskInMemoryRepo{
		tasks: make(map[string]*models.Task),
	}
}

func (r *taskInMemoryRepo) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *taskInMemoryRepo) FindByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *taskInMemoryRepo) FindAll(status string) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := []*models.Task{}
	for _, task := range r.tasks {
		//Фильтруем по статусу, если он задан
		if status == "" || task.Status == status {
			result = append(result, task)
		}
	}
	return result, nil
}
