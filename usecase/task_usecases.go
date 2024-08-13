package usecase

import (
	"context"
	"task/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: repo}
}

func (uc *TaskUsecase) CreateTask(ctx context.Context, task *domain.Task) error {
	return uc.TaskRepo.CreateTask(ctx, task)
}

func (uc *TaskUsecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return uc.TaskRepo.GetAllTasks(ctx)
}

func (uc *TaskUsecase) GetTasksByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	return uc.TaskRepo.GetTaskByID(ctx, id)
}

func (uc *TaskUsecase) UpdateTask(ctx context.Context, task *domain.Task) error {
	return uc.TaskRepo.UpdateTask(ctx, task)
}

func (uc *TaskUsecase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	return uc.TaskRepo.DeleteTask(ctx, id)
}
