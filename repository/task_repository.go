package repository

import (
	"context"
	"errors"
	"task/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		collection: db.Collection("tasks"),
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	_, err := r.collection.InsertOne(ctx, task)
	return err
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}
func (r *TaskRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"duedate":     task.DueDate,
			"status":      task.Status,
			"userID":      task.UserID,
		},
	}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": task.ID}, update)
	return err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
