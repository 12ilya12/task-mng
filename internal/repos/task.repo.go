package repos

import (
	"github.com/12ilya12/task-mng/internal/models"
)

type TaskRepo interface {
	Create(task *models.Task) error
	FindByID(id string) (*models.Task, error)
	FindAll(status string) ([]*models.Task, error)
}
