package repository

import (
	"GoDatabase"
	"GoDatabase/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentsRepository(t *testing.T) {
	db, ctx := GoDatabase.GetConnection()
	commentsRepository := NewCommentsRepository(db)
	t.Run("Test Insert", func(t *testing.T) {
		comment := entity.Comments{
			Title:   "TEST INSERT WOY",
			Comment: "TEST INSERT COMMENT WOY!",
		}
		result, err := commentsRepository.Insert(ctx, comment)
		if err != nil {
			panic(err.Error())
		}
		fmt.Print("Insert Success! ID:", result)
	})
	t.Run("FindById", func(t *testing.T) {
		comment, err := commentsRepository.FindById(ctx, 42)
		if err != nil {
			panic(err.Error())
		}
		fmt.Print("FindById Success! Comment:", comment)
	})
	t.Run("FindAll", func(t *testing.T) {
		comments, err := commentsRepository.FindAll(ctx)
		if err != nil {
			panic(err.Error())
		}
		for _, comment := range comments { // loop through the comments
			fmt.Println("FindAll Success! Comment:", comment)
		}
	})
}
