package repository

import (
	"GoDatabase/entity"
	"context"
	"database/sql"
	"errors"
)

type CommentsRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentsRepository(db *sql.DB) CommentsRepository {
	return &CommentsRepositoryImpl{DB: db}
}

func (repository *CommentsRepositoryImpl) Insert(ctx context.Context, comments entity.Comments) (entity.Comments, error) {
	sqlExec := "INSERT INTO comments (title, comment) VALUES (?, ?)"                         // ? is placeholder
	result, err := repository.DB.ExecContext(ctx, sqlExec, comments.Title, comments.Comment) // exec query
	if err != nil {
		return comments, errors.New("error when insert data") // return error
	}
	id, err := result.LastInsertId() // get last insert id
	if err != nil {
		return comments, errors.New("error when get last insert id") // return error
	}
	comments.Id = int32(id) // set id to comments
	return comments, nil    // return comments
}

func (repository *CommentsRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	sqlExec := "SELECT id, title, comment FROM comments WHERE id = ?" // ? is placeholder
	result, err := repository.DB.QueryContext(ctx, sqlExec, id)       // exec query
	if err != nil {
		return entity.Comments{}, errors.New("error when find data by id")
	}
	defer func(result *sql.Rows) { // close result
		err := result.Close()
		if err != nil {
			panic(err.Error())
		}
	}(result)
	var comments entity.Comments // create comments
	if result.Next() {
		err := result.Scan(&comments.Id, &comments.Title, &comments.Comment) // scan result to comments
		if err != nil {
			return entity.Comments{}, errors.New("error when scan data to comments")
		}
	}
	return comments, nil // return comments
}

func (repository *CommentsRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	sqlQuery := "SELECT id, title, comment FROM comments"
	result, err := repository.DB.QueryContext(ctx, sqlQuery) // exec query
	if err != nil {
		return nil, errors.New("error when find all data")
	}
	defer func(result *sql.Rows) { // close result
		err := result.Close()
		if err != nil {
			panic(err.Error())
		}
	}(result)

	var comments []entity.Comments // create comments
	for result.Next() {
		var comment entity.Comments                                       // create comment
		err := result.Scan(&comment.Id, &comment.Title, &comment.Comment) // scan result to comment
		if err != nil {
			return nil, errors.New("error when scan data to comment")
		}
		comments = append(comments, comment) // append comment to comments
	}
	return comments, nil // return comments
}
