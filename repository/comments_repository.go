package repository

import (
	"GoDatabase/entity"
	"context"
)

type CommentsRepository interface {
	Insert(ctx context.Context, comments entity.Comments) (entity.Comments, error) // insert data, return data with ID
	FindById(ctx context.Context, id int32) (entity.Comments, error)               // find data by ID
	FindAll(ctx context.Context) ([]entity.Comments, error)                        // find all data
}
