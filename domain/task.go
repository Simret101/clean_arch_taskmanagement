package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string

const (
	TaskStatusComplete   TaskStatus = "complete"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusStarted    TaskStatus = "started"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"duedate" bson:"duedate"`
	Status      TaskStatus         `json:"status" bson:"status"`
	UserID      int                `json:"userID" bson:"userID"`
}

type TaskRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(ctx context.Context, id primitive.ObjectID) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
}
